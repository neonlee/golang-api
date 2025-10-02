package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"petApi/internal/requests"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllersClientes struct {
	repository repositories.ClientesRepository
}

func NewClientsController(connection repositories.ClientesRepository) *ControllersClientes {
	return &ControllersClientes{repository: connection}
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
func (p *ControllersClientes) UpdateClient(ctx *gin.Context) {

	var cliente models.Clientes
	if err := ctx.BindJSON(&cliente); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	err := p.repository.Update(&cliente)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, cliente)
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
func (p *ControllersClientes) GetClient(ctx *gin.Context) {
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

// GetClients godoc
//
//	@Summary		Lista todos os clientes
//	@Description	Retorna todos os clientes cadastrados
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Client
//	@Failure		500	{object}	map[string]string
//	@Router			/clients [get]
func (p *ControllersClientes) GetClients(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	var filter = requests.ClientesFilter{}
	if err := ctx.BindJSON(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	result, err := p.repository.ListByEmpresa(uint(user), filter)

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
func (p *ControllersClientes) CreateClients(ctx *gin.Context) {
	var client models.Clientes
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
func (p *ControllersClientes) DeleteClient(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	err = p.repository.Delete(uint(user))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"deleted": true})
}

func (p *ControllersClientes) SearchClients(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	termo := ctx.Query("termo")
	result, err := p.repository.Search(uint(user), termo)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, result)
}

func (p *ControllersClientes) GetTotalClients(ctx *gin.Context) {
	id := ctx.Param("EmpresaID")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	total, err := p.repository.GetTotalClientes(uint(user))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, gin.H{"total": total})
}

func (p *ControllersClientes) GetNewClients(ctx *gin.Context) {
	mesStr := ctx.Query("mes")
	anoStr := ctx.Query("ano")
	mes, err := strconv.Atoi(mesStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Mês inválido"})
		return
	}
	ano, err := strconv.Atoi(anoStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Ano inválido"})
		return
	}
	id := ctx.Param("EmpresaID")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	result, err := p.repository.GetClientesNovos(uint(user), mes, ano)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, result)
}
