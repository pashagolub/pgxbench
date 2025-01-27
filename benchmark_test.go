package main

import (
	"context"
	"errors"
	"testing"

	pgx "github.com/jackc/pgx/v5"
	pgxmock "github.com/pashagolub/pgxmock/v4"
)

func TestInsertSimple(t *testing.T) {
	ctx := context.Background()
	conn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}

	t.Run("Successful InsertSimple", func(t *testing.T) {
		*numberOfRows = 10
		conn.ExpectExec(`INSERT.*`).
			WithArgs(1, "John", 42, `{"role": "developer"}`).
			WillReturnResult(pgxmock.NewResult("INSERT", 1)).
			Times(10)
		if err = InsertSimple(ctx, conn); err != nil {
			t.Fatalf("Failed to execute InsertSimple: %v\n", err)
		}
		if err = conn.ExpectationsWereMet(); err != nil {
			t.Fatalf("Expectations were not met: %v\n", err)
		}
	})

	t.Run("Failed InsertSimple", func(t *testing.T) {
		conn.ExpectExec(`INSERT.*`).
			WithArgs(1, "John", 42, `{"role": "developer"}`).
			WillReturnError(errors.New("failed to insert"))
		if err = InsertSimple(ctx, conn); err == nil {
			t.Fatalf("Expecting error, got nil\n")
		}
		if err = conn.ExpectationsWereMet(); err != nil {
			t.Fatalf("Expectations were not met: %v\n", err)
		}
	})
}

func TestInsertBatch(t *testing.T) {
	ctx := context.Background()
	conn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}

	t.Run("Successful InsertBatch", func(t *testing.T) {
		*numberOfRows = 1
		*batchSize = 1
		conn.ExpectBatch().
			ExpectExec(`INSERT.*`).
			WithArgs(1, "John", 42, `{"role": "developer"}`).
			WillReturnResult(pgxmock.NewResult("INSERT", 100))
		if err = InsertBatch(ctx, conn); err != nil {
			t.Fatalf("Failed to execute InsertBatch: %v\n", err)
		}
		if err = conn.ExpectationsWereMet(); err != nil {
			t.Fatalf("Expectations were not met: %v\n", err)
		}
	})

	t.Run("Failed InsertBatch", func(t *testing.T) {
		*numberOfRows = 1
		*batchSize = 1
		conn.ExpectBatch().
			ExpectExec(`INSERT.*`).
			WithArgs(1, "John", 42, `{"role": "developer"}`).
			WillReturnError(errors.New("failed to insert"))
		if err = InsertBatch(ctx, conn); err == nil {
			t.Fatalf("Expecting error, got nil\n")
		}
		if err = conn.ExpectationsWereMet(); err != nil {
			t.Fatalf("Expectations were not met: %v\n", err)
		}
	})
}

func TestInsertCopy(t *testing.T) {
	ctx := context.Background()
	conn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}

	*numberOfRows = 1000
	conn.ExpectCopyFrom(pgx.Identifier{"test"}, []string{"id", "name", "age", "meta"}).WillReturnResult(1000)
	if err = InsertCopy(ctx, conn); err != nil {
		t.Fatalf("Failed to execute InsertCopy: %v\n", err)
	}
	if err = conn.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations were not met: %v\n", err)
	}
}

func TestFetchSelectScan(t *testing.T) {
	ctx := context.Background()
	conn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}

	t.Run("Successful FetchSelectScan", func(t *testing.T) {
		*numberOfRows = 2
		rows := pgxmock.NewRows([]string{"id", "name", "age", "meta"}).
			AddRow(1, "John", 42, `{"role": "developer"}`).
			AddRow(2, "Jane", 24, `{"role": "manager"}`)
		conn.ExpectQuery(`SELECT.*`).
			WithArgs(2).
			WillReturnRows(rows)
		if err = FetchSelectScan(ctx, conn); err != nil {
			t.Fatalf("Failed to execute FetchSelectScan: %v\n", err)
		}
		if err = conn.ExpectationsWereMet(); err != nil {
			t.Fatalf("Expectations were not met: %v\n", err)
		}
	})

	t.Run("Failed FetchSelectScan", func(t *testing.T) {
		*numberOfRows = 2
		conn.ExpectQuery(`SELECT.*`).
			WithArgs(2).
			WillReturnError(errors.New("failed to fetch"))
		if err = FetchSelectScan(ctx, conn); err == nil {
			t.Fatalf("Expecting error, got nil\n")
		}
		if err = conn.ExpectationsWereMet(); err != nil {
			t.Fatalf("Expectations were not met: %v\n", err)
		}
	})

	t.Run("Failed FetchSelectScan rows scan", func(t *testing.T) {
		*numberOfRows = 2
		rows := pgxmock.NewRows([]string{"id", "name", "age", "meta"}).
			AddRow(1, "John", 42, `{"role": "developer"}`).
			AddRow(2, "Jane", 24, `{"role": "manager"}`).RowError(1, errors.New("failed to scan"))
		conn.ExpectQuery(`SELECT.*`).
			WithArgs(2).
			WillReturnRows(rows)
		if err = FetchSelectScan(ctx, conn); err == nil {
			t.Fatalf("Expecting error, got nil\n")
		}
		if err = conn.ExpectationsWereMet(); err != nil {
			t.Fatalf("Expectations were not met: %v\n", err)
		}
	})
}

func TestFetchSelectCollect(t *testing.T) {
	ctx := context.Background()
	conn, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}

	t.Run("Successful FetchSelectScan", func(t *testing.T) {
		*numberOfRows = 2
		rows := pgxmock.NewRows([]string{"id", "name", "age", "meta"}).
			AddRow(1, "John", 42, `{"role": "developer"}`).
			AddRow(2, "Jane", 24, `{"role": "manager"}`)
		conn.ExpectQuery(`SELECT.*`).
			WithArgs(2).
			WillReturnRows(rows)
		if err = FetchSelectCollect(ctx, conn); err != nil {
			t.Fatalf("Failed to execute FetchSelectScan: %v\n", err)
		}
		if err = conn.ExpectationsWereMet(); err != nil {
			t.Fatalf("Expectations were not met: %v\n", err)
		}
	})

	t.Run("Failed FetchSelectScan", func(t *testing.T) {
		*numberOfRows = 2
		conn.ExpectQuery(`SELECT.*`).
			WithArgs(2).
			WillReturnError(errors.New("failed to fetch"))
		if err = FetchSelectCollect(ctx, conn); err == nil {
			t.Fatalf("Expecting error, got nil\n")
		}
		if err = conn.ExpectationsWereMet(); err != nil {
			t.Fatalf("Expectations were not met: %v\n", err)
		}
	})

	t.Run("Failed FetchSelectScan rows scan", func(t *testing.T) {
		*numberOfRows = 2
		rows := pgxmock.NewRows([]string{"id", "name", "age", "meta"}).
			AddRow(1, "John", 42, `{"role": "developer"}`).
			AddRow(2, "Jane", 24, `{"role": "manager"}`).RowError(1, errors.New("failed to scan"))
		conn.ExpectQuery(`SELECT.*`).
			WithArgs(2).
			WillReturnRows(rows)
		if err = FetchSelectCollect(ctx, conn); err == nil {
			t.Fatalf("Expecting error, got nil\n")
		}
		if err = conn.ExpectationsWereMet(); err != nil {
			t.Fatalf("Expectations were not met: %v\n", err)
		}
	})
}
