package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	connection *gorm.DB
}

func NewCategoryRepository(connection *gorm.DB) CategoryRepository {
	return CategoryRepository{connection: connection}
}

func (r *CategoryRepository) Create(user models.Category) (*models.Category, error) {

	err := r.connection.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *CategoryRepository) GetCategorys() (*[]models.Category, error) {
	var clientes []models.Category

	err := r.connection.Find(&clientes).Error
	if err != nil {
		return nil, err
	}

	return &clientes, nil
}

func (r *CategoryRepository) GetCategory(id int) (*models.Category, error) {
	servico := &models.Category{}

	err := r.connection.Where("id = ?", id).First(&servico).Error
	if err != nil {
		return nil, err
	}

	return servico, nil
}

func (r *CategoryRepository) UpdateCategorys(id int, services models.Category) (*models.Category, error) {
	err := r.connection.Model(&models.Category{}).Where("id = ?", id).Updates(services).Error
	if err != nil {
		return nil, err
	}

	// Busca o cliente atualizado
	var updatedClient models.Category
	err = r.connection.First(&updatedClient, id).Error
	if err != nil {
		return nil, err
	}

	return &updatedClient, nil
}

func (r *CategoryRepository) DeleteCategorys(id int) (bool, error) {
	err := r.connection.Delete(&models.Pet{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
