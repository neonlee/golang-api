package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type ProntuarioRepository interface {
	Create(prontuario *models.Prontuarios) error
	GetByID(id uint) (*models.Prontuarios, error)
	Update(prontuario *models.Prontuarios) error
	GetByPet(petID uint) ([]models.Prontuarios, error)
	GetByVeterinario(veterinarioID uint, inicio, fim string) ([]models.Prontuarios, error)
	GetUltimoProntuario(petID uint) (*models.Prontuarios, error)
	RegistrarVacina(vacina *models.Vacinas) error
	GetVacinasPorPet(petID uint) ([]models.Vacinas, error)
	GetVacinasVencidas(empresaID uint) ([]models.Vacinas, error)
}
type prontuarioRepository struct {
	db *gorm.DB
}

func NewProntuarioRepository(db *gorm.DB) ProntuarioRepository {
	return &prontuarioRepository{db: db}
}

func (r *prontuarioRepository) Create(prontuario *models.Prontuarios) error {
	return r.db.Create(prontuario).Error
}

func (r *prontuarioRepository) GetByID(id uint) (*models.Prontuarios, error) {
	var prontuario models.Prontuarios
	err := r.db.
		Preload("Pet").
		Preload("Pet.Clientes").
		Preload("Veterinario").
		First(&prontuario, id).Error

	return &prontuario, err
}

func (r *prontuarioRepository) Update(prontuario *models.Prontuarios) error {
	return r.db.Save(prontuario).Error
}

func (r *prontuarioRepository) GetByPet(petID uint) ([]models.Prontuarios, error) {
	var prontuarios []models.Prontuarios

	err := r.db.
		Where("pet_id = ?", petID).
		Preload("Veterinario").
		Order("data_consulta DESC").
		Find(&prontuarios).Error

	return prontuarios, err
}

func (r *prontuarioRepository) GetByVeterinario(veterinarioID uint, inicio, fim string) ([]models.Prontuarios, error) {
	var prontuarios []models.Prontuarios

	err := r.db.
		Where("veterinario_id = ? AND DATE(data_consulta) BETWEEN ? AND ?",
			veterinarioID, inicio, fim).
		Preload("Pet").
		Preload("Pet.Clientes").
		Order("data_consulta DESC").
		Find(&prontuarios).Error

	return prontuarios, err
}

func (r *prontuarioRepository) GetUltimoProntuario(petID uint) (*models.Prontuarios, error) {
	var prontuario models.Prontuarios

	err := r.db.
		Where("pet_id = ?", petID).
		Preload("Veterinario").
		Order("data_consulta DESC").
		First(&prontuario).Error

	return &prontuario, err
}

func (r *prontuarioRepository) RegistrarVacina(vacina *models.Vacinas) error {
	return r.db.Create(vacina).Error
}

func (r *prontuarioRepository) GetVacinasPorPet(petID uint) ([]models.Vacinas, error) {
	var vacinas []models.Vacinas

	err := r.db.
		Where("pet_id = ?", petID).
		Preload("Veterinario").
		Order("data_aplicacao DESC").
		Find(&vacinas).Error

	return vacinas, err
}

func (r *prontuarioRepository) GetVacinasVencidas(empresaID uint) ([]models.Vacinas, error) {
	var vacinas []models.Vacinas

	// Subquery para buscar pets da empresa
	subquery := r.db.Model(&models.Pets{}).
		Select("pets.id").
		Joins("JOIN clientes ON pets.cliente_id = clientes.id").
		Where("clientes.empresa_id = ?", empresaID)

	err := r.db.
		Where("pet_id IN (?) AND data_proxima < NOW()", subquery).
		Preload("Pet").
		Preload("Pet.Clientes").
		Preload("Veterinario").
		Order("data_proxima ASC").
		Find(&vacinas).Error

	return vacinas, err
}
