package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"petApi/internal/requests"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllersSuppliers struct {
	repository repositories.FornecedoresRepository
}

func NewSuppliersController(connection repositories.FornecedoresRepository) *ControllersSuppliers {
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

	var Fornecedores models.Fornecedores
	if err := ctx.BindJSON(&Fornecedores); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	err := p.repository.Update(&Fornecedores)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, &Fornecedores)
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
	client, err := p.repository.GetByID(uint(user))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, client)
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
	var client models.Fornecedores
	err := ctx.BindJSON(&client)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = p.repository.Create(&client)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, &client)
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

	fornecedorId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	err = p.repository.Delete(uint(fornecedorId))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"deleted": true})
}

func (p *ControllersSuppliers) Search(ctx *gin.Context) {
	empresaIDParam := ctx.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID da empresa inválido"})
		return
	}
	termo := ctx.Query("termo")
	fornecedores, err := p.repository.Search(uint(empresaID), termo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, fornecedores)
}

func (p *ControllersSuppliers) GetTotalFornecedores(ctx *gin.Context) {
	empresaIDParam := ctx.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID da empresa inválido"})
		return
	}
	total, err := p.repository.GetTotalFornecedores(uint(empresaID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"total": total})
}

func (p *ControllersSuppliers) ListByEmpresa(ctx *gin.Context) {
	empresaIDParam := ctx.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID da empresa inválido"})
		return
	}
	var filters requests.FornecedorFilter
	if err := ctx.BindQuery(&filters); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Filtros inválidos"})
		return
	}

	fornecedores, err := p.repository.ListByEmpresa(uint(empresaID), filters)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, fornecedores)
}
