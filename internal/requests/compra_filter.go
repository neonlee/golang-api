package requests

type CompraFilter struct {
	FornecedorID *uint
	DataInicio   string // Formato "YYYY-MM-DD"
	DataFim      string // Formato "YYYY-MM-DD"
	Status       string
}
