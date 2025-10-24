package requests

type UpdateSenhaRequest struct {
	UsuarioID     uint   `json:"usuario_id"`
	SenhaAtual    string `json:"senha_atual"`
	NovaSenha     string `json:"nova_senha"`
	ConfirmaSenha string `json:"confirma_senha"`
}
