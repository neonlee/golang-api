package requests

type VendaFilter struct {
	DataInicio     string `json:"data_inicio"`
	DataFim        string `json:"data_fim"`
	ClientesID     *uint  `json:"cliente_id,omitempty"`
	Status         string `json:"status,omitempty"`
	UsuarioID      *uint  `json:"usuario_id,omitempty"`
	Pagina         int    `json:"pagina,omitempty"`
	ItensPorPagina int    `json:"itens_por_pagina,omitempty"`
	OrdenarPor     string `json:"ordenar_por,omitempty"`
	Ordem          string `json:"ordem,omitempty"` // "asc" ou "desc"
}
