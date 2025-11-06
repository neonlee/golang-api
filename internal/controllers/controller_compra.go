package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"petApi/internal/requests"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompraController struct {
	Repo repositories.ComprasRepository
}

func NewCompraController(repo repositories.ComprasRepository) *CompraController {
	return &CompraController{Repo: repo}
}

func (cc *CompraController) CreateCompra(c *gin.Context) {
	var compra models.Compras
	if err := c.ShouldBindJSON(&compra); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	compras, err := cc.Repo.Create(&compra, compra.Itens)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, compras)
}
func (cc *CompraController) GetCompraByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	compra, err := cc.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, compra)
}
func (cc *CompraController) CancelarCompra(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var body struct {
		Motivo string `json:"motivo"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.Repo.Cancelar(uint(id), body.Motivo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Compra cancelada com sucesso"})
}

func (cc *CompraController) UpdateCompra(c *gin.Context) {
	var compra models.Compras
	if err := c.ShouldBindJSON(&compra); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.Repo.Update(&compra); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, compra)
}

func (cc *CompraController) ListByEmpresa(c *gin.Context) {
	var filtro requests.CompraFilter
	if err := c.ShouldBindQuery(&filtro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	empresaIDParam := c.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empresa inválido"})
		return
	}
	compras, err := cc.Repo.ListByEmpresa(uint(empresaID), filtro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, compras)
}

func (cc *CompraController) ListByFornecedor(c *gin.Context) {
	var filtro requests.CompraFilter
	if err := c.ShouldBindQuery(&filtro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fornecedorIDParam := c.Param("fornecedor_id")
	fornecedorID, err := strconv.Atoi(fornecedorIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de fornecedor inválido"})
		return
	}
	compras, err := cc.Repo.ListByFornecedor(uint(fornecedorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, compras)
}

func (cc *CompraController) GetComprassPorPeriodo(c *gin.Context) {
	var filtro requests.PeriodoFilter
	if err := c.ShouldBindQuery(&filtro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	empresaIDParam := c.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empresa inválido"})
		return
	}
	compras, err := cc.Repo.GetComprassPorPeriodo(uint(empresaID), filtro.Inicio, filtro.Fim)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, compras)
}

func (cc *CompraController) GetResumoComprass(c *gin.Context) {
	var filtro requests.ResumoFilter
	if err := c.ShouldBindQuery(&filtro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	empresaIDParam := c.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empresa inválido"})
		return
	}
	resumo, err := cc.Repo.GetResumoComprass(uint(empresaID), filtro.Periodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resumo)
}

func (cc *CompraController) GetItensByCompraID(c *gin.Context) {
	idParam := c.Param("compra_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	itens, err := cc.Repo.GetItensByCompraID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, itens)
}
