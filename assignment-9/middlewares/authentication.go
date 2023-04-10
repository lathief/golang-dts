package middlewares

import (
	"assignment-9/database"
	"assignment-9/helpers"
	"assignment-9/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helpers.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			return
		}

		c.Set("userData", claims)
		c.Next()
	}
}

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productID, err := strconv.Atoi(c.Param("productID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Unauthorized",
				"error":   "Invalid product ID data type",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		product := models.Product{}

		err = db.Select("user_id").First(&product, uint(productID)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unauthorized",
				"error":   "Failed to find product",
			})
			return
		}

		if product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
				"error":   "You are not allowed to access this product",
			})
			return
		}

		c.Next()
	}
}

func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userData").(jwt.MapClaims)
		userAdmin := userData["email"].(string)
		if userAdmin != "admin@admin.com" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
				"error":   "You are not allowed to access this endpoint",
			})
			return
		}
		c.Next()
	}
}
