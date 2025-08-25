package models

import (
	"time"
)

type Category struct {
	PetshopId int       `gorm:"column:petshop_id" json:"petshop_id"`
	Id        int       `gorm:"column:id;primaryKey" json:"id"`
	Nome      string    `gorm:"column:nome" json:"name"`
	Descricao string    `gorm:"column:descricao" json:"description"`
	CreatedAt time.Time `gorm:"column:create_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
