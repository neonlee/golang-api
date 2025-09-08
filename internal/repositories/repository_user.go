package repositories

import (
	"errors"
	"petApi/internal/models"

	"gorm.io/gorm"
)

// UserRepository defines the interface for user repository methods.
type UserRepository interface {
	Create(user *models.Users) error
	FindByID(id int) (*models.Users, error)
	GetAll() (*[]models.Users, error)
	FindByEmail(email string) (*models.Users, error)
	Update(user models.Users, id int) (*models.Users, error)
	Delete(id uint) error
}

// userRepository implements UserRepository using GORM.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.Users) error {
	return r.db.Create(user).Error
}

// GetAll implements UserRepository.
func (r *userRepository) GetAll() (*[]models.Users, error) {
	var clientes []models.Users

	err := r.db.
		Find(&clientes).Error
	if err != nil {
		return nil, err
	}

	return &clientes, nil
}
func (r *userRepository) FindByID(id int) (*models.Users, error) {
	var user models.Users
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.Users, error) {
	var user models.Users
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user models.Users, id int) (*models.Users, error) {
	err := r.db.Model(&models.Supplier{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return nil, err
	}

	// Busca o cliente atualizado
	var model models.Users
	err = r.db.First(&model, id).Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (r *userRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Users{}, id)
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return result.Error
}
