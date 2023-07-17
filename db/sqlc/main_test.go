package backend_masterclass

import (
	"backend_masterclass/util"

	"database/sql"

	"log"

	"os"

	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB


// The TestMain function connects to a database, initializes test queries, and runs the tests.
func TestMain(m *testing.M) { 
	config, err := util.LoadConfig("../..")
	if err != nil{
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource )
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
