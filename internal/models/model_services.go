package models

import (
	"time"
)

type Services struct {
	PetshopId int       `json:"petshop_id"`
	Id        int       `json:"id"`
	Nome      string    `json:"name"`
	Descricao string    `json:"description"`
	Preco     float64   `json:"price"`
	Tipo      string    `json:"type"`
	Duracao   time.Time `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
