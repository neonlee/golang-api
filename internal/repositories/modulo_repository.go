package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type ModuloRepository interface {
	GetByID(id uint) (*models.Modulo, error)
	ListAll() ([]models.Modulo, error)
	ListByCategoria(categoria string) ([]models.Modulo, error)
	GetByNome(nome string) (*models.Modulo, error)
}

type moduloRepository struct {
	db *gorm.DB // Assuming you're using GORM for database operations
}

func NewModuloRepository(db *gorm.DB) ModuloRepository {
	return &moduloRepository{db: db}
}

func (r *moduloRepository) GetByID(id uint) (*models.Modulo, error) {
	var modulo models.Modulo
	if err := r.db.First(&modulo, id).Error; err != nil {
		return nil, err
	}
	return &modulo, nil
}

func (r *moduloRepository) ListAll() ([]models.Modulo, error) {
	var modulos []models.Modulo
	if err := r.db.Find(&modulos).Error; err != nil {
		return nil, err
	}
	return modulos, nil
}

func (r *moduloRepository) ListByCategoria(categoria string) ([]models.Modulo, error) {
	var modulos []models.Modulo
	if err := r.db.Where("categoria = ?", categoria).Find(&modulos).Error; err != nil {
		return nil, err
	}
	return modulos, nil
}

func (r *moduloRepository) GetByNome(nome string) (*models.Modulo, error) {
	var modulo models.Modulo
	if err := r.db.Where("nome = ?", nome).First(&modulo).Error; err != nil {
		return nil, err
	}
	return &modulo, nil
}
