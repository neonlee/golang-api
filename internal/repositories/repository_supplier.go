package repositories

import (
	"petApi/internal/models"
	"petApi/internal/requests"

	"gorm.io/gorm"
)

type FornecedorRepository interface {
	Create(fornecedor *models.Fornecedor) error
	GetByID(id uint) (*models.Fornecedor, error)
	Update(fornecedor *models.Fornecedor) error
	Delete(id uint) error
	ListByEmpresa(empresaID uint, filters requests.FornecedorFilter) ([]models.Fornecedor, error)
	Search(empresaID uint, termo string) ([]models.Fornecedor, error)
	GetTotalFornecedores(empresaID uint) (int64, error)
}
type supplierRepository struct {
	connection *gorm.DB
}

func NewSupplierRepository(connection *gorm.DB) FornecedorRepository {
	return &supplierRepository{connection: connection}
}
func (r *supplierRepository) Create(fornecedor *models.Fornecedor) error {
	return r.connection.Create(fornecedor).Error
}
func (r *supplierRepository) GetByID(id uint) (*models.Fornecedor, error) {
	var fornecedor models.Fornecedor
	err := r.connection.First(&fornecedor, id).Error
	return &fornecedor, err
}
func (r *supplierRepository) Update(fornecedor *models.Fornecedor) error {
	return r.connection.Save(fornecedor).Error
}
func (r *supplierRepository) Delete(id uint) error {
	return r.connection.Delete(&models.Fornecedor{}, id).Error
}
func (r *supplierRepository) ListByEmpresa(empresaID uint, filters requests.FornecedorFilter) ([]models.Fornecedor, error) {
	var fornecedores []models.Fornecedor
	query := r.connection.Where("empresa_id = ?", empresaID)
	if filters.Nome != "" {
		query = query.Where("nome ILIKE ?", "%"+filters.Nome+"%")
	}
	if filters.Cidade != "" {
		query = query.Where("cidade ILIKE ?", "%"+filters.Cidade+"%")
	}
	err := query.Find(&fornecedores).Error
	return fornecedores, err
}
func (r *supplierRepository) Search(empresaID uint, termo string) ([]models.Fornecedor, error) {
	var fornecedores []models.Fornecedor
	err := r.connection.
		Where("empresa_id = ? AND (nome ILIKE ? OR email ILIKE ? OR telefone ILIKE ?)",
			empresaID, "%"+termo+"%", "%"+termo+"%", "%"+termo+"%").
		Limit(10).
		Find(&fornecedores).Error
	return fornecedores, err
}
func (r *supplierRepository) GetTotalFornecedores(empresaID uint) (int64, error) {
	var count int64
	err := r.connection.Model(&models.Fornecedor{}).
		Where("empresa_id = ?", empresaID).
		Count(&count).Error
	return count, err
}
