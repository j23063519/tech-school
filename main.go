package main

import (
	"database/sql"
	"log"

	"github.com/j23063519/tech-school/api"
	db "github.com/j23063519/tech-school/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://user:pwd@localhost:5432/tech_school?sslmode=disable"
	serverAddress = "0.0.0.0:8123"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can't connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("connot start versern", err)
	}
}
