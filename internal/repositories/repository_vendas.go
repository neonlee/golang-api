// repositories/venda_repository.go
package repositories

import (
	"petApi/internal/models"
	"petApi/internal/requests"
	"petApi/internal/responses"
	"time"

	"gorm.io/gorm"
)

type VendaRepository interface {
	Create(venda *models.Venda, itens []models.VendaItem) error
	GetByID(id uint) (*models.Venda, error)
	UpdateStatus(id uint, status string) error
	CancelarVenda(id uint, motivo string) error
	ListByEmpresa(empresaID uint, filters requests.VendaFilter) ([]models.Venda, error)
	GetVendasDoDia(empresaID uint, data string) ([]models.Venda, error)
	GetVendasPorPeriodo(empresaID uint, inicio, fim string) ([]models.Venda, error)
	GetResumoVendas(empresaID uint, periodo string) (*responses.ResumoVendas, error)
	GetProdutosMaisVendidos(empresaID uint, limite int) ([]responses.ProdutoVendas, error)
	GetVendasPorFormaPagamento(empresaID uint, inicio, fim string) (map[string]float64, error)
}

type vendaRepository struct {
	db *gorm.DB
}

func NewVendaRepository(db *gorm.DB) VendaRepository {
	return &vendaRepository{db: db}
}

func (r *vendaRepository) Create(venda *models.Venda, itens []models.VendaItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Criar a venda
		if err := tx.Create(venda).Error; err != nil {
			return err
		}

		// Criar os itens da venda
		for i := range itens {
			itens[i].VendaID = venda.ID
			if err := tx.Create(&itens[i]).Error; err != nil {
				return err
			}

			// Se for produto, atualizar estoque
			if itens[i].ProdutoID != nil {
				movimentacao := models.MovimentacaoEstoque{
					ProdutoID:        *itens[i].ProdutoID,
					TipoMovimentacao: "saida",
					Quantidade:       itens[i].Quantidade,
					Motivo:           "Venda #" + string(venda.ID),
					UsuarioID:        venda.UsuarioID,
				}
				if err := tx.Create(&movimentacao).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func (r *vendaRepository) GetByID(id uint) (*models.Venda, error) {
	var venda models.Venda
	err := r.db.
		Preload("Cliente").
		Preload("Usuario").
		Preload("Itens").
		Preload("Itens.Produto").
		Preload("Itens.TipoServico").
		First(&venda, id).Error

	return &venda, err
}

func (r *vendaRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&models.Venda{}).Where("id = ?", id).Update("status", status).Error
}

func (r *vendaRepository) CancelarVenda(id uint, motivo string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Buscar venda e itens
		var venda models.Venda
		if err := tx.Preload("Itens").First(&venda, id).Error; err != nil {
			return err
		}

		// Reverter estoque dos produtos
		for _, item := range venda.Itens {
			if item.ProdutoID != nil {
				movimentacao := models.MovimentacaoEstoque{
					ProdutoID:        *item.ProdutoID,
					TipoMovimentacao: "entrada",
					Quantidade:       item.Quantidade,
					Motivo:           "Cancelamento venda #" + string(venda.ID) + " - " + motivo,
					UsuarioID:        venda.UsuarioID,
				}
				if err := tx.Create(&movimentacao).Error; err != nil {
					return err
				}
			}
		}

		// Atualizar status da venda
		return tx.Model(&venda).Updates(map[string]interface{}{
			"status": "cancelado",
		}).Error
	})
}

func (r *vendaRepository) ListByEmpresa(empresaID uint, filters requests.VendaFilter) ([]models.Venda, error) {
	var vendas []models.Venda

	query := r.db.Where("empresa_id = ?", empresaID)

	if filters.ClienteID != nil {
		query = query.Where("cliente_id = ?", *filters.ClienteID)
	}

	if filters.DataInicio != "" && filters.DataFim != "" {
		query = query.Where("DATE(data_venda) BETWEEN ? AND ?", filters.DataInicio, filters.DataFim)
	}

	if filters.Status != "" {
		query = query.Where("status = ?", filters.Status)
	}

	if filters.UsuarioID != nil {
		query = query.Where("usuario_id = ?", *filters.UsuarioID)
	}

	err := query.
		Preload("Cliente").
		Preload("Usuario").
		Order("data_venda DESC").
		Find(&vendas).Error

	return vendas, err
}

func (r *vendaRepository) GetVendasDoDia(empresaID uint, data string) ([]models.Venda, error) {
	var vendas []models.Venda

	err := r.db.
		Where("empresa_id = ? AND DATE(data_venda) = ?", empresaID, data).
		Preload("Cliente").
		Preload("Usuario").
		Preload("Itens").
		Order("data_venda ASC").
		Find(&vendas).Error

	return vendas, err
}

func (r *vendaRepository) GetResumoVendas(empresaID uint, periodo string) (*responses.ResumoVendas, error) {
	var resumo responses.ResumoVendas

	now := time.Now()
	var startDate time.Time

	switch periodo {
	case "hoje":
		startDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	case "semana":
		startDate = now.AddDate(0, 0, -7)
	case "mes":
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	default:
		startDate = now.AddDate(0, 0, -30)
	}

	// Total de vendas no período
	err := r.db.
		Model(&models.Venda{}).
		Where("empresa_id = ? AND data_venda >= ? AND status = ?", empresaID, startDate, "pago").
		Select("COUNT(*) as quantidade_mes, COALESCE(SUM(valor_final), 0) as total_mes").
		Scan(&resumo).Error

	if err != nil {
		return nil, err
	}

	// Calcular totais da semana e hoje
	var totalSemana, totalHoje float64
	r.db.Model(&models.Venda{}).
		Where("empresa_id = ? AND data_venda >= ? AND status = ?",
			empresaID, now.AddDate(0, 0, -7), "pago").
		Select("COALESCE(SUM(valor_final), 0)").
		Scan(&totalSemana)

	r.db.Model(&models.Venda{}).
		Where("empresa_id = ? AND DATE(data_venda) = ? AND status = ?",
			empresaID, now.Format("2006-01-02"), "pago").
		Select("COALESCE(SUM(valor_final), 0)").
		Scan(&totalHoje)

	resumo.TotalSemana = totalSemana
	resumo.TotalHoje = totalHoje

	return &resumo, nil
}

func (r *vendaRepository) GetProdutosMaisVendidos(empresaID uint, limite int) ([]responses.ProdutoVendas, error) {
	var resultados []responses.ProdutoVendas

	err := r.db.
		Table("venda_itens").
		Select("produtos.nome, SUM(venda_itens.quantidade) as quantidade_vendida, SUM(venda_itens.valor_total) as total_vendido").
		Joins("JOIN vendas ON venda_itens.venda_id = vendas.id").
		Joins("JOIN produtos ON venda_itens.produto_id = produtos.id").
		Where("vendas.empresa_id = ? AND venda_itens.produto_id IS NOT NULL", empresaID).
		Group("produtos.id, produtos.nome").
		Order("quantidade_vendida DESC").
		Limit(limite).
		Scan(&resultados).Error

	return resultados, err
}

func (r *vendaRepository) GetVendasPorFormaPagamento(empresaID uint, inicio, fim string) (map[string]float64, error) {
	resultados := make(map[string]float64)
	err := r.db.
		Table("vendas").
		Select("forma_pagamento, COALESCE(SUM(valor_final), 0) as total").
		Where("empresa_id = ? AND DATE(data_venda) BETWEEN ? AND ? AND status = ?", empresaID, inicio, fim, "pago").
		Group("forma_pagamento").
		Scan(&resultados).Error
	return resultados, err
}

func (r *vendaRepository) GetVendasPorPeriodo(empresaID uint, inicio, fim string) ([]models.Venda, error) {
	var vendas []models.Venda
	err := r.db.
		Where("empresa_id = ? AND DATE(data_venda) BETWEEN ? AND ?", empresaID, inicio, fim).
		Preload("Cliente").
		Preload("Usuario").
		Preload("Itens").
		Order("data_venda DESC").
		Find(&vendas).Error
	return vendas, err
}
