package responses

type RelatorioEstoque struct {
	ProdutoID         uint   `json:"produto_id"`
	Nome              string `json:"nome"`
	Categoria         string `json:"categoria"`
	Fornecedor        string `json:"fornecedor"`
	QuantidadeEstoque int    `json:"quantidade_estoque"`
}
