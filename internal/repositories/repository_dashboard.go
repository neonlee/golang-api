package repositories

import (
	"encoding/json"
	"petApi/internal/models"
	"petApi/internal/responses"
	"time"

	"gorm.io/gorm"
)

type DashboardRepository interface {
	GetResumoVendas(empresaID uint) (*responses.ResumoVendas, error)
	GetResumoFinanceiro(empresaID uint) (*responses.ResumoFinanceiro, error)
	GetProximosAgendamentos(empresaID uint) ([]models.Agendamentos, error)
	GetAlertasEstoque(empresaID uint) ([]responses.AlertaEstoque, error)
	GetMetricasGerais(empresaID uint) (*responses.MetricasGerais, error)
}

type dashboardRepository struct {
	db *gorm.DB // Assuming you're using GORM for database operations
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db: db}
}

func (r *dashboardRepository) GetResumoVendas(empresaID uint) (*responses.ResumoVendas, error) {
	// Implement the logic to fetch sales summary from the database
	// This is a placeholder implementation
	var resumo responses.ResumoVendas
	query := `
 WITH vendas_por_mes AS (
        SELECT 
            EXTRACT(YEAR FROM data_venda) AS ano,
            EXTRACT(MONTH FROM data_venda) AS mes,
            COUNT(*) as total_vendas_mes,
            SUM(valor_total) as total_valor_mes
        FROM vendas
        GROUP BY ano, mes
        ORDER BY ano, mes
    ),
    resumo_geral AS (
        SELECT 
            COUNT(*) as total_vendas,
            SUM(valor_total) as total_valor,
            COUNT(DISTINCT TO_CHAR(data_venda, 'YYYY-MM')) as quantidade_mes
        FROM vendas
    ),
    resumo_mes_atual AS (
        SELECT 
            COALESCE(SUM(valor_total), 0) as total_mes
        FROM vendas
        WHERE EXTRACT(YEAR FROM data_venda) = EXTRACT(YEAR FROM CURRENT_DATE)
          AND EXTRACT(MONTH FROM data_venda) = EXTRACT(MONTH FROM CURRENT_DATE)
    ),
    resumo_semana AS (
        SELECT 
            COALESCE(SUM(valor_total), 0) as total_semana
        FROM vendas
        WHERE data_venda >= DATE_TRUNC('week', CURRENT_DATE)
    ),
    resumo_hoje AS (
        SELECT 
            COALESCE(SUM(valor_total), 0) as total_hoje
        FROM vendas
        WHERE data_venda::DATE = CURRENT_DATE
    ),
    produtos_mais_vendidos AS (
         SELECT 
            COUNT(v.id) as quantidade_vendida,
            SUM(v.valor_total) as total_vendido
        FROM vendas v
        INNER JOIN venda_itens p ON v.id = p.id
        LIMIT 10
    )
    SELECT 
        (SELECT total_vendas FROM resumo_geral) as total_vendas,
        (SELECT total_valor FROM resumo_geral) as total_valor,
        (SELECT quantidade_mes FROM resumo_geral) as quantidade_mes,
        (SELECT total_mes FROM resumo_mes_atual) as total_mes,
        (SELECT total_semana FROM resumo_semana) as total_semana,
        (SELECT total_hoje FROM resumo_hoje) as total_hoje
    `

	err := r.db.Raw(query).Scan(&resumo).Error
	if err != nil {
		return nil, err
	}

	resumoCompleto := &responses.ResumoVendas{
		TotalVendas:   resumo.TotalVendas,
		TotalValor:    resumo.TotalValor,
		QuantidadeMes: resumo.QuantidadeMes,
		TotalMes:      resumo.TotalMes,
		TotalSemana:   resumo.TotalSemana,
		TotalHoje:     resumo.TotalHoje,
	}

	return resumoCompleto, nil
}
func (r *dashboardRepository) GetResumoFinanceiro(empresaID uint) (*responses.ResumoFinanceiro, error) {
	var resumoRow struct {
		ReceitaTotal      float64 `gorm:"column:receita_total"`
		DespesaTotal      float64 `gorm:"column:despesa_total"`
		LucroBruto        float64 `gorm:"column:lucro_bruto"`
		MargemLucro       float64 `gorm:"column:margem_lucro"`
		ReceitaMesAtual   float64 `gorm:"column:receita_mes_atual"`
		DespesaMesAtual   float64 `gorm:"column:despesa_mes_atual"`
		LucroMesAtual     float64 `gorm:"column:lucro_mes_atual"`
		ReceitaAnoAtual   float64 `gorm:"column:receita_ano_atual"`
		DespesaAnoAtual   float64 `gorm:"column:despesa_ano_atual"`
		MediaMensal       float64 `gorm:"column:media_mensal"`
		ProjecaoAnual     float64 `gorm:"column:projecao_anual"`
		CategoriasDespesa string  `gorm:"column:categorias_despesa"`
		FluxoMensal       string  `gorm:"column:fluxo_mensal"`
	}
	query := `
    WITH receita_total AS (
        SELECT COALESCE(SUM(valor_total), 0) as total
        FROM vendas
    ),
    despesa_total AS (
        SELECT COALESCE(SUM(valor), 0) as total
        FROM despesas
    ),
    receita_mes_atual AS (
        SELECT COALESCE(SUM(valor_total), 0) as total
        FROM vendas
        WHERE EXTRACT(YEAR FROM data_venda) = EXTRACT(YEAR FROM CURRENT_DATE)
          AND EXTRACT(MONTH FROM data_venda) = EXTRACT(MONTH FROM CURRENT_DATE)
    ),
    despesa_mes_atual AS (
        SELECT COALESCE(SUM(valor), 0) as total
        FROM despesas
        WHERE EXTRACT(YEAR FROM data_despesa) = EXTRACT(YEAR FROM CURRENT_DATE)
          AND EXTRACT(MONTH FROM data_despesa) = EXTRACT(MONTH FROM CURRENT_DATE)
    ),
    receita_ano_atual AS (
        SELECT COALESCE(SUM(valor_total), 0) as total
        FROM vendas
        WHERE EXTRACT(YEAR FROM data_venda) = EXTRACT(YEAR FROM CURRENT_DATE)
    ),
    despesa_ano_atual AS (
        SELECT COALESCE(SUM(valor), 0) as total
        FROM despesas
        WHERE EXTRACT(YEAR FROM data_despesa) = EXTRACT(YEAR FROM CURRENT_DATE)
    ),
    media_mensal AS (
        SELECT COALESCE(AVG(receita_mensal), 0) as media
        FROM (
            SELECT 
                TO_CHAR(data_venda, 'YYYY-MM') as mes,
                SUM(valor_total) as receita_mensal
            FROM vendas
            GROUP BY TO_CHAR(data_venda, 'YYYY-MM')
        ) meses
    ),
    categorias_despesa AS (
        SELECT 
            categoria,
            SUM(valor) as total,
            ROUND((SUM(valor) / (SELECT total FROM despesa_total) * 100), 2) as percentual
        FROM despesas
        GROUP BY categoria
        ORDER BY total DESC
    ),
    fluxo_mensal AS (
        SELECT 
            TO_CHAR(periodo.mes, 'YYYY-MM') as mes_ano,
            COALESCE(v.receita, 0) as receita,
            COALESCE(d.despesa, 0) as despesa,
            COALESCE(v.receita, 0) - COALESCE(d.despesa, 0) as lucro
        FROM (
            SELECT DATE_TRUNC('month', generate_series(
                DATE_TRUNC('month', CURRENT_DATE) - INTERVAL '11 months',
                DATE_TRUNC('month', CURRENT_DATE),
                '1 month'
            )) as mes
        ) periodo
        LEFT JOIN (
            SELECT 
                DATE_TRUNC('month', data_venda) as mes,
                SUM(valor_total) as receita
            FROM vendas
            GROUP BY DATE_TRUNC('month', data_venda)
        ) v ON periodo.mes = v.mes
        LEFT JOIN (
            SELECT 
                DATE_TRUNC('month', data_despesa) as mes,
                SUM(valor) as despesa
            FROM despesas
            GROUP BY DATE_TRUNC('month', data_despesa)
        ) d ON periodo.mes = d.mes
        ORDER BY periodo.mes
    )
    SELECT 
        rt.total as receita_total,
        dt.total as despesa_total,
        (rt.total - dt.total) as lucro_bruto,
        CASE 
            WHEN rt.total > 0 THEN ROUND(((rt.total - dt.total) / rt.total * 100), 2)
            ELSE 0 
        END as margem_lucro,
        rma.total as receita_mes_atual,
        dma.total as despesa_mes_atual,
        (rma.total - dma.total) as lucro_mes_atual,
        rya.total as receita_ano_atual,
        dya.total as despesa_ano_atual,
        mm.media as media_mensal,
        (mm.media * 12) as projecao_anual,
        (SELECT JSON_AGG(
            JSON_BUILD_OBJECT(
                'categoria', categoria,
                'total', total,
                'percentual', percentual
            )
        ) FROM categorias_despesa) as categorias_despesa,
        (SELECT JSON_AGG(
            JSON_BUILD_OBJECT(
                'mes_ano', mes_ano,
                'receita', receita,
                'despesa', despesa,
                'lucro', lucro
            )
        ) FROM fluxo_mensal) as fluxo_mensal
    FROM receita_total rt, despesa_total dt, receita_mes_atual rma,
         despesa_mes_atual dma, receita_ano_atual rya, despesa_ano_atual dya,
         media_mensal mm;
		 `
	err := r.db.Raw(query).Scan(&resumoRow).Error
	if err != nil {
		return nil, err
	}
	// Converter JSON strings para structs
	var categorias []responses.CategoriaDespesa
	var fluxoMensal []responses.FluxoMensal

	if resumoRow.CategoriasDespesa != "" {
		err = json.Unmarshal([]byte(resumoRow.CategoriasDespesa), &categorias)
		if err != nil {
			return nil, err
		}
	}

	if resumoRow.FluxoMensal != "" {
		err = json.Unmarshal([]byte(resumoRow.FluxoMensal), &fluxoMensal)
		if err != nil {
			return nil, err
		}
	}

	resumo := &responses.ResumoFinanceiro{
		ReceitaTotal:      resumoRow.ReceitaTotal,
		DespesaTotal:      resumoRow.DespesaTotal,
		LucroBruto:        resumoRow.LucroBruto,
		MargemLucro:       resumoRow.MargemLucro,
		ReceitaMesAtual:   resumoRow.ReceitaMesAtual,
		DespesaMesAtual:   resumoRow.DespesaMesAtual,
		LucroMesAtual:     resumoRow.LucroMesAtual,
		ReceitaAnoAtual:   resumoRow.ReceitaAnoAtual,
		DespesaAnoAtual:   resumoRow.DespesaAnoAtual,
		MediaMensal:       resumoRow.MediaMensal,
		ProjecaoAnual:     resumoRow.ProjecaoAnual,
		CategoriasDespesa: categorias,
		FluxoMensal:       fluxoMensal,
	}

	return resumo, nil
}
func (r *dashboardRepository) GetProximosAgendamentos(empresaID uint) ([]models.Agendamentos, error) {
	var agendamentos []models.Agendamentos

	hoje := time.Now()
	inicioDia := time.Date(hoje.Year(), hoje.Month(), hoje.Day(), 0, 0, 0, 0, hoje.Location())
	fimDia := time.Date(hoje.Year(), hoje.Month(), hoje.Day(), 23, 59, 59, 0, hoje.Location())

	err := r.db.
		Preload("Clientes", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, nome, telefone")
		}).
		Preload("TipoServico", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, nome")
		}).
		Where("data_agendamento BETWEEN ? AND ?", inicioDia, fimDia).
		Where("status IN ?", []string{"agendado", "confirmado"}).
		Order("data_agendamento ASC").
		Find(&agendamentos).Error

	return agendamentos, err
}

func (r *dashboardRepository) GetAlertasEstoque(empresaID uint) ([]responses.AlertaEstoque, error) {
	var alertas []responses.AlertaEstoque
	if err := r.db.Raw("SELECT e.id, p.nome, quantidade FROM movimentacao_estoques e, produtos p WHERE empresa_id = ? AND e.quantidade < p.estoque_minimo", empresaID).Scan(&alertas).Error; err != nil {
		return nil, err
	}
	return alertas, nil
}
func (r *dashboardRepository) GetMetricasGerais(empresaID uint) (*responses.MetricasGerais, error) {
	var metricas responses.MetricasGerais
	if err := r.db.Raw("SELECT COUNT(*) as total_clientes, COUNT(*) as total_produtos FROM clientes WHERE empresa_id = ?", empresaID).Scan(&metricas).Error; err != nil {
		return nil, err
	}
	return &metricas, nil
}
