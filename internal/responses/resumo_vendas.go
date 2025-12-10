package responses

type ResumoVendas struct {
	TotalVendas   int     `json:"total_vendas"`
	TotalValor    float64 `json:"total_valor"`
	QuantidadeMes int     `json:"quantidade_mes"`
	TotalMes      float64 `json:"total_mes"`
	TotalSemana   float64 `json:"total_semana"`
	TotalHoje     float64 `json:"total_hoje"`
}
