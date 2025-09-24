package repositories

import (
	"petApi/internal/models"
	"petApi/internal/requests"

	"gorm.io/gorm"
)

// PetRepository interface
type PetRepository interface {
	Create(pet *models.Pet) error
	GetByID(id uint) (*models.Pet, error)
	Update(pet *models.Pet) error
	Delete(id uint) error
	GetByCliente(clienteID uint) ([]models.Pet, error)
	GetWithCliente(id uint) (*models.Pet, error)
	ListByEmpresa(empresaID uint, filters requests.PetFilter) ([]models.Pet, error)
	GetTotalPets(empresaID uint) (int64, error)
	GetPetsPorEspecie(empresaID uint) (map[string]int64, error)
}

// repositories/pet_repository.go

type petRepository struct {
	db *gorm.DB
}

func NewPetRepository(db *gorm.DB) PetRepository {
	return &petRepository{db: db}
}

func (r *petRepository) Create(pet *models.Pet) error {
	return r.db.Create(pet).Error
}

func (r *petRepository) GetByID(id uint) (*models.Pet, error) {
	var pet models.Pet
	err := r.db.Preload("Cliente").First(&pet, id).Error
	return &pet, err
}

func (r *petRepository) Update(pet *models.Pet) error {
	return r.db.Save(pet).Error
}

func (r *petRepository) Delete(id uint) error {
	return r.db.Delete(&models.Pet{}, id).Error
}

func (r *petRepository) GetByCliente(clienteID uint) ([]models.Pet, error) {
	var pets []models.Pet
	err := r.db.Where("cliente_id = ?", clienteID).Order("nome ASC").Find(&pets).Error
	return pets, err
}

func (r *petRepository) GetWithCliente(id uint) (*models.Pet, error) {
	var pet models.Pet
	err := r.db.Preload("Cliente").First(&pet, id).Error
	return &pet, err
}

func (r *petRepository) ListByEmpresa(empresaID uint, filters requests.PetFilter) ([]models.Pet, error) {
	var pets []models.Pet

	query := r.db.
		Joins("JOIN clientes ON pets.cliente_id = clientes.id").
		Where("clientes.empresa_id = ?", empresaID)

	if filters.Nome != "" {
		query = query.Where("pets.nome ILIKE ?", "%"+filters.Nome+"%")
	}

	if filters.Especie != "" {
		query = query.Where("pets.especie = ?", filters.Especie)
	}

	if filters.Raca != "" {
		query = query.Where("pets.raca ILIKE ?", "%"+filters.Raca+"%")
	}

	err := query.Preload("Cliente").Order("pets.nome ASC").Find(&pets).Error
	return pets, err
}

func (r *petRepository) GetTotalPets(empresaID uint) (int64, error) {
	var count int64
	err := r.db.
		Model(&models.Pet{}).
		Joins("JOIN clientes ON pets.cliente_id = clientes.id").
		Where("clientes.empresa_id = ?", empresaID).
		Count(&count).Error

	return count, err
}

func (r *petRepository) GetPetsPorEspecie(empresaID uint) (map[string]int64, error) {
	type Result struct {
		Especie string
		Count   int64
	}

	var results []Result

	err := r.db.
		Model(&models.Pet{}).
		Select("especie, COUNT(*) as count").
		Joins("JOIN clientes ON pets.cliente_id = clientes.id").
		Where("clientes.empresa_id = ?", empresaID).
		Group("especie").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	resultMap := make(map[string]int64)
	for _, result := range results {
		resultMap[result.Especie] = result.Count
	}

	return resultMap, nil
}
