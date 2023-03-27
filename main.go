package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Desc   string
}

var dataBukuArr = []Book{
	{1, "Matematika", "Upin", "Buku matematika"},
	{2, "Bahasa Indonesia", "Susanti", "Buku Bhs.Indonesia"},
}

func incrementID(ID *int) {
	*ID = dataBukuArr[len(dataBukuArr)-1].ID + 1
}
func RemoveData(s []Book, index int) []Book {
	return append(s[:index], s[index+1:]...)
}

func getPostBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		books := dataBukuArr
		dataBuku, err := json.Marshal(books)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(dataBuku)
		return
	}
	if r.Method == "POST" {
		var buku Book
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&buku); err != nil {
				log.Fatal(err)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		incrementID(&buku.ID)
		dataBukuArr = append(dataBukuArr, buku)
		json.NewEncoder(w).Encode("Created")
		return
	}

	http.Error(w, "NOT FOUND", http.StatusMethodNotAllowed)
	return
}

func ById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var index int
	param := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if r.Method == "GET" {
		var book Book
		for _, buku := range dataBukuArr {
			if buku.ID == id {
				book = buku
			}
		}
		dataBuku, err := json.Marshal(book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(dataBuku)
		return
	}
	if r.Method == "PUT" {
		var buku Book
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&buku); err != nil {
				log.Fatal(err)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for ind, item := range dataBukuArr {
			if item.ID == id {
				index = ind
			}
		}
		if index == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Id book : " + param + " Not Found")
			return
		}
		w.WriteHeader(http.StatusOK)
		dataBukuArr[index].Author = buku.Author
		dataBukuArr[index].Title = buku.Title
		dataBukuArr[index].Desc = buku.Desc
		json.NewEncoder(w).Encode("Updated")
		return
	}
	if r.Method == "DELETE" {
		for ind, item := range dataBukuArr {
			if item.ID == id {
				index = ind
			}
		}
		if index == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Id book : " + param + " Not Found")
			return
		}
		dataBukuArr = RemoveData(dataBukuArr, index)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Deleted")
		return
	}

	http.Error(w, "NOT FOUND", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/books", getPostBook)
	http.HandleFunc("/books/", ById)

	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
