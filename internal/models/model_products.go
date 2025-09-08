package models

import (
	"time"
)

type Product struct {
	Id                int       `gorm:"primaryKey;autoIncrement" validate:"required,min=2,max=100" json:"id"`
	PetshopID         int       `gorm:"column:petshop_id;not null;index" validate:"required" json:"petshop_id"`
	Sku               string    `gorm:"column:sku;type:varchar(100);uniqueIndex;not null" validate:"required,min=14,max=20" json:"sku"`
	Nome              string    `gorm:"column:nome;type:varchar(255);not null" validate:"required,min=2,max=100" json:"name"`
	Photo             string    `gorm:"column:foto;type:text" json:"photo"`
	Descricao         string    `gorm:"column:descricao;type:text" json:"description"`
	PrecoCusto        float64   `gorm:"column:preco_custo;type:decimal(10,2);not null check:price_custo > 0" validate:"required" json:"price_custo"`
	PrecoVenda        float64   `gorm:"column:preco_venda;type:decimal(10,2);not null check:price_venda > 0" validate:"required" json:"price_venda"`
	SupplierID        int       `gorm:"column:supplier_id;not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"required" json:"supplier_id"`
	Supplier          Supplier  `gorm:"foreignKey:id;references:SupplierID" json:"supplier,omitempty"`
	CategoryID        int       `gorm:"column:category_id;not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"required" json:"category_id"`
	Category          Category  `gorm:"foreignKey:Id;references:CategoryID" json:"category,omitempty"`
	QuantidadeEstoque int       `gorm:"column:quantidade_estoque;not null;default:0" validate:"required" json:"quantidade_estoque"`
	QuantidadeMinima  int       `gorm:"column:quantidade_minima;not null;default:1" validate:"required" json:"quantidade_minima"`
	DataValidade      string    `gorm:"column:data_validade;type:date" validate:"required" json:"data_validade"`
	CreatedAt         time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
