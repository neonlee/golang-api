// models/entities/cliente.go
package models

import (
	"time"
)

type Clientes struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	EmpresaID      uint       `gorm:"not null;index" json:"empresa_id"`
	Nome           string     `gorm:"size:100;not null" json:"nome"`
	CPFCNPJ        string     `gorm:"column:cpf_cnpj;size:18" json:"cpf_cnpj"`
	Telefone       string     `gorm:"size:15" json:"telefone"`
	Email          string     `gorm:"size:100" json:"email"`
	Endereco       JSON       `gorm:"type:jsonb" json:"endereco"`
	DataNascimento *time.Time `gorm:"type:date" json:"data_nascimento"`
	Observacoes    string     `gorm:"type:text" json:"observacoes"`
	Ativo          bool       `gorm:"default:true" json:"ativo"`
	CreatedAt      time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"index" json:"deleted_at,omitempty"`

	// Relacionamentos
	Pets   []Pets   `gorm:"foreignKey:ClientesID" json:"pets,omitempty"`
	Vendas []Vendas `gorm:"foreignKey:ClientesID" json:"vendas,omitempty"`
}
