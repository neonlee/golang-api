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
	query, err := r.connection.Prepare("INSERT INTO clientes (petshop_id, nome, descricao, preco, tipo)" +
		" VALUES ($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		return nil, err
	}

	err = query.QueryRow(user.PetshopId, user.Nome, user.Descricao, user.Preco, user.Tipo).Scan(
		&user.Id,
	)
	if err != nil {
		return nil, err
	}

	// query.Close()
	return &user, nil
}

func (r *ServicesRepository) GetServices() (*[]models.Services, error) {

	rows, err := r.connection.Query("SELECT * FROM servicos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []models.Services
	for rows.Next() {
		var cliente models.Services
		if err := rows.Scan(&cliente.Id, &cliente.Nome, &cliente.Descricao, &cliente.Preco); err != nil {
			return nil, err
		}
		clientes = append(clientes, cliente)
	}
	return &clientes, nil
}

func (r *ServicesRepository) GetService(id int) (*models.Services, error) {
	servico := &models.Services{}
	query := `Select * from services where id = $1`

	rows, err := r.connection.Prepare(query)
	if err != nil {
		return nil, err
	}

	err = rows.QueryRow(id).Scan(
		&servico.Id,
		&servico.PetshopId,
		&servico.Nome,
		&servico.Descricao,
		&servico.Preco,
		&servico.Tipo,
	)
	if err != nil {
		return nil, err
	}
	return servico, nil
}

func (r *ServicesRepository) UpdateServices(id int, cliente models.Services) (*models.Services, error) {
	query := "UPDATE clientes SET nome = $1, descricao = $2, preco = $3, tipo = $4 WHERE id = $5"
	rows, err := r.connection.Exec(query, cliente.Nome, cliente.Descricao, cliente.Preco, cliente.Tipo, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := rows.RowsAffected()

	if rowsAffected == 0 {
		return nil, err
	}
	client, err := r.GetService(id)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (r *ServicesRepository) DeleteServices(id int) (bool, error) {
	query := "DELETE FROM servicos WHERE id = $1"
	rows, err := r.connection.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, _ := rows.RowsAffected()

	if rowsAffected == 0 {
		return false, err
	}

	return true, nil
}
