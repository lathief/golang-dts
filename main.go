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
		var bookResponse []model.BookResponse
		books, err := database.GetAllBooks(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		for _, item := range books {
			bookResponse = append(bookResponse, model.BookResponse{
				ID:        item.ID,
				Name_Book: item.Name_Book,
				Author:    item.Author,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
			})
		}
		dataBuku, err := json.Marshal(bookResponse)
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
		var bukures model.BookResponse
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
		if err := database.CreateBook(db, &buku); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusCreated)
		bukures.Author = buku.Author
		bukures.Name_Book = buku.Name_Book
		bukures.ID = buku.ID
		bukures.CreatedAt = buku.CreatedAt
		bukures.UpdatedAt = buku.UpdatedAt
		json.NewEncoder(w).Encode(bukures)
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
		var bookResponse model.BookResponse
		book, err := database.GetBooksByID(db, param)
		bookResponse.Author = book.Author
		bookResponse.Name_Book = book.Name_Book
		bookResponse.ID = book.ID
		bookResponse.CreatedAt = book.CreatedAt
		bookResponse.UpdatedAt = book.UpdatedAt
		dataBuku, err := json.Marshal(bookResponse)
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
		var bukures model.BookResponse
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
		buku, _ = database.GetBooksByID(db, param)
		bukures.Author = buku.Author
		bukures.Name_Book = buku.Name_Book
		bukures.ID = buku.ID
		bukures.CreatedAt = buku.CreatedAt
		bukures.UpdatedAt = buku.UpdatedAt
		json.NewEncoder(w).Encode(bukures)
		return
	}
	if r.Method == "DELETE" {
		if err := database.DeleteBook(db, param); err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Id book : " + param + " Not Found")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Book deleted successfully",
		})
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
