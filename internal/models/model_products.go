package models

import (
	"time"
)

type Products struct {
	PetshopId         int       `json:"petshop_id"`
	Id                int       `gorm:"column:id;primaryKey" json:"id"`
	Sku               string    `gorm:"column:sku;index;not null" json:"sku"`
	Nome              string    `gorm:"column:nome" json:"name"`
	Photo             string    `gorm:"column:foto" json:"photo"`
	Descricao         string    `gorm:"column:descricao" json:"description"`
	PrecoCusto        float64   `gorm:"column:preco_custo" json:"price_custer"`
	PrecoVenda        float64   `gorm:"column:preco_venda" json:"price_venda"`
	IdCategory        int       `gorm:"column:id_category;index;not null" json:"category_id"`
	IdSupplier        int       `gorm:"column:id_supplier;index;not null" json:"supplier_id"`
	QuantidadeEstoque int       `gorm:"column:quantidade_estoque" json:"quantidade_estoque"`
	QuantidadeMinima  int       `gorm:"column:quantidade_minima" json:"quantidade_minima"`
	DataValidade      time.Time `gorm:"column:data_validade" json:"data_validade"`
	CreatedAt         time.Time `gorm:"column:create_at" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`
}
