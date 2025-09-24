package responses

type AlertaEstoque struct {
	ProdutoID   uint   `json:"produto_id"`
	ProdutoNome string `json:"produto_nome"`
	Quantidade  int    `json:"quantidade"`
}
