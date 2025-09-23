package models

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	ID             uint       `gorm:"primaryKey" json:"id"`
	ClienteID      uint       `gorm:"not null;index" json:"cliente_id"`
	Nome           string     `gorm:"size:50;not null" json:"nome"`
	Especie        string     `gorm:"size:20" json:"especie"`
	Raca           string     `gorm:"size:50" json:"raca"`
	Sexo           string     `gorm:"size:1" json:"sexo"`
	DataNascimento *time.Time `gorm:"type:date" json:"data_nascimento"`
	Peso           float64    `gorm:"type:decimal(5,2)" json:"peso"`
	Cor            string     `gorm:"size:30" json:"cor"`
	Observacoes    string     `gorm:"type:text" json:"observacoes"`
	FotoURL        string     `gorm:"size:255" json:"foto_url"`

	// Relacionamentos
	Cliente      Cliente       `gorm:"foreignKey:ClienteID" json:"cliente,omitempty"`
	Agendamentos []Agendamento `gorm:"foreignKey:PetID" json:"agendamentos,omitempty"`
	Prontuarios  []Prontuario  `gorm:"foreignKey:PetID" json:"prontuarios,omitempty"`
}
