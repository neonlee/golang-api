package models

import (
	"time"
)

type Supplier struct {
	ID             int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PetshopID      int       `gorm:"column:petshop_id" json:"petshop_id"`
	Name           string    `gorm:"column:nome" json:"name"`
	Cnpj           string    `gorm:"column:cnpj;unique" json:"cnpj"`
	Phone          string    `gorm:"column:telefone" json:"phone"`
	Email          string    `gorm:"column:email" json:"email"`
	NumeroEndereco string    `gorm:"column:numero" json:"numero"`
	Bairro         string    `gorm:"column:bairro" json:"bairro"`
	Cidade         string    `gorm:"column:cidade" json:"cidade"`
	Estado         string    `gorm:"column:estado" json:"estado"`
	Address        string    `gorm:"column:endereco" json:"address"`
	Complemento    string    `gorm:"column:complemento" json:"complemento"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
