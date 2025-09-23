// models/entities/log.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type LogSistema struct {
	gorm.Model
	ID              uint      `gorm:"primaryKey" json:"id"`
	EmpresaID       *uint     `gorm:"index" json:"empresa_id"`
	UsuarioID       *uint     `gorm:"index" json:"usuario_id"`
	Modulo          string    `gorm:"size:50;not null" json:"modulo"`
	Acao            string    `gorm:"size:50;not null" json:"acao"`
	Descricao       string    `gorm:"type:text" json:"descricao"`
	TabelaAfetada   string    `gorm:"size:50" json:"tabela_afetada"`
	RegistroID      *uint     `gorm:"index" json:"registro_id"`
	DadosAnteriores JSON      `gorm:"type:jsonb" json:"dados_anteriores"`
	DadosNovos      JSON      `gorm:"type:jsonb" json:"dados_novos"`
	IPCliente       string    `gorm:"size:45" json:"ip_cliente"`
	UserAgent       string    `gorm:"type:text" json:"user_agent"`
	NivelLog        string    `gorm:"size:20;default:'INFO'" json:"nivel_log"`
	CreatedAt       time.Time `json:"created_at"`

	// Relacionamentos
	Empresa *Empresa `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	Usuario *Usuario `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`
}
