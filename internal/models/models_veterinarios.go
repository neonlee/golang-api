package models

import "time"

type MedicoVeterinario struct {
	ID                   uint      `gorm:"primaryKey" json:"id"`
	FuncionarioID        uint      `gorm:"not null" json:"funcionario_id"`
	CRMV                 string    `gorm:"size:20;not null;unique" json:"crmv"`
	CRMVUF               string    `gorm:"size:2;not null" json:"crmv_uf"`
	Especialidade        string    `gorm:"size:100" json:"especialidade"`
	FormacaoGraduacao    string    `gorm:"size:100" json:"formacao_graduacao"`
	InstituicaoGraduacao string    `gorm:"size:100" json:"instituicao_graduacao"`
	AnoFormacao          int       `json:"ano_formacao"`
	PosGraduacao         string    `gorm:"type:text" json:"pos_graduacao"`
	AreasAtuacao         string    `gorm:"type:text" json:"areas_atuacao"`
	Procedimentos        string    `gorm:"type:text" json:"procedimentos"`
	ConsultaValor        float64   `gorm:"type:numeric(8,2)" json:"consulta_valor"`
	PlantaoValor         float64   `gorm:"type:numeric(8,2)" json:"plantao_valor"`
	CirurgiaValor        float64   `gorm:"type:numeric(8,2)" json:"cirurgia_valor"`
	DisponivelPlantao    bool      `gorm:"default:true" json:"disponivel_plantao"`
	DisponivelConsulta   bool      `gorm:"default:true" json:"disponivel_consulta"`
	DisponivelCirurgia   bool      `gorm:"default:true" json:"disponivel_cirurgia"`
	FotoURL              string    `gorm:"size:255" json:"foto_url"`
	Observacoes          string    `gorm:"type:text" json:"observacoes"`
	Ativo                bool      `gorm:"default:true" json:"ativo"`
	CreatedAt            time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relacionamentos
	Funcionario      Funcionarios            `gorm:"foreignKey:FuncionarioID" json:"funcionario"`
	Especialidades   []MedicoEspecialidade   `gorm:"foreignKey:MedicoID" json:"especialidades,omitempty"`
	Disponibilidades []MedicoDisponibilidade `gorm:"foreignKey:MedicoID" json:"disponibilidades,omitempty"`
}

type MedicoEspecialidade struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	MedicoID      uint      `gorm:"not null" json:"medico_id"`
	Especialidade string    `gorm:"size:100;not null" json:"especialidade"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type MedicoDisponibilidade struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	MedicoID    uint   `gorm:"not null" json:"medico_id"`
	DiaSemana   string `gorm:"size:10;not null;check:dia_semana IN ('segunda','terca','quarta','quinta','sexta','sabado','domingo')" json:"dia_semana"`
	HoraInicio  string `gorm:"type:time;not null" json:"hora_inicio"`
	HoraFim     string `gorm:"type:time;not null" json:"hora_fim"`
	TipoServico string `gorm:"size:10;not null;check:tipo_servico IN ('consulta','cirurgia','plantao')" json:"tipo_servico"`
	Ativo       bool   `gorm:"default:true" json:"ativo"`
}
