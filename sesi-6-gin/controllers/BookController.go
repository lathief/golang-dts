package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var Books = []Book{
	{1, "Matematika", "Upin", "Buku matematika"},
	{2, "Bahasa Indonesia", "Susanti", "Buku Bhs.Indonesia"},
}

func GetAllBook(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Books": Books,
	})
}

func GetBookByID(ctx *gin.Context) {
	condition := false
	var buku Book
	BookID, err := strconv.Atoi(ctx.Param("bookID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": "Enter with correct number id",
		})
		return
	}
	for _, Book := range Books {
		if Book.BookID == BookID {
			buku = Book
			condition = true
		}
	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %v not found", BookID),
		})
		return
	}
	ctx.JSON(http.StatusOK, buku)
}

func CreateBook(ctx *gin.Context) {
	var newBook Book
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newBook.BookID = len(Books) + 1
	Books = append(Books, newBook)

	ctx.JSON(http.StatusCreated, "Created")
}

func UpdateBook(ctx *gin.Context) {
	BookID, err := strconv.Atoi(ctx.Param("bookID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": "Enter with correct number id",
		})
		return
	}
	condition := false
	var updatedBook Book
	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	for i, Book := range Books {
		if BookID == Book.BookID {
			condition = true
			Books[i] = updatedBook
			Books[i].BookID = BookID
			break
		}
	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %v not found", BookID),
		})
		return
	}
	ctx.JSON(http.StatusOK, "Updated")
}

func DeleteBook(ctx *gin.Context) {
	BookID, err := strconv.Atoi(ctx.Param("bookID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": "Enter with correct number id",
		})
		return
	}
	condition := false
	var BookIndex int
	for i, Book := range Books {
		if BookID == Book.BookID {
			condition = true
			BookIndex = i
			break
		}
	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found", "error_message": fmt.Sprintf("Book with id %v not found", BookID),
		})
		return
	}
	copy(Books[BookIndex:], Books[BookIndex+1:])
	Books[len(Books)-1] = Book{}
	Books = Books[:len(Books)-1]
	ctx.JSON(http.StatusOK, "Deleted")
}
