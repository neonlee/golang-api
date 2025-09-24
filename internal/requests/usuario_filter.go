package requests

type UsuarioFilter struct {
	Nome  string `form:"nome"`
	Email string `form:"email"`
	Ativo *bool  `form:"ativo"`
	Cargo string `form:"cargo"`
}
