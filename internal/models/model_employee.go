package models

import (
	"encoding/json"
	"time"
)

type Funcionarios struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	Nome         string     `json:"nome"`
	CPF          string     `json:"cpf"`
	Telefone     string     `json:"telefone"`
	Cargo        string     `json:"cargo"`
	DataAdmissao time.Time  `json:"data_admissao"`
	DataDemissao *time.Time `json:"data_demissao,omitempty"` // Use pointer para campos opcionais
	Salario      float64    `json:"salario"`
	Comissao     float64    `json:"comissao"`
	Ativo        bool       `json:"ativo"`
	CreatedAt    time.Time  `json:"created_at"`
	UsuarioID    uint       `gorm:"not null;index" json:"usuario_id"`
}

// Custom JSON unmarshal para DataAdmissao
func (f *Funcionarios) UnmarshalJSON(data []byte) error {
	type Alias Funcionarios
	aux := &struct {
		DataAdmissao string `json:"data_admissao"`
		DataDemissao string `json:"data_demissao,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(f),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse da data de admissão
	if aux.DataAdmissao != "" {
		parsedTime, err := time.Parse("2006-01-02", aux.DataAdmissao)
		if err != nil {
			return err
		}
		f.DataAdmissao = parsedTime
	}

	// Parse da data de demissão (opcional)
	if aux.DataDemissao != "" {
		parsedTime, err := time.Parse("2006-01-02", aux.DataDemissao)
		if err != nil {
			return err
		}
		f.DataDemissao = &parsedTime
	}

	return nil
}
