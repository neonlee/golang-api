// models/entities/empresa.go
package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Empresas struct {
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
	Plano    Planos     `gorm:"foreignKey:PlanoID" json:"plano,omitempty"`
	Usuarios []Usuarios `gorm:"foreignKey:EmpresaID" json:"usuarios,omitempty"`
}

type Planos struct {
	ID                 uint    `gorm:"primaryKey" json:"id"`
	Nome               string  `gorm:"size:50;not null" json:"nome"`
	Descricao          string  `gorm:"type:text" json:"descricao"`
	ValorMensal        float64 `gorm:"type:decimal(10,2)" json:"valor_mensal"`
	ModulosDisponiveis JSON    `gorm:"type:jsonb" json:"modulos_disponiveis"`
	LimiteUsuarios     int     `gorm:"default:1" json:"limite_usuarios"`
	LimiteEmpresas     int     `gorm:"default:1" json:"limite_empresas"`
	Ativo              bool    `gorm:"default:true" json:"ativo"`

	Empresas []Empresas `gorm:"foreignKey:PlanoID" json:"empresas,omitempty"`
}
type JSON json.RawMessage

// Value implementa driver.Valuer
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 || string(j) == "null" {
		return nil, nil
	}
	return string(j), nil
}

// Scan implementa sql.Scanner
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		if len(v) > 0 {
			*j = append((*j)[0:0], v...)
		} else {
			*j = nil
		}
		return nil
	case string:
		if v != "" {
			*j = JSON(v)
		} else {
			*j = nil
		}
		return nil
	default:
		return errors.New("invalid scan source for JSON")
	}
}

// MarshalJSON implementa json.Marshaler
func (j JSON) MarshalJSON() ([]byte, error) {
	if len(j) == 0 {
		return []byte("null"), nil
	}
	return j, nil
}

// UnmarshalJSON implementa json.Unmarshaler - CORREÇÃO AQUI
func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("JSON receiver is nil")
	}

	// Verifica se é null
	if string(data) == "null" {
		*j = nil
		return nil
	}

	// Valida se o JSON é válido
	if !json.Valid(data) {
		return errors.New("invalid JSON")
	}

	// Faz uma cópia dos dados
	*j = append((*j)[0:0], data...)
	return nil
}

// Helper methods
func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}

func (j JSON) String() string {
	return string(j)
}
