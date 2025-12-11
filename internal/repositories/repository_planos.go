package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

// PerfilRepository interface
type PlanoRepository interface {
	Create(plano *models.Planos) error
	GetByID(id uint) (*models.Planos, error)
	Update(plano *models.Planos) error
	ListAll() ([]models.Planos, error)
}

type planoRepository struct {
	db *gorm.DB
}

func NewPlanosRepository(db *gorm.DB) PlanoRepository {
	return &planoRepository{db: db}
}

func (r *planoRepository) Create(plano *models.Planos) error {
	return r.db.Create(plano).Error
}

func (r *planoRepository) GetByID(id uint) (*models.Planos, error) {
	var plano models.Planos
	err := r.db.First(&plano, id).Error
	return &plano, err
}

func (r *planoRepository) Update(plano *models.Planos) error {
	return r.db.Save(plano).Error
}

func (r *planoRepository) ListAll() ([]models.Planos, error) {
	var planos []models.Planos
	err := r.db.Find(&planos).Error
	return planos, err
}
