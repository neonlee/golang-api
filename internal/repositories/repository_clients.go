package repositories

import (
	"petApi/internal/models"
	"petApi/internal/requests"

	"gorm.io/gorm"
)

type ClienteRepository interface {
	Create(cliente *models.Cliente) error
	GetByID(id uint) (*models.Cliente, error)
	Update(cliente *models.Cliente) error
	Delete(id uint) error
	ListByEmpresa(empresaID uint, filters requests.ClienteFilter) ([]models.Cliente, error)
	Search(empresaID uint, termo string) ([]models.Cliente, error)
	GetWithPets(id uint) (*models.Cliente, error)
	GetTotalClientes(empresaID uint) (int64, error)
	GetClientesNovos(empresaID uint, mes int, ano int) ([]models.Cliente, error)
}

type clienteRepository struct {
	db *gorm.DB
}

func NewClienteRepository(db *gorm.DB) ClienteRepository {
	return &clienteRepository{db: db}
}

func (r *clienteRepository) Create(cliente *models.Cliente) error {
	return r.db.Create(cliente).Error
}

func (r *clienteRepository) GetByID(id uint) (*models.Cliente, error) {
	var cliente models.Cliente
	err := r.db.Preload("Pets").First(&cliente, id).Error
	return &cliente, err
}

func (r *clienteRepository) Update(cliente *models.Cliente) error {
	return r.db.Save(cliente).Error
}

func (r *clienteRepository) Delete(id uint) error {
	return r.db.Delete(&models.Cliente{}, id).Error
}

func (r *clienteRepository) ListByEmpresa(empresaID uint, filters requests.ClienteFilter) ([]models.Cliente, error) {
	var clientes []models.Cliente

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

func (r *clienteRepository) Search(empresaID uint, termo string) ([]models.Cliente, error) {
	var clientes []models.Cliente

	err := r.db.
		Where("empresa_id = ? AND (nome ILIKE ? OR email ILIKE ? OR telefone ILIKE ?)",
			empresaID, "%"+termo+"%", "%"+termo+"%", "%"+termo+"%").
		Preload("Pets").
		Order("nome ASC").
		Limit(50).
		Find(&clientes).Error

	return clientes, err
}

func (r *clienteRepository) GetWithPets(id uint) (*models.Cliente, error) {
	var cliente models.Cliente
	err := r.db.
		Preload("Pets").
		First(&cliente, id).Error

	return &cliente, err
}

func (r *clienteRepository) GetTotalClientes(empresaID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Cliente{}).Where("empresa_id = ?", empresaID).Count(&count).Error
	return count, err
}

func (r *clienteRepository) GetClientesNovos(empresaID uint, mes int, ano int) ([]models.Cliente, error) {
	var clientes []models.Cliente

	err := r.db.
		Where("empresa_id = ? AND EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?",
			empresaID, mes, ano).
		Order("created_at DESC").
		Find(&clientes).Error

	return clientes, err
}
