package repositories

import (
	"fmt"
	"petApi/internal/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db: db}
}

func (r *ProductRepository) Create(product models.Produto) (*models.Produto, error) {
	if product.EstoqueAtual < 0 {
		return nil, fmt.Errorf("estoque inválido para produto %s: %d (deve ser >= 0)", product.Nome, product.EstoqueAtual)
	}
	if product.EstoqueMinimo < 0 {
		return nil, fmt.Errorf("quantidade minima inválido para produto %s: %d (deve ser >= 0)", product.Nome, product.EstoqueMinimo)
	}
	if product.EstoqueAtual < product.EstoqueMinimo {
		return nil, fmt.Errorf("quantidade em estoque deve ser maior que a quantidade minima")
	}
	var supplier models.Supplier
	var category models.Category
	if err := r.db.First(&supplier, product.FornecedorID).Error; err != nil {
		return nil, fmt.Errorf("supplier com ID %d não existe", product.FornecedorID)
	}

	if err := r.db.First(&category, product.CategoriaID).Error; err != nil {
		return nil, fmt.Errorf("categoria com ID %d não existe", product.CategoriaID)
	}

	err := r.db.Create(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) GetByID(id uint) (*models.Produto, error) {
	var product models.Produto
	err := r.db.Preload("Supplier").Preload("Category").First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) GetAll() ([]models.Produto, error) {
	var products []models.Produto
	err := r.db.Preload("Supplier").Preload("Category").Find(&products).Error
	return products, err
}

func (r *ProductRepository) Update(product models.Produto, id int) (*models.Produto, error) {
	err := r.db.Model(&models.Produto{}).Where("id = ?", id).Updates(product).Error
	if err != nil {
		return nil, err
	}

	// Busca o cliente atualizado
	var model models.Produto
	err = r.db.First(&model, id).Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Produto{}, id).Error
}
