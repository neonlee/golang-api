package models

import (
	"time"
)

type Category struct {
	PetshopId int       `json:"petshop_id"`
	Id        int       `json:"id"`
	Nome      string    `json:"name"`
	Descricao string    `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
