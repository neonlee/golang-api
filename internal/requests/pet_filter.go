package requests

type PetFilter struct {
	Nome       string
	Tipo       string
	Raca       string
	ClientesID uint
	EmpresaID  uint
	Pagina     int
	TamanhoPag int
	Especie    string
}
