package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project-2/database"
	"project-2/model"
	"strings"

	"gorm.io/gorm"
)

var db *gorm.DB

var dataBukuArr = []model.Book{}

func getPostBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		books, err := database.GetAllBooks(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		dataBuku, err := json.Marshal(books)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(dataBuku)
		return
	}
	if r.Method == "POST" {
		var buku model.Book
		var bukureq model.BookRequest
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&bukureq); err != nil {
				log.Fatal(err)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		buku.Author = bukureq.Author
		buku.Name_Book = bukureq.Name_Book
		if err := database.CreateBook(db, buku); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Created")
		return
	}

	http.Error(w, "NOT FOUND", http.StatusMethodNotAllowed)
	return
}

func ById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := strings.TrimPrefix(r.URL.Path, "/books/")

	if r.Method == "GET" {
		var book model.Book
		book, err := database.GetBooksByID(db, param)
		dataBuku, err := json.Marshal(book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if book.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Id book : " + param + " Not Found")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(dataBuku)
		return
	}
	if r.Method == "PUT" {
		var buku model.Book
		var bukureq model.BookRequest
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&bukureq); err != nil {
				log.Fatal(err)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		buku.Author = bukureq.Author
		buku.Name_Book = bukureq.Name_Book
		if err := database.UpdateBook(db, param, buku); err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Id book : " + param + " Not Found")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Updated")
		return
	}
	if r.Method == "DELETE" {
		if err := database.DeleteBook(db, param); err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Id book : " + param + " Not Found")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Deleted")
		return
	}

	http.Error(w, "NOT FOUND", http.StatusMethodNotAllowed)
}

func main() {
	var err error
	http.HandleFunc("/books", getPostBook)
	http.HandleFunc("/books/", ById)
	db, err = database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
