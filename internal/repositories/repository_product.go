package repositories

import (
	"petApi/internal/models"
	"petApi/internal/requests"

	"gorm.io/gorm"
)

// ProdutoRepository interface
type ProdutoRepository interface {
	Create(produto *models.Produtos) error
	GetByID(id uint) (*models.Produtos, error)
	Update(produto *models.Produtos) error
	Delete(id uint) error
	ListByEmpresa(empresaID uint, filters requests.ProdutoFilter) ([]models.Produtos, error)
	GetByCategoria(categoriaID uint) ([]models.Produtos, error)
	Search(empresaID uint, termo string) ([]models.Produtos, error)
	GetProdutosBaixoEstoque(empresaID uint) ([]models.Produtos, error)
	UpdateEstoque(produtoID uint, quantidade int) error
	GetProdutoComEstoque(id uint) (*models.Produtos, error)
	GetProdutosProximosVencimento(id uint) ([]models.Produtos, error)
	GetProdutosVencimentoHoje(id uint) ([]models.Produtos, error)
	GetProdutosSemEstoque(id uint) ([]models.Produtos, error)
	GetProdutosVencidos(id uint) ([]models.Produtos, error)
}

type produtoRepository struct {
	db *gorm.DB
}

func NewProdutoRepository(db *gorm.DB) ProdutoRepository {
	return &produtoRepository{db: db}
}

func (r *produtoRepository) Create(produto *models.Produtos) error {
	return r.db.Create(produto).Error
}

func (r *produtoRepository) GetByID(id uint) (*models.Produtos, error) {
	var produto models.Produtos
	err := r.db.
		Preload("Categoria").
		Preload("Fornecedor").
		First(&produto, id).Error

	return &produto, err
}

func (r *produtoRepository) Update(produto *models.Produtos) error {
	return r.db.Save(produto).Error
}

func (r *produtoRepository) Delete(id uint) error {
	return r.db.Delete(&models.Produtos{}, id).Error
}

func (r *produtoRepository) ListByEmpresa(empresaID uint, filters requests.ProdutoFilter) ([]models.Produtos, error) {
	var produtos []models.Produtos

	query := r.db.Where("empresa_id = ?", empresaID)

	if filters.Nome != "" {
		query = query.Where("nome ILIKE ?", "%"+filters.Nome+"%")
	}

	if filters.CategoriaID != nil {
		query = query.Where("categoria_id = ?", *filters.CategoriaID)
	}

	if filters.Ativo != nil {
		query = query.Where("ativo = ?", *filters.Ativo)
	}

	if filters.EspecieDestinada != "" {
		query = query.Where("especie_destinada = ?", filters.EspecieDestinada)
	}

	err := query.
		Preload("Categoria").
		Preload("Fornecedor").
		Order("nome ASC").
		Find(&produtos).Error

	return produtos, err
}

func (r *produtoRepository) GetByCategoria(categoriaID uint) ([]models.Produtos, error) {
	var produtos []models.Produtos
	err := r.db.
		Where("categoria_id = ?", categoriaID).
		Preload("Categoria").
		Preload("Fornecedor").
		Order("nome ASC").
		Find(&produtos).Error

	return produtos, err
}

func (r *produtoRepository) Search(empresaID uint, termo string) ([]models.Produtos, error) {
	var produtos []models.Produtos

	err := r.db.
		Where("empresa_id = ? AND (nome ILIKE ? OR codigo_barras ILIKE ?)",
			empresaID, "%"+termo+"%", "%"+termo+"%").
		Preload("Categoria").
		Preload("Fornecedor").
		Order("nome ASC").
		Limit(50).
		Find(&produtos).Error

	return produtos, err
}

func (r *produtoRepository) GetProdutosBaixoEstoque(empresaID uint) ([]models.Produtos, error) {
	var produtos []models.Produtos

	// Subquery para obter estoque atual
	subquery := r.db.Model(&models.MovimentacaoEstoque{}).
		Select("produto_id, quantidade_atual").
		Where("id IN (SELECT MAX(id) FROM movimentacao_estoque GROUP BY produto_id)")

	err := r.db.
		Joins("JOIN (?) AS estoque ON produtos.id = estoque.produto_id", subquery).
		Where("produtos.empresa_id = ? AND estoque.quantidade_atual <= produtos.estoque_minimo", empresaID).
		Preload("Categoria").
		Preload("Fornecedor").
		Order("estoque.quantidade_atual ASC").
		Find(&produtos).Error

	return produtos, err
}

