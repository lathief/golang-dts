package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sesi-7/database"
	"sesi-7/model"
	"strconv"
	"strings"
)

var DB *sql.DB

var dataBukuArr = []model.Book{}

func getPostBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		books := database.GetAllBooks(DB)
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
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&buku); err != nil {
				log.Fatal(err)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := database.CreateBook(DB, buku); err != nil {
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
	id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if r.Method == "GET" {
		var book model.Book = database.GetBooksByID(DB, param)
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
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&buku); err != nil {
				log.Fatal(err)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := database.UpdateBook(DB, id, buku); err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Id book : " + param + " Not Found")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Updated")
		return
	}
	if r.Method == "DELETE" {
		if err := database.DeleteBook(DB, id); err != nil {
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
	DB, err = database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
