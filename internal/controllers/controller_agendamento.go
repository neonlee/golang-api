package controllers

import (
	"fmt"
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AgendamentoController struct {
	agendamentoService repositories.AgendamentoRepository
}

func NewAgendamentoController(agendamentoService repositories.AgendamentoRepository) *AgendamentoController {
	return &AgendamentoController{agendamentoService: agendamentoService}
}

func (ac *AgendamentoController) Create(c *gin.Context) {
	var agendamento models.Agendamentos
	if err := c.ShouldBindJSON(&agendamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ac.agendamentoService.Create(&agendamento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Agendamento created successfully"})
}

func (ac *AgendamentoController) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	agendamento, err := ac.agendamentoService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, agendamento)
}

func (ac *AgendamentoController) Update(c *gin.Context) {
	var agendamento models.Agendamentos
	if err := c.ShouldBindJSON(&agendamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ac.agendamentoService.Update(&agendamento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Agendamento updated successfully"})
}
func (ac *AgendamentoController) Cancelar(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var payload struct {
		Motivo string `json:"motivo"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ac.agendamentoService.Cancelar(id, payload.Motivo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Agendamento canceled successfully"})
}
func (ac *AgendamentoController) ListByData(c *gin.Context) {
	empresaIDParam := c.Query("empresa_id")
	var empresaID uint
	_, err := fmt.Sscan(empresaIDParam, &empresaID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid empresa_id"})
		return
	}
	dateParam := c.Query("data")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})

		return
	}
	agendamentos, err := ac.agendamentoService.ListByData(empresaID, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, agendamentos)
}

func (ac *AgendamentoController) ListByPeriodo(c *gin.Context) {
	empresaIDParam := c.Query("empresa_id")
	var empresaID uint
	_, err := fmt.Sscan(empresaIDParam, &empresaID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid empresa_id"})
		return
	}
	inicioParam := c.Query("inicio")
	inicio, err := time.Parse("2006-01-02", inicioParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inicio date format. Use YYYY-MM-DD"})
		return
	}
	fimParam := c.Query("fim")
	fim, err := time.Parse("2006-01-02", fimParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fim date format. Use YYYY-MM-DD"})
		return
	}
	agendamentos, err := ac.agendamentoService.ListByPeriodo(empresaID, inicio, fim)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, agendamentos)
}
func (ac *AgendamentoController) ListByPet(c *gin.Context) {
	petIDParam := c.Query("pet_id")
	var petID uint
	_, err := fmt.Sscan(petIDParam, &petID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pet_id"})
		return
	}
	agendamentos, err := ac.agendamentoService.ListByPet(petID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, agendamentos)
}

func (ac *AgendamentoController) VerificarDisponibilidade(c *gin.Context) {
	empresaIDParam := c.Query("empresa_id")
	var empresaID uint
	_, err := fmt.Sscan(empresaIDParam, &empresaID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid empresa_id"})
		return
	}
	dataHoraParam := c.Query("data_hora")
	dataHora, err := time.Parse("2006-01-02T15:04", dataHoraParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data_hora format. Use YYYY-MM-DDTHH:MM"})
		return
	}
	servicoIDParam := c.Query("servico_id")
	var servicoID uint
	_, err = fmt.Sscan(servicoIDParam, &servicoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid servico_id"})
		return
	}
	disp, err := ac.agendamentoService.VerificarDisponibilidade(empresaID, dataHora, servicoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"disponivel": disp})
}
func (c *AgendamentoController) GetHorariosDisponiveis(ctx *gin.Context) {
	empresaID := uint(ctx.GetInt("empresa_id"))
	medicoID, _ := strconv.ParseUint(ctx.Query("medico_id"), 10, 32)
	dataStr := ctx.Query("data")
	tipoServico := ctx.Query("tipo_servico")

	// Validar parâmetros
	if dataStr == "" || tipoServico == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetros data e tipo_servico são obrigatórios"})
		return
	}

	// Converter data
	data, err := time.Parse("2006-01-02", dataStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data inválido. Use YYYY-MM-DD"})
		return
	}
	servico, errTipo := strconv.Atoi(tipoServico)
	if errTipo != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data inválido. Use YYYY-MM-DD"})
		return
	}
	// Buscar horários disponíveis
	horarios, err := c.agendamentoService.GetHorariosDisponiveis(empresaID, uint(medicoID), data, servico)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"medico_id":            medicoID,
		"data":                 data.Format("2006-01-02"),
		"tipo_servico":         tipoServico,
		"horarios_disponiveis": horarios,
		"total_horarios":       len(horarios),
	})
}
