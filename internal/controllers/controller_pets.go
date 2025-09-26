package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
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
// func (p *ControllersPets) GetPets(ctx *gin.Context) {
// 	result, err := p.repository.GetPets()

// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, err)
// 	}

// 	ctx.JSON(http.StatusOK, result)
// }

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
	err := ctx.BindJSON(&pet)
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
