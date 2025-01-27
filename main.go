package main

import (
	"context"
	"flag"
	"log"
)

var connStr = flag.String("c", "", "connection string to PostgreSQL database")

func main() {
	ctx := context.Background()
	flag.Parse()
	conn, err := ConnectToDB(ctx, *connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer CloseDB(ctx, conn)
	if err = InitDB(ctx, conn); err != nil {
		log.Fatalf("Unable to initialise test schema: %v\n", err)
	}
}
