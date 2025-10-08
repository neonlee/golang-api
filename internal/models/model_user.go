// models/entities/usuario.go
package models

import (
	"time"
)

type Usuarios struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	EmpresaID   uint       `gorm:"not null;index" json:"empresa_id"`
	Nome        string     `gorm:"size:100;not null" json:"nome"`
	Email       string     `gorm:"size:100;uniqueIndex;not null" json:"email"`
	SenhaHash   string     `gorm:"size:255;not null" json:"-"`
	Telefone    string     `gorm:"size:15" json:"telefone"`
	Cargo       string     `gorm:"size:50" json:"cargo"`
	FotoURL     string     `gorm:"size:255" json:"foto_url"`
	Ativo       bool       `gorm:"default:true" json:"ativo"`
	UltimoLogin *time.Time `json:"ultimo_login"`

	// Relacionamentos
	Empresa       Empresas               `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	UsuarioPerfis []UsuarioPerfis        `gorm:"foreignKey:UsuarioID" json:"usuario_perfis,omitempty"`
	Vendas        []Vendas               `gorm:"foreignKey:UsuarioID" json:"vendas,omitempty"`
	Movimentacoes []MovimentacaoEstoques `gorm:"foreignKey:UsuarioID" json:"movimentacoes,omitempty"`
}

type Perfis struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	EmpresaID uint   `gorm:"not null;index" json:"empresa_id"`
	Nome      string `gorm:"size:50;not null" json:"nome"`
	Descricao string `gorm:"type:text" json:"descricao"`
	IsAdmin   bool   `gorm:"default:false" json:"is_admin"`

	// Relacionamentos
	Empresa       Empresas        `gorm:"foreignKey:EmpresaID" json:"empresa,omitempty"`
	Permissoes    []Permissoes    `gorm:"foreignKey:PerfilID" json:"permissoes,omitempty"`
	UsuarioPerfis []UsuarioPerfis `gorm:"foreignKey:PerfilID" json:"usuario_perfis,omitempty"`
}

type Permissoes struct {
	ID                 uint `gorm:"primaryKey" json:"id"`
	PerfilID           uint `gorm:"not null;index" json:"perfil_id"`
	ModuloID           uint `gorm:"not null;index" json:"modulo_id"`
	PodeVisualizar     bool `gorm:"default:false" json:"pode_visualizar"`
	PodeEditar         bool `gorm:"default:false" json:"pode_editar"`
	PodeExcluir        bool `gorm:"default:false" json:"pode_excluir"`
	PodeGerarRelatorio bool `gorm:"default:false" json:"pode_gerar_relatorio"`

	// Relacionamentos
	Perfil Perfis `gorm:"foreignKey:PerfilID" json:"perfil,omitempty"`
	Modulo Modulo `gorm:"foreignKey:ModuloID" json:"modulo,omitempty"`
}

type UsuarioPerfis struct {
	UsuarioID uint      `gorm:"primaryKey" json:"usuario_id"`
	PerfilID  uint      `gorm:"primaryKey" json:"perfil_id"`
	CreatedAt time.Time `json:"created_at"`

	// Relacionamentos
	Usuario Usuarios `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`
	Perfil  Perfis   `gorm:"foreignKey:PerfilID" json:"perfil,omitempty"`
}

type Modulo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Nome      string `gorm:"size:50;not null" json:"nome"`
	Descricao string `gorm:"type:text" json:"descricao"`
	Categoria string `gorm:"size:30" json:"categoria"`
	Icone     string `gorm:"size:30" json:"icone"`
	Ordem     int    `gorm:"default:0" json:"ordem"`
	Ativo     bool   `gorm:"default:true" json:"ativo"`

	Permissoes []Permissoes `gorm:"foreignKey:ModuloID" json:"permissoes,omitempty"`
}
