// repositories/agendamento_repository.go
package repositories

import (
	"fmt"
	"petApi/internal/models"
	"time"

	"gorm.io/gorm"
)

// AgendamentoRepository interface
type AgendamentoRepository interface {
	Create(agendamento *models.Agendamentos) error
	GetByID(id uint) (*models.Agendamentos, error)
	Update(agendamento *models.Agendamentos) error
	Cancelar(id uint, motivo string) error
	ListByData(empresaID uint, data time.Time) ([]models.Agendamentos, error)
	ListByPeriodo(empresaID uint, inicio, fim time.Time) ([]models.Agendamentos, error)
	ListByPet(petID uint) ([]models.Agendamentos, error)
	VerificarDisponibilidade(empresaID uint, dataHora time.Time, servicoID uint) (bool, error)
	GetHorariosDisponiveis(empresaID uint, medicoID uint, data time.Time, tipoServico int) ([]time.Time, error)
}

type agendamentoRepository struct {
	db *gorm.DB
}

func NewAgendamentoRepository(db *gorm.DB) AgendamentoRepository {
	return &agendamentoRepository{db: db}
}

func (r *agendamentoRepository) Create(agendamento *models.Agendamentos) error {
	return r.db.Create(agendamento).Error
}

func (r *agendamentoRepository) GetByID(id uint) (*models.Agendamentos, error) {
	var agendamento models.Agendamentos
	err := r.db.
		Preload("Clientes").
		Preload("Pet").
		Preload("TipoServico").
		Preload("Usuario").
		First(&agendamento, id).Error

	return &agendamento, err
}

func (r *agendamentoRepository) Update(agendamento *models.Agendamentos) error {
	return r.db.Save(agendamento).Error
}

func (r *agendamentoRepository) Cancelar(id uint, motivo string) error {
	return r.db.Model(&models.Agendamentos{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":      "cancelado",
			"observacoes": motivo,
		}).Error
}

func (r *agendamentoRepository) ListByData(empresaID uint, data time.Time) ([]models.Agendamentos, error) {
	var agendamentos []models.Agendamentos

	dataStr := data.Format("2006-01-02")

	err := r.db.
		Where("empresa_id = ? AND DATE(data_agendamento) = ?", empresaID, dataStr).
		Preload("Clientes").
		Preload("Pet").
		Preload("TipoServico").
		Preload("Usuario").
		Order("data_agendamento ASC").
		Find(&agendamentos).Error

	return agendamentos, err
}

func (r *agendamentoRepository) ListByPeriodo(empresaID uint, inicio, fim time.Time) ([]models.Agendamentos, error) {
	var agendamentos []models.Agendamentos

	err := r.db.
		Where("empresa_id = ? AND data_agendamento BETWEEN ? AND ?", empresaID, inicio, fim).
		Preload("Clientes").
		Preload("Pet").
		Preload("TipoServico").
		Preload("Usuario").
		Order("data_agendamento ASC").
		Find(&agendamentos).Error

	return agendamentos, err
}

func (r *agendamentoRepository) ListByPet(petID uint) ([]models.Agendamentos, error) {
	var agendamentos []models.Agendamentos

	err := r.db.
		Where("pet_id = ?", petID).
		Preload("Clientes").
		Preload("TipoServico").
		Preload("Usuario").
		Order("data_agendamento DESC").
		Find(&agendamentos).Error

	return agendamentos, err
}

func (r *agendamentoRepository) VerificarDisponibilidade(empresaID uint, dataHora time.Time, servicoID uint) (bool, error) {
	var count int64

	// Buscar duração do serviço
	var servico models.TiposServicos
	if err := r.db.First(&servico, servicoID).Error; err != nil {
		return false, err
	}

	fim := dataHora.Add(time.Duration(servico.DuracaoMinutos) * time.Minute)

	err := r.db.Model(&models.Agendamentos{}).
		Where("empresa_id = ? AND tipo_servico_id = ? AND status NOT IN ('cancelado')", empresaID, servicoID).
		Where("(data_agendamento BETWEEN ? AND ?)", dataHora, fim).
		Count(&count).Error

	return count == 0, err
}

