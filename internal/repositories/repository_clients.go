package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type ClientsRepository struct {
	connection *gorm.DB
}

func NewClientsRepository(connection *gorm.DB) ClientsRepository {
	return ClientsRepository{connection: connection}
}

func (r *ClientsRepository) Create(user models.Cliente) (*models.Cliente, error) {
	err := r.connection.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// query.Close()
	return &user, nil
}

func (r *ClientsRepository) GetClients() (*[]models.Cliente, error) {
	var clientes []models.Cliente

	err := r.connection.
		Preload("Pets").
		Find(&clientes).Error
	if err != nil {
		return nil, err
	}

	return &clientes, nil
}

func (r *ClientsRepository) GetClient(id int) (*models.Cliente, error) {
	var client models.Cliente
	err := r.connection.
		Preload("Pets").
		First(&client, id).Error

	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *ClientsRepository) UpdateClient(id int, cliente models.Cliente) (*models.Cliente, error) {
	// Aplica o update direto pelo ID
	err := r.connection.Model(&models.Cliente{}).Where("id = ?", id).Updates(cliente).Error
	if err != nil {
		return nil, err
	}

	// Busca o cliente atualizado
	var updatedClient models.Cliente
	err = r.connection.First(&updatedClient, id).Error
	if err != nil {
		return nil, err
	}

	return &updatedClient, nil
}

func (r *ClientsRepository) DeleteClient(id int) (bool, error) {
	err := r.connection.Delete(&models.Cliente{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
