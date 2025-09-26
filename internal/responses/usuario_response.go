package responses

import "petApi/internal/models"

type UsuarioResponse struct {
	Usuario   models.Usuarios `json:"usuario"`
	ID        uint            `json:"id"`
	Nome      string          `json:"nome"`
	Email     string          `json:"email"`
	Ativo     bool            `json:"ativo"`
	Cargo     string          `json:"cargo"`
	EmpresaID uint            `json:"empresa_id"`
	Empresa   string          `json:"empresa"`
	Perfis    []models.Perfil `json:"perfis"`
}
