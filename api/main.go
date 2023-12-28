package main

import (
	"api/models"
	"api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	models.ConnectDatabase()

	router := r.Group("users")
	{
		router.GET("", services.FindAllUsers)
		router.GET(":id", services.FindOneUser)
		router.POST("", services.CreateUser)
		router.PATCH(":id", services.UpdateUser)
		router.PUT(":id/password", services.UpdatePasswordUser)
		router.DELETE(":id", services.RemoveUser)
	}

	r.Run(":8080")
}
