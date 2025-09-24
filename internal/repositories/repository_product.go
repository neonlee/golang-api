package repositories

import (
	"petApi/internal/models"
	"petApi/internal/requests"

	"gorm.io/gorm"
)

// ProdutoRepository interface
type ProdutoRepository interface {
	Create(produto *models.Produto) error
	GetByID(id uint) (*models.Produto, error)
	Update(produto *models.Produto) error
	Delete(id uint) error
	ListByEmpresa(empresaID uint, filters requests.ProdutoFilter) ([]models.Produto, error)
	GetByCategoria(categoriaID uint) ([]models.Produto, error)
	Search(empresaID uint, termo string) ([]models.Produto, error)
	GetProdutosBaixoEstoque(empresaID uint) ([]models.Produto, error)
	UpdateEstoque(produtoID uint, quantidade int) error
	GetProdutoComEstoque(id uint) (*models.Produto, error)
}

type produtoRepository struct {
	db *gorm.DB
}

func NewProdutoRepository(db *gorm.DB) ProdutoRepository {
	return &produtoRepository{db: db}
}

func (r *produtoRepository) Create(produto *models.Produto) error {
	return r.db.Create(produto).Error
}

func (r *produtoRepository) GetByID(id uint) (*models.Produto, error) {
	var produto models.Produto
	err := r.db.
		Preload("Categoria").
		Preload("Fornecedor").
		First(&produto, id).Error

	return &produto, err
}

func (r *produtoRepository) Update(produto *models.Produto) error {
	return r.db.Save(produto).Error
}

func (r *produtoRepository) Delete(id uint) error {
	return r.db.Delete(&models.Produto{}, id).Error
}

func (r *produtoRepository) ListByEmpresa(empresaID uint, filters requests.ProdutoFilter) ([]models.Produto, error) {
	var produtos []models.Produto

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

func (r *produtoRepository) GetByCategoria(categoriaID uint) ([]models.Produto, error) {
	var produtos []models.Produto
	err := r.db.
		Where("categoria_id = ?", categoriaID).
		Preload("Categoria").
		Preload("Fornecedor").
		Order("nome ASC").
		Find(&produtos).Error

	return produtos, err
}

func (r *produtoRepository) Search(empresaID uint, termo string) ([]models.Produto, error) {
	var produtos []models.Produto

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

func (r *produtoRepository) GetProdutosBaixoEstoque(empresaID uint) ([]models.Produto, error) {
	var produtos []models.Produto

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

func (r *produtoRepository) GetProdutoComEstoque(id uint) (*models.Produto, error) {
	var produto models.Produto

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
