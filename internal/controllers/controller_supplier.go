package Controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllersSuppliers struct {
	repository *repositories.SupplierRepository
}

func NewSuppliersController(connection *repositories.SupplierRepository) *ControllersSuppliers {
	return &ControllersSuppliers{repository: connection}
}

// UpdateClient godoc
//
//	@Summary		atualiza o cliente
//	@Description	atualiza o cliente
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Client
//	@Failure		500	{object}	map[string]string
//	@Router			/clients [get]
func (p *ControllersSuppliers) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	var cliente models.Supplier
	if err := ctx.BindJSON(&cliente); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	client, err := p.repository.UpdateSupplier(user, cliente)

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
//	@Success		200	{array}		models.Client
//	@Failure		500	{object}	map[string]string
//	@Router			/client/:id [get]
func (p *ControllersSuppliers) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	client, err := p.repository.GetSupplier(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, client)
}

// GetSuppliers godoc
//
//	@Summary		Lista todos os clientes
//	@Description	Retorna todos os clientes cadastrados
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Client
//	@Failure		500	{object}	map[string]string
//	@Router			/clients [get]
func (p *ControllersSuppliers) GetSuppliers(ctx *gin.Context) {
	result, err := p.repository.GetSuppliers()

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
//	@Success		200	{array}		models.Client
//	@Failure		500	{object}	map[string]string
//	@Router			/client [get]
func (p *ControllersSuppliers) Create(ctx *gin.Context) {
	var client models.Supplier
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

// DeleteClient godoc
//
//	@Summary		Deleta um cliente
//	@Description	deleta um cliente
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		bool
//	@Failure		500	{object}	map[string]string
//	@Router			/client [get]
func (p *ControllersSuppliers) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	client, err := p.repository.DeleteSupplier(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, client)
}
