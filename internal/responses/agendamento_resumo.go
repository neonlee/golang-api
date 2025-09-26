package responses

type AgendamentoResumo struct {
	ID           uint   `json:"id"`
	ClientesNome string `json:"cliente_nome"`
	DataHora     string `json:"data_hora"`
}
