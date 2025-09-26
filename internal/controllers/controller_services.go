package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllersServices struct {
	repository repositories.ServicoRepository
}

func NewServicesController(connection repositories.ServicoRepository) *ControllersServices {
	return &ControllersServices{repository: connection}
}

// UpdateService godoc
//
//	@Summary		atualiza o serviço
//	@Description	atualiza o serviço
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Service
//	@Failure		500	{object}	map[string]string
//	@Router			/clients [get]
func (p *ControllersServices) UpdateService(ctx *gin.Context) {

	var servico models.TipoServico
	if err := ctx.BindJSON(&servico); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	err := p.repository.UpdateTipoServico(&servico)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, &servico)
}

// Getclient godoc
//
//	@Summary		Lista um cliente
//	@Description	Retorna um cliente
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Service
//	@Failure		500	{object}	map[string]string
//	@Router			/client/:id [get]
func (p *ControllersServices) GetService(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	client, err := p.repository.GetTipoServicoByID(uint(user))

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
//	@Success		200	{array}		models.Service
//	@Failure		500	{object}	map[string]string
//	@Router			/clients [get]
// func (p *ControllersServices) GetServices(ctx *gin.Context) {
// 	result, err := p.repository.GetServices()

// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, err)
// 	}

// 	ctx.JSON(http.StatusOK, result)
// }

// PostPet godoc
//
//	@Summary		Cria um cliente
//	@Description	Cria um cliente
//	@Tags			clients
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Service
//	@Failure		500	{object}	map[string]string
//	@Router			/client [get]
func (p *ControllersServices) CreateServices(ctx *gin.Context) {
	var client models.TipoServico
	err := ctx.BindJSON(&client)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = p.repository.CreateTipoServico(&client)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, &client)
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
func (p *ControllersServices) DeleteService(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	err = p.repository.DeleteTipoServico(uint(user))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"deleted": true})
}
