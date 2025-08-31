package models

import (
	"time"
)

type Employee struct {
	PetshopId int       `gorm:"column:petshop_id" json:"petshop_id"`
	Id        int       `gorm:"column:id;primaryKey" json:"id"`
	Nome      string    `gorm:"column:nome; not null" json:"name"`
	Cpf       string    `gorm:"column:cpf; not null" json:"cpf"`
	Telefone  string    `gorm:"column:telefone" json:"cellphone"`
	Ativo     string    `gorm:"column:ativo;default:true" json:"active"`
	UserID    int       `gorm:"column:user_id;foreignKey:UserID;references:id" json:"user_id"`
	CreatedAt time.Time `gorm:"column:create_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
