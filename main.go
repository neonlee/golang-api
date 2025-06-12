package main

import (
	Controllers "petApi/internal/controllers"
	"petApi/internal/repositories"
	"petApi/pkg/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	server := gin.Default()

	dbConnection := database.ConnectDB()

	//Camada de repository
	userRepo := repositories.NewUserRepository(dbConnection)
	//Camada de controllers
	ProductController := Controllers.NewUserHandler(&userRepo)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.POST("/product", ProductController.CreateUser)
	server.GET("/product/:productId", ProductController.GetUser)

	server.Run(":8000")

}
