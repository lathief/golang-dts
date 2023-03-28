package routers

import (
	"sesi-7-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.GET("/books", controllers.GetAllBook)
	router.GET("/books/:bookID", controllers.GetBookByID)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:bookID", controllers.UpdateBook)
	router.DELETE("/books/:bookID", controllers.DeleteBook)

	return router
}
