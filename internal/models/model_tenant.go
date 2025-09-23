// models/entities/empresa.go
package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Empresa struct {
	gorm.Model
	ID             uint           `gorm:"primaryKey" json:"id"`
	NomeFantasia   string         `gorm:"size:100;not null" json:"nome_fantasia"`
	RazaoSocial    string         `gorm:"size:100" json:"razao_social"`
	CNPJ           string         `gorm:"size:18;uniqueIndex" json:"cnpj"`
	Telefone       string         `gorm:"size:15" json:"telefone"`
	Email          string         `gorm:"size:100" json:"email"`
	Endereco       JSON           `gorm:"type:jsonb" json:"endereco"`
	PlanoID        uint           `gorm:"not null" json:"plano_id"`
	DataAssinatura time.Time      `gorm:"type:date" json:"data_assinatura"`
	Status         string         `gorm:"size:20;default:'ativo'" json:"status"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relacionamentos
	Plano    Plano     `gorm:"foreignKey:PlanoID" json:"plano,omitempty"`
	Usuarios []Usuario `gorm:"foreignKey:EmpresaID" json:"usuarios,omitempty"`
}

type Plano struct {
	gorm.Model
	ID                 uint    `gorm:"primaryKey" json:"id"`
	Nome               string  `gorm:"size:50;not null" json:"nome"`
	Descricao          string  `gorm:"type:text" json:"descricao"`
	ValorMensal        float64 `gorm:"type:decimal(10,2)" json:"valor_mensal"`
	ModulosDisponiveis JSON    `gorm:"type:jsonb" json:"modulos_disponiveis"`
	LimiteUsuarios     int     `gorm:"default:1" json:"limite_usuarios"`
	LimiteEmpresas     int     `gorm:"default:1" json:"limite_empresas"`
	Ativo              bool    `gorm:"default:true" json:"ativo"`

	Empresas []Empresa `gorm:"foreignKey:PlanoID" json:"empresas,omitempty"`
}

// JSON type para campos JSONB
type JSON json.RawMessage

func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return string(j), nil
}

func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid scan source for JSON")
	}
	*j = append((*j)[0:0], s...)
	return nil
}
