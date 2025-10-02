package repositories

import (
	"petApi/internal/models"
	"petApi/internal/requests"

	"gorm.io/gorm"
)

type FornecedoresRepository interface {
	Create(fornecedor *models.Fornecedores) error
	GetByID(id uint) (*models.Fornecedores, error)
	Update(fornecedor *models.Fornecedores) error
	Delete(id uint) error
	ListByEmpresa(empresaID uint, filters requests.FornecedorFilter) ([]models.Fornecedores, error)
	Search(empresaID uint, termo string) ([]models.Fornecedores, error)
	GetTotalFornecedores(empresaID uint) (int64, error)
}
type supplierRepository struct {
	connection *gorm.DB
}

func NewSupplierRepository(connection *gorm.DB) FornecedoresRepository {
	return &supplierRepository{connection: connection}
}
func (r *supplierRepository) Create(fornecedor *models.Fornecedores) error {
	return r.connection.Create(fornecedor).Error
}
func (r *supplierRepository) GetByID(id uint) (*models.Fornecedores, error) {
	var fornecedor models.Fornecedores
	err := r.connection.First(&fornecedor, id).Error
	return &fornecedor, err
}
func (r *supplierRepository) Update(fornecedor *models.Fornecedores) error {
	return r.connection.Save(fornecedor).Error
}
func (r *supplierRepository) Delete(id uint) error {
	return r.connection.Delete(&models.Fornecedores{}, id).Error
}
func (r *supplierRepository) ListByEmpresa(empresaID uint, filters requests.FornecedorFilter) ([]models.Fornecedores, error) {
	var fornecedores []models.Fornecedores
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
func (r *supplierRepository) Search(empresaID uint, termo string) ([]models.Fornecedores, error) {
	var fornecedores []models.Fornecedores
	err := r.connection.
		Where("empresa_id = ? AND (nome ILIKE ? OR email ILIKE ? OR telefone ILIKE ?)",
			empresaID, "%"+termo+"%", "%"+termo+"%", "%"+termo+"%").
		Limit(10).
		Find(&fornecedores).Error
	return fornecedores, err
}
func (r *supplierRepository) GetTotalFornecedores(empresaID uint) (int64, error) {
	var count int64
	err := r.connection.Model(&models.Fornecedores{}).
		Where("empresa_id = ?", empresaID).
		Count(&count).Error
	return count, err
}
