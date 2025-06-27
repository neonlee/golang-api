package Controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllersPets struct {
	repository *repositories.PetsRepository
}

func NewPetsController(connection *repositories.PetsRepository) *ControllersPets {
	return &ControllersPets{repository: connection}
}

// GetPets godoc
//
//	@Summary		Lista todos os pets
//	@Description	Retorna todos os pets cadastrados
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		Pet
//	@Failure		500	{object}	map[string]string
//	@Router			/pets [get]
func (p *ControllersPets) GetPets(ctx *gin.Context) {
	result, err := p.repository.GetPets()

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
//	@Success		200	{array}		Pet
//	@Failure		500	{object}	map[string]string
//	@Router			/pets [get]
func (p *ControllersPets) GetPet(ctx *gin.Context) {
	id := ctx.Param("id")

	pet, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	result, err := p.repository.GetPet(pet)

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
//	@Success		200	{array}		Pet
//	@Failure		500	{object}	map[string]string
//	@Router			/pets [get]
func (p *ControllersPets) UpdatePet(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	var pet models.Pet
	if err := ctx.BindJSON(&pet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	client, err := p.repository.UpdatePet(user, pet)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, client)
}

// createClients godoc
//
//	@Summary		Lista todos os pets
//	@Description	Retorna todos os pets cadastrados
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		Pet
//	@Failure		500	{object}	map[string]string
//	@Router			/pets [get]
func (p *ControllersPets) CreatePets(ctx *gin.Context) {
	var pet models.Pet
	err := ctx.BindJSON(&pet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := p.repository.Create(pet)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, result)
}

// DeleteClient godoc
//
//	@Summary		Lista todos os pets
//	@Description	Retorna todos os pets cadastrados
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		Pet
//	@Failure		500	{object}	map[string]string
//	@Router			/pets [get]
func (p *ControllersPets) DeletePet(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	client, err := p.repository.DeletePet(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, client)
}
