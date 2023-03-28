package database

import (
	"database/sql"
	"errors"
	"fmt"

	"sesi-7/model"

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
	db  *sql.DB
	err error
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to database")
	return db, nil
}
func GetAllBooks(db *sql.DB) []model.Book {
	var results = []model.Book{}

	sqlStatement := `SELECT * from books`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var book = model.Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			panic(err)
		}
		results = append(results, book)
	}
	fmt.Println("Books datas:", results)
	return results
}
func GetBooksByID(db *sql.DB, index string) model.Book {
	var result = model.Book{}
	sqlStatement := `SELECT * from books where id = $1`
	rows, err := db.Query(sqlStatement, index)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Title, &result.Author, &result.Desc)
		if err != nil {
			panic(err)
		}
	}
	return result
}

func CreateBook(db *sql.DB, insertBook model.Book) error {
	sqlStatement := `
	INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	`
	errs := db.QueryRow(sqlStatement, insertBook.Title, insertBook.Author, insertBook.Desc)
	return errs.Err()
}

func UpdateBook(db *sql.DB, id int, updateBook model.Book) error {
	sqlStatement := `
	UPDATE books
	SET title = $2, author = $3, description = $4
	WHERE id = $1;
	`
	res, err := db.Exec(sqlStatement, id, updateBook.Title, updateBook.Author, updateBook.Desc)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil || count == 0 {
		return errors.New("No Rows Affected")
	}
	return err
}

func DeleteBook(db *sql.DB, id int) error {
	sqlStatement := `
	DELETE from books
	WHERE id = $1;
	`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil || count == 0 {
		return errors.New("No Rows Affected")
	}
	return err
}
