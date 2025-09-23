package models

import (
	"time"
)

type Service struct {
	PetshopId       int       `gorm:"column:petshop_id;index;not null" json:"petshop_id"`
	Id              int       `gorm:"column:id;primaryKey" json:"id"`
	Nome            string    `gorm:"column:nome" json:"name"`
	Descricao       string    `gorm:"column:descricao" json:"description"`
	Active          bool      `gorm:"column:ativo" json:"active"`
	Preco           float64   `gorm:"column:preco" json:"price"`
	ComissaoTecnico float64   `gorm:"column:comissao_tecnico" json:"technician_commission"`
	Duracao         string    `gorm:"column:duracao;type:interval" json:"duration"`
	CreatedAt       time.Time `gorm:"column:create_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}
