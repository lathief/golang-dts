package router

import (
	"assignment-9/controllers"
	"assignment-9/middlewares"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
	}

	productRouter := r.Group("products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.GET("/", middlewares.AdminAuthorization(), controllers.ReadAllProduct)
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/:productID", middlewares.ProductAuthorization(), controllers.ReadProductById)
		productRouter.PUT("/:productID", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productID", middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	return r
}
