package models

import (
	"time"
)

type Clients struct {
	PetshopId   int       `json:"petshop_id"`
	Id          int       `json:"id"`
	Nome        string    `json:"name"`
	Observacoes string    `json:"observacoes"`
	Telefone    string    `json:"telefone"`
	Email       string    `json:"email"`
	Endereco    string    `json:"endereco"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Pets        []Pet     `json:"pets,omitempty"`
}
