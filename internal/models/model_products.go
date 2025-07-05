package models

import (
	"time"
)

type Products struct {
	PetshopId         int       `json:"petshop_id"`
	Id                int       `json:"id"`
	Nome              string    `json:"name"`
	Marca             string    `json:"marca"`
	Photo             string    `json:"photo"`
	Descricao         string    `json:"description"`
	PrecoCusto        float64   `json:"price_custer"`
	PrecoVenda        float64   `json:"price_venda"`
	IdCategory        int       `json:"category_id"`
	QuantidadeEstoque int       `json:"quantidade_estoque"`
	QuantidadeMinima  int       `json:"quantidade_minima"`
	DataValidade      time.Time `json:"data_validade"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
