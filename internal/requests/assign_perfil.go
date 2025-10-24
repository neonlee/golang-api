package requests

type AssignPerfilRequest struct {
	UsuarioID uint   `json:"usuario_id"`
	Perfil    []uint `json:"perfil"`
}
