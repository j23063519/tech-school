package main

import (
	"database/sql"
	"log"

	"github.com/j23063519/tech-school/api"
	db "github.com/j23063519/tech-school/db/sqlc"
	"github.com/j23063519/tech-school/util"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load config:", err)
	}

	dbSource := "postgresql://" + config.POSTGRESUSER + ":" + config.POSTGRESPASSWORD + "@localhost:" + config.POSTGRESPORT + "/" + config.POSTGRESDB + "?sslmode=disable"

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can't connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start("0.0.0.0:" + config.APPPORT)
	if err != nil {
		log.Fatal("connot start versern", err)
	}
}
