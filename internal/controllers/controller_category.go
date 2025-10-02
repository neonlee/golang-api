package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoriaController struct {
	Repo repositories.CategoriaRepository
}

func NewCategoriaController(repo repositories.CategoriaRepository) *CategoriaController {
	return &CategoriaController{Repo: repo}
}

func (c *CategoriaController) GetAll(ctx *gin.Context) {
	categorias, err := c.Repo.GetWithProdutos(1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categorias)
}

func (c *CategoriaController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	categoria, err := c.Repo.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categoria)
}

func (c *CategoriaController) Create(ctx *gin.Context) {
	var categoria models.CategoriasProdutos
	if err := ctx.ShouldBindJSON(&categoria); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.Repo.Create(&categoria); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, categoria)
}

func (c *CategoriaController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var categoria models.CategoriasProdutos
	if err := ctx.ShouldBindJSON(&categoria); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	categoria.ID = uint(id)
	if err := c.Repo.Update(&categoria); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categoria)
}

func (c *CategoriaController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.Repo.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
