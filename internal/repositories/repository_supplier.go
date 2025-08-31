package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type SupplierRepository struct {
	connection *gorm.DB
}

func NewSupplierRepository(connection *gorm.DB) SupplierRepository {
	return SupplierRepository{connection: connection}
}

func (r *SupplierRepository) Create(user models.Supplier) (*models.Supplier, error) {
	err := r.connection.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// query.Close()
	return &user, nil
}

func (r *SupplierRepository) GetSuppliers() (*[]models.Supplier, error) {
	var clientes []models.Supplier

	err := r.connection.
		Preload("Products").
		Find(&clientes).Error
	if err != nil {
		return nil, err
	}

	return &clientes, nil
}

func (r *SupplierRepository) GetSupplier(id int) (*models.Supplier, error) {
	var client models.Supplier
	err := r.connection.
		Preload("Products").
		First(&client, id).Error

	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *SupplierRepository) UpdateSupplier(id int, cliente models.Supplier) (*models.Supplier, error) {
	// Aplica o update direto pelo ID
	err := r.connection.Model(&models.Supplier{}).Where("id = ?", id).Updates(cliente).Error
	if err != nil {
		return nil, err
	}

	// Busca o cliente atualizado
	var updatedSupplier models.Supplier
	err = r.connection.First(&updatedSupplier, id).Error
	if err != nil {
		return nil, err
	}

	return &updatedSupplier, nil
}

func (r *SupplierRepository) DeleteSupplier(id int) (bool, error) {
	err := r.connection.Delete(&models.Supplier{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
