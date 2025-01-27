package main

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	pgxmock "github.com/pashagolub/pgxmock/v4"
)

func TestConnectToDB(t *testing.T) {
	conn, err := pgxmock.NewPool()
	ctx := context.Background()

	t.Run("Failed default NewConn", func(t *testing.T) {
		if _, e := ConnectToDB(ctx, "mailformed connection string"); e == nil {
			t.Fatalf("Expecting error, got nil\n")
		}
	})

	t.Run("Successful ConnectToDB", func(t *testing.T) {
		NewConn = func(ctx context.Context, connStr string) (Database, error) {
			return conn, err
		}
		conn.ExpectQuery("SELECT version()").
			WillReturnRows(pgxmock.NewRows([]string{"version"}).
				AddRow("PostgreSQL 17.2"))
		if _, err = ConnectToDB(ctx, "postgresql://localhost/test"); err != nil {
			t.Fatalf("Failed to execute ConnectToDB: %v\n", err)
		}
		if err = conn.ExpectationsWereMet(); err != nil {
			t.Fatalf("Expectations were not met: %v\n", err)
		}
	})

	t.Run("Failed ConnectToDB", func(t *testing.T) {
		NewConn = func(ctx context.Context, connStr string) (Database, error) {
			return nil, errors.New("failed to connect")
		}
		if _, err = ConnectToDB(ctx, "postgresql://localhost/test"); err == nil {
			t.Fatalf("Expecting error, got nil\n")
		}
	})
}

func TestInitDB(t *testing.T) {
	conn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}
	ctx := context.Background()

	t.Run("Successful InitDB", func(t *testing.T) {
		conn.ExpectExec("CREATE TABLE IF NOT EXISTS test").
			WillReturnResult(pgxmock.NewResult("CREATE", 1))
		conn.ExpectBegin()
		conn.ExpectExec("INSERT INTO test").
			WillReturnResult(pgxmock.NewResult("INSERT", 1))
		conn.ExpectRollback()
		if err = InitDB(ctx, conn); err != nil {
			t.Fatalf("Failed to execute InitDB: %v\n", err)
		}
		if err = conn.ExpectationsWereMet(); err != nil {
			t.Fatalf("Expectations were not met: %v\n", err)
		}
	})

	t.Run("Failed CREATE TABLE", func(t *testing.T) {
		conn.ExpectExec("CREATE TABLE IF NOT EXISTS test").
			WillReturnError(errors.New("failed to create table"))
		if err = InitDB(ctx, conn); err == nil {
			t.Fatalf("Expecting error, got nil\n")
		}
	})

	t.Run("Failed BEGIN", func(t *testing.T) {
		conn.ExpectExec("CREATE TABLE IF NOT EXISTS test").
			WillReturnResult(pgxmock.NewResult("CREATE", 1))
		conn.ExpectBegin().WillReturnError(errors.New("failed to begin"))
		if err = InitDB(ctx, conn); err == nil {
			t.Fatalf("Expecting error, got nil\n")
		}
	})
}

func TestCloseDB(t *testing.T) {
	conn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}
	ctx := context.Background()

	t.Run("Successful CloseDB", func(t *testing.T) {
		conn.ExpectExec("DROP TABLE test").
			WillReturnResult(pgxmock.NewResult("DROP", 1))
		if err = CloseDB(ctx, conn); err != nil {
			t.Fatalf("Failed to execute CloseDB: %v\n", err)
		}
		if err = conn.ExpectationsWereMet(); err != nil {
			t.Fatalf("Expectations were not met: %v\n", err)
		}
	})
}

func TestRunBenchmarks(t *testing.T) {
	conn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}
	ctx := context.Background()

	*numberOfRows = 1
	*batchSize = 1
	conn.ExpectExec("INSERT INTO test").
		WithArgs(1, "John", 42, `{"role": "developer"}`).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))
	conn.ExpectBatch().
		ExpectExec("INSERT INTO test").
		WithArgs(1, "John", 42, `{"role": "developer"}`).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))
	conn.ExpectCopyFrom(pgx.Identifier{"test"}, []string{"id", "name", "age", "meta"}).
		WillReturnResult(1)
	conn.ExpectQuery("SELECT id, name, age, meta FROM test").
		WithArgs(1).
		Times(2).
		WillReturnError(errors.New("failed to select"))
	if err = RunBenchmarks(ctx, conn); err == nil {
		t.Fatal("Expecting error, got nil\n")
	}
	if err = conn.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations were not met: %v\n", err)
	}

}
