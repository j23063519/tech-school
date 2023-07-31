package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/j23063519/tech-school/util"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("can't load config:", err)
	}

	dbSource := "postgresql://" + config.POSTGRESUSER + ":" + config.POSTGRESPASSWORD + "@localhost:" + config.POSTGRESPORT + "/" + config.POSTGRESDB + "?sslmode=disable"

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can't connect to database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
