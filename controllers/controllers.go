package controllers

import "github.com/gin-gonic/gin"

func HashPassword(password string) string {
	return ""
}

func VerifyPassword(userPassword, givenPassword string) (bool, string) {
	return false, ""
}

func SignUp() gin.HandlerFunc {
	return nil
}

func Login() gin.HandlerFunc {
	return nil
}

func ProductViewerAdmin() gin.HandlerFunc {
	return nil
}

func SearchProduct() gin.HandlerFunc {
	return nil
}

func SearchProductByQuery() gin.HandlerFunc {
	return nil
}
