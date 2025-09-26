// models/entities/conta_bancaria.go
package models

import (
	"time"
)

type ContaBancarias struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	EmpresaID    uint    `gorm:"not null;index" json:"empresa_id"`
	Nome         string  `gorm:"size:50;not null" json:"nome"`
	Banco        string  `gorm:"size:50" json:"banco"`
	Agencia      string  `gorm:"size:10" json:"agencia"`
	Conta        string  `gorm:"size:15" json:"conta"`
	SaldoInicial float64 `gorm:"type:decimal(10,2);default:0" json:"saldo_inicial"`
	Tipo         string  `gorm:"size:20" json:"tipo"` // corrente, poupanca
	Ativo        bool    `gorm:"default:true" json:"ativo"`

	// Relacionamentos
	Empresa       Empresa                 `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	Movimentacoes []MovimentacaoBancarias `gorm:"foreignKey:ContaBancariaID" json:"movimentacoes,omitempty"`
}

type MovimentacaoBancarias struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	ContaBancariaID  uint      `gorm:"not null;index" json:"conta_bancaria_id"`
	TipoMovimentacao string    `gorm:"size:20" json:"tipo_movimentacao"` // entrada, saida, transferencia
	Valor            float64   `gorm:"type:decimal(10,2)" json:"valor"`
	DataMovimentacao time.Time `gorm:"type:date" json:"data_movimentacao"`
	Descricao        string    `gorm:"size:100" json:"descricao"`
	Categoria        string    `gorm:"size:50" json:"categoria"`
	ContaReceberID   *uint     `gorm:"index" json:"conta_receber_id"`
	ContaPagarID     *uint     `gorm:"index" json:"conta_pagar_id"`
	Observacoes      string    `gorm:"type:text" json:"observacoes"`
	UsuarioID        uint      `gorm:"not null;index" json:"usuario_id"`

	// Relacionamentos
	ContaBancaria ContaBancarias `gorm:"foreignKey:ContaBancariaID" json:"conta_bancaria,omitempty"`
	ContaReceber  ContaReceber   `gorm:"foreignKey:ContaReceberID" json:"conta_receber,omitempty"`
	ContaPagar    ContaPagar     `gorm:"foreignKey:ContaPagarID" json:"conta_pagar,omitempty"`
	Usuario       Usuarios       `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`
}

type FechamentoCaixas struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	EmpresaID      uint      `gorm:"not null;index" json:"empresa_id"`
	UsuarioID      uint      `gorm:"not null;index" json:"usuario_id"`
	DataFechamento time.Time `gorm:"type:date" json:"data_fechamento"`
	ValorInicial   float64   `gorm:"type:decimal(10,2)" json:"valor_inicial"`
	ValorFinal     float64   `gorm:"type:decimal(10,2)" json:"valor_final"`
	ValorVendas    float64   `gorm:"type:decimal(10,2)" json:"valor_vendas"`
	ValorRetiradas float64   `gorm:"type:decimal(10,2)" json:"valor_retiradas"`
	ValorDiferenca float64   `gorm:"type:decimal(10,2)" json:"valor_diferenca"`
	Observacoes    string    `gorm:"type:text" json:"observacoes"`
	Status         string    `gorm:"size:20;default:'aberto'" json:"status"`

	// Relacionamentos
	Empresa Empresa  `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	Usuario Usuarios `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`
}
