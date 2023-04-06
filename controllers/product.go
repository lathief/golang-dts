package controllers

import (
	"assinment-8/database"
	"assinment-8/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(product)

	product.UserID = uint(userData["id"].(float64))

	err = db.Create(&product).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	product.UserID = uint(userData["id"].(float64))

	err = db.Model(&product).Where("id=?", productID).Updates(models.Product{Title: product.Title, Description: product.Description}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func ReadProductById(ctx *gin.Context) {
	db := database.GetDB()
	product := models.Product{}
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	if err := db.Where("id = ?", productID).First(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, product)
	}
}

func ReadAllProduct(ctx *gin.Context) {
	db := database.GetDB()
	product := []models.Product{}
	if err := db.Find(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, product)
	}
}

func DeleteProduct(ctx *gin.Context) {
	db := database.GetDB()
	product := models.Product{}
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	if err := db.Where("id =?", productID).Delete(&product).RowsAffected; err == 0 {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, "product deleted")
}
