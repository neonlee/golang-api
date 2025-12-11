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
	agendamentoRepo := repositories.NewAgendamentoRepository(dbConnection)
	compraRepo := repositories.NewComprasRepository(dbConnection)
	estoqueRepo := repositories.NewEstoqueRepository(dbConnection)
	financeiroRepo := repositories.NewFinanceiroRepository(dbConnection)
	planosRepo := repositories.NewPlanosRepository(dbConnection)
	logRepo := repositories.NewLogRepository(dbConnection)

	logController := Controllers.NewLogsController(logRepo)
	planoController := Controllers.NewPlanosController(planosRepo)
	financeiroController := Controllers.NewFinanceiroController(financeiroRepo)
	estoqueController := Controllers.NewEstoqueController(estoqueRepo)
	agendamentoController := Controllers.NewAgendamentoController(agendamentoRepo)
	comprasController := Controllers.NewCompraController(compraRepo)
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
	dashboardRepository := repositories.NewDashboardRepository(dbConnection)
	dashboardController := Controllers.NewDashboardController(dashboardRepository)

	logs := server.Group("/logs")
	{
		logs.POST("/", logController.CreateLog)
		logs.GET("/:id", logController.GetLogByID)
		logs.GET("/usuario/:usuario_id", logController.GetLogsByUsuario)
		logs.GET("/modulo", logController.GetLogsByModulo)
		logs.GET("/erros/:empresa_id", logController.GetLogsErro)
	}

	planos := server.Group("/planos")
	{
		planos.POST("/", planoController.CreatePlano)
		planos.GET("/:id", planoController.GetPlanoByID)
		planos.PUT("/:id", planoController.UpdatePlano)
		planos.GET("/", planoController.ListAllPlanos)
	}

	financeiro := server.Group("/financeiro")
	{
		financeiro.POST("/contas-receber", financeiroController.CreateContaReceber)
		financeiro.GET("/contas-receber/:id", financeiroController.GetContaReceberByID)
		financeiro.PUT("/contas-receber/baixar/:id", financeiroController.BaixarContaReceber)
		// Additional financeiro routes can be added here
	}

	dashboard := server.Group("/dashboard")
	{
		dashboard.GET("/resumo-vendas/:empresa_id", dashboardController.GetResumoVendas)
		dashboard.GET("/resumo-financeiro/:empresa_id", dashboardController.GetResumoFinanceiro)
		dashboard.GET("/proximos-agendamentos/:empresa_id", dashboardController.GetProximosAgendamentos)
		dashboard.GET("/alerta-estoque/:empresa_id", dashboardController.GetAlertasEstoque)
	}

	estoque := server.Group("/estoque")
	{
		estoque.POST("/movimentar", estoqueController.MovimentarEstoque)
		estoque.GET("/historico/:produto_id", estoqueController.GetHistoricoEstoque)
		estoque.POST("/ajustar", estoqueController.AjustarEstoque)
		estoque.GET("/saldo/:produto_id", estoqueController.GetSaldoAtual)
		estoque.POST("/transferir", estoqueController.TransferirEstoque)
		estoque.GET("/relatorio/:empresa_id", estoqueController.GetRelatorioEstoque)
		estoque.GET("/movimentacoes/periodo/:empresa_id", estoqueController.GetMovimentacoesPorPeriodo)
	}

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
		category.GET("/get-with-products/:id", categoryController.GetWithProdutos)
		category.GET("/:id", categoryController.GetByID)
		category.POST("/", categoryController.Create)
		category.PUT("/:id", categoryController.Update)
		category.DELETE("/:id", categoryController.Delete)
		category.GET("/empresa/:empresa_id", categoryController.ListByEmpresa)
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
		medico.POST("/especialidade", medicoController.AddEspecialidade)
		medico.DELETE("/disponibilidade/:id", medicoController.DeleteDisponibilidade)
		medico.PUT("/especialidade/:id", medicoController.UpdateDisponibilidade)
		medico.DELETE("/especialidade/:id", medicoController.DeleteEspecialidade)
		medico.GET("/list", medicoController.ListarMedicosComEspecialidadesEDisponibilidades)

	}

	agendamento := server.Group("/agendamentos")
	{
		agendamento.POST("/", agendamentoController.Create)
		agendamento.GET("/:id", agendamentoController.GetByID)
		agendamento.PUT("/", agendamentoController.Update)
		agendamento.PUT("/cancelar/:id", agendamentoController.Cancelar)
		agendamento.GET("/list-by-data", agendamentoController.ListByData)
		agendamento.GET("/list-by-periodo", agendamentoController.ListByPeriodo)
		agendamento.GET("/list-by-pet", agendamentoController.ListByPet)
		agendamento.GET("/verificar-disponibilidade", agendamentoController.VerificarDisponibilidade)
		agendamento.GET("/horarios-disponiveis", agendamentoController.GetHorariosDisponiveis)
	}

	compra := server.Group("/compras")
	{
		compra.POST("/", comprasController.CreateCompra)
		compra.GET("/:id", comprasController.GetCompraByID)
		compra.PUT("/", comprasController.UpdateCompra)
		compra.PUT("/cancelar/:id", comprasController.CancelarCompra)
		compra.GET("/empresa/:empresa_id", comprasController.ListByEmpresa)
		compra.GET("/itens/:compra_id", comprasController.GetItensByCompraID)
		compra.GET("/fornecedor/:fornecedor_id", comprasController.ListByFornecedor)
		compra.GET("/relatorio-periodo", comprasController.GetComprassPorPeriodo)
		compra.GET("/resumo/:empresa_id", comprasController.GetResumoComprass)
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8000")

}
