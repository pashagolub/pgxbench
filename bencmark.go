package main

import (
	"context"

	pgx "github.com/jackc/pgx/v5"
)

func InsertSimple(ctx context.Context, conn *pgx.Conn) error {
	for i := 0; i < *numberOfRows; i++ {
		_, err := conn.Exec(ctx, `INSERT INTO test (id, name, age, meta) VALUES ($1, $2, $3, $4)`,
			TestUser.Id, TestUser.Name, TestUser.Age, TestUser.Meta)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertBatch(ctx context.Context, conn *pgx.Conn) error {
	for i := 0; i < *numberOfRows; i = i + *batchSize {
		batch := &pgx.Batch{}
		for b := 0; b < *batchSize; b++ {
			batch.Queue(`INSERT INTO test (id, name, age, meta) VALUES ($1, $2, $3, $4)`,
				TestUser.Id, TestUser.Name, TestUser.Age, TestUser.Meta)
		}
		br := conn.SendBatch(ctx, batch)
		if err := br.Close(); err != nil {
			return err
		}
	}
	return nil
}

func InsertCopy(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.CopyFrom(ctx, pgx.Identifier{"test"},
		[]string{"id", "name", "age", "meta"},
		pgx.CopyFromSlice(*numberOfRows, func(i int) ([]any, error) {
			return []any{TestUser.Id, TestUser.Name, TestUser.Age, TestUser.Meta}, nil
		}))
	return err
}

func FetchSelectScan(ctx context.Context, conn *pgx.Conn) error {
	rows, err := conn.Query(ctx, "SELECT id, name, age, meta FROM test LIMIT $1", *numberOfRows)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var user DbUser
		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Meta)
		if err != nil {
			return err
		}
	}
	return nil
}

func FetchSelectCollect(ctx context.Context, conn *pgx.Conn) error {
	rows, err := conn.Query(ctx, "SELECT id, name, age, meta FROM test LIMIT $1", *numberOfRows)
	if err != nil {
		return err
	}
	defer rows.Close()
	_, err = pgx.CollectRows[DbUser](rows, pgx.RowToStructByPos)
	if err != nil {
		return err
	}
	return nil
}
