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
	UserID    int       `gorm:"column:user_id;not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_id"`
	Supplier  Users     `gorm:"foreignKey:id;references:UserID" json:"supplier,omitempty"`
	CreatedAt time.Time `gorm:"column:create_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
