package models

import (
	"time"
)

type Supplier struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PetshopID int       `gorm:"column:petshop_id" json:"petshop_id"`
	Name      string    `gorm:"column:nome" json:"name"`
	Cnpj      string    `gorm:"column:cnpj;unique" json:"cnpj"`
	Phone     string    `gorm:"column:telefone" json:"phone"`
	Email     string    `gorm:"column:email" json:"email"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
