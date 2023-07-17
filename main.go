package main

import (
	"database/sql"

	"log"

	"backend_masterclass/api"

	db "backend_masterclass/db/sqlc"

	"backend_masterclass/util"

	_ "github.com/lib/pq"
)


func main() {
	config, err := util.LoadConfig(".")
	if err != nil{
		log.Fatal("cannot load config:", err)
	}


	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
} 