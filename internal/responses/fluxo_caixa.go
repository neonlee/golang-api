package responses

import "time"

type FluxoCaixa struct {
	TotalEntradas float64   `json:"total_entradas"`
	TotalSaidas   float64   `json:"total_saidas"`
	SaldoFinal    float64   `json:"saldo_final"`
	PeriodoInicio time.Time `json:"periodo_inicio"`
	PeriodoFim    time.Time `json:"periodo_fim"`
	Receitas      float64   `json:"receitas"`
	Despesas      float64   `json:"despesas"`
	Saldo         float64   `json:"saldo"`
}
