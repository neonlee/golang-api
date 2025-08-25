package Controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllersCategory struct {
	repository *repositories.CategoryRepository
}

func NewCategoryController(connection *repositories.CategoryRepository) *ControllersCategory {
	return &ControllersCategory{repository: connection}
}

// UpdateService godoc
//
//	@Summary		atualiza o serviço
//	@Description	atualiza o serviço
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Services
//	@Failure		500	{object}	map[string]string
//	@Router			/clients [get]
func (p *ControllersCategory) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	var cliente models.Category
	if err := ctx.BindJSON(&cliente); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	client, err := p.repository.UpdateCategorys(user, cliente)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, client)
}

// Getclient godoc
//
//	@Summary		Lista um cliente
//	@Description	Retorna um cliente
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Services
//	@Failure		500	{object}	map[string]string
//	@Router			/client/:id [get]
func (p *ControllersCategory) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	client, err := p.repository.GetCategory(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, client)
}

// GetServices godoc
//
//	@Summary		Lista todos os clientes
//	@Description	Retorna todos os clientes cadastrados
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Services
//	@Failure		500	{object}	map[string]string
//	@Router			/clients [get]
func (p *ControllersCategory) GetCategorys(ctx *gin.Context) {
	result, err := p.repository.GetCategorys()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, result)
}

// PostPet godoc
//
//	@Summary		Cria um cliente
//	@Description	Cria um cliente
//	@Tags			clients
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Services
//	@Failure		500	{object}	map[string]string
//	@Router			/client [get]
func (p *ControllersCategory) Create(ctx *gin.Context) {
	var client models.Category
	err := ctx.BindJSON(&client)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := p.repository.Create(client)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, result)
}

// DeleteService godoc
//
//	@Summary		Deleta um cliente
//	@Description	deleta um cliente
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		bool
//	@Failure		500	{object}	map[string]string
//	@Router			/client [get]
func (p *ControllersCategory) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	client, err := p.repository.DeleteCategorys(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, client)
}
