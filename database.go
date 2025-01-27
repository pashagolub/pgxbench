package main

import (
	"context"
	"log"

	pgx "github.com/jackc/pgx/v5"
)

// ConnectToDB connects to the database and tries to execute empty query to check the connection
func ConnectToDB(ctx context.Context, connStr string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	var version string
	err = conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	log.Printf("Connected to database %s on %s", conn.Config().Database, version)
	return conn, err
}

// InitDB creates a table and try to insert a row to check it the schema is correct
func InitDB(ctx context.Context, conn *pgx.Conn) error {
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
func CloseDB(ctx context.Context, conn *pgx.Conn) error {
	defer conn.Close(ctx)
	_, err := conn.Exec(ctx, "DROP TABLE test")
	return err
}

// RunBenchmarks will run benchmarks available and output results
func RunBenchmarks(ctx context.Context, conn *pgx.Conn) error {
	return conn.Ping(ctx)
}
