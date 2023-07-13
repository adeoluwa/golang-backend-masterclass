package backend_masterclass

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_"github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/postgres12?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB


// The TestMain function connects to a database, initializes test queries, and runs the tests.
func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource )
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
