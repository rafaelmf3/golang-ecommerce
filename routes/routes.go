package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelmf3/golang-ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/product", controllers.AddProduct())
	incomingRoutes.GET("/users/product", controllers.GetProduct())
	incomingRoutes.GET("/users/search", controllers.Search())
}
