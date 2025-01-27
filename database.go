package main

import (
	"context"
	"errors"
	"log"
	"time"

	pgx "github.com/jackc/pgx/v5"
	pgconn "github.com/jackc/pgx/v5/pgconn"
	pgpool "github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) (br pgx.BatchResults)
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

type DbUser struct {
	Id   int
	Name string
	Age  int
	Meta string
}

var TestUser DbUser = DbUser{Id: 1, Name: "John", Age: 42, Meta: `{"role": "developer"}`}

var NewConn = func(ctx context.Context, connStr string) (Database, error) {
	return pgpool.New(ctx, connStr)
}

// ConnectToDB connects to the database and tries to execute empty query to check the connection
func ConnectToDB(ctx context.Context, connStr string) (Database, error) {
	conn, err := NewConn(ctx, connStr)
	if err != nil {
		return nil, err
	}
	var version string
	err = conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	log.Printf("Connected to database on %s", version)
	return conn, err
}

// InitDB creates a table and try to insert a row to check it the schema is correct
func InitDB(ctx context.Context, conn Database) error {
	_, err := conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS test (id bigint, name text, age int, meta jsonb)")
	if err != nil {
		return err
	}
	var tx pgx.Tx
	if tx, err = conn.Begin(ctx); err != nil {
		return err
	}
	_, err = tx.Exec(ctx, `INSERT INTO test (id, name, age, meta) VALUES (1, 'John', 25, '{"role": "dveloper"}')`)
	defer tx.Rollback(ctx)
	return err
}

// CloseDB drops a test table and closes the connetion to database
func CloseDB(ctx context.Context, conn Database) error {
	defer func() {
		switch c := conn.(type) {
		case *pgpool.Pool:
			c.Close()
		case *pgx.Conn:
			c.Close(ctx)
		}
	}()
	_, err := conn.Exec(ctx, "DROP TABLE test")
	return err
}

// RunBenchmarks will run benchmarks available and output results
func RunBenchmarks(ctx context.Context, conn Database) (err error) {

	report := func(name string, f func(context.Context, Database) error) {
		t := time.Now()
		log.Println("Starting", name)
		e := f(ctx, conn)
		d := time.Since(t)
		if e != nil {
			err = errors.Join(err, e)
			log.Printf(`Error running "%s": %v`, name, err)
			return
		}
		log.Printf("Finished %s in %dms\n", name, d.Milliseconds())
	}

	report("Insert row by row", InsertSimple)
	report("Insert in batch", InsertBatch)
	report("Insert using copy", InsertCopy)
	report("Select, then Scan()", FetchSelectScan)
	report("Select, then CollectRows()", FetchSelectCollect)
	return
}
