package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type ServicesRepository struct {
	connection *gorm.DB
}

func NewServicesRepository(connection *gorm.DB) ServicesRepository {
	return ServicesRepository{connection: connection}
}

func (r *ServicesRepository) Create(user models.Services) (*models.Services, error) {
	err := r.connection.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *ServicesRepository) GetServices() (*[]models.Services, error) {
	var clientes []models.Services

	err := r.connection.Find(&clientes).Error
	if err != nil {
		return nil, err
	}

	return &clientes, nil
}

func (r *ServicesRepository) GetService(id int) (*models.Services, error) {
	servico := &models.Services{}

	err := r.connection.Where("id = ?", id).First(&servico).Error
	if err != nil {
		return nil, err
	}

	return servico, nil
}

func (r *ServicesRepository) UpdateServices(id int, services models.Services) (*models.Services, error) {
	err := r.connection.Model(&models.Services{}).Where("id = ?", id).Updates(services).Error
	if err != nil {
		return nil, err
	}

	// Busca o cliente atualizado
	var updatedClient models.Services
	err = r.connection.First(&updatedClient, id).Error
	if err != nil {
		return nil, err
	}

	return &updatedClient, nil
}

func (r *ServicesRepository) DeleteServices(id int) (bool, error) {
	err := r.connection.Delete(&models.Pet{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
