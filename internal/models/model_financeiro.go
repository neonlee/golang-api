// models/entities/financeiro.go
package models

import (
	"time"
)

type Compras struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	EmpresaID        uint      `gorm:"not null;index" json:"empresa_id"`
	FornecedorID     uint      `gorm:"not null;index" json:"fornecedor_id"`
	NumeroNotaFiscal string    `gorm:"size:50" json:"numero_nota_fiscal"`
	DataCompra       time.Time `gorm:"type:date" json:"data_compra"`
	DataEntrada      time.Time `gorm:"type:date" json:"data_entrada"`
	ValorTotal       float64   `gorm:"type:decimal(10,2)" json:"valor_total"`
	ValorFrete       float64   `gorm:"type:decimal(10,2);default:0" json:"valor_frete"`
	ValorDesconto    float64   `gorm:"type:decimal(10,2);default:0" json:"valor_desconto"`
	Status           string    `gorm:"size:20;default:'pendente'" json:"status"`
	Observacoes      string    `gorm:"type:text" json:"observacoes"`
	UsuarioID        uint      `gorm:"not null;index" json:"usuario_id"`

	// Relacionamentos
	Empresa    Empresa       `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	Fornecedor Fornecedores  `gorm:"foreignKey:FornecedorID" json:"fornecedor,omitempty"`
	Usuario    Usuarios      `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`
	Itens      []CompraItens `gorm:"foreignKey:CompraID" json:"itens,omitempty"`
}

type CompraItens struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	CompraID      uint    `gorm:"not null;index" json:"compra_id"`
	ProdutoID     uint    `gorm:"not null;index" json:"produto_id"`
	Quantidade    int     `gorm:"not null" json:"quantidade"`
	ValorUnitario float64 `gorm:"type:decimal(10,2)" json:"valor_unitario"`
	ValorTotal    float64 `gorm:"type:decimal(10,2)" json:"valor_total"`

	// Relacionamentos
	Compra  Compras  `gorm:"foreignKey:CompraID" json:"compra,omitempty"`
	Produto Produtos `gorm:"foreignKey:ProdutoID" json:"produto,omitempty"`
}

type ContaReceber struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	VendaID        *uint      `gorm:"index" json:"venda_id"`
	ClientesID     uint       `gorm:"not null;index" json:"cliente_id"`
	Descricao      string     `gorm:"size:100;not null" json:"descricao"`
	Valor          float64    `gorm:"type:decimal(10,2)" json:"valor"`
	DataVencimento time.Time  `gorm:"type:date" json:"data_vencimento"`
	DataPagamento  *time.Time `gorm:"type:date" json:"data_pagamento"`
	Status         string     `gorm:"size:20" json:"status"`
	FormaPagamento string     `gorm:"size:30" json:"forma_pagamento"`

	// Relacionamentos
	Venda    Vendas   `gorm:"foreignKey:VendaID" json:"venda,omitempty"`
	Clientes Clientes `gorm:"foreignKey:ClientesID" json:"cliente,omitempty"`
}

type ContaPagar struct {
	ID                 uint       `gorm:"primaryKey" json:"id"`
	EmpresaID          uint       `gorm:"not null;index" json:"empresa_id"`
	FornecedorID       *uint      `gorm:"index" json:"fornecedor_id"`
	CategoriaDespesaID uint       `gorm:"not null;index" json:"categoria_despesa_id"`
	Descricao          string     `gorm:"size:100;not null" json:"descricao"`
	NumeroDocumento    string     `gorm:"size:50" json:"numero_documento"`
	ValorOriginal      float64    `gorm:"type:decimal(10,2)" json:"valor_original"`
	ValorJuros         float64    `gorm:"type:decimal(10,2);default:0" json:"valor_juros"`
	ValorMulta         float64    `gorm:"type:decimal(10,2);default:0" json:"valor_multa"`
	ValorDesconto      float64    `gorm:"type:decimal(10,2);default:0" json:"valor_desconto"`
	ValorFinal         float64    `gorm:"type:decimal(10,2)" json:"valor_final"`
	DataEmissao        time.Time  `gorm:"type:date" json:"data_emissao"`
	DataVencimento     time.Time  `gorm:"type:date" json:"data_vencimento"`
	DataPagamento      *time.Time `gorm:"type:date" json:"data_pagamento"`
	Status             string     `gorm:"size:20;default:'pendente'" json:"status"`
	FormaPagamento     string     `gorm:"size:30" json:"forma_pagamento"`
	Observacoes        string     `gorm:"type:text" json:"observacoes"`
	UsuarioID          uint       `gorm:"not null;index" json:"usuario_id"`

	// Relacionamentos
	Empresa          Empresa          `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	Fornecedor       Fornecedores     `gorm:"foreignKey:FornecedorID" json:"fornecedor,omitempty"`
	CategoriaDespesa CategoriaDespesa `gorm:"foreignKey:CategoriaDespesaID" json:"categoria_despesa,omitempty"`
	Usuario          Usuarios         `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`
}

type CategoriaDespesa struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	EmpresaID uint   `gorm:"not null;index" json:"empresa_id"`
	Nome      string `gorm:"size:50;not null" json:"nome"`
	Descricao string `gorm:"type:text" json:"descricao"`
	Tipo      string `gorm:"size:20" json:"tipo"`
	Ativo     bool   `gorm:"default:true" json:"ativo"`

	// Relacionamentos
	Empresa     Empresa      `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	ContasPagar []ContaPagar `gorm:"foreignKey:CategoriaDespesaID" json:"contas_pagar,omitempty"`
}
