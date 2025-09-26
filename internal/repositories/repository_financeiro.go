package repositories

import (
	"petApi/internal/models"
	"petApi/internal/responses"
	"time"

	"gorm.io/gorm"
)

// FinanceiroRepository interface
type FinanceiroRepository interface {
	CreateContaReceber(conta *models.ContaReceber) error
	GetContaReceberByID(id uint) (*models.ContaReceber, error)
	BaixarContaReceber(id uint, dataPagamento string, formaPagamento string) error
	CreateContaPagar(conta *models.ContaPagar) error
	GetContaPagarByID(id uint) (*models.ContaPagar, error)
	PagarContaPagar(id uint, dataPagamento string, formaPagamento string) error
	GetFluxoCaixa(empresaID uint, inicio time.Time, fim time.Time) (*responses.FluxoCaixa, error)
	GetContasVencidas(empresaID uint) ([]models.ContaReceber, []models.ContaPagar, error)
	GetDemonstrativo(empresaID uint, mes, ano int) (*responses.DemonstrativoFinanceiro, error)
}

type financeiroRepository struct {
	db *gorm.DB
}

func NewFinanceiroRepository(db *gorm.DB) FinanceiroRepository {
	return &financeiroRepository{db: db}
}

func (r *financeiroRepository) CreateContaReceber(conta *models.ContaReceber) error {
	return r.db.Create(conta).Error
}

func (r *financeiroRepository) GetContaReceberByID(id uint) (*models.ContaReceber, error) {
	var conta models.ContaReceber
	err := r.db.Preload("Clientes").Preload("Venda").First(&conta, id).Error
	return &conta, err
}

func (r *financeiroRepository) BaixarContaReceber(id uint, dataPagamento string, formaPagamento string) error {
	return r.db.Model(&models.ContaReceber{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"data_pagamento":  dataPagamento,
			"forma_pagamento": formaPagamento,
			"status":          "pago",
		}).Error
}

func (r *financeiroRepository) CreateContaPagar(conta *models.ContaPagar) error {
	return r.db.Create(conta).Error
}

func (r *financeiroRepository) GetContaPagarByID(id uint) (*models.ContaPagar, error) {
	var conta models.ContaPagar
	err := r.db.Preload("Fornecedor").Preload("CategoriaDespesa").First(&conta, id).Error
	return &conta, err
}

func (r *financeiroRepository) PagarContaPagar(id uint, dataPagamento string, formaPagamento string) error {
	return r.db.Model(&models.ContaPagar{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"data_pagamento":  dataPagamento,
			"forma_pagamento": formaPagamento,
			"status":          "pago",
		}).Error
}

func (r *financeiroRepository) GetFluxoCaixa(empresaID uint, inicio time.Time, fim time.Time) (*responses.FluxoCaixa, error) {
	var fluxo responses.FluxoCaixa

	// Receitas do período
	var receitas float64
	r.db.Model(&models.ContaReceber{}).
		Where("empresa_id = ? AND data_pagamento BETWEEN ? AND ? AND status = ?",
			empresaID, inicio, fim, "pago").
		Select("COALESCE(SUM(valor), 0)").
		Scan(&receitas)

	// Despesas do período
	var despesas float64
	r.db.Model(&models.ContaPagar{}).
		Where("empresa_id = ? AND data_pagamento BETWEEN ? AND ? AND status = ?",
			empresaID, inicio, fim, "pago").
		Select("COALESCE(SUM(valor_final), 0)").
		Scan(&despesas)

	fluxo.Receitas = receitas
	fluxo.Despesas = despesas
	fluxo.Saldo = receitas - despesas
	fluxo.PeriodoInicio = inicio
	fluxo.PeriodoFim = fim

	return &fluxo, nil
}

func (r *financeiroRepository) GetContasVencidas(empresaID uint) ([]models.ContaReceber, []models.ContaPagar, error) {
	hoje := time.Now().Format("2006-01-02")

	var contasReceber []models.ContaReceber
	err := r.db.
		Where("empresa_id = ? AND data_vencimento < ? AND status = ?",
			empresaID, hoje, "pendente").
		Preload("Clientes").
		Find(&contasReceber).Error
	if err != nil {
		return nil, nil, err
	}

	var contasPagar []models.ContaPagar
	err = r.db.
		Where("empresa_id = ? AND data_vencimento < ? AND status = ?",
			empresaID, hoje, "pendente").
		Preload("Fornecedor").
		Find(&contasPagar).Error

	return contasReceber, contasPagar, err
}

func (r *financeiroRepository) GetDemonstrativo(empresaID uint, mes, ano int) (*responses.DemonstrativoFinanceiro, error) {
	var demonstrativo responses.DemonstrativoFinanceiro

	inicio := time.Date(ano, time.Month(mes), 1, 0, 0, 0, 0, time.UTC)
	fim := inicio.AddDate(0, 1, -1)

	// Receitas do mês
	var receitas float64
	r.db.Model(&models.ContaReceber{}).
		Where("empresa_id = ? AND data_pagamento BETWEEN ? AND ? AND status = ?",
			empresaID, inicio, fim, "pago").
		Select("COALESCE(SUM(valor), 0)").
		Scan(&receitas)

	// Despesas do mês
	var despesas float64
	r.db.Model(&models.ContaPagar{}).
		Where("empresa_id = ? AND data_pagamento BETWEEN ? AND ? AND status = ?",
			empresaID, inicio, fim, "pago").
		Select("COALESCE(SUM(valor_final), 0)").
		Scan(&despesas)

	demonstrativo.Mes = mes
	demonstrativo.Ano = ano
	demonstrativo.TotalReceitas = receitas
	demonstrativo.TotalDespesas = despesas
	demonstrativo.Resultado = receitas - despesas

	return &demonstrativo, nil
}
