package models

import (
	"time"
)

type Users struct {
	ID           int        `gorm:"primaryKey" json:"id"`
	Username     string     `gorm:"unique;not null" json:"username"`
	Password     string     `gorm:"not null" json:"password"`
	Email        string     `gorm:"unique" json:"email"`
	TipoAcessoID int        `json:"tipo_acesso_id"`
	TipoAcesso   TipoAcesso `json:"tipo_acesso"`
	Ativo        bool       `gorm:"default:true" json:"ativo"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
