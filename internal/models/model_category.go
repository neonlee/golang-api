package models

import (
	"time"
)

type CategoriaTipo string

const (
	CategoriaTipoProduto       CategoriaTipo = "produto"
	CategoriaTipoServico       CategoriaTipo = "servico"
	CategoriaTipoFornecedor    CategoriaTipo = "fornecedor"
	CategoriaTipoDespesa       CategoriaTipo = "despesa"
	CategoriaTipoProblemaSaude CategoriaTipo = "problema_saude"
	CategoriaTipoVacina        CategoriaTipo = "vacina"
	CategoriaTipoExame         CategoriaTipo = "exame"
)

type Categoryd struct {
	PetshopId int           `gorm:"column:petshop_id" json:"petshop_id"`
	Id        int           `gorm:"column:id;primaryKey" json:"id"`
	Order     int           `gorm:"column:ordem" json:"order"`
	Nome      string        `gorm:"column:nome" json:"name"`
	Ativo     bool          `gorm:"default:true" json:"active"`
	Tipo      CategoriaTipo `gorm:"type:varchar(20);not null" json:"tipo"`
	Descricao string        `gorm:"column:descricao" json:"description"`
	CreatedAt time.Time     `gorm:"column:create_at" json:"created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at" json:"updated_at"`
}
