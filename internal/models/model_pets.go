package models

import (
	"time"
)

type Pet struct {
	ID          int       `gorm:"column:id;primaryKey" json:"id"`
	ClientID    int       `gorm:"column:client_id;index;not null" json:"client_id"`
	Name        string    `gorm:"column:nome" json:"name"`
	Race        string    `gorm:"column:raca" json:"race"`
	Specie      string    `gorm:"column:especie" json:"specie"`
	Color       string    `gorm:"column:cor" json:"color"`
	Size        string    `gorm:"column:porte" json:"size"`
	Weight      float64   `gorm:"column:peso" json:"weight"`
	Age         int       `gorm:"column:idade" json:"age"`
	Observation string    `gorm:"column:observacao" json:"observation"`
	Sex         string    `gorm:"column:sexo" json:"sex"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
