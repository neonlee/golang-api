package controllers

import (
	"net/http"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	Repo repositories.DashboardRepository
}

func NewDashboardController(repo repositories.DashboardRepository) *DashboardController {
	return &DashboardController{Repo: repo}
}

func (dc *DashboardController) GetResumoVendas(c *gin.Context) {

	empresaIDParam := c.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	data, err := dc.Repo.GetResumoVendas(uint(empresaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)

}

func (dc *DashboardController) GetResumoFinanceiro(c *gin.Context) {
	empresaIDParam := c.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	data, err := dc.Repo.GetResumoFinanceiro(uint(empresaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
func (dc *DashboardController) GetProximosAgendamentos(c *gin.Context) {
	empresaIDParam := c.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	data, err := dc.Repo.GetProximosAgendamentos(uint(empresaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
func (dc *DashboardController) GetAlertasEstoque(c *gin.Context) {
	empresaIDParam := c.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	data, err := dc.Repo.GetAlertasEstoque(uint(empresaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
func (dc *DashboardController) GetMetricasGerais(c *gin.Context) {
	empresaIDParam := c.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaIDParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	data, err := dc.Repo.GetMetricasGerais(uint(empresaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
