package models

import (
	"time"

	"gorm.io/gorm"
)

type Despesa struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	Descricao       string     `gorm:"type:varchar(255);not null" json:"descricao"`
	Categoria       string     `gorm:"type:varchar(100);not null" json:"categoria"`
	Valor           float64    `gorm:"type:decimal(10,2);not null" json:"valor"`
	DataDespesa     time.Time  `gorm:"type:date;not null;default:CURRENT_DATE" json:"data_despesa"`
	DataVencimento  *time.Time `gorm:"type:date" json:"data_vencimento,omitempty"`
	DataPagamento   *time.Time `gorm:"type:date" json:"data_pagamento,omitempty"`
	Status          string     `gorm:"type:varchar(20);default:'pendente'" json:"status"`
	FormaPagamento  string     `gorm:"type:varchar(50)" json:"forma_pagamento,omitempty"`
	Observacoes     string     `gorm:"type:text" json:"observacoes,omitempty"`
	Fornecedor      string     `gorm:"type:varchar(255)" json:"fornecedor,omitempty"`
	NumeroDocumento string     `gorm:"type:varchar(100)" json:"numero_documento,omitempty"`
	TipoDespesa     string     `gorm:"type:varchar(50)" json:"tipo_despesa,omitempty"`
	CentroCusto     string     `gorm:"type:varchar(100)" json:"centro_custo,omitempty"`
	Recorrente      bool       `gorm:"default:false" json:"recorrente"`
	DiaVencimento   *int       `gorm:"check:dia_vencimento>=1 AND dia_vencimento<=31" json:"dia_vencimento,omitempty"`
	UsuarioID       *uint      `json:"usuario_id,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`

	// Relacionamentos
	Usuario *Usuarios `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`
}

// TableName especifica o nome da tabela
func (Despesa) TableName() string {
	return "despesas"
}

// Hooks do GORM
func (d *Despesa) BeforeCreate(tx *gorm.DB) error {
	if d.DataDespesa.IsZero() {
		d.DataDespesa = time.Now()
	}
	return nil
}

func (d *Despesa) BeforeUpdate(tx *gorm.DB) error {
	// Atualizar status baseado nas datas
	if d.DataVencimento != nil && d.DataPagamento == nil {
		if d.DataVencimento.Before(time.Now()) {
			d.Status = "atrasado"
		} else {
			d.Status = "pendente"
		}
	}
	if d.DataPagamento != nil {
		d.Status = "pago"
	}
	return nil
}
