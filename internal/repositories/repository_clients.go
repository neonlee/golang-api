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

func (r *ClientsRepository) Create(user models.Client) (*models.Client, error) {
	err := r.connection.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// query.Close()
	return &user, nil
}

func (r *ClientsRepository) GetClients() (*[]models.Client, error) {
	var clientes []models.Client

	err := r.connection.
		Preload("Pets").
		Find(&clientes).Error
	if err != nil {
		return nil, err
	}

	return &clientes, nil
}

func (r *ClientsRepository) GetClient(id int) (*models.Client, error) {
	var client models.Client
	err := r.connection.
		Preload("Pets").
		First(&client, id).Error

	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *ClientsRepository) UpdateClient(id int, cliente models.Client) (*models.Client, error) {
	// Aplica o update direto pelo ID
	err := r.connection.Model(&models.Client{}).Where("id = ?", id).Updates(cliente).Error
	if err != nil {
		return nil, err
	}

	// Busca o cliente atualizado
	var updatedClient models.Client
	err = r.connection.First(&updatedClient, id).Error
	if err != nil {
		return nil, err
	}

	return &updatedClient, nil
}

func (r *ClientsRepository) DeleteClient(id int) (bool, error) {
	err := r.connection.Delete(&models.Client{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
