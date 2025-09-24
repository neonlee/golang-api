package responses

type ProdutoVendas struct {
	ProdutoID    uint    `json:"produto_id"`
	Nome         string  `json:"nome"`
	TotalVendido int     `json:"total_vendido"`
	TotalValor   float64 `json:"total_valor"`
}
