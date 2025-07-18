package repositories

import (
	"database/sql"
	"petApi/internal/models"

	"gorm.io/gorm"
)

type ClientsRepository struct {
	connection *gorm.DB
}

func NewClientsRepository(connection *gorm.DB) ClientsRepository {
	return ClientsRepository{connection: connection}
}

func (r *ClientsRepository) Create(user models.Clients) (*models.Clients, error) {
	err := r.connection.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// query.Close()
	return &user, nil
}

func (r *ClientsRepository) GetClients() (*[]models.Clients, error) {

	rows, err := r.connection.First("SELECT id, nome, email FROM clientes")
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
	user := &models.Clients{
		Pets: []models.Pet{},
	}

	query := `
		SELECT 
			c.id, c.petshop_id, c.nome, c.telefone, c.email, c.endereco,
			p.id, p.nome
		FROM clientes c
		LEFT JOIN pets p ON p.cliente_id = c.id
		WHERE c.id = $1
	`

	rows, err := r.connection.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	first := true

	for rows.Next() {
		var (
			clienteID int
			petshopID int
			nome      string
			telefone  string
			email     string
			endereco  string
			petID     sql.NullInt64
			petNome   sql.NullString
		)

		if err := rows.Scan(&clienteID, &petshopID, &nome, &telefone, &email, &endereco, &petID, &petNome); err != nil {
			return nil, err
		}

		if first {
			user.Id = clienteID
			user.PetshopId = petshopID
			user.Nome = nome
			user.Telefone = telefone
			user.Email = email
			user.Endereco = endereco
			first = false
		}

		if petID.Valid {
			user.Pets = append(user.Pets, models.Pet{
				IdPet:    int(petID.Int64),
				Name:     petNome.String,
				ClientId: clienteID,
			})
		}
	}

	return user, nil
}

func (r *ClientsRepository) UpdateClient(id int, cliente models.Clients) (*models.Clients, error) {
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
