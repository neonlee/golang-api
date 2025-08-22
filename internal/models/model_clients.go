package models

import (
	"time"
)

type Client struct { // Mudando para singular (boa prática)
	PetshopID    int       `gorm:"column:petshop_id" json:"petshop_id"`
	ID           int       `gorm:"column:id;primaryKey" json:"id"`
	Name         string    `gorm:"column:nome" json:"name"`
	Observations string    `gorm:"column:observacoes" json:"observations"`
	Phone        string    `gorm:"column:telefone" json:"phone"`
	Email        string    `gorm:"column:email" json:"email"`
	Address      string    `gorm:"column:endereco" json:"address"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"` // Corrigido create_at para created_at
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	Pets         []Pet     `gorm:"foreignKey:ClientID;references:id" json:"pets,omitempty"`
}
