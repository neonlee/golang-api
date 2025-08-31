package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID                int       `gorm:"primaryKey;autoIncrement" json:"id"`
	PetshopID         int       `gorm:"column:petshop_id;not null;index" json:"petshop_id"`
	Sku               string    `gorm:"column:sku;type:varchar(100);uniqueIndex;not null" json:"sku"`
	Nome              string    `gorm:"column:nome;type:varchar(255);not null" json:"name"`
	Photo             string    `gorm:"column:foto;type:text" json:"photo"`
	Descricao         string    `gorm:"column:descricao;type:text" json:"description"`
	PrecoCusto        float64   `gorm:"column:preco_custo;type:decimal(10,2);not null" json:"price_custo"`
	PrecoVenda        float64   `gorm:"column:preco_venda;type:decimal(10,2);not null" json:"price_venda"`
	SupplierID        int       `gorm:"column:supplier_id;not null;index;foreignKey:SupplierId;references:id" json:"supplier_id"`
	CategoryID        int       `gorm:"column:category_id;not null;indexforeignKey:CategoryId;references:id" json:"category_id"`
	QuantidadeEstoque int       `gorm:"column:quantidade_estoque;not null;default:0" json:"quantidade_estoque"`
	QuantidadeMinima  int       `gorm:"column:quantidade_minima;not null;default:1" json:"quantidade_minima"`
	DataValidade      time.Time `gorm:"column:data_validade" json:"data_validade"`
	CreatedAt         time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
