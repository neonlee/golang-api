package models

import (
	"time"

	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model
	PetshopID int       `gorm:"column:petshop_id" json:"petshop_id"`
	ID        int       `gorm:"column:id;primaryKey" json:"id"`
	Name      string    `gorm:"column:nome" json:"name"`
	Cnpj      string    `gorm:"column:cnpj;unique" json:"cnpj"`
	Phone     string    `gorm:"column:telefone" json:"phone"`
	Email     string    `gorm:"column:email" json:"email"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // Corrigido create_at para created_at
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Products  []Product `gorm:"foreignKey:SupplierID;references:id" json:"products,omitempty"`
}
