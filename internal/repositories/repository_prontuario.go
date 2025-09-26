package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type ProntuarioRepository interface {
	Create(prontuario *models.Prontuario) error
	GetByID(id uint) (*models.Prontuario, error)
	Update(prontuario *models.Prontuario) error
	GetByPet(petID uint) ([]models.Prontuario, error)
	GetByVeterinario(veterinarioID uint, inicio, fim string) ([]models.Prontuario, error)
	GetUltimoProntuario(petID uint) (*models.Prontuario, error)
	RegistrarVacina(vacina *models.Vacina) error
	GetVacinasPorPet(petID uint) ([]models.Vacina, error)
	GetVacinasVencidas(empresaID uint) ([]models.Vacina, error)
}
type prontuarioRepository struct {
	db *gorm.DB
}

func NewProntuarioRepository(db *gorm.DB) ProntuarioRepository {
	return &prontuarioRepository{db: db}
}

func (r *prontuarioRepository) Create(prontuario *models.Prontuario) error {
	return r.db.Create(prontuario).Error
}

func (r *prontuarioRepository) GetByID(id uint) (*models.Prontuario, error) {
	var prontuario models.Prontuario
	err := r.db.
		Preload("Pet").
		Preload("Pet.Clientes").
		Preload("Veterinario").
		First(&prontuario, id).Error

	return &prontuario, err
}

func (r *prontuarioRepository) Update(prontuario *models.Prontuario) error {
	return r.db.Save(prontuario).Error
}

func (r *prontuarioRepository) GetByPet(petID uint) ([]models.Prontuario, error) {
	var prontuarios []models.Prontuario

	err := r.db.
		Where("pet_id = ?", petID).
		Preload("Veterinario").
		Order("data_consulta DESC").
		Find(&prontuarios).Error

	return prontuarios, err
}

func (r *prontuarioRepository) GetByVeterinario(veterinarioID uint, inicio, fim string) ([]models.Prontuario, error) {
	var prontuarios []models.Prontuario

	err := r.db.
		Where("veterinario_id = ? AND DATE(data_consulta) BETWEEN ? AND ?",
			veterinarioID, inicio, fim).
		Preload("Pet").
		Preload("Pet.Clientes").
		Order("data_consulta DESC").
		Find(&prontuarios).Error

	return prontuarios, err
}

func (r *prontuarioRepository) GetUltimoProntuario(petID uint) (*models.Prontuario, error) {
	var prontuario models.Prontuario

	err := r.db.
		Where("pet_id = ?", petID).
		Preload("Veterinario").
		Order("data_consulta DESC").
		First(&prontuario).Error

	return &prontuario, err
}

func (r *prontuarioRepository) RegistrarVacina(vacina *models.Vacina) error {
	return r.db.Create(vacina).Error
}

func (r *prontuarioRepository) GetVacinasPorPet(petID uint) ([]models.Vacina, error) {
	var vacinas []models.Vacina

	err := r.db.
		Where("pet_id = ?", petID).
		Preload("Veterinario").
		Order("data_aplicacao DESC").
		Find(&vacinas).Error

	return vacinas, err
}

func (r *prontuarioRepository) GetVacinasVencidas(empresaID uint) ([]models.Vacina, error) {
	var vacinas []models.Vacina

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
