package models

import (
	"time"
)

// TipoAcesso define os tipos de acesso (Admin, Gerente, Operador, etc.)
type TipoAcesso struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Nome      string    `gorm:"unique;not null" json:"nome"`
	Descricao string    `json:"descricao"`
	Modulos   []Modulo  `gorm:"many2many:tipo_acesso_modulos;joinForeignKey:TipoAcessoID;joinReferences:ModuloID" json:"modulos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Modulo representa um módulo do sistema
type Modulo struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Nome      string    `gorm:"unique;not null" json:"nome"`
	Descricao string    `json:"descricao"`
	Rota      string    `json:"rota"`                   // Rota principal do módulo
	Icone     string    `json:"icone"`                  // Ícone para o menu
	Ordem     int       `gorm:"default:0" json:"ordem"` // Ordem no menu
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TipoAcessoModulo é a tabela de junção com permissões específicas
type TipoAcessoModulo struct {
	TipoAcessoID int       `gorm:"primaryKey" json:"tipo_acesso_id"`
	ModuloID     int       `gorm:"primaryKey" json:"modulo_id"`
	Ler          bool      `gorm:"default:true" json:"ler"`
	Criar        bool      `gorm:"default:false" json:"criar"`
	Editar       bool      `gorm:"default:false" json:"editar"`
	Excluir      bool      `gorm:"default:false" json:"excluir"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
