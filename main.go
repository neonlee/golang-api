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
	employeeRepo := repositories.NewFuncionariosRepository(dbConnection)
	supplierRepo := repositories.NewSupplierRepository(dbConnection)
	productRepo := repositories.NewProdutoRepository(dbConnection)
	authRepo := repositories.NewAuthRepository(dbConnection, "token", time.Hour*24)
	userRepo := repositories.NewUsuarioRepository(dbConnection)
	medicoRepo := repositories.NewMedicoVeterinarioRepository(dbConnection)

	clientsController := Controllers.NewClientsController(ClientsRepo)
	petsController := Controllers.NewPetsController(petsRepo)
	servicesController := Controllers.NewServicesController(servicesRepo)
	categoryController := Controllers.NewCategoriaController(categoryRepo)
	employeeController := Controllers.NewEmployeeController(employeeRepo)
	supplierController := Controllers.NewSuppliersController(supplierRepo)
	productController := Controllers.NewProductController(productRepo)
	authController := Controllers.NewAuthController(authRepo)
	userController := Controllers.NewUsuarioController(userRepo)
	medicoController := Controllers.NewMedicoVeterinarioController(medicoRepo)

	server.POST("/auth/login", authController.Login)

	supplier := server.Group("/suppliers")
	{
		supplier.GET("/:id", supplierController.Get)
		supplier.POST("/", supplierController.Create)
		supplier.PUT("/:id", supplierController.Update)
		supplier.DELETE("/:id", supplierController.Delete)
		supplier.GET("/ListByEmpresa/:empresa_id", supplierController.ListByEmpresa)
		supplier.GET("/search/:empresa_id", supplierController.Search)
		supplier.GET("/count/:empresa_id", supplierController.GetTotalFornecedores)
	}

	clients := server.Group("/clients")
	{
		clients.POST("/ListByEmpresa/:id", clientsController.GetClients)
		clients.GET("/:id", clientsController.GetClient)
		clients.POST("/", clientsController.CreateClients)
		clients.PUT("/:id", clientsController.UpdateClient)
		clients.DELETE("/:id", clientsController.DeleteClient)
		clients.GET("/:id/search", clientsController.SearchClients)
		clients.GET("/novos_clientes/:empresa_id", clientsController.GetNewClients)
		clients.GET("/list", clientsController.ListByEmpresa)
		clients.GET("/count/:empresa_id", clientsController.GetTotalClients)

	}
	employee := server.Group("/employee")
	{
		employee.GET("/", employeeController.GetAll)
		employee.GET("/:empresaID/:funcionarioID", employeeController.Get)
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
		services.GET("/list", servicesController.GetServices)
		services.GET("/:id", servicesController.GetService)
		services.POST("/", servicesController.CreateServices)
		services.PUT("/:id", servicesController.UpdateService)
		services.DELETE("/:id", servicesController.DeleteService)
		services.GET("/most-used", servicesController.GetServicosMaisUtilizados)
	}

	pets := server.Group("/pets")
	{
		pets.GET("/GetByClientes/:cliente_id", petsController.GetByClientes)
		pets.POST("/", petsController.CreatePets)
		pets.PUT("/:id", petsController.UpdatePet)
		pets.DELETE("/:id", petsController.DeletePet)
		pets.GET("/GetPet/:id", petsController.GetPet)
		pets.GET("/list/:empresa_id", petsController.ListByEmpresa)
		pets.GET("/count/:empresa_id", petsController.GetTotalPets)
		pets.GET("/get-especie/:empresa_id", petsController.GetPetsPorEspecie)
	}
	product := server.Group("/products")
	{
		product.POST("/", productController.CreateProduct)
		product.PUT("/:id", productController.UpdateProduct)
		product.DELETE("/:id", productController.DeleteProduct)
		product.GET("/:id", productController.GetProductByID)
		product.GET("/", productController.ListByEmpresa)
		product.GET("/search", productController.SearchProducts)
		product.GET("/low-stock/:empresa_id", productController.GetProdutosBaixoEstoque)
		product.PATCH("/update-stock/:id", productController.UpdateEstoque)
		product.GET("/with-stock/:id", productController.GetProdutoComEstoque)
		product.GET("/expiring-soon/:id", productController.GetProdutosProximosVencimento)
		product.GET("/expiring-today/:id", productController.GetProdutosVencimentoHoje)
		product.GET("/out-of-stock/:id", productController.GetProdutosSemEstoque)
		product.GET("/expired/:id", productController.GetProdutosVencidos)
	}

	user := server.Group("/users")
	{
		user.GET("/:id", userController.GetByID)
		user.PUT("/", userController.Update)
		user.DELETE("/:id", userController.Delete)
		user.PUT("/update-senha", userController.UpdateSenha)
		user.GET("/empresa/:empresa_id", userController.ListByEmpresa)
		user.PUT("/assign-perfis", userController.AssignPerfis)
		user.GET("/has-permission/:id", userController.HasPermission)
		user.GET("/check-empresa-access/:usuario_id/:empresa_id", userController.CheckEmpresaAccess)
		user.GET("/permissoes/:id", userController.GetPermissoes)
		user.POST("/", userController.CreateUser)
		user.GET("/with-perfis", userController.GetWithPerfis)
		user.GET("/get-by-email", userController.GetByEmail)
	}

	medico := server.Group("/medico-veterinario")
	{
		medico.POST("/veterinario", medicoController.CreateVeterinario)
		medico.POST("/disponibilidade", medicoController.AddDisponibilidade)
		medico.GET("/especialidade", medicoController.AddEspecialidade)
		medico.DELETE("/disponibilidade/:id", medicoController.DeleteDisponibilidade)
		medico.PUT("/especialidade/:id", medicoController.UpdateDisponibilidade)
		medico.DELETE("/especialidade/:id", medicoController.DeleteEspecialidade)
		medico.GET("/list", medicoController.ListarMedicosComEspecialidadesEDisponibilidades)

	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8000")

}
