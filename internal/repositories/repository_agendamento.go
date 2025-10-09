// repositories/agendamento_repository.go
package repositories

import (
	"petApi/internal/models"
	"time"

	"gorm.io/gorm"
)

// AgendamentoRepository interface
type AgendamentoRepository interface {
	Create(agendamento *models.Agendamento) error
	GetByID(id uint) (*models.Agendamento, error)
	Update(agendamento *models.Agendamento) error
	Cancelar(id uint, motivo string) error
	ListByData(empresaID uint, data time.Time) ([]models.Agendamento, error)
	ListByPeriodo(empresaID uint, inicio, fim time.Time) ([]models.Agendamento, error)
	ListByPet(petID uint) ([]models.Agendamento, error)
	VerificarDisponibilidade(empresaID uint, dataHora time.Time, servicoID uint) (bool, error)
	GetHorariosDisponiveis(empresaID uint, data time.Time, servicoID uint) ([]time.Time, error)
}

type agendamentoRepository struct {
	db *gorm.DB
}

func NewAgendamentoRepository(db *gorm.DB) AgendamentoRepository {
	return &agendamentoRepository{db: db}
}

func (r *agendamentoRepository) Create(agendamento *models.Agendamento) error {
	return r.db.Create(agendamento).Error
}

func (r *agendamentoRepository) GetByID(id uint) (*models.Agendamento, error) {
	var agendamento models.Agendamento
	err := r.db.
		Preload("Clientes").
		Preload("Pet").
		Preload("TipoServico").
		Preload("Usuario").
		First(&agendamento, id).Error

	return &agendamento, err
}

func (r *agendamentoRepository) Update(agendamento *models.Agendamento) error {
	return r.db.Save(agendamento).Error
}

func (r *agendamentoRepository) Cancelar(id uint, motivo string) error {
	return r.db.Model(&models.Agendamento{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":      "cancelado",
			"observacoes": motivo,
		}).Error
}

func (r *agendamentoRepository) ListByData(empresaID uint, data time.Time) ([]models.Agendamento, error) {
	var agendamentos []models.Agendamento

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

func (r *agendamentoRepository) ListByPeriodo(empresaID uint, inicio, fim time.Time) ([]models.Agendamento, error) {
	var agendamentos []models.Agendamento

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

func (r *agendamentoRepository) ListByPet(petID uint) ([]models.Agendamento, error) {
	var agendamentos []models.Agendamento

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

	err := r.db.Model(&models.Agendamento{}).
		Where("empresa_id = ? AND tipo_servico_id = ? AND status NOT IN ('cancelado')", empresaID, servicoID).
		Where("(data_agendamento BETWEEN ? AND ?) OR (data_agendamento + INTERVAL '1 minute' * duracao_minutos BETWEEN ? AND ?)",
			dataHora, fim, dataHora, fim).
		Count(&count).Error

	return count == 0, err
}

func (r *agendamentoRepository) GetHorariosDisponiveis(empresaID uint, data time.Time, servicoID uint) ([]time.Time, error) {
	var horariosDisponiveis []time.Time

	// Horário de funcionamento (exemplo: 8h às 18h)
	horaInicio := time.Date(data.Year(), data.Month(), data.Day(), 8, 0, 0, 0, data.Location())
	horaFim := time.Date(data.Year(), data.Month(), data.Day(), 18, 0, 0, 0, data.Location())

	// Buscar agendamentos do dia
	agendamentos, err := r.ListByData(empresaID, data)
	if err != nil {
		return nil, err
	}

	// Gerar horários disponíveis (a cada 30 minutos)
	for hora := horaInicio; hora.Before(horaFim); hora = hora.Add(30 * time.Minute) {
		disponivel := true

		// Verificar conflito com agendamentos existentes
		for _, agendamento := range agendamentos {
			fimAgendamento := agendamento.DataAgendamento.Add(
				time.Duration(agendamento.TipoServico.DuracaoMinutos) * time.Minute)

			if (hora.After(agendamento.DataAgendamento) && hora.Before(fimAgendamento)) ||
				(hora.Equal(agendamento.DataAgendamento)) {
				disponivel = false
				break
			}
		}

		if disponivel {
			horariosDisponiveis = append(horariosDisponiveis, hora)
		}
	}

	return horariosDisponiveis, nil
}
