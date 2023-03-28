package controllers

import (
	"fmt"
	"net/http"
	"sesi-7-gin/database"
	"sesi-7-gin/model"
	"sesi-7-gin/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Books = []model.Book{}

func GetAllBook(ctx *gin.Context) {
	books := repository.GetAllBooks(database.Db)
	ctx.JSON(http.StatusOK, books)
}

func GetBookByID(ctx *gin.Context) {
	var buku model.Book
	BookID, err := strconv.Atoi(ctx.Param("bookID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": "Enter with correct number id",
		})
		return
	}
	buku = repository.GetBooksByID(database.Db, ctx.Param("bookID"))
	if buku.BookID == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %v not found", BookID),
		})
		return
	}
	ctx.JSON(http.StatusOK, buku)
}

func CreateBook(ctx *gin.Context) {
	var newBook model.Book
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := repository.CreateBook(database.Db, newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Server Error",
			"error_message": "Internal Server Error",
		})
		return
	}
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
	var updatedBook model.Book
	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = repository.UpdateBook(database.Db, BookID, updatedBook)
	if err != nil {
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
	err = repository.DeleteBook(database.Db, BookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found", "error_message": fmt.Sprintf("Book with id %v not found", BookID),
		})
		return
	}
	ctx.JSON(http.StatusOK, "Deleted")
}
