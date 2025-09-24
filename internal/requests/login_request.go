package requests

type LoginRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Senha     string `json:"senha" binding:"required,min=6"`
	EmpresaID uint   `json:"empresa_id" binding:"required"`
}
