package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VendasController struct {
	vendasRepo repositories.VendaRepository
}

func NewVendasController(vendasRepo repositories.VendaRepository) *VendasController {
	return &VendasController{vendasRepo: vendasRepo}
}

func (c *VendasController) CreateVenda(ctx *gin.Context) {
	var venda models.Vendas
	var itens []models.VendaItem
	if err := ctx.ShouldBindJSON(&venda); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(&itens); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.vendasRepo.Create(&venda, itens); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Venda criada com sucesso", "venda_id": venda.ID})
}

func (c *VendasController) GetVendaByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	venda, err := c.vendasRepo.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, venda)
}
func (c *VendasController) UpdateVendaStatus(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var statusUpdate struct {
		Status string `json:"status"`
	}
	if err := ctx.ShouldBindJSON(&statusUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.vendasRepo.UpdateStatus(uint(id), statusUpdate.Status); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Status da venda atualizado com sucesso"})
}

func (c *VendasController) ListVendas(ctx *gin.Context) {
	// Implementar listagem de vendas conforme necessário
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "Listagem de vendas não implementada"})

}
