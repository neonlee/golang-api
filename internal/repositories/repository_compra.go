package repositories

import (
	"petApi/internal/models"
	"petApi/internal/requests"
	"petApi/internal/responses"
	"strconv"

	"gorm.io/gorm"
)

type CompraRepository interface {
	Create(compra *models.Compra, itens []models.CompraItem) error
	GetByID(id uint) (*models.Compra, error)
	Update(compra *models.Compra) error
	Cancelar(id uint, motivo string) error
	ListByEmpresa(empresaID uint, filters requests.CompraFilter) ([]models.Compra, error)
	ListByFornecedor(fornecedorID uint) ([]models.Compra, error)
	GetComprasPorPeriodo(empresaID uint, inicio, fim string) ([]models.Compra, error)
	GetResumoCompras(empresaID uint, periodo string) (*responses.ResumoCompras, error)
}

type compraRepository struct {
	db *gorm.DB
}

func NewCompraRepository(db *gorm.DB) CompraRepository {
	return &compraRepository{db: db}
}

func (r *compraRepository) Create(compra *models.Compra, itens []models.CompraItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Criar a compra
		if err := tx.Create(compra).Error; err != nil {
			return err
		}

		// Criar os itens da compra
		for i := range itens {
			itens[i].CompraID = compra.ID
			if err := tx.Create(&itens[i]).Error; err != nil {
				return err
			}

			// Atualizar estoque
			movimentacao := models.MovimentacaoEstoque{
				ProdutoID:        itens[i].ProdutoID,
				TipoMovimentacao: "entrada",
				Quantidade:       itens[i].Quantidade,
				Motivo:           "Compra #" + strconv.FormatUint(uint64(compra.ID), 10),
				UsuarioID:        compra.UsuarioID,
			}
			if err := tx.Create(&movimentacao).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *compraRepository) GetByID(id uint) (*models.Compra, error) {
	var compra models.Compra
	err := r.db.
		Preload("Fornecedor").
		Preload("Usuario").
		Preload("Itens").
		Preload("Itens.Produto").
		First(&compra, id).Error

	return &compra, err
}

func (r *compraRepository) Update(compra *models.Compra) error {
	return r.db.Save(compra).Error
}

func (r *compraRepository) Cancelar(id uint, motivo string) error {
	return r.db.Model(&models.Compra{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":      "cancelado",
			"observacoes": motivo,
		}).Error
}

func (r *compraRepository) ListByEmpresa(empresaID uint, filters requests.CompraFilter) ([]models.Compra, error) {
	var compras []models.Compra

	query := r.db.Where("empresa_id = ?", empresaID)

	if filters.FornecedorID != nil {
		query = query.Where("fornecedor_id = ?", *filters.FornecedorID)
	}

	if filters.DataInicio != "" && filters.DataFim != "" {
		query = query.Where("DATE(data_compra) BETWEEN ? AND ?", filters.DataInicio, filters.DataFim)
	}

	if filters.Status != "" {
		query = query.Where("status = ?", filters.Status)
	}

	err := query.
		Preload("Fornecedor").
		Preload("Usuario").
		Order("data_compra DESC").
		Find(&compras).Error

	return compras, err
}

func (r *compraRepository) ListByFornecedor(fornecedorID uint) ([]models.Compra, error) {
	var compras []models.Compra

	err := r.db.
		Where("fornecedor_id = ?", fornecedorID).
		Preload("Fornecedor").
		Preload("Usuario").
		Order("data_compra DESC").
		Find(&compras).Error

	return compras, err
}

func (r *compraRepository) GetComprasPorPeriodo(empresaID uint, inicio, fim string) ([]models.Compra, error) {
	var compras []models.Compra

	err := r.db.
		Where("empresa_id = ? AND DATE(data_compra) BETWEEN ? AND ?", empresaID, inicio, fim).
		Preload("Fornecedor").
		Preload("Usuario").
		Preload("Itens").
		Order("data_compra DESC").
		Find(&compras).Error

	return compras, err
}

func (r *compraRepository) GetResumoCompras(empresaID uint, periodo string) (*responses.ResumoCompras, error) {
	var resumo responses.ResumoCompras

	// Implementar lógica de resumo por período
	// Exemplo: total de compras, valor total, etc.

	return &resumo, nil
}
