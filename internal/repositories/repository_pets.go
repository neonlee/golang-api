package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type PetsRepository struct {
	connection *gorm.DB
}

func NewPetsRepository(connection *gorm.DB) PetsRepository {
	return PetsRepository{connection: connection}
}

func (r *PetsRepository) GetPets() (*[]models.Pet, error) {

	var petsList []models.Pet

	err := r.connection.Find(&petsList).Error

	return &petsList, err
}

func (r *PetsRepository) GetPet(id int) (*models.Pet, error) {
	var pet models.Pet
	err := r.connection.Where("id = ?", id).First(&pet).Error

	return &pet, err
}
func (r *PetsRepository) Create(user models.Pet) (*models.Pet, error) {
	err := r.connection.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PetsRepository) DeletePet(id int) (bool, error) {
	err := r.connection.Delete(&models.Pet{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
func (r *PetsRepository) UpdatePet(id int, pet models.Pet) (*models.Pet, error) {
	err := r.connection.Model(&models.Pet{}).Where("id = ?", id).Updates(pet).Error
	if err != nil {
		return nil, err
	}

	// Busca o cliente atualizado
	var updatedClient models.Pet
	err = r.connection.First(&updatedClient, id).Error
	if err != nil {
		return nil, err
	}

	return &updatedClient, nil
}
