package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LogsController struct {
	logRepo repositories.LogRepository
}

func NewLogsController(logRepo repositories.LogRepository) *LogsController {
	return &LogsController{logRepo: logRepo}
}

func (c *LogsController) CreateLog(ctx *gin.Context) {
	var log models.LogSistema
	if err := ctx.ShouldBindJSON(&log); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.logRepo.Create(&log); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Log criado com sucesso"})
}
func (c *LogsController) GetLogByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	log, err := c.logRepo.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, log)
}
func (c *LogsController) GetLogsByUsuario(ctx *gin.Context) {
	usuarioIDParam := ctx.Param("usuario_id")
	usuarioID, err := strconv.ParseUint(usuarioIDParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuário inválido"})
		return
	}
	limiteParam := ctx.DefaultQuery("limite", "10")
	limite, err := strconv.Atoi(limiteParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Limite inválido"})
		return
	}
	logs, err := c.logRepo.GetLogsByUsuario(uint(usuarioID), limite)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, logs)
}
func (c *LogsController) GetLogsByModulo(ctx *gin.Context) {
	empresaIDParam := ctx.Param("empresa_id")
	empresaID, err := strconv.ParseUint(empresaIDParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de empresa inválido"})
		return
	}
	modulo := ctx.Param("modulo")
	inicio := ctx.Query("inicio")
	fim := ctx.Query("fim")
	logs, err := c.logRepo.GetLogsByModulo(uint(empresaID), modulo, inicio, fim)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, logs)
}
func (c *LogsController) GetLogsErro(ctx *gin.Context) {
	empresaIDParam := ctx.Param("empresa_id")
	empresaID, err := strconv.ParseUint(empresaIDParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de empresa inválido"})
		return
	}
	diasParam := ctx.DefaultQuery("dias", "7")
	dias, err := strconv.Atoi(diasParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Número de dias inválido"})
		return
	}
	logs, err := c.logRepo.GetLogsErro(uint(empresaID), dias)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, logs)
}
