package responses

type DemonstrativoFinanceiro struct {
	Mes           int     `json:"mes"`
	Ano           int     `json:"ano"`
	TotalReceitas float64 `json:"total_receitas"`
	TotalDespesas float64 `json:"total_despesas"`
	Resultado     float64 `json:"resultados"`
	LucroLiquido  float64 `json:"lucro_liquido"`
	PeriodoInicio string  `json:"periodo_inicio"`
	PeriodoFim    string  `json:"periodo_fim"`
}
