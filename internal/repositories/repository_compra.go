package repositories

import (
	"errors"
	"fmt"
	"math"
	"petApi/internal/models"
	"petApi/internal/requests"
	"petApi/internal/responses"

	"gorm.io/gorm"
)

type ComprasRepository interface {
	Create(compra *models.Compras, itens []models.CompraItens) (*models.Compras, error)
	GetByID(id uint) (*models.Compras, error)
	Update(compra *models.Compras) error
	Cancelar(id uint, motivo string) error
	ListByEmpresa(empresaID uint, filters requests.CompraFilter) ([]models.Compras, error)
	ListByFornecedor(fornecedorID uint) ([]models.Compras, error)
	GetComprassPorPeriodo(empresaID uint, inicio, fim string) ([]models.Compras, error)
	GetResumoComprass(empresaID uint, periodo string) (*responses.ResumoCompras, error)
	GetItensByCompraID(compraID uint) ([]models.CompraItens, error)
}

type compraRepository struct {
	db *gorm.DB
}

func NewComprasRepository(db *gorm.DB) ComprasRepository {
	return &compraRepository{db: db}
}

func (r *compraRepository) Create(compra *models.Compras, itens []models.CompraItens) (*models.Compras, error) {
	// Iniciar transação
	tx := r.db.Begin()
	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	// Validar dados básicos
	if err := r.validarCompra(compra); err != nil {
		tx.Rollback()
		return nil, err
	}

	// Verificar se número da nota fiscal já existe (se fornecido)
	if compra.NumeroNotaFiscal != "" {
		var existingCompra models.Compras
		if err := tx.Where("empresa_id = ? AND numero_nota_fiscal = ?", compra.EmpresaID, compra.NumeroNotaFiscal).First(&existingCompra).Error; err == nil {
			tx.Rollback()
			return nil, errors.New("número da nota fiscal já existe para esta empresa")
		}
	}

	// Calcular valor total baseado nos itens se não fornecido
	if compra.ValorTotal == 0 && len(compra.Itens) > 0 {
		compra.ValorTotal = r.calcularValorTotal(compra.Itens)
	}

	// Definir status padrão se não informado
	if compra.Status == "" {
		compra.Status = "pendente"
	}

	// Criar a compra
	if err := tx.Create(compra).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Carregar relacionamentos
	if err := tx.Preload("Empresa").
		Preload("Fornecedor").
		Preload("Usuario").
		Preload("Itens.Produto").
		First(compra, compra.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit da transação
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return compra, nil
}

func (r *compraRepository) GetByID(id uint) (*models.Compras, error) {
	var compra models.Compras
	err := r.db.
		Preload("Fornecedor").
		Preload("Empresa").
		Preload("Usuario").
		Preload("Itens").
		Preload("Itens.Produto").
		First(&compra, id).Error

	return &compra, err
}

func (r *compraRepository) Update(compra *models.Compras) error {
	return r.db.Save(compra).Error
}

func (r *compraRepository) Cancelar(id uint, motivo string) error {
	return r.db.Model(&models.Compras{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":      "cancelado",
			"observacoes": motivo,
		}).Error
}

func (r *compraRepository) ListByEmpresa(empresaID uint, filters requests.CompraFilter) ([]models.Compras, error) {
	var compras []models.Compras

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

func (r *compraRepository) ListByFornecedor(fornecedorID uint) ([]models.Compras, error) {
	var compras []models.Compras

	err := r.db.
		Where("fornecedor_id = ?", fornecedorID).
		Preload("Fornecedor").
		Preload("Usuario").
		Order("data_compra DESC").
		Find(&compras).Error

	return compras, err
}

func (r *compraRepository) GetComprassPorPeriodo(empresaID uint, inicio, fim string) ([]models.Compras, error) {
	var compras []models.Compras

	err := r.db.
		Where("empresa_id = ? AND DATE(data_compra) BETWEEN ? AND ?", empresaID, inicio, fim).
		Preload("Fornecedor").
		Preload("Usuario").
		Preload("Itens").
		Order("data_compra DESC").
		Find(&compras).Error

	return compras, err
}

func (r *compraRepository) GetResumoComprass(empresaID uint, periodo string) (*responses.ResumoCompras, error) {
	var resumo responses.ResumoCompras

	// Implementar lógica de resumo por período
	// Exemplo: total de compras, valor total, etc.

	return &resumo, nil
}
func (r *compraRepository) GetItensByCompraID(compraID uint) ([]models.CompraItens, error) {
	var itens []models.CompraItens
	err := r.db.
		Where("compra_id = ?", compraID).
		Preload("Produto").
		Find(&itens).Error
	return itens, err
}

// Métodos auxiliares privados
func (r *compraRepository) validarCompra(compra *models.Compras) error {
	// Validar empresa
	var empresa models.Empresas
	if err := r.db.First(&empresa, compra.EmpresaID).Error; err != nil {
		return errors.New("empresa não encontrada")
	}

	// Validar fornecedor
	var fornecedor models.Fornecedores
	if err := r.db.First(&fornecedor, compra.FornecedorID).Error; err != nil {
		return errors.New("fornecedor não encontrado")
	}

	// Validar usuário
	var usuario models.Usuarios
	if err := r.db.First(&usuario, compra.UsuarioID).Error; err != nil {
		return errors.New("usuário não encontrado")
	}

	// Validar datas
	if compra.DataCompra.IsZero() {
		return errors.New("data da compra é obrigatória")
	}

	if compra.DataEntrada.IsZero() {
		compra.DataEntrada = compra.DataCompra
	}

	if compra.DataEntrada.Before(compra.DataCompra) {
		return errors.New("data de entrada não pode ser anterior à data da compra")
	}

	// Validar valores
	if compra.ValorTotal < 0 {
		return errors.New("valor total não pode ser negativo")
	}

	if compra.ValorFrete < 0 {
		return errors.New("valor do frete não pode ser negativo")
	}

	if compra.ValorDesconto < 0 {
		return errors.New("valor do desconto não pode ser negativo")
	}

	// Validar status
	statusValidos := map[string]bool{
		"pendente":     true,
		"em_andamento": true,
		"finalizada":   true,
		"cancelada":    true,
		"devolvida":    true,
	}
	if !statusValidos[compra.Status] {
		return errors.New("status inválido")
	}

	// Validar itens se existirem
	if len(compra.Itens) > 0 {
		for i, item := range compra.Itens {
			if err := r.validarItemCompra(&item); err != nil {
				return fmt.Errorf("erro no item %d: %v", i+1, err)
			}
			// Calcular valor total do item se não informado
			if item.ValorTotal == 0 {
				item.ValorTotal = float64(item.Quantidade) * item.ValorUnitario
			}
			compra.Itens[i] = item
		}
	}

	return nil
}
func (r *compraRepository) validarItemCompra(item *models.CompraItens) error {
	// Validar produto
	var produto models.Produtos
	if err := r.db.First(&produto, item.ProdutoID).Error; err != nil {
		return errors.New("produto não encontrado")
	}

	// Validar quantidade
	if item.Quantidade <= 0 {
		return errors.New("quantidade deve ser maior que zero")
	}

	// Validar valores
	if item.ValorUnitario <= 0 {
		return errors.New("valor unitário deve ser maior que zero")
	}

	if item.ValorTotal <= 0 {
		return errors.New("valor total deve ser maior que zero")
	}

	// Verificar consistência dos valores
	valorCalculado := float64(item.Quantidade) * item.ValorUnitario
	if math.Abs(valorCalculado-item.ValorTotal) > 0.01 { // Tolerância para float
		return errors.New("valor total inconsistente com quantidade e valor unitário")
	}

	return nil
}

func (r *compraRepository) calcularValorTotal(itens []models.CompraItens) float64 {
	var total float64
	for _, item := range itens {
		total += item.ValorTotal
	}
	return total
}
