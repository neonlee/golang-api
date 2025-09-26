// models/entities/prontuario.go
package models

import (
	"time"
)

type Prontuario struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	PetID         uint      `gorm:"not null;index" json:"pet_id"`
	VeterinarioID uint      `gorm:"not null;index" json:"veterinario_id"`
	DataConsulta  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"data_consulta"`
	Anamnese      string    `gorm:"type:text" json:"anamnese"`
	Diagnostico   string    `gorm:"type:text" json:"diagnostico"`
	Prescricao    JSON      `gorm:"type:jsonb" json:"prescricao"` // Array de medicamentos
	Observacoes   string    `gorm:"type:text" json:"observacoes"`
	PesoAtual     float64   `gorm:"type:decimal(5,2)" json:"peso_atual"`
	Temperatura   float64   `gorm:"type:decimal(4,2)" json:"temperatura"`

	// Relacionamentos
	Pet         Pets     `gorm:"foreignKey:PetID" json:"pet,omitempty"`
	Veterinario Usuarios `gorm:"foreignKey:VeterinarioID" json:"veterinario,omitempty"`
}

// Estrutura para a prescrição médica
type PrescricaoMedica struct {
	Medicamento string `json:"medicamento"`
	Dose        string `json:"dose"`
	Frequencia  string `json:"frequencia"`
	Duracao     string `json:"duracao"`
	Observacoes string `json:"observacoes,omitempty"`
}

// Estrutura para vacinas
type Vacina struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	PetID         uint      `gorm:"not null;index" json:"pet_id"`
	Nome          string    `gorm:"size:100;not null" json:"nome"`
	DataAplicacao time.Time `gorm:"type:date" json:"data_aplicacao"`
	DataProxima   time.Time `gorm:"type:date" json:"data_proxima"`
	VeterinarioID uint      `gorm:"not null;index" json:"veterinario_id"`
	Lote          string    `gorm:"size:50" json:"lote"`
	Fabricante    string    `gorm:"size:100" json:"fabricante"`
	Observacoes   string    `gorm:"type:text" json:"observacoes"`

	// Relacionamentos
	Pet         Pets     `gorm:"foreignKey:PetID" json:"pet,omitempty"`
	Veterinario Usuarios `gorm:"foreignKey:VeterinarioID" json:"veterinario,omitempty"`
}
