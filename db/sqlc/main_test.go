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


// The TestMain function connects to a database, initializes test queries, and runs the tests.
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource )
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