func (r *produtoRepository) UpdateEstoque(produtoID uint, quantidade int) error {
	// Obter estoque atual
	var movimentacao models.MovimentacaoEstoque
	err := r.db.
		Where("produto_id = ?", produtoID).
		Order("created_at DESC").
		First(&movimentacao).Error

	quantidadeAnterior := 0
	if err == nil {
		quantidadeAnterior = movimentacao.QuantidadeAtual
	}

	// Criar nova movimentação
	novaMovimentacao := models.MovimentacaoEstoque{
		ProdutoID:          produtoID,
		TipoMovimentacao:   "ajuste",
		Quantidade:         quantidade,
		QuantidadeAnterior: quantidadeAnterior,
		QuantidadeAtual:    quantidade,
		Motivo:             "Ajuste manual de estoque",
		UsuarioID:          1, // TODO: Obter do usuário logado
	}

	return r.db.Create(&novaMovimentacao).Error
}

func (r *produtoRepository) GetProdutoComEstoque(id uint) (*models.Produtos, error) {
	var produto models.Produtos

	// Subquery para estoque atual
	subquery := r.db.Model(&models.MovimentacaoEstoque{}).
		Select("produto_id, quantidade_atual").
		Where("id IN (SELECT MAX(id) FROM movimentacao_estoque GROUP BY produto_id)")

	err := r.db.
		Joins("JOIN (?) AS estoque ON produtos.id = estoque.produto_id", subquery).
		Preload("Categoria").
		Preload("Fornecedor").
		First(&produto, id).Error

	return &produto, err
}

// GetProdutosProximosVencimento implements ProdutoRepository.
func (r *produtoRepository) GetProdutosProximosVencimento(id uint) ([]models.Produtos, error) {
	var produtos []models.Produtos
	err := r.db.
		Where("empresa_id = ? AND data_validade IS NOT NULL AND data_validade > NOW() AND data_validade <= NOW() + INTERVAL '7 days'", id).
		Preload("Categoria").
		Preload("Fornecedor").
		Order("data_validade ASC").
		Find(&produtos).Error
	return produtos, err
}

// GetProdutosSemEstoque implements ProdutoRepository.
func (r *produtoRepository) GetProdutosSemEstoque(id uint) ([]models.Produtos, error) {
	var produtos []models.Produtos
	err := r.db.
		Joins("LEFT JOIN (SELECT produto_id, quantidade_atual FROM movimentacao_estoque WHERE id IN (SELECT MAX(id) FROM movimentacao_estoque GROUP BY produto_id)) AS estoque ON produtos.id = estoque.produto_id").
		Where("produtos.empresa_id = ? AND (estoque.quantidade_atual IS NULL OR estoque.quantidade_atual <= 0)", id).
		Preload("Categoria").
		Preload("Fornecedor").
		Order("produtos.nome ASC").
		Find(&produtos).Error
	return produtos, err
}

// GetProdutosVencidos implements ProdutoRepository.
func (r *produtoRepository) GetProdutosVencidos(id uint) ([]models.Produtos, error) {
	var produtos []models.Produtos
	err := r.db.
		Where("empresa_id = ? AND data_validade IS NOT NULL AND data_validade < NOW()", id).
		Preload("Categoria").
		Preload("Fornecedor").
		Order("data_validade ASC").
		Find(&produtos).Error
	return produtos, err
}

// GetProdutosVencimentoHoje implements ProdutoRepository.
func (r *produtoRepository) GetProdutosVencimentoHoje(id uint) ([]models.Produtos, error) {
	var produtos []models.Produtos
	err := r.db.
		Where("empresa_id = ? AND data_validade IS NOT NULL AND data_validade = CURRENT_DATE", id).
		Preload("Categoria").
		Preload("Fornecedor").
		Order("data_validade ASC").
		Find(&produtos).Error
	return produtos, err
}
