package main

import (
	"context"
	"fmt"
	"log"

	pgx "github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgresql://pasha@localhost:5432/Test")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())
	var version string
	_ = conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	fmt.Printf("Connected to database %s on %s", conn.Config().Database, version)
}
