package repositories

import (
	"database/sql"
	"petApi/internal/models"
)

type ClientsRepository struct {
	connection *sql.DB
}

func NewClientsRepository(connection *sql.DB) ClientsRepository {
	return ClientsRepository{connection: connection}
}

func (r *ClientsRepository) Create(user models.Clients) (*models.Clients, error) {
	query, err := r.connection.Prepare("INSERT INTO clientes (petshop_id, nome, telefone, email, endereco)" +
		" VALUES ($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		return nil, err
	}

	err = query.QueryRow(user.PetshopId, user.Nome, user.Telefone, user.Email, user.Endereco).Scan(
		&user.Id,
	)
	if err != nil {
		return nil, err
	}

	// query.Close()
	return &user, nil
}

func (r *ClientsRepository) GetClients() (*[]models.Clients, error) {

	rows, err := r.connection.Query("SELECT id, nome, email FROM clientes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []models.Clients
	for rows.Next() {
		var cliente models.Clients
		if err := rows.Scan(&cliente.Id, &cliente.Nome, &cliente.Email); err != nil {
			return nil, err
		}
		clientes = append(clientes, cliente)
	}
	return &clientes, nil
}

func (r *ClientsRepository) GetClient(id int) (*models.Clients, error) {
	user := &models.Clients{}
	query := `SELECT * FROM clientes WHERE id = $1`
	rows, err := r.connection.Prepare(query)
	if err != nil {
		return nil, err
	}

	err = rows.QueryRow(id).Scan(
		&user.PetshopId,
		&user.Id,
		&user.Nome,
		&user.Telefone,
		&user.Email,
		&user.Endereco,
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

func (r *ClientsRepository) UpdateClient(id int, cliente models.Clients) (*models.Clients, error) {
	query := "UPDATE clientes SET nome = $1, email = $2, endereco = $3, telefone = $4 WHERE id = $5"
	rows, err := r.connection.Exec(query, cliente.Nome, cliente.Email, cliente.Endereco, cliente.Telefone, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := rows.RowsAffected()

	if rowsAffected == 0 {
		return nil, err
	}
	client, err := r.GetClient(id)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (r *ClientsRepository) DeleteClient(id int) (bool, error) {
	query := "DELETE FROM clientes WHERE id = $1"
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
