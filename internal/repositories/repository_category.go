package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

// CategoriaRepository interface
type CategoriaRepository interface {
	Create(categoria *models.CategoriasProdutos) error
	GetByID(id uint) (*models.CategoriasProdutos, error)
	Update(categoria *models.CategoriasProdutos) error
	Delete(id uint) error
	ListByEmpresa(empresaID uint) ([]models.CategoriasProdutos, error)
	GetWithProdutos(id uint) (*models.CategoriasProdutos, error)
}

type categoryRepository struct {
	connection *gorm.DB
}

func NewCategoryRepository(connection *gorm.DB) CategoriaRepository {
	return &categoryRepository{connection: connection}
}

func (r *categoryRepository) Create(categoria *models.CategoriasProdutos) error {
	return r.connection.Create(categoria).Error
}

func (r *categoryRepository) GetByID(id uint) (*models.CategoriasProdutos, error) {
	var categoria models.CategoriasProdutos
	if err := r.connection.First(&categoria, id).Error; err != nil {
		return nil, err
	}
	return &categoria, nil
}

func (r *categoryRepository) Update(categoria *models.CategoriasProdutos) error {
	return r.connection.Save(categoria).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return r.connection.Delete(&models.CategoriasProdutos{}, id).Error
}

func (r *categoryRepository) ListByEmpresa(empresaID uint) ([]models.CategoriasProdutos, error) {
	var categorias []models.CategoriasProdutos
	if err := r.connection.Where("empresa_id = ?", empresaID).Find(&categorias).Error; err != nil {
		return nil, err
	}
	return categorias, nil
}

func (r *categoryRepository) GetWithProdutos(id uint) (*models.CategoriasProdutos, error) {
	var categoria models.CategoriasProdutos
	if err := r.connection.Preload("Produtos").First(&categoria, id).Error; err != nil {
		return nil, err
	}
	return &categoria, nil
}
