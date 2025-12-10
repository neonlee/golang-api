// repositories/estoque_repository.go
package repositories

import (
	"petApi/internal/models"
	"petApi/internal/responses"

	"gorm.io/gorm"
)

// EstoqueRepository interface
type EstoqueRepository interface {
	MovimentarEstoque(movimentacao *models.MovimentacaoEstoques) error
	GetHistoricoEstoque(produtoID uint) ([]models.MovimentacaoEstoques, error)
	GetSaldoAtual(produtoID uint) (int, error)
	AjustarEstoque(produtoID uint, novaQuantidade int, motivo string, usuarioID uint) error
	TransferirEstoque(origemID, destinoID uint, quantidade int, usuarioID uint) error
	GetRelatorioEstoque(empresaID uint) ([]responses.RelatorioEstoque, error)
	GetMovimentacoesPorPeriodo(empresaID uint, inicio, fim string) ([]models.MovimentacaoEstoques, error)
}
type estoqueRepository struct {
	db *gorm.DB
}

func NewEstoqueRepository(db *gorm.DB) EstoqueRepository {
	return &estoqueRepository{db: db}
}

func (r *estoqueRepository) MovimentarEstoque(movimentacao *models.MovimentacaoEstoques) error {
	// Obter estoque atual
	var ultimaMovimentacao models.MovimentacaoEstoques
	err := r.db.
		Where("produto_id = ?", movimentacao.ProdutoID).
		Order("created_at DESC").
		First(&ultimaMovimentacao).Error

	quantidadeAnterior := 0
	if err == nil {
		quantidadeAnterior = ultimaMovimentacao.QuantidadeAtual
	}

	movimentacao.QuantidadeAnterior = quantidadeAnterior

	// Calcular nova quantidade
	switch movimentacao.TipoMovimentacao {
	case "entrada":
		movimentacao.QuantidadeAtual = quantidadeAnterior + movimentacao.Quantidade
	case "saida", "ajuste":
		movimentacao.QuantidadeAtual = quantidadeAnterior - movimentacao.Quantidade
	}

	return r.db.Create(movimentacao).Error
}

func (r *estoqueRepository) GetHistoricoEstoque(produtoID uint) ([]models.MovimentacaoEstoques, error) {
	var movimentacoes []models.MovimentacaoEstoques

	err := r.db.
		Where("produto_id = ?", produtoID).
		Preload("Usuario").
		Preload("Produto").
		Order("created_at DESC").
		Find(&movimentacoes).Error

	return movimentacoes, err
}

func (r *estoqueRepository) GetSaldoAtual(produtoID uint) (int, error) {
	var movimentacao models.MovimentacaoEstoques

	err := r.db.
		Where("produto_id = ?", produtoID).
		Order("created_at DESC").
		First(&movimentacao).Error

	if err != nil {
		return 0, err
	}

	return movimentacao.QuantidadeAtual, nil
}

func (r *estoqueRepository) AjustarEstoque(produtoID uint, novaQuantidade int, motivo string, usuarioID uint) error {
	// Obter estoque atual
	saldoAtual, err := r.GetSaldoAtual(produtoID)
	if err != nil {
		saldoAtual = 0
	}

	diferenca := novaQuantidade - saldoAtual
	tipoMovimentacao := "ajuste"
	if diferenca > 0 {
		tipoMovimentacao = "entrada"
	} else if diferenca < 0 {
		tipoMovimentacao = "saida"
		diferenca = -diferenca
	} else {
		// Não há alteração
		return nil
	}

	movimentacao := models.MovimentacaoEstoques{
		ProdutoID:          produtoID,
		TipoMovimentacao:   tipoMovimentacao,
		Quantidade:         diferenca,
		QuantidadeAnterior: saldoAtual,
		QuantidadeAtual:    novaQuantidade,
		Motivo:             motivo,
		UsuarioID:          usuarioID,
	}

	return r.MovimentarEstoque(&movimentacao)
}

func (r *estoqueRepository) TransferirEstoque(origemID, destinoID uint, quantidade int, usuarioID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Saída do produto origem
		movimentacaoSaida := models.MovimentacaoEstoques{
			ProdutoID:        origemID,
			TipoMovimentacao: "saida",
			Quantidade:       quantidade,
			Motivo:           "Transferência entre produtos",
			UsuarioID:        usuarioID,
		}
		if err := r.MovimentarEstoque(&movimentacaoSaida); err != nil {
			return err
		}

		// Entrada no produto destino
		movimentacaoEntrada := models.MovimentacaoEstoques{
			ProdutoID:        destinoID,
			TipoMovimentacao: "entrada",
			Quantidade:       quantidade,
			Motivo:           "Transferência entre produtos",
			UsuarioID:        usuarioID,
		}
		return r.MovimentarEstoque(&movimentacaoEntrada)
	})
}

func (r *estoqueRepository) GetRelatorioEstoque(empresaID uint) ([]responses.RelatorioEstoque, error) {
	var relatorios []responses.RelatorioEstoque

	// Subquery para estoque atual
	subquery := r.db.Model(&models.MovimentacaoEstoques{}).
		Select("produto_id, quantidade_atual").
		Where("id IN (SELECT MAX(id) FROM movimentacao_estoque GROUP BY produto_id)")

	err := r.db.
		Table("produtos").
		Select("produtos.id, produtos.nome, produtos.estoque_minimo, estoque.quantidade_atual as estoque_atual").
		Joins("JOIN (?) AS estoque ON produtos.id = estoque.produto_id", subquery).
		Where("produtos.empresa_id = ?", empresaID).
		Order("produtos.nome ASC").
		Scan(&relatorios).Error

	return relatorios, err
}

func (r *estoqueRepository) GetMovimentacoesPorPeriodo(empresaID uint, inicio, fim string) ([]models.MovimentacaoEstoques, error) {
	var movimentacoes []models.MovimentacaoEstoques

	err := r.db.
		Joins("JOIN produtos ON movimentacao_estoque.produto_id = produtos.id").
		Where("produtos.empresa_id = ? AND DATE(movimentacao_estoque.created_at) BETWEEN ? AND ?",
			empresaID, inicio, fim).
		Preload("Produto").
		Preload("Usuario").
		Order("movimentacao_estoque.created_at DESC").
		Find(&movimentacoes).Error

	return movimentacoes, err
}
