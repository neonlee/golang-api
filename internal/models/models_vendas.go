// models/entities/venda.go
package models

import (
	"time"
)

type Vendas struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	EmpresaID      uint      `gorm:"not null;index" json:"empresa_id"`
	ClientesID     uint      `gorm:"column:cliente_id;not null;index" json:"cliente_id"`
	UsuarioID      uint      `gorm:"column:usuario_id;not null;index" json:"usuario_id"`
	DataVenda      time.Time `gorm:"not null" json:"data_venda"`
	TipoVenda      string    `gorm:"size:20" json:"tipo_venda"`
	Status         string    `gorm:"size:20" json:"status"`
	ValorTotal     float64   `gorm:"type:decimal(10,2)" json:"valor_total"`
	Desconto       float64   `gorm:"type:decimal(10,2)" json:"desconto"`
	ValorFinal     float64   `gorm:"type:decimal(10,2)" json:"valor_final"`
	FormaPagamento string    `gorm:"size:30" json:"forma_pagamento"`
	Observacoes    string    `gorm:"type:text" json:"observacoes"`

	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`

	// Relacionamentos
	Empresa       Empresa        `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	Clientes      Clientes       `gorm:"foreignKey:ClientesID" json:"cliente,omitempty"`
	Usuario       Usuarios       `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`
	Itens         []VendaItem    `gorm:"foreignKey:VendaID" json:"itens,omitempty"`
	ContasReceber []ContaReceber `gorm:"foreignKey:VendaID" json:"contas_receber,omitempty"`
}

type VendaItem struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	VendaID       uint    `gorm:"not null;index" json:"venda_id"`
	ProdutoID     *uint   `gorm:"index" json:"produto_id"`
	TipoServicoID *uint   `gorm:"index" json:"tipo_servico_id"`
	Quantidade    int     `gorm:"not null" json:"quantidade"`
	ValorUnitario float64 `gorm:"type:decimal(10,2)" json:"valor_unitario"`
	ValorTotal    float64 `gorm:"type:decimal(10,2)" json:"valor_total"`
	TipoItem      string  `gorm:"size:10" json:"tipo_item"`

	// Relacionamentos
	Venda       Vendas      `gorm:"foreignKey:VendaID" json:"venda,omitempty"`
	Produto     Produtos    `gorm:"foreignKey:ProdutoID" json:"produto,omitempty"`
	TipoServico TipoServico `gorm:"foreignKey:TipoServicoID" json:"tipo_servico,omitempty"`
}

type TipoServico struct {
	ID             uint    `gorm:"primaryKey" json:"id"`
	EmpresaID      uint    `gorm:"not null;index" json:"empresa_id"`
	Nome           string  `gorm:"size:50;not null" json:"nome"`
	Categoria      string  `gorm:"size:30" json:"categoria"`
	DuracaoMinutos int     `json:"duracao_minutos"`
	Valor          float64 `gorm:"type:decimal(10,2)" json:"valor"`
	Descricao      string  `gorm:"type:text" json:"descricao"`
	Ativo          bool    `gorm:"default:true" json:"ativo"`

	// Relacionamentos
	Empresa      Empresa       `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	Agendamentos []Agendamento `gorm:"foreignKey:TipoServicoID" json:"agendamentos,omitempty"`
	VendaItens   []VendaItem   `gorm:"foreignKey:TipoServicoID" json:"venda_itens,omitempty"`
}

type Agendamento struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	EmpresaID       uint      `gorm:"not null;index" json:"empresa_id"`
	ClientesID      uint      `gorm:"not null;index" json:"cliente_id"`
	PetID           uint      `gorm:"not null;index" json:"pet_id"`
	TipoServicoID   uint      `gorm:"not null;index" json:"tipo_servico_id"`
	DataAgendamento time.Time `gorm:"not null" json:"data_agendamento"`
	Status          string    `gorm:"size:20" json:"status"`
	Observacoes     string    `gorm:"type:text" json:"observacoes"`
	ValorEstimado   float64   `gorm:"type:decimal(10,2)" json:"valor_estimado"`
	UsuarioID       uint      `gorm:"not null;index" json:"usuario_id"`

	// Relacionamentos
	Empresa     Empresa     `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	Clientes    Clientes    `gorm:"foreignKey:ClientesID" json:"cliente,omitempty"`
	Pet         Pets        `gorm:"foreignKey:PetID" json:"pet,omitempty"`
	TipoServico TipoServico `gorm:"foreignKey:TipoServicoID" json:"tipo_servico,omitempty"`
	Usuario     Usuarios    `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`
}
