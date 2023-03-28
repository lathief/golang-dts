package main

import (
	"database/sql"
	"log"
	"sesi-7-gin/database"
	"sesi-7-gin/routers"
)

var DB *sql.DB

func main() {
	var err error
	var PORT = ":8080"
	DB, err = database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	routers.StartServer().Run(PORT)
}
