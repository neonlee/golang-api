package requests

type ProdutoFilter struct {
	Nome             string `json:"nome"`
	Categoria        uint   `json:"categoria"`
	Fornecedor       uint   `json:"fornecedor"`
	Ativo            *bool  `json:"ativo"`
	EspecieDestinada string `json:"especie_destinada"`
	CategoriaID      *uint  `json:"categoria_id"`
}
