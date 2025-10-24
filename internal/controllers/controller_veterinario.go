package controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllersVeterinario struct {
	repository repositories.MedicoVeterinarioRepository
}

func NewMedicoVeterinarioController(connection repositories.MedicoVeterinarioRepository) *ControllersVeterinario {
	return &ControllersVeterinario{repository: connection}
}

func (v *ControllersVeterinario) CreateVeterinario(ctx *gin.Context) {
	var medico models.MedicoVeterinario
	if err := ctx.BindJSON(&medico); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	err := v.repository.CreateVeterinario(medico)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Médico veterinário criado com sucesso"})
}
func (v *ControllersVeterinario) AddEspecialidade(ctx *gin.Context) {
	var especialidade models.MedicoEspecialidade
	if err := ctx.BindJSON(&especialidade); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	err := v.repository.AddEspecialidade(especialidade)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Especialidade adicionada com sucesso"})
}
func (v *ControllersVeterinario) DeleteEspecialidade(ctx *gin.Context) {
	medicoIDParam := ctx.Param("medico_id")
	especialidadeIDParam := ctx.Param("especialidade_id")
	medicoID, err := strconv.ParseUint(medicoIDParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID do médico inválido"})
		return
	}
	especialidadeID, err := strconv.ParseUint(especialidadeIDParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID da especialidade inválido"})
		return
	}
	err = v.repository.DeleteEspecialidade(uint(medicoID), uint(especialidadeID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Especialidade deletada com sucesso"})
}

func (v *ControllersVeterinario) AddDisponibilidade(ctx *gin.Context) {
	var disponibilidade models.MedicoDisponibilidade
	if err := ctx.BindJSON(&disponibilidade); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	err := v.repository.AddDisponibilidade(disponibilidade)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Disponibilidade adicionada com sucesso"})
}
func (v *ControllersVeterinario) UpdateDisponibilidade(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	var disponibilidade models.MedicoDisponibilidade
	if err := ctx.BindJSON(&disponibilidade); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	err = v.repository.UpdateDisponibilidade(uint(id), disponibilidade)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Disponibilidade atualizada com sucesso"})
}
func (v *ControllersVeterinario) DeleteDisponibilidade(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	err = v.repository.DeleteDisponibilidade(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Disponibilidade deletada com sucesso"})
}

func (v *ControllersVeterinario) ListarMedicosComEspecialidadesEDisponibilidades(ctx *gin.Context) {
	medicos, err := v.repository.ListarMedicosComEspecialidadesEDisponibilidades()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao listar médicos"})
		return
	}
	ctx.JSON(http.StatusOK, medicos)
}
