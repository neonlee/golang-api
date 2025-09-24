package responses

type ServicoUtilizado struct {
	ServicoID     uint    `json:"servico_id"`
	Nome          string  `json:"nome"`
	Quantidade    int     `json:"quantidade"`
	TotalFaturado float64 `json:"total_faturado"`
}
