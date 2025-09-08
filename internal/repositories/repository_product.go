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

func (r *ProductRepository) Create(product models.Product) (*models.Product, error) {
	if product.QuantidadeEstoque < 0 {
		return nil, fmt.Errorf("estoque inválido para produto %s: %d (deve ser >= 0)", product.Nome, product.QuantidadeEstoque)
	}
	if product.QuantidadeMinima < 0 {
		return nil, fmt.Errorf("quantidade minima inválido para produto %s: %d (deve ser >= 0)", product.Nome, product.QuantidadeMinima)
	}
	if product.QuantidadeEstoque < product.QuantidadeMinima {
		return nil, fmt.Errorf("quantidade em estoque deve ser maior que a quantidade minima")
	}
	var supplier models.Supplier
	var category models.Category
	if err := r.db.First(&supplier, product.SupplierID).Error; err != nil {
		return nil, fmt.Errorf("supplier com ID %d não existe", product.SupplierID)
	}

	if err := r.db.First(&category, product.CategoryID).Error; err != nil {
		return nil, fmt.Errorf("categoria com ID %d não existe", product.CategoryID)
	}

	err := r.db.Create(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Supplier").Preload("Category").First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Supplier").Preload("Category").Find(&products).Error
	return products, err
}

func (r *ProductRepository) Update(product models.Product, id int) (*models.Product, error) {
	err := r.db.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
	if err != nil {
		return nil, err
	}

	// Busca o cliente atualizado
	var model models.Product
	err = r.db.First(&model, id).Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
