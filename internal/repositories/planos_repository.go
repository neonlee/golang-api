package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

// PerfilRepository interface
type PlanoRepository interface {
	Create(plano *models.Plano) error
	GetByID(id uint) (*models.Plano, error)
	Update(plano *models.Plano) error
	ListAll() ([]models.Plano, error)
	GetByNome(nome string) (*models.Plano, error)
}

type planoRepository struct {
	db *gorm.DB
}

func NewPlanoRepository(db *gorm.DB) PlanoRepository {
	return &planoRepository{db: db}
}

func (r *planoRepository) Create(plano *models.Plano) error {
	return r.db.Create(plano).Error
}

func (r *planoRepository) GetByID(id uint) (*models.Plano, error) {
	var plano models.Plano
	err := r.db.First(&plano, id).Error
	return &plano, err
}

func (r *planoRepository) Update(plano *models.Plano) error {
	return r.db.Save(plano).Error
}

func (r *planoRepository) ListAll() ([]models.Plano, error) {
	var planos []models.Plano
	err := r.db.Find(&planos).Error
	return planos, err
}
func (r *planoRepository) GetByNome(nome string) (*models.Plano, error) {
	var plano models.Plano
	err := r.db.Where("nome = ?", nome).First(&plano).Error
	return &plano, err
}
