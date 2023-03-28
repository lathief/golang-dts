package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "book-go"
)

var (
	Db  *sql.DB
	err error
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = Db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to database")
	return Db, nil
}
