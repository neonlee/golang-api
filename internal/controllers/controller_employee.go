package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	Repo repositories.FuncionariosRepository
}

func NewEmployeeController(repository repositories.FuncionariosRepository) *EmployeeController {
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

	var employee models.Funcionarios
	if err := ctx.BindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	err = p.Repo.UpdateFuncionarios(employeeID, &employee)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &employee)
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
	id := ctx.Param("empresaID")

	empresaID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	idFuncionario := ctx.Param("funcionarioID")

	funcionarioId, err := strconv.Atoi(idFuncionario)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	employee, err := p.Repo.GetFuncionarios(funcionarioId, uint(empresaID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

// GetFuncionarios godoc
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
	id := ctx.Query("empresa_id")

	employeeID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	employee, err := p.Repo.GetAllFuncionarios(uint(employeeID))
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
	var employee models.Funcionarios
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

	success, err := p.Repo.DeleteFuncionarios(employeeID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"deleted": success})
}
