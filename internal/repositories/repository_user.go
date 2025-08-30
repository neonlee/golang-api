package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepository {
	return UserRepository{connection: connection}
}

// func (repository *UserRepository) Update(user *models.User) (models.User, error) {
// 	query := "SELECT id, product_name, price FROM product"
// 	rows, err := repository.connection.Query(query)
// 	if err != nil {
// 		return models.User{}, err
// 	}

// 	var productList []models.User
// 	var productObj models.User
// }

// func (r *UserRepository) Delete(id int) error {
// 	query := `DELETE FROM users WHERE id = $1`
// 	_, err := r.db.Exec(query, id)
// 	return err
// }

// Adicione role nas consultas SQL
func (r *UserRepository) Create(user models.Users) (*models.Users, error) {
	users := &models.Users{}

	return users, nil
}

func (repository *UserRepository) GetByID(id int) (*models.Users, error) {
	user := &models.Users{}

	return user, nil

}