func (r *agendamentoRepository) GetHorariosDisponiveis(empresaID uint, medicoID uint, data time.Time, tipoServico int) ([]time.Time, error) {
	var horariosDisponiveis []time.Time
	getServico, err := r.getDuracaoServico(tipoServico)
	if err != nil {
		return nil, err
	}
	// 1. Buscar disponibilidade do médico para o tipo de serviço
	if medicoID != 0 {
		var disponibilidade models.MedicoDisponibilidade
		err := r.db.Where("medico_id = ?", medicoID).First(&disponibilidade).Error
		if disponibilidade.TipoServico != getServico.Categoria {
			return nil, fmt.Errorf("erro ao buscar disponibilidade do médico: %v", err)
		}
		if err != nil {
			return nil, fmt.Errorf("erro ao buscar disponibilidade do médico: %v", err)
		}
		if !disponibilidade.EstaDisponivel(data.Weekday()) {
			return nil, fmt.Errorf("erro ao buscar agendamentos: %v", err)
		}
	}
	agendamentos, err := r.ListAgendamentoByDataAndTipoServico(empresaID, tipoServico, data)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar agendamentos: %v", err)
	}

	// 3. Obter duração do serviço

	HoraInicio, err := time.Parse("15:04:05", "08:00:00")
	if err != nil {
		return nil, fmt.Errorf("erro ao converter hora início: %v", err)
	}
	HoraFim, err := time.Parse("15:04:05", "18:00:00")
	if err != nil {
		return nil, fmt.Errorf("erro ao converter hora fim: %v", err)
	}
	// 4. Gerar horários baseado na disponibilidade
	currentTime := time.Date(data.Year(), data.Month(), data.Day(),
		HoraInicio.Hour(), HoraInicio.Minute(), 0, 0, data.Location())

	endTime := time.Date(data.Year(), data.Month(), data.Day(),
		HoraFim.Hour(), HoraFim.Minute(), 0, 0, data.Location())

	for currentTime.Before(endTime) {
		horaFimServico := currentTime.Add(time.Duration(getServico.DuracaoMinutos) * time.Minute)

		// Verificar se o horário cabe na janela de disponibilidade
		if horaFimServico.After(endTime) {
			break
		}

		disponivel := true

		// Verificar conflito com agendamentos existentes
		for _, agendamento := range agendamentos {
			fimAgendamento := agendamento.DataAgendamento.Add(time.Duration(getServico.DuracaoMinutos) * time.Minute)

			if temConflitoHorario(currentTime, horaFimServico, agendamento.DataAgendamento, fimAgendamento) {
				disponivel = false
				break
			}
		}

		if disponivel {
			horariosDisponiveis = append(horariosDisponiveis, currentTime)
		}

		currentTime = currentTime.Add(time.Duration(getServico.DuracaoMinutos) * time.Minute)
	}

	return horariosDisponiveis, nil
}

// Buscar agendamentos do servico na data
func (r *agendamentoRepository) ListAgendamentoByDataAndTipoServico(empresaID uint, tipoServico int, data time.Time) ([]models.Agendamentos, error) {
	var agendamentos []models.Agendamentos
	dataInicio := time.Date(data.Year(), data.Month(), data.Day(), 0, 0, 0, 0, data.Location())
	dataFim := dataInicio.Add(24 * time.Hour)

	err := r.db.Where("empresa_id = ? AND tipo_servico_id = ? AND data_agendamento BETWEEN ? AND ? AND status NOT IN ('cancelado', 'finalizado')",
		empresaID, tipoServico, dataInicio, dataFim).
		Find(&agendamentos).Error
	return agendamentos, err
}

func (r *agendamentoRepository) getDuracaoServico(tipoServico int) (*models.TiposServicos, error) {
	var servico *models.TiposServicos
	err := r.db.First(&servico, tipoServico).Error
	if err != nil {
		return nil, fmt.Errorf("erro ao obter tipo servico: %v", err)
	}
	return servico, nil
}

// Verificar conflito de horários
func temConflitoHorario(inicio1, fim1, inicio2, fim2 time.Time) bool {
	return (inicio1.Before(fim2) && fim1.After(inicio2))
}
