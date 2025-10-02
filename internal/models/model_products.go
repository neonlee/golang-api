// models/entities/produto.go
package models

import (
	"gorm.io/gorm"
)

type CategoriaProdutos struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	EmpresaID uint   `gorm:"not null;index" json:"empresa_id"`
	Nome      string `gorm:"size:50;not null" json:"nome"`
	Descricao string `gorm:"type:text" json:"descricao"`
	Ativo     bool   `gorm:"default:true" json:"ativo"`

	// Relacionamentos
	Empresa  Empresa    `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	Produtos []Produtos `gorm:"foreignKey:CategoriaID" json:"produtos,omitempty"`
}

type Fornecedores struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	EmpresaID    uint   `gorm:"not null;index" json:"empresa_id"`
	NomeFantasia string `gorm:"size:100;not null" json:"nome_fantasia"`
	RazaoSocial  string `gorm:"size:100" json:"razao_social"`
	CNPJ         string `gorm:"size:18" json:"cnpj"`
	Telefone     string `gorm:"size:15" json:"telefone"`
	Email        string `gorm:"size:100" json:"email"`
	Endereco     JSON   `gorm:"type:jsonb" json:"endereco"`
	Ativo        bool   `gorm:"default:true" json:"ativo"`

	// Relacionamentos
	Empresa  Empresa    `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	Produtos []Produtos `gorm:"foreignKey:FornecedorID" json:"produtos,omitempty"`
	Compras  []Compras  `gorm:"foreignKey:FornecedorID" json:"compras,omitempty"`
}

type Produtos struct {
	ID               uint    `gorm:"primaryKey" json:"id"`
	EmpresaID        uint    `gorm:"not null;index" json:"empresa_id"`
	CategoriaID      uint    `gorm:"not null;index" json:"categoria_id"`
	FornecedorID     uint    `gorm:"index" json:"fornecedor_id"`
	CodigoBarras     string  `gorm:"size:50" json:"codigo_barras"`
	Nome             string  `gorm:"size:100;not null" json:"nome"`
	Descricao        string  `gorm:"type:text" json:"descricao"`
	Tipo             string  `gorm:"size:20" json:"tipo"`
	EspecieDestinada string  `gorm:"size:20" json:"especie_destinada"`
	PesoKg           float64 `gorm:"type:decimal(8,3)" json:"peso_kg"`
	UnidadeMedida    string  `gorm:"size:10" json:"unidade_medida"`
	PrecoCusto       float64 `gorm:"type:decimal(10,2)" json:"preco_custo"`
	PrecoVenda       float64 `gorm:"type:decimal(10,2)" json:"preco_venda"`
	EstoqueMinimo    int     `gorm:"default:0" json:"estoque_minimo"`
	EstoqueAtual     int     `gorm:"default:0" json:"estoque_atual"`
	Ativo            bool    `gorm:"default:true" json:"ativo"`

	// Relacionamentos
	Empresa       Empresa               `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	Categoria     CategoriaProdutos     `gorm:"foreignKey:CategoriaID" json:"categoria,omitempty"`
	Fornecedor    Fornecedores          `gorm:"foreignKey:FornecedorID" json:"fornecedor,omitempty"`
	Movimentacoes []MovimentacaoEstoque `gorm:"foreignKey:ProdutoID" json:"movimentacoes,omitempty"`
	VendaItens    []VendaItem           `gorm:"foreignKey:ProdutoID" json:"venda_itens,omitempty"`
	CompraItens   []CompraItens         `gorm:"foreignKey:ProdutoID" json:"compra_itens,omitempty"`
}

type MovimentacaoEstoque struct {
	gorm.Model
	ID                 uint   `gorm:"primaryKey" json:"id"`
	ProdutoID          uint   `gorm:"not null;index" json:"produto_id"`
	TipoMovimentacao   string `gorm:"size:20;not null" json:"tipo_movimentacao"`
	Quantidade         int    `gorm:"not null" json:"quantidade"`
	QuantidadeAnterior int    `json:"quantidade_anterior"`
	QuantidadeAtual    int    `json:"quantidade_atual"`
	Motivo             string `gorm:"size:100" json:"motivo"`
	UsuarioID          uint   `gorm:"not null;index" json:"usuario_id"`

	// Relacionamentos
	Produto Produtos `gorm:"foreignKey:ProdutoID" json:"produto,omitempty"`
	Usuario Usuarios `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`
}
