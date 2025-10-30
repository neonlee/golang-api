package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type MedicoVeterinarioRepository interface {
	CreateVeterinario(medico models.MedicosVeterinarios) error
	AddEspecialidade(especialidade models.MedicoEspecialidade) error
	DeleteEspecialidade(idMedico uint, idEspecialidade uint) error
	AddDisponibilidade(disponibilidade models.MedicoDisponibilidade) error
	UpdateDisponibilidade(id uint, disponibilidade models.MedicoDisponibilidade) error
	DeleteDisponibilidade(id uint) error
	ListarMedicosComEspecialidadesEDisponibilidades() ([]models.MedicosVeterinarios, error)
}

type medicoVeterinarioRepository struct {
	db *gorm.DB
}

func NewMedicoVeterinarioRepository(connection *gorm.DB) MedicoVeterinarioRepository {
	return &medicoVeterinarioRepository{db: connection}
}

func (m *medicoVeterinarioRepository) CreateVeterinario(medico models.MedicosVeterinarios) error {
	return m.db.Create(&medico).Error
}

func (m *medicoVeterinarioRepository) AddEspecialidade(especialidade models.MedicoEspecialidade) error {
	return m.db.Create(&especialidade).Error
}
func (m *medicoVeterinarioRepository) AddDisponibilidade(disponibilidade models.MedicoDisponibilidade) error {
	return m.db.Create(&disponibilidade).Error
}
func (m *medicoVeterinarioRepository) UpdateDisponibilidade(id uint, disponibilidade models.MedicoDisponibilidade) error {
	return m.db.Model(&models.MedicoDisponibilidade{}).Where("id = ?", id).Updates(disponibilidade).Error
}
func (m *medicoVeterinarioRepository) DeleteDisponibilidade(id uint) error {
	return m.db.Delete(&models.MedicoDisponibilidade{}, id).Error
}
func (m *medicoVeterinarioRepository) DeleteEspecialidade(idMedico uint, idEspecialidade uint) error {
	return m.db.Where("medico_id = ? AND id = ?", idMedico, idEspecialidade).Delete(&models.MedicoEspecialidade{}).Error
}
func (m *medicoVeterinarioRepository) ListarMedicosComEspecialidadesEDisponibilidades() ([]models.MedicosVeterinarios, error) {
	var medicos []models.MedicosVeterinarios
	err := m.db.Preload("Especialidades").Preload("Disponibilidades").Preload("Funcionarios").Find(&medicos).Error
	return medicos, err
}
