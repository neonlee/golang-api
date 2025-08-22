package main

import (
	Controllers "petApi/internal/controllers"
	"petApi/internal/repositories"
	"petApi/migrations"
	"petApi/pkg/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	server := gin.Default()

	dbConnection := database.ConnectDB()

	m := migrations.NewMigrationsDB(dbConnection)
	m.RunMigrations()

	petsRepo := repositories.NewPetsRepository(dbConnection)
	userRepo := repositories.NewUserRepository(dbConnection)
	ClientsRepo := repositories.NewClientsRepository(dbConnection)
	servicesRepo := repositories.NewServicesRepository(dbConnection)

	userController := Controllers.NewUserHandler(&userRepo)
	clientsController := Controllers.NewClientsController(&ClientsRepo)
	petsController := Controllers.NewPetsController(&petsRepo)
	servicesController := Controllers.NewServicesController(&servicesRepo)

	clients := server.Group("/clients")
	{
		clients.GET("/", clientsController.GetClients)
		clients.GET("/:id", clientsController.GetClient)
		clients.POST("/", clientsController.CreateClients)
		clients.PUT("/:id", clientsController.UpdateClient)
		clients.DELETE("/:id", clientsController.DeleteClient)
	}

	services := server.Group("/services")
	{
		services.GET("/", servicesController.GetServices)
		services.GET("/:id", servicesController.GetService)
		services.POST("/", servicesController.CreateServices)
		services.PUT("/:id", servicesController.UpdateService)
		services.DELETE("/:id", servicesController.DeleteService)
	}

	pets := server.Group("/pets")
	{
		pets.GET("/", petsController.GetPets)
		pets.POST("/", petsController.CreatePets)
		pets.PUT("/:id", petsController.UpdatePet)
		pets.DELETE("/:id", petsController.DeletePet)
		pets.GET("/:id", petsController.GetPet)
	}

	user := server.Group("/users")
	{
		user.POST("/", userController.CreateUser)
		user.GET("/:userId", userController.GetUser)
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8000")

}
