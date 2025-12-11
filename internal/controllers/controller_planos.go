package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlanoController struct {
	planoRepo repositories.PlanoRepository
}

func NewPlanosController(planoRepo repositories.PlanoRepository) *PlanoController {
	return &PlanoController{planoRepo: planoRepo}
}
func (c *PlanoController) GetPlanoByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	plano, err := c.planoRepo.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, plano)
}

func (c *PlanoController) ListAllPlanos(ctx *gin.Context) {
	planos, err := c.planoRepo.ListAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, planos)
}
func (c *PlanoController) CreatePlano(ctx *gin.Context) {
	var plano models.Planos
	if err := ctx.ShouldBindJSON(&plano); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.planoRepo.Create(&plano); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Plano criado com sucesso"})
}

func (c *PlanoController) UpdatePlano(ctx *gin.Context) {
	var plano models.Planos
	if err := ctx.ShouldBindJSON(&plano); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.planoRepo.Update(&plano); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Plano atualizado com sucesso"})
}
