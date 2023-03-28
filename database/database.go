package database

import (
	"errors"
	"fmt"
	"project-2/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "book-gorm"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.Debug().AutoMigrate(model.Book{})

	fmt.Println("Successfully connected to database")

	return db, nil
}
func GetAllBooks(db *gorm.DB) ([]model.Book, error) {
	var books []model.Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	} else {
		return books, nil
	}
}
func GetBooksByID(db *gorm.DB, index string) (model.Book, error) {
	var book model.Book
	if err := db.Where("id = ?", index).First(&book).Error; err != nil {
		return book, err
	} else {
		return book, nil
	}
}
func CreateBook(db *gorm.DB, insertBook model.Book) error {
	if err := db.Create(&insertBook).Error; err != nil {
		return err
	} else {
		return nil
	}
}
func UpdateBook(db *gorm.DB, id string, updateBook model.Book) error {
	if err := db.Model(model.Book{}).Where("id =?", id).Updates(map[string]interface{}{"author": updateBook.Author, "name_book": updateBook.Name_Book}).RowsAffected; err == 0 {
		return errors.New("Book not Found")
	} else {
		return nil
	}
}
func DeleteBook(db *gorm.DB, id string) error {
	if err := db.Where("id =?", id).Delete(&model.Book{}).RowsAffected; err == 0 {
		return errors.New("Book not Found")
	} else {
		return nil
	}
}
