// models/entities/cliente.go
package models

import (
	"time"
)

type Clientes struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	EmpresaID      uint      `gorm:"not null;index" json:"empresa_id"`
	Nome           string    `gorm:"size:100;not null" json:"nome"`
	CPFCNPJ        string    `gorm:"column:cpf_cnpj;size:18" json:"cpf_cnpj"`
	Telefone       string    `gorm:"size:15" json:"telefone"`
	Email          string    `gorm:"size:100" json:"email"`
	DataNascimento string    `gorm:"type:date" json:"data_nascimento"`
	Observacoes    string    `gorm:"type:text" json:"observacoes"`
	Ativo          bool      `gorm:"default:true" json:"ativo"`
	Logradouro     string    `gorm:"type:text" json:"logradouro"`
	Numero         string    `gorm:"type:text" json:"numero"`
	Bairro         string    `gorm:"type:text" json:"bairro"`
	Cidade         string    `gorm:"type:text" json:"cidade"`
	Estado         string    `gorm:"type:text" json:"estado"`
	CEP            string    `gorm:"type:text" json:"cep"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relacionamentos
	Pets   []Pets   `gorm:"foreignKey:ClientesID" json:"pets,omitempty"`
	Vendas []Vendas `gorm:"foreignKey:ClientesID" json:"vendas,omitempty"`
}
