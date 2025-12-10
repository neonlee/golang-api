package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EstoqueController struct {
	estoqueRepo repositories.EstoqueRepository
}

func NewEstoqueController(estoqueRepo repositories.EstoqueRepository) *EstoqueController {
	return &EstoqueController{estoqueRepo: estoqueRepo}
}

func (c *EstoqueController) MovimentarEstoque(ctx *gin.Context) {
	var movimentacao models.MovimentacaoEstoques
	if err := ctx.ShouldBindJSON(&movimentacao); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.estoqueRepo.MovimentarEstoque(&movimentacao); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Movimentação de estoque registrada com sucesso"})
}

func (c *EstoqueController) GetHistoricoEstoque(ctx *gin.Context) {
	produtoIDParam := ctx.Param("produto_id")
	produtoID, err := strconv.ParseUint(produtoIDParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de produto inválido"})
		return
	}

	movimentacoes, err := c.estoqueRepo.GetHistoricoEstoque(uint(produtoID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movimentacoes)
}

func (c *EstoqueController) GetSaldoAtual(ctx *gin.Context) {
	produtoIDParam := ctx.Param("produto_id")
	produtoID, err := strconv.ParseUint(produtoIDParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de produto inválido"})
		return
	}
	saldo, err := c.estoqueRepo.GetSaldoAtual(uint(produtoID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"saldo_atual": saldo})
}

func (c *EstoqueController) AjustarEstoque(ctx *gin.Context) {
	var ajuste struct {
		ProdutoID      uint   `json:"produto_id"`
		NovaQuantidade int    `json:"nova_quantidade"`
		UsuarioID      uint   `json:"usuario_id"`
		Motivo         string `json:"motivo"`
	}
	if err := ctx.ShouldBindJSON(&ajuste); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.estoqueRepo.AjustarEstoque(ajuste.ProdutoID, ajuste.NovaQuantidade, ajuste.Motivo, ajuste.UsuarioID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Estoque ajustado com sucesso"})
}
func (c *EstoqueController) TransferirEstoque(ctx *gin.Context) {
	var transferencia struct {
		OrigemID   uint `json:"origem_id"`
		DestinoID  uint `json:"destino_id"`
		Quantidade int  `json:"quantidade"`
		UsuarioID  uint `json:"usuario_id"`
	}

	if err := ctx.ShouldBindJSON(&transferencia); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.estoqueRepo.TransferirEstoque(transferencia.OrigemID, transferencia.DestinoID, transferencia.Quantidade, transferencia.UsuarioID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Estoque transferido com sucesso"})
}

func (c *EstoqueController) GetRelatorioEstoque(ctx *gin.Context) {
	empresaIDParam := ctx.Param("empresa_id")
	empresaID, err := strconv.ParseUint(empresaIDParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de empresa inválido"})
		return

	}
	relatorio, err := c.estoqueRepo.GetRelatorioEstoque(uint(empresaID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, relatorio)
}

func (c *EstoqueController) GetMovimentacoesPorPeriodo(ctx *gin.Context) {
	produtoIDParam := ctx.Param("produto_id")
	produtoID, err := strconv.ParseUint(produtoIDParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de produto inválido"})
		return
	}
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")
	movimentacoes, err := c.estoqueRepo.GetMovimentacoesPorPeriodo(uint(produtoID), startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movimentacoes)
}
