package responses

type ResumoFinanceiro struct {
	TotalReceitas float64 `json:"total_receitas"`
	TotalDespesas float64 `json:"total_despesas"`
	Saldo         float64 `json:"saldo"`
}
