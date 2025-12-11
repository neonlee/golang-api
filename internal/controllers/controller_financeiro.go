package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FinanceiroController struct {
	financeiroRepo repositories.FinanceiroRepository
}

func NewFinanceiroController(financeiroRepo repositories.FinanceiroRepository) *FinanceiroController {
	return &FinanceiroController{financeiroRepo: financeiroRepo}
}
func (c *FinanceiroController) CreateContaReceber(ctx *gin.Context) {
	var conta models.ContaReceber
	if err := ctx.ShouldBindJSON(&conta); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.financeiroRepo.CreateContaReceber(&conta); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Conta a receber criada com sucesso"})
}

func (c *FinanceiroController) GetContaReceberByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	conta, err := c.financeiroRepo.GetContaReceberByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, conta)
}
func (c *FinanceiroController) BaixarContaReceber(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var payload struct {
		DataPagamento  string `json:"data_pagamento"`
		FormaPagamento string `json:"forma_pagamento"`
	}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.financeiroRepo.BaixarContaReceber(uint(id), payload.DataPagamento, payload.FormaPagamento); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Conta a receber baixada com sucesso"})
}
func (c *FinanceiroController) CreateContaPagar(ctx *gin.Context) {
	var conta models.ContaPagar
	if err := ctx.ShouldBindJSON(&conta); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.financeiroRepo.CreateContaPagar(&conta); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Conta a pagar criada com sucesso"})
}
func (c *FinanceiroController) GetContaPagarByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	conta, err := c.financeiroRepo.GetContaPagarByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, conta)
}
func (c *FinanceiroController) PagarContaPagar(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var payload struct {
		DataPagamento  string `json:"data_pagamento"`
		FormaPagamento string `json:"forma_pagamento"`
	}
	if err := ctx.ShouldBindJSON(&payload); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.financeiroRepo.PagarContaPagar(uint(id), payload.DataPagamento, payload.FormaPagamento); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Conta a pagar paga com sucesso"})
}
func (c *FinanceiroController) GetFluxoCaixa(ctx *gin.Context) {
	empresaIDParam := ctx.Param("empresa_id")
	empresaID, err := strconv.ParseUint(empresaIDParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de empresa inválido"})
		return
	}
	inicioParam := ctx.Query("inicio")
	fimParam := ctx.Query("fim")
	fluxo, err := c.financeiroRepo.GetFluxoCaixa(uint(empresaID), inicioParam, fimParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, fluxo)
}
