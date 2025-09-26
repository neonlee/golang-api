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

	petsRepo := repositories.NewPetRepository(dbConnection)
	ClientsRepo := repositories.NewClientesRepository(dbConnection)
	servicesRepo := repositories.NewServicoRepository(dbConnection)
	categoryRepo := repositories.NewCategoryRepository(dbConnection)
	employeeRepo := repositories.NewEmployeesRepository(dbConnection)
	supplierRepo := repositories.NewSupplierRepository(dbConnection)
	productRepo := repositories.NewProdutoRepository(dbConnection)

	clientsController := Controllers.NewClientsController(ClientsRepo)
	petsController := Controllers.NewPetsController(petsRepo)
	servicesController := Controllers.NewServicesController(servicesRepo)
	categoryController := Controllers.NewCategoriaController(categoryRepo)
	employeeController := Controllers.NewEmployeeController(employeeRepo)
	supplierController := Controllers.NewSuppliersController(supplierRepo)
	productController := Controllers.NewProductController(productRepo)

	supplier := server.Group("/suppliers")
	{
		supplier.GET("/", supplierController.GetSuppliers)
		supplier.GET("/:id", supplierController.Get)
		supplier.POST("/", supplierController.Create)
		supplier.PUT("/:id", supplierController.Update)
		supplier.DELETE("/:id", supplierController.Delete)
	}

	clients := server.Group("/clients")
	{
		clients.GET("/", clientsController.GetClients)
		clients.GET("/:id", clientsController.GetClient)
		clients.POST("/", clientsController.CreateClients)
		clients.PUT("/:id", clientsController.UpdateClient)
		clients.DELETE("/:id", clientsController.DeleteClient)
	}
	employee := server.Group("/employee")
	{
		employee.GET("/", employeeController.Get)
		employee.GET("/:id", employeeController.Get)
		employee.POST("/", employeeController.Create)
		employee.PUT("/:id", employeeController.Update)
		employee.DELETE("/:id", employeeController.Delete)
	}

	category := server.Group("/category")
	{
		// category.GET("/", categoryController.GetCategorys)
		// category.GET("/:id", categoryController.Get)
		category.POST("/", categoryController.Create)
		category.PUT("/:id", categoryController.Update)
		category.DELETE("/:id", categoryController.Delete)
	}

	services := server.Group("/services")
	{
		// services.GET("/", servicesController.GetServices)
		services.GET("/:id", servicesController.GetService)
		services.POST("/", servicesController.CreateServices)
		services.PUT("/:id", servicesController.UpdateService)
		services.DELETE("/:id", servicesController.DeleteService)
	}

	pets := server.Group("/pets")
	{
		// pets.GET("/", petsController.GetPets)
		pets.POST("/", petsController.CreatePets)
		pets.PUT("/:id", petsController.UpdatePet)
		pets.DELETE("/:id", petsController.DeletePet)
		pets.GET("/:id", petsController.GetPet)
	}
	product := server.Group("/products")
	{
		// product.GET("/", productController.GetAllProducts)
		product.POST("/", productController.CreateProduct)
		product.PUT("/:id", productController.UpdateProduct)
		product.DELETE("/:id", productController.DeleteProduct)
		product.GET("/:id", productController.GetProductByID)
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8000")

}
