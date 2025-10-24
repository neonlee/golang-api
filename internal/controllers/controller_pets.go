package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"petApi/internal/requests"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllersPets struct {
	repository repositories.PetRepository
}

func NewPetsController(connection repositories.PetRepository) *ControllersPets {
	return &ControllersPets{repository: connection}
}

// GetPets godoc
//
//	@Summary		Lista todos os pets
//	@Description	Retorna todos os pets cadastrados
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Pet
//	@Failure		500	{object}	map[string]string
//	@Router			/pets [get]
func (p *ControllersPets) GetByClientes(ctx *gin.Context) {
	id := ctx.Param("cliente_id")

	clienteID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	result, err := p.repository.GetByClientes(uint(clienteID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, result)
}

// GetPet godoc
//
//	@Summary		Lista todos os pets
//	@Description	Retorna todos os pets cadastrados
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Pet
//	@Failure		500	{object}	map[string]string
//	@Router			/pets [get]
func (p *ControllersPets) GetPet(ctx *gin.Context) {
	id := ctx.Param("id")

	pet, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	result, err := p.repository.GetByID(uint(pet))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, result)
}

// UpdatePet godoc
//
//	@Summary		Lista todos os pets
//	@Description	Retorna todos os pets cadastrados
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Pet
//	@Failure		500	{object}	map[string]string
//	@Router			/pets [get]
func (p *ControllersPets) UpdatePet(ctx *gin.Context) {

	var pet models.Pets
	if err := ctx.BindJSON(&pet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	err := p.repository.Update(&pet)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, &pet)
}

// createClients godoc
//
//	@Summary		Lista todos os pets
//	@Description	Retorna todos os pets cadastrados
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Pet
//	@Failure		500	{object}	map[string]string
//	@Router			/pets [get]
func (p *ControllersPets) CreatePets(ctx *gin.Context) {
	var pet models.Pets
	err := ctx.ShouldBindJSON(&pet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = p.repository.Create(&pet)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, &pet)
}

// DeleteClient godoc
//
//	@Summary		Lista todos os pets
//	@Description	Retorna todos os pets cadastrados
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		bool
//	@Failure		500	{object}	map[string]string
//	@Router			/pets [get]
func (p *ControllersPets) DeletePet(ctx *gin.Context) {
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

	ctx.JSON(http.StatusOK, &user)
}

func (p *ControllersPets) ListByEmpresa(ctx *gin.Context) {
	id := ctx.Param("empresa_id")
	filters := requests.PetFilter{}
	empresaID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	filters.Especie = ctx.Query("especie")
	filters.Nome = ctx.Query("nome")

	result, err := p.repository.ListByEmpresa(uint(empresaID), filters)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (p *ControllersPets) GetTotalPets(ctx *gin.Context) {
	id := ctx.Param("empresa_id")
	empresaID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	result, err := p.repository.GetTotalPets(uint(empresaID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"total_pets": result})
}
func (p *ControllersPets) GetPetsPorEspecie(ctx *gin.Context) {
	id := ctx.Param("empresa_id")
	empresaID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	result, err := p.repository.GetPetsPorEspecie(uint(empresaID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
