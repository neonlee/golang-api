package responses

type ResumoCompras struct {
	TotalCompras      float64 `json:"total_compras"`
	TotalItens        int     `json:"total_itens"`
	TotalFornecedores int     `json:"total_fornecedores"`
	Periodo           string  `json:"periodo"`
}
