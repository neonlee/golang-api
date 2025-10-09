package repositories

import (
	"petApi/internal/models"
	"petApi/internal/requests"

	"gorm.io/gorm"
)

type ClientesRepository interface {
	Create(cliente *models.Clientes) error
	GetByID(id uint) (*models.Clientes, error)
	Update(cliente *models.Clientes) error
	Delete(id uint) error
	ListByEmpresa(empresaID uint, filters requests.ClientesFilter) ([]models.Clientes, error)
	Search(empresaID uint, termo string) ([]models.Clientes, error)
	GetTotalClientes(empresaID uint) (int64, error)
	GetClientesNovos(empresaID uint, mes int, ano int) ([]models.Clientes, error)
}

type clienteRepository struct {
	db *gorm.DB
}

func NewClientesRepository(db *gorm.DB) ClientesRepository {
	return &clienteRepository{db: db}
}

func (r *clienteRepository) Create(cliente *models.Clientes) error {
	return r.db.Create(cliente).Error
}

func (r *clienteRepository) GetByID(id uint) (*models.Clientes, error) {
	var cliente models.Clientes
	err := r.db.Preload("Pets").Preload("Vendas").First(&cliente, id).Error
	return &cliente, err
}

func (r *clienteRepository) Update(cliente *models.Clientes) error {
	return r.db.Save(cliente).Error
}

func (r *clienteRepository) Delete(id uint) error {
	return r.db.Delete(&models.Clientes{}, id).Error
}

func (r *clienteRepository) ListByEmpresa(empresaID uint, filters requests.ClientesFilter) ([]models.Clientes, error) {
	var clientes []models.Clientes

	query := r.db.Where("empresa_id = ?", empresaID)

	if filters.Nome != "" {
		query = query.Where("nome ILIKE ?", "%"+filters.Nome+"%")
	}

	if filters.Email != "" {
		query = query.Where("email ILIKE ?", "%"+filters.Email+"%")
	}

	if filters.Ativo != nil {
		query = query.Where("ativo = ?", *filters.Ativo)
	}

	err := query.Preload("Pets").Order("nome ASC").Find(&clientes).Error
	return clientes, err
}

func (r *clienteRepository) Search(empresaID uint, termo string) ([]models.Clientes, error) {
	var clientes []models.Clientes

	err := r.db.
		Where("empresa_id = ? AND (nome ILIKE ? OR email ILIKE ? OR telefone ILIKE ?)",
			empresaID, "%"+termo+"%", "%"+termo+"%", "%"+termo+"%").
		Preload("Pets").
		Order("nome ASC").
		Limit(50).
		Find(&clientes).Error

	return clientes, err
}

func (r *clienteRepository) GetTotalClientes(empresaID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Clientes{}).Where("empresa_id = ?", empresaID).Count(&count).Error
	return count, err
}

func (r *clienteRepository) GetClientesNovos(empresaID uint, mes int, ano int) ([]models.Clientes, error) {
	var clientes []models.Clientes

	err := r.db.
		Where("empresa_id = ? AND EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?",
			empresaID, mes, ano).
		Order("created_at DESC").
		Find(&clientes).Error

	return clientes, err
}
