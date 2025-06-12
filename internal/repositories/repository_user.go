package repositories

import (
	"database/sql"
	"petApi/internal/models"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
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
func (r *UserRepository) Create(user models.User) (*models.User, error) {
	users := &models.User{}
	query, err := r.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		" VALUES ($1, $2) RETURNING id")
	if err != nil {
		return nil, err
	}

	err = query.QueryRow(user.Name, user.Email, user.Name, user.Password, user.Role).Scan(
		&users.ID,
		&users.Name,
		&users.Role,
	)
	if err != nil {
		return nil, err
	}

	query.Close()
	return users, nil
}

func (repository *UserRepository) GetByID(id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, email, role, created_at, updated_at FROM users WHERE id = $1`
	rows, err := repository.connection.Prepare(query)
	if err != nil {
		return nil, err
	}

	err = rows.QueryRow(id).Scan(
		&user.ID,
		&user.Name,
		&user.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	rows.Close()
	return user, nil

}
