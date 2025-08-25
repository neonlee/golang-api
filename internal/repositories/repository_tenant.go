package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type TenantRepository struct {
	connection *gorm.DB
}

func NewTenantRepository(connection *gorm.DB) TenantRepository {
	return TenantRepository{connection: connection}
}

func (r *TenantRepository) Create(user models.Tenant) (*models.Tenant, error) {

	err := r.connection.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *TenantRepository) GetTenants() (*[]models.Tenant, error) {
	var clientes []models.Tenant

	err := r.connection.Find(&clientes).Error
	if err != nil {
		return nil, err
	}

	return &clientes, nil
}

func (r *TenantRepository) GetTenant(id int) (*models.Tenant, error) {
	servico := &models.Tenant{}

	err := r.connection.Where("id = ?", id).First(&servico).Error
	if err != nil {
		return nil, err
	}

	return servico, nil
}

func (r *TenantRepository) UpdateTenants(id int, services models.Tenant) (*models.Tenant, error) {
	err := r.connection.Model(&models.Tenant{}).Where("id = ?", id).Updates(services).Error
	if err != nil {
		return nil, err
	}

	// Busca o cliente atualizado
	var updatedClient models.Tenant
	err = r.connection.First(&updatedClient, id).Error
	if err != nil {
		return nil, err
	}

	return &updatedClient, nil
}

func (r *TenantRepository) DeleteTenants(id int) (bool, error) {
	err := r.connection.Delete(&models.Pet{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
