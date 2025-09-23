// models/entities/cliente.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	ID             uint       `gorm:"primaryKey" json:"id"`
	EmpresaID      uint       `gorm:"not null;index" json:"empresa_id"`
	Nome           string     `gorm:"size:100;not null" json:"nome"`
	CPFCNPJ        string     `gorm:"size:18" json:"cpf_cnpj"`
	Telefone       string     `gorm:"size:15" json:"telefone"`
	Email          string     `gorm:"size:100" json:"email"`
	Endereco       JSON       `gorm:"type:jsonb" json:"endereco"`
	DataNascimento *time.Time `gorm:"type:date" json:"data_nascimento"`
	Observacoes    string     `gorm:"type:text" json:"observacoes"`
	Ativo          bool       `gorm:"default:true" json:"ativo"`

	// Relacionamentos
	Empresa Empresa `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	Pets    []Pet   `gorm:"foreignKey:ClienteID" json:"pets,omitempty"`
	Vendas  []Venda `gorm:"foreignKey:ClienteID" json:"vendas,omitempty"`
}
