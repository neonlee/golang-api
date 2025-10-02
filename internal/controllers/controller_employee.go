package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	Repo repositories.EmployeesRepository
}

func NewEmployeeController(repository repositories.EmployeesRepository) *EmployeeController {
	return &EmployeeController{Repo: repository}
}

// UpdateEmployee godoc
//
//	@Summary		Atualiza um funcionário
//	@Description	Atualiza os dados de um funcionário
//	@Tags			employees
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"ID do Funcionário"
//	@Param			employee	body		models.Employee	true	"Dados do Funcionário"
//	@Success		200		{object}	models.Employee
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/employees/{id} [put]
func (p *EmployeeController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	employeeID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	var employee models.Employees
	if err := ctx.BindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	updatedEmployee, err := p.Repo.UpdateEmployees(employeeID, employee)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedEmployee)
}

// GetEmployee godoc
//
//	@Summary		Obtém um funcionário
//	@Description	Retorna um funcionário específico pelo ID
//	@Tags			employees
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ID do Funcionário"
//	@Success		200	{object}	models.Employee
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Router			/employees/{id} [get]
func (p *EmployeeController) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	employeeID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	employee, err := p.Repo.GetEmployees(employeeID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

// GetEmployees godoc
//
//	@Summary		Lista todos os funcionários
//	@Description	Retorna todos os funcionários cadastrados
//	@Tags			employees
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Employee
//	@Failure		500	{object}	map[string]string
//	@Router			/employees [get]
func (p *EmployeeController) GetAll(ctx *gin.Context) {
	employee, err := p.Repo.GetAllEmployees()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

// CreateEmployee godoc
//
//	@Summary		Cria um funcionário
//	@Description	Cria um novo funcionário
//	@Tags			employees
//	@Accept			json
//	@Produce		json
//	@Param			employee	body		models.Employee	true	"Dados do Funcionário"
//	@Success		201		{object}	models.Employee
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/employees [post]
func (p *EmployeeController) Create(ctx *gin.Context) {
	var employee models.Employees
	err := ctx.BindJSON(&employee)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := p.Repo.Create(employee)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

// DeleteEmployee godoc
//
//	@Summary		Deleta um funcionário
//	@Description	Remove um funcionário do sistema
//	@Tags			employees
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ID do Funcionário"
//	@Success		200	{object}	map[string]bool
//	@Failure		400	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/employees/{id} [delete]
func (p *EmployeeController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	employeeID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	success, err := p.Repo.DeleteEmployees(employeeID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"deleted": success})
}
