package repositories

import (
	"database/sql"
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

	rows, err := r.connection.Query("SELECT * FROM pets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var petsList []models.Pet
	for rows.Next() {
		var pets models.Pet
		if err := rows.Scan(&pets.IdPet,
			&pets.ClientId,
			&pets.Name,
			&pets.Race,
			&pets.Port,
			&pets.Age,
			&pets.Weight); err != nil {
			return nil, err
		}
		petsList = append(petsList, pets)
	}
	return &petsList, nil
}

func (r *PetsRepository) GetPet(id int) (*models.Pet, error) {
	pet := &models.Pet{}
	query := `SELECT * FROM pets WHERE id = $1`
	rows, err := r.connection.Prepare(query)
	if err != nil {
		return nil, err
	}

	err = rows.QueryRow(id).Scan(
		&pet.IdPet,
		&pet.ClientId,
		&pet.Name,
		&pet.Race,
		&pet.Port,
		&pet.Age,
		&pet.Weight,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	rows.Close()
	return pet, nil
}
func (r *PetsRepository) Create(user models.Pet) (*models.Pet, error) {
	query, err := r.connection.Prepare("INSERT INTO pets (cliente_id, nome, especie, raca, idade, peso)" +
		" VALUES ($1, $2, $3, $4, $5, $6) RETURNING id")
	if err != nil {
		return nil, err
	}

	err = query.QueryRow(user.ClientId, user.Name, user.Specie, user.Race, user.Age, user.Weight).Scan(
		&user.IdPet,
	)
	if err != nil {
		return nil, err
	}

	query.Close()
	return &user, nil
}

func (r *PetsRepository) DeletePet(id int) (bool, error) {
	query := "DELETE FROM pets WHERE id = $1"
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
func (r *PetsRepository) UpdatePet(id int, pet models.Pet) (*models.Pet, error) {
	query := "UPDATE pets SET nome = $1, idade = $2, raca = $3, peso = $4, especie = $5 WHERE id = $6"
	rows, err := r.connection.Exec(query, pet.Name, pet.Age, pet.Race, pet.Weight, pet.Specie, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := rows.RowsAffected()

	if rowsAffected == 0 {
		return nil, err
	}
	pets, err := r.GetPet(id)

	if err != nil {
		return nil, err
	}

	return pets, nil
}
