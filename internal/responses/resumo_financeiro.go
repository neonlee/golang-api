package responses

type ResumoFinanceiro struct {
	ReceitaTotal      float64            `json:"receita_total"`
	DespesaTotal      float64            `json:"despesa_total"`
	LucroBruto        float64            `json:"lucro_bruto"`
	MargemLucro       float64            `json:"margem_lucro"`
	ReceitaMesAtual   float64            `json:"receita_mes_atual"`
	DespesaMesAtual   float64            `json:"despesa_mes_atual"`
	LucroMesAtual     float64            `json:"lucro_mes_atual"`
	ReceitaAnoAtual   float64            `json:"receita_ano_atual"`
	DespesaAnoAtual   float64            `json:"despesa_ano_atual"`
	MediaMensal       float64            `json:"media_mensal"`
	ProjecaoAnual     float64            `json:"projecao_anual"`
	CategoriasDespesa []CategoriaDespesa `json:"categorias_despesa"`
	FluxoMensal       []FluxoMensal      `json:"fluxo_mensal"`
}

type CategoriaDespesa struct {
	Categoria  string  `json:"categoria"`
	Total      float64 `json:"total"`
	Percentual float64 `json:"percentual"`
}

type FluxoMensal struct {
	MesAno  string  `json:"mes_ano"`
	Receita float64 `json:"receita"`
	Despesa float64 `json:"despesa"`
	Lucro   float64 `json:"lucro"`
}
