package controllers

import (
	"net/http"
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
func (c *FinanceiroController) GetResumoFinanceiro(ctx *gin.Context) {
	empresaIDParam := ctx.Param("empresa_id")
	empresaID, err := strconv.ParseUint(empresaIDParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de empresa inválido"})
		return
	}
	resumo, err := c.financeiroRepo.GetResumoFinanceiro(uint(empresaID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resumo)
}
