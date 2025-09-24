package responses

import "petApi/internal/models"

type LoginResponse struct {
	Token      string         `json:"token"`
	UsuarioID  uint           `json:"usuario_id"`
	Usuario    models.Usuario `json:"usuario"`
	ExpiraEm   int64          `json:"expira_em"`
	Empresa    models.Empresa `json:"empresa"`
	Permissoes []string       `json:"permissoes"`
}
