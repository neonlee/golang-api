package main

import (
	Controllers "petApi/internal/controllers"
	"petApi/internal/repositories"
	"petApi/migrations"
	"petApi/pkg/database"
	"time"

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
	authRepo := repositories.NewAuthRepository(dbConnection, "token", time.Hour*24)

	clientsController := Controllers.NewClientsController(ClientsRepo)
	petsController := Controllers.NewPetsController(petsRepo)
	servicesController := Controllers.NewServicesController(servicesRepo)
	categoryController := Controllers.NewCategoriaController(categoryRepo)
	employeeController := Controllers.NewEmployeeController(employeeRepo)
	supplierController := Controllers.NewSuppliersController(supplierRepo)
	productController := Controllers.NewProductController(productRepo)
	authController := Controllers.NewAuthController(authRepo)
	server.POST("/auth/login", authController.Login)

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
		clients.GET("/search", clientsController.SearchClients)
		clients.GET("/new-clients", clientsController.GetNewClients)
		clients.GET("/list", clientsController.ListByEmpresa)
		clients.GET("/count", clientsController.GetTotalClients)

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
		product.POST("/", productController.CreateProduct)
		product.PUT("/:id", productController.UpdateProduct)
		product.DELETE("/:id", productController.DeleteProduct)
		product.GET("/:id", productController.GetProductByID)
		product.GET("/", productController.ListByEmpresa)
		product.GET("/search", productController.SearchProducts)
		product.GET("/low-stock", productController.GetProdutosBaixoEstoque)
		product.PATCH("/update-stock/:id", productController.UpdateEstoque)
		product.GET("/with-stock/:id", productController.GetProdutoComEstoque)
		product.GET("/expiring-soon/:id", productController.GetProdutosProximosVencimento)
		product.GET("/expiring-today/:id", productController.GetProdutosVencimentoHoje)
		product.GET("/out-of-stock/:id", productController.GetProdutosSemEstoque)
		product.GET("/expired/:id", productController.GetProdutosVencidos)
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8000")

}
