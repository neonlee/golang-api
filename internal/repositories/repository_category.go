package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	connection *gorm.DB
}

func NewCategoryRepository(connection *gorm.DB) CategoryRepository {
	return CategoryRepository{connection: connection}
}

func (r *CategoryRepository) Create(user models.Category) (*models.Category, error) {

	err := r.connection.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *CategoryRepository) GetCategorys() (*[]models.Category, error) {
	var clientes []models.Category

	err := r.connection.Find(&clientes).Error
	if err != nil {
		return nil, err
	}

	return &clientes, nil
}

func (r *CategoryRepository) GetCategory(id int) (*models.Category, error) {
	servico := &models.Category{}

	err := r.connection.Where("id = ?", id).First(&servico).Error
	if err != nil {
		return nil, err
	}

	return servico, nil
}

func (r *CategoryRepository) UpdateCategorys(id int, services models.Category) (*models.Category, error) {
	err := r.connection.Model(&models.Category{}).Where("id = ?", id).Updates(services).Error
	if err != nil {
		return nil, err
	}

	// Busca o cliente atualizado
	var updatedClient models.Category
	err = r.connection.First(&updatedClient, id).Error
	if err != nil {
		return nil, err
	}

	return &updatedClient, nil
}

func (r *CategoryRepository) DeleteCategorys(id int) (bool, error) {
	err := r.connection.Delete(&models.Pet{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

type servicoRepository struct {
	db *gorm.DB
}

func NewServicoRepository(db *gorm.DB) ServicoRepository {
	return &servicoRepository{db: db}
}

func (r *servicoRepository) CreateTipoServico(servico *entities.TipoServico) error {
	return r.db.Create(servico).Error
}

func (r *servicoRepository) GetTipoServicoByID(id uint) (*entities.TipoServico, error) {
	var servico entities.TipoServico
	err := r.db.First(&servico, id).Error
	return &servico, err
}

func (r *servicoRepository) UpdateTipoServico(servico *entities.TipoServico) error {
	return r.db.Save(servico).Error
}

func (r *servicoRepository) DeleteTipoServico(id uint) error {
	return r.db.Delete(&entities.TipoServico{}, id).Error
}

func (r *servicoRepository) ListTiposServico(empresaID uint, categoria string) ([]entities.TipoServico, error) {
	var servicos []entities.TipoServico

	query := r.db.Where("empresa_id = ?", empresaID)

	if categoria != "" {
		query = query.Where("categoria = ?", categoria)
	}

	err := query.Where("ativo = ?", true).Order("nome ASC").Find(&servicos).Error
	return servicos, err
}

func (r *servicoRepository) GetServicosMaisUtilizados(empresaID uint, limite int) ([]responses.ServicoUtilizado, error) {
	var resultados []responses.ServicoUtilizado

	err := r.db.
		Table("venda_itens").
		Select("tipos_servicos.nome, COUNT(*) as quantidade_utilizada, SUM(venda_itens.valor_total) as total_faturado").
		Joins("JOIN vendas ON venda_itens.venda_id = vendas.id").
		Joins("JOIN tipos_servicos ON venda_itens.tipo_servico_id = tipos_servicos.id").
		Where("vendas.empresa_id = ? AND venda_itens.tipo_servico_id IS NOT NULL", empresaID).
		Group("tipos_servicos.id, tipos_servicos.nome").
		Order("quantidade_utilizada DESC").
		Limit(limite).
		Scan(&resultados).Error

	return resultados, err
}

type compraRepository struct {
	db *gorm.DB
}

func NewCompraRepository(db *gorm.DB) CompraRepository {
	return &compraRepository{db: db}
}

func (r *compraRepository) Create(compra *entities.Compra, itens []entities.CompraItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Criar a compra
		if err := tx.Create(compra).Error; err != nil {
			return err
		}

		// Criar os itens da compra
		for i := range itens {
			itens[i].CompraID = compra.ID
			if err := tx.Create(&itens[i]).Error; err != nil {
				return err
			}

			// Atualizar estoque
			movimentacao := entities.MovimentacaoEstoque{
				ProdutoID:        itens[i].ProdutoID,
				TipoMovimentacao: "entrada",
				Quantidade:       itens[i].Quantidade,
				Motivo:           "Compra #" + string(compra.ID),
				UsuarioID:        compra.UsuarioID,
			}
			if err := tx.Create(&movimentacao).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *compraRepository) GetByID(id uint) (*entities.Compra, error) {
	var compra entities.Compra
	err := r.db.
		Preload("Fornecedor").
		Preload("Usuario").
		Preload("Itens").
		Preload("Itens.Produto").
		First(&compra, id).Error

	return &compra, err
}

func (r *compraRepository) Update(compra *entities.Compra) error {
	return r.db.Save(compra).Error
}

func (r *compraRepository) Cancelar(id uint, motivo string) error {
	return r.db.Model(&entities.Compra{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":      "cancelado",
			"observacoes": motivo,
		}).Error
}

func (r *compraRepository) ListByEmpresa(empresaID uint, filters requests.CompraFilter) ([]entities.Compra, error) {
	var compras []entities.Compra

	query := r.db.Where("empresa_id = ?", empresaID)

	if filters.FornecedorID != nil {
		query = query.Where("fornecedor_id = ?", *filters.FornecedorID)
	}

	if filters.DataInicio != "" && filters.DataFim != "" {
		query = query.Where("DATE(data_compra) BETWEEN ? AND ?", filters.DataInicio, filters.DataFim)
	}

	if filters.Status != "" {
		query = query.Where("status = ?", filters.Status)
	}

	err := query.
		Preload("Fornecedor").
		Preload("Usuario").
		Order("data_compra DESC").
		Find(&compras).Error

	return compras, err
}

func (r *compraRepository) ListByFornecedor(fornecedorID uint) ([]entities.Compra, error) {
	var compras []entities.Compra

	err := r.db.
		Where("fornecedor_id = ?", fornecedorID).
		Preload("Fornecedor").
		Preload("Usuario").
		Order("data_compra DESC").
		Find(&compras).Error

	return compras, err
}

func (r *compraRepository) GetComprasPorPeriodo(empresaID uint, inicio, fim string) ([]entities.Compra, error) {
	var compras []entities.Compra

	err := r.db.
		Where("empresa_id = ? AND DATE(data_compra) BETWEEN ? AND ?", empresaID, inicio, fim).
		Preload("Fornecedor").
		Preload("Usuario").
		Preload("Itens").
		Order("data_compra DESC").
		Find(&compras).Error

	return compras, err
}

func (r *compraRepository) GetResumoCompras(empresaID uint, periodo string) (*responses.ResumoCompras, error) {
	var resumo responses.ResumoCompras

	// Implementar lógica de resumo por período
	// Exemplo: total de compras, valor total, etc.

	return &resumo, nil
}

type prontuarioRepository struct {
	db *gorm.DB
}

func NewProntuarioRepository(db *gorm.DB) ProntuarioRepository {
	return &prontuarioRepository{db: db}
}

func (r *prontuarioRepository) Create(prontuario *entities.Prontuario) error {
	return r.db.Create(prontuario).Error
}

func (r *prontuarioRepository) GetByID(id uint) (*entities.Prontuario, error) {
	var prontuario entities.Prontuario
	err := r.db.
		Preload("Pet").
		Preload("Pet.Cliente").
		Preload("Veterinario").
		First(&prontuario, id).Error

	return &prontuario, err
}

func (r *prontuarioRepository) Update(prontuario *entities.Prontuario) error {
	return r.db.Save(prontuario).Error
}

func (r *prontuarioRepository) GetByPet(petID uint) ([]entities.Prontuario, error) {
	var prontuarios []entities.Prontuario

	err := r.db.
		Where("pet_id = ?", petID).
		Preload("Veterinario").
		Order("data_consulta DESC").
		Find(&prontuarios).Error

	return prontuarios, err
}

func (r *prontuarioRepository) GetByVeterinario(veterinarioID uint, inicio, fim string) ([]entities.Prontuario, error) {
	var prontuarios []entities.Prontuario

	err := r.db.
		Where("veterinario_id = ? AND DATE(data_consulta) BETWEEN ? AND ?",
			veterinarioID, inicio, fim).
		Preload("Pet").
		Preload("Pet.Cliente").
		Order("data_consulta DESC").
		Find(&prontuarios).Error

	return prontuarios, err
}

func (r *prontuarioRepository) GetUltimoProntuario(petID uint) (*entities.Prontuario, error) {
	var prontuario entities.Prontuario

	err := r.db.
		Where("pet_id = ?", petID).
		Preload("Veterinario").
		Order("data_consulta DESC").
		First(&prontuario).Error

	return &prontuario, err
}

func (r *prontuarioRepository) RegistrarVacina(vacina *entities.Vacina) error {
	return r.db.Create(vacina).Error
}

func (r *prontuarioRepository) GetVacinasPorPet(petID uint) ([]entities.Vacina, error) {
	var vacinas []entities.Vacina

	err := r.db.
		Where("pet_id = ?", petID).
		Preload("Veterinario").
		Order("data_aplicacao DESC").
		Find(&vacinas).Error

	return vacinas, err
}

func (r *prontuarioRepository) GetVacinasVencidas(empresaID uint) ([]entities.Vacina, error) {
	var vacinas []entities.Vacina

	// Subquery para buscar pets da empresa
	subquery := r.db.Model(&entities.Pet{}).
		Select("pets.id").
		Joins("JOIN clientes ON pets.cliente_id = clientes.id").
		Where("clientes.empresa_id = ?", empresaID)

	err := r.db.
		Where("pet_id IN (?) AND data_proxima < NOW()", subquery).
		Preload("Pet").
		Preload("Pet.Cliente").
		Preload("Veterinario").
		Order("data_proxima ASC").
		Find(&vacinas).Error

	return vacinas, err
}

type logRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepository{db: db}
}

func (r *logRepository) Create(log *entities.LogSistema) error {
	return r.db.Create(log).Error
}

func (r *logRepository) GetByID(id uint) (*entities.LogSistema, error) {
	var log entities.LogSistema
	err := r.db.
		Preload("Usuario").
		Preload("Empresa").
		First(&log, id).Error

	return &log, err
}

func (r *logRepository) GetLogsByUsuario(usuarioID uint, limite int) ([]entities.LogSistema, error) {
	var logs []entities.LogSistema

	err := r.db.
		Where("usuario_id = ?", usuarioID).
		Preload("Empresa").
		Order("created_at DESC").
		Limit(limite).
		Find(&logs).Error

	return logs, err
}

func (r *logRepository) GetLogsByModulo(empresaID uint, modulo string, inicio, fim string) ([]entities.LogSistema, error) {
	var logs []entities.LogSistema

	err := r.db.
		Where("empresa_id = ? AND modulo = ? AND DATE(created_at) BETWEEN ? AND ?",
			empresaID, modulo, inicio, fim).
		Preload("Usuario").
		Order("created_at DESC").
		Find(&logs).Error

	return logs, err
}

func (r *logRepository) GetLogsErro(empresaID uint, dias int) ([]entities.LogSistema, error) {
	var logs []entities.LogSistema

	err := r.db.
		Where("empresa_id = ? AND nivel_log = ? AND created_at >= NOW() - INTERVAL '? days'",
			empresaID, "ERROR", dias).
		Preload("Usuario").
		Order("created_at DESC").
		Find(&logs).Error

	return logs, err
}

func (r *logRepository) GetEstatisticasUso(empresaID uint, periodo string) (*responses.EstatisticasUso, error) {
	var estatisticas responses.EstatisticasUso

	// Implementar lógica de estatísticas de uso
	// Exemplo: total de logs, logs por módulo, etc.

	return &estatisticas, nil
}



type financeiroRepository struct {
	db *gorm.DB
}

func NewFinanceiroRepository(db *gorm.DB) FinanceiroRepository {
	return &financeiroRepository{db: db}
}

func (r *financeiroRepository) CreateContaReceber(conta *entities.ContaReceber) error {
	return r.db.Create(conta).Error
}

func (r *financeiroRepository) GetContaReceberByID(id uint) (*entities.ContaReceber, error) {
	var conta entities.ContaReceber
	err := r.db.Preload("Cliente").Preload("Venda").First(&conta, id).Error
	return &conta, err
}

func (r *financeiroRepository) BaixarContaReceber(id uint, dataPagamento time.Time, formaPagamento string) error {
	return r.db.Model(&entities.ContaReceber{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"data_pagamento": dataPagamento,
			"forma_pagamento": formaPagamento,
			"status": "pago",
		}).Error
}

func (r *financeiroRepository) CreateContaPagar(conta *entities.ContaPagar) error {
	return r.db.Create(conta).Error
}

func (r *financeiroRepository) GetContaPagarByID(id uint) (*entities.ContaPagar, error) {
	var conta entities.ContaPagar
	err := r.db.Preload("Fornecedor").Preload("CategoriaDespesa").First(&conta, id).Error
	return &conta, err
}

func (r *financeiroRepository) PagarContaPagar(id uint, dataPagamento time.Time, formaPagamento string) error {
	return r.db.Model(&entities.ContaPagar{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"data_pagamento": dataPagamento,
			"forma_pagamento": formaPagamento,
			"status": "pago",
		}).Error
}

func (r *financeiroRepository) GetFluxoCaixa(empresaID uint, inicio, fim time.Time) (*responses.FluxoCaixa, error) {
	var fluxo responses.FluxoCaixa
	
	// Receitas do período
	var receitas float64
	r.db.Model(&entities.ContaReceber{}).
		Where("empresa_id = ? AND data_pagamento BETWEEN ? AND ? AND status = ?", 
			empresaID, inicio, fim, "pago").
		Select("COALESCE(SUM(valor), 0)").
		Scan(&receitas)
	
	// Despesas do período
	var despesas float64
	r.db.Model(&entities.ContaPagar{}).
		Where("empresa_id = ? AND data_pagamento BETWEEN ? AND ? AND status = ?", 
			empresaID, inicio, fim, "pago").
		Select("COALESCE(SUM(valor_final), 0)").
		Scan(&despesas)
	
	fluxo.Receitas = receitas
	fluxo.Despesas = despesas
	fluxo.Saldo = receitas - despesas
	fluxo.PeriodoInicio = inicio
	fluxo.PeriodoFim = fim
	
	return &fluxo, nil
}

func (r *financeiroRepository) GetContasVencidas(empresaID uint) ([]entities.ContaReceber, []entities.ContaPagar, error) {
	hoje := time.Now().Format("2006-01-02")
	
	var contasReceber []entities.ContaReceber
	err := r.db.
		Where("empresa_id = ? AND data_vencimento < ? AND status = ?", 
			empresaID, hoje, "pendente").
		Preload("Cliente").
		Find(&contasReceber).Error
	if err != nil {
		return nil, nil, err
	}
	
	var contasPagar []entities.ContaPagar
	err = r.db.
		Where("empresa_id = ? AND data_vencimento < ? AND status = ?", 
			empresaID, hoje, "pendente").
		Preload("Fornecedor").
		Find(&contasPagar).Error
	
	return contasReceber, contasPagar, err
}

func (r *financeiroRepository) GetDemonstrativo(empresaID uint, mes, ano int) (*responses.DemonstrativoFinanceiro, error) {
	var demonstrativo responses.DemonstrativoFinanceiro
	
	inicio := time.Date(ano, time.Month(mes), 1, 0, 0, 0, 0, time.UTC)
	fim := inicio.AddDate(0, 1, -1)
	
	// Receitas do mês
	var receitas float64
	r.db.Model(&entities.ContaReceber{}).
		Where("empresa_id = ? AND data_pagamento BETWEEN ? AND ? AND status = ?", 
			empresaID, inicio, fim, "pago").
		Select("COALESCE(SUM(valor), 0)").
		Scan(&receitas)
	
	// Despesas do mês
	var despesas float64
	r.db.Model(&entities.ContaPagar{}).
		Where("empresa_id = ? AND data_pagamento BETWEEN ? AND ? AND status = ?", 
			empresaID, inicio, fim, "pago").
		Select("COALESCE(SUM(valor_final), 0)").
		Scan(&despesas)
	
	demonstrativo.Mes = mes
	demonstrativo.Ano = ano
	demonstrativo.Receitas = receitas
	demonstrativo.Despesas = despesas
	demonstrativo.Resultado = receitas - despesas
	
	return &demonstrativo, nil
}

// repositories/agendamento_repository.go
package repositories

import (
	"time"
	"gorm.io/gorm"
	"seu-projeto/models/entities"
)

type agendamentoRepository struct {
	db *gorm.DB
}

func NewAgendamentoRepository(db *gorm.DB) AgendamentoRepository {
	return &agendamentoRepository{db: db}
}

func (r *agendamentoRepository) Create(agendamento *entities.Agendamento) error {
	return r.db.Create(agendamento).Error
}

func (r *agendamentoRepository) GetByID(id uint) (*entities.Agendamento, error) {
	var agendamento entities.Agendamento
	err := r.db.
		Preload("Cliente").
		Preload("Pet").
		Preload("TipoServico").
		Preload("Usuario").
		First(&agendamento, id).Error
	
	return &agendamento, err
}

func (r *agendamentoRepository) Update(agendamento *entities.Agendamento) error {
	return r.db.Save(agendamento).Error
}

func (r *agendamentoRepository) Cancelar(id uint, motivo string) error {
	return r.db.Model(&entities.Agendamento{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": "cancelado",
			"observacoes": motivo,
		}).Error
}

func (r *agendamentoRepository) ListByData(empresaID uint, data time.Time) ([]entities.Agendamento, error) {
	var agendamentos []entities.Agendamento
	
	dataStr := data.Format("2006-01-02")
	
	err := r.db.
		Where("empresa_id = ? AND DATE(data_agendamento) = ?", empresaID, dataStr).
		Preload("Cliente").
		Preload("Pet").
		Preload("TipoServico").
		Preload("Usuario").
		Order("data_agendamento ASC").
		Find(&agendamentos).Error
	
	return agendamentos, err
}

func (r *agendamentoRepository) ListByPeriodo(empresaID uint, inicio, fim time.Time) ([]entities.Agendamento, error) {
	var agendamentos []entities.Agendamento
	
	err := r.db.
		Where("empresa_id = ? AND data_agendamento BETWEEN ? AND ?", empresaID, inicio, fim).
		Preload("Cliente").
		Preload("Pet").
		Preload("TipoServico").
		Preload("Usuario").
		Order("data_agendamento ASC").
		Find(&agendamentos).Error
	
	return agendamentos, err
}

func (r *agendamentoRepository) ListByPet(petID uint) ([]entities.Agendamento, error) {
	var agendamentos []entities.Agendamento
	
	err := r.db.
		Where("pet_id = ?", petID).
		Preload("Cliente").
		Preload("TipoServico").
		Preload("Usuario").
		Order("data_agendamento DESC").
		Find(&agendamentos).Error
	
	return agendamentos, err
}

func (r *agendamentoRepository) VerificarDisponibilidade(empresaID uint, dataHora time.Time, servicoID uint) (bool, error) {
	var count int64
	
	// Buscar duração do serviço
	var servico entities.TipoServico
	if err := r.db.First(&servico, servicoID).Error; err != nil {
		return false, err
	}
	
	fim := dataHora.Add(time.Duration(servico.DuracaoMinutos) * time.Minute)
	
	err := r.db.Model(&entities.Agendamento{}).
		Where("empresa_id = ? AND tipo_servico_id = ? AND status NOT IN ('cancelado')", empresaID, servicoID).
		Where("(data_agendamento BETWEEN ? AND ?) OR (data_agendamento + INTERVAL '1 minute' * duracao_minutos BETWEEN ? AND ?)", 
			dataHora, fim, dataHora, fim).
		Count(&count).Error
	
	return count == 0, err
}

func (r *agendamentoRepository) GetHorariosDisponiveis(empresaID uint, data time.Time, servicoID uint) ([]time.Time, error) {
	var horariosDisponiveis []time.Time
	
	// Horário de funcionamento (exemplo: 8h às 18h)
	horaInicio := time.Date(data.Year(), data.Month(), data.Day(), 8, 0, 0, 0, data.Location())
	horaFim := time.Date(data.Year(), data.Month(), data.Day(), 18, 0, 0, 0, data.Location())
	
	// Buscar agendamentos do dia
	agendamentos, err := r.ListByData(empresaID, data)
	if err != nil {
		return nil, err
	}
	
	// Gerar horários disponíveis (a cada 30 minutos)
	for hora := horaInicio; hora.Before(horaFim); hora = hora.Add(30 * time.Minute) {
		disponivel := true
		
		// Verificar conflito com agendamentos existentes
		for _, agendamento := range agendamentos {
			fimAgendamento := agendamento.DataAgendamento.Add(
				time.Duration(agendamento.TipoServico.DuracaoMinutos) * time.Minute)
			
			if (hora.After(agendamento.DataAgendamento) && hora.Before(fimAgendamento)) ||
			   (hora.Equal(agendamento.DataAgendamento)) {
				disponivel = false
				break
			}
		}
		
		if disponivel {
			horariosDisponiveis = append(horariosDisponiveis, hora)
		}
	}
	
	return horariosDisponiveis, nil
}

// repositories/estoque_repository.go
package repositories

import (
	"gorm.io/gorm"
	"seu-projeto/models/entities"
	"seu-projeto/models/responses"
)

type estoqueRepository struct {
	db *gorm.DB
}

func NewEstoqueRepository(db *gorm.DB) EstoqueRepository {
	return &estoqueRepository{db: db}
}

func (r *estoqueRepository) MovimentarEstoque(movimentacao *entities.MovimentacaoEstoque) error {
	// Obter estoque atual
	var ultimaMovimentacao entities.MovimentacaoEstoque
	err := r.db.
		Where("produto_id = ?", movimentacao.ProdutoID).
		Order("created_at DESC").
		First(&ultimaMovimentacao).Error
	
	quantidadeAnterior := 0
	if err == nil {
		quantidadeAnterior = ultimaMovimentacao.QuantidadeAtual
	}
	
	movimentacao.QuantidadeAnterior = quantidadeAnterior
	
	// Calcular nova quantidade
	switch movimentacao.TipoMovimentacao {
	case "entrada":
		movimentacao.QuantidadeAtual = quantidadeAnterior + movimentacao.Quantidade
	case "saida", "ajuste":
		movimentacao.QuantidadeAtual = quantidadeAnterior - movimentacao.Quantidade
	}
	
	return r.db.Create(movimentacao).Error
}

func (r *estoqueRepository) GetHistoricoEstoque(produtoID uint) ([]entities.MovimentacaoEstoque, error) {
	var movimentacoes []entities.MovimentacaoEstoque
	
	err := r.db.
		Where("produto_id = ?", produtoID).
		Preload("Usuario").
		Order("created_at DESC").
		Find(&movimentacoes).Error
	
	return movimentacoes, err
}

func (r *estoqueRepository) GetSaldoAtual(produtoID uint) (int, error) {
	var movimentacao entities.MovimentacaoEstoque
	
	err := r.db.
		Where("produto_id = ?", produtoID).
		Order("created_at DESC").
		First(&movimentacao).Error
	
	if err != nil {
		return 0, err
	}
	
	return movimentacao.QuantidadeAtual, nil
}

func (r *estoqueRepository) AjustarEstoque(produtoID uint, novaQuantidade int, motivo string, usuarioID uint) error {
	// Obter estoque atual
	saldoAtual, err := r.GetSaldoAtual(produtoID)
	if err != nil {
		saldoAtual = 0
	}
	
	diferenca := novaQuantidade - saldoAtual
	tipoMovimentacao := "ajuste"
	if diferenca > 0 {
		tipoMovimentacao = "entrada"
	} else if diferenca < 0 {
		tipoMovimentacao = "saida"
		diferenca = -diferenca
	} else {
		// Não há alteração
		return nil
	}
	
	movimentacao := entities.MovimentacaoEstoque{
		ProdutoID:        produtoID,
		TipoMovimentacao: tipoMovimentacao,
		Quantidade:       diferenca,
		QuantidadeAnterior: saldoAtual,
		QuantidadeAtual:  novaQuantidade,
		Motivo:          motivo,
		UsuarioID:       usuarioID,
	}
	
	return r.MovimentarEstoque(&movimentacao)
}

func (r *estoqueRepository) TransferirEstoque(origemID, destinoID uint, quantidade int, usuarioID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Saída do produto origem
		movimentacaoSaida := entities.MovimentacaoEstoque{
			ProdutoID:        origemID,
			TipoMovimentacao: "saida",
			Quantidade:       quantidade,
			Motivo:          "Transferência entre produtos",
			UsuarioID:       usuarioID,
		}
		if err := r.MovimentarEstoque(&movimentacaoSaida); err != nil {
			return err
		}
		
		// Entrada no produto destino
		movimentacaoEntrada := entities.MovimentacaoEstoque{
			ProdutoID:        destinoID,
			TipoMovimentacao: "entrada",
			Quantidade:       quantidade,
			Motivo:          "Transferência entre produtos",
			UsuarioID:       usuarioID,
		}
		return r.MovimentarEstoque(&movimentacaoEntrada)
	})
}

func (r *estoqueRepository) GetRelatorioEstoque(empresaID uint) ([]responses.RelatorioEstoque, error) {
	var relatorios []responses.RelatorioEstoque
	
	// Subquery para estoque atual
	subquery := r.db.Model(&entities.MovimentacaoEstoque{}).
		Select("produto_id, quantidade_atual").
		Where("id IN (SELECT MAX(id) FROM movimentacao_estoque GROUP BY produto_id)")
	
	err := r.db.
		Table("produtos").
		Select("produtos.id, produtos.nome, produtos.estoque_minimo, estoque.quantidade_atual as estoque_atual").
		Joins("JOIN (?) AS estoque ON produtos.id = estoque.produto_id", subquery).
		Where("produtos.empresa_id = ?", empresaID).
		Order("produtos.nome ASC").
		Scan(&relatorios).Error
	
	return relatorios, err
}

func (r *estoqueRepository) GetMovimentacoesPorPeriodo(empresaID uint, inicio, fim string) ([]entities.MovimentacaoEstoque, error) {
	var movimentacoes []entities.MovimentacaoEstoque
	
	err := r.db.
		Joins("JOIN produtos ON movimentacao_estoque.produto_id = produtos.id").
		Where("produtos.empresa_id = ? AND DATE(movimentacao_estoque.created_at) BETWEEN ? AND ?", 
			empresaID, inicio, fim).
		Preload("Produto").
		Preload("Usuario").
		Order("movimentacao_estoque.created_at DESC").
		Find(&movimentacoes).Error
	
	return movimentacoes, err
}

// repositories/produto_repository.go
package repositories

import (
	"gorm.io/gorm"
	"seu-projeto/models/entities"
	"seu-projeto/models/requests"
)

type produtoRepository struct {
	db *gorm.DB
}

func NewProdutoRepository(db *gorm.DB) ProdutoRepository {
	return &produtoRepository{db: db}
}

func (r *produtoRepository) Create(produto *entities.Produto) error {
	return r.db.Create(produto).Error
}

func (r *produtoRepository) GetByID(id uint) (*entities.Produto, error) {
	var produto entities.Produto
	err := r.db.
		Preload("Categoria").
		Preload("Fornecedor").
		First(&produto, id).Error
	
	return &produto, err
}

func (r *produtoRepository) Update(produto *entities.Produto) error {
	return r.db.Save(produto).Error
}

func (r *produtoRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Produto{}, id).Error
}

func (r *produtoRepository) ListByEmpresa(empresaID uint, filters requests.ProdutoFilter) ([]entities.Produto, error) {
	var produtos []entities.Produto
	
	query := r.db.Where("empresa_id = ?", empresaID)
	
	if filters.Nome != "" {
		query = query.Where("nome ILIKE ?", "%"+filters.Nome+"%")
	}
	
	if filters.CategoriaID != nil {
		query = query.Where("categoria_id = ?", *filters.CategoriaID)
	}
	
	if filters.Ativo != nil {
		query = query.Where("ativo = ?", *filters.Ativo)
	}
	
	if filters.EspecieDestinada != "" {
		query = query.Where("especie_destinada = ?", filters.EspecieDestinada)
	}
	
	err := query.
		Preload("Categoria").
		Preload("Fornecedor").
		Order("nome ASC").
		Find(&produtos).Error
	
	return produtos, err
}

func (r *produtoRepository) GetByCategoria(categoriaID uint) ([]entities.Produto, error) {
	var produtos []entities.Produto
	err := r.db.
		Where("categoria_id = ?", categoriaID).
		Preload("Categoria").
		Preload("Fornecedor").
		Order("nome ASC").
		Find(&produtos).Error
	
	return produtos, err
}

func (r *produtoRepository) Search(empresaID uint, termo string) ([]entities.Produto, error) {
	var produtos []entities.Produto
	
	err := r.db.
		Where("empresa_id = ? AND (nome ILIKE ? OR codigo_barras ILIKE ?)", 
			empresaID, "%"+termo+"%", "%"+termo+"%").
		Preload("Categoria").
		Preload("Fornecedor").
		Order("nome ASC").
		Limit(50).
		Find(&produtos).Error
	
	return produtos, err
}

func (r *produtoRepository) GetProdutosBaixoEstoque(empresaID uint) ([]entities.Produto, error) {
	var produtos []entities.Produto
	
	// Subquery para obter estoque atual
	subquery := r.db.Model(&entities.MovimentacaoEstoque{}).
		Select("produto_id, quantidade_atual").
		Where("id IN (SELECT MAX(id) FROM movimentacao_estoque GROUP BY produto_id)")
	
	err := r.db.
		Joins("JOIN (?) AS estoque ON produtos.id = estoque.produto_id", subquery).
		Where("produtos.empresa_id = ? AND estoque.quantidade_atual <= produtos.estoque_minimo", empresaID).
		Preload("Categoria").
		Preload("Fornecedor").
		Order("estoque.quantidade_atual ASC").
		Find(&produtos).Error
	
	return produtos, err
}

func (r *produtoRepository) UpdateEstoque(produtoID uint, quantidade int) error {
	// Obter estoque atual
	var movimentacao entities.MovimentacaoEstoque
	err := r.db.
		Where("produto_id = ?", produtoID).
		Order("created_at DESC").
		First(&movimentacao).Error
	
	quantidadeAnterior := 0
	if err == nil {
		quantidadeAnterior = movimentacao.QuantidadeAtual
	}
	
	// Criar nova movimentação
	novaMovimentacao := entities.MovimentacaoEstoque{
		ProdutoID:         produtoID,
		TipoMovimentacao:  "ajuste",
		Quantidade:        quantidade,
		QuantidadeAnterior: quantidadeAnterior,
		QuantidadeAtual:   quantidade,
		Motivo:           "Ajuste manual de estoque",
		UsuarioID:        1, // TODO: Obter do usuário logado
	}
	
	return r.db.Create(&novaMovimentacao).Error
}

func (r *produtoRepository) GetProdutoComEstoque(id uint) (*entities.Produto, error) {
	var produto entities.Produto
	
	// Subquery para estoque atual
	subquery := r.db.Model(&entities.MovimentacaoEstoque{}).
		Select("produto_id, quantidade_atual").
		Where("id IN (SELECT MAX(id) FROM movimentacao_estoque GROUP BY produto_id)")
	
	err := r.db.
		Joins("JOIN (?) AS estoque ON produtos.id = estoque.produto_id", subquery).
		Preload("Categoria").
		Preload("Fornecedor").
		First(&produto, id).Error
	
	return &produto, err
}

// repositories/pet_repository.go
package repositories

import (
	"gorm.io/gorm"
	"seu-projeto/models/entities"
	"seu-projeto/models/requests"
)

type petRepository struct {
	db *gorm.DB
}

func NewPetRepository(db *gorm.DB) PetRepository {
	return &petRepository{db: db}
}

func (r *petRepository) Create(pet *entities.Pet) error {
	return r.db.Create(pet).Error
}

func (r *petRepository) GetByID(id uint) (*entities.Pet, error) {
	var pet entities.Pet
	err := r.db.Preload("Cliente").First(&pet, id).Error
	return &pet, err
}

func (r *petRepository) Update(pet *entities.Pet) error {
	return r.db.Save(pet).Error
}

func (r *petRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Pet{}, id).Error
}

func (r *petRepository) GetByCliente(clienteID uint) ([]entities.Pet, error) {
	var pets []entities.Pet
	err := r.db.Where("cliente_id = ?", clienteID).Order("nome ASC").Find(&pets).Error
	return pets, err
}

func (r *petRepository) GetWithCliente(id uint) (*entities.Pet, error) {
	var pet entities.Pet
	err := r.db.Preload("Cliente").First(&pet, id).Error
	return &pet, err
}

func (r *petRepository) ListByEmpresa(empresaID uint, filters requests.PetFilter) ([]entities.Pet, error) {
	var pets []entities.Pet
	
	query := r.db.
		Joins("JOIN clientes ON pets.cliente_id = clientes.id").
		Where("clientes.empresa_id = ?", empresaID)
	
	if filters.Nome != "" {
		query = query.Where("pets.nome ILIKE ?", "%"+filters.Nome+"%")
	}
	
	if filters.Especie != "" {
		query = query.Where("pets.especie = ?", filters.Especie)
	}
	
	if filters.Raca != "" {
		query = query.Where("pets.raca ILIKE ?", "%"+filters.Raca+"%")
	}
	
	err := query.Preload("Cliente").Order("pets.nome ASC").Find(&pets).Error
	return pets, err
}

func (r *petRepository) GetTotalPets(empresaID uint) (int64, error) {
	var count int64
	err := r.db.
		Model(&entities.Pet{}).
		Joins("JOIN clientes ON pets.cliente_id = clientes.id").
		Where("clientes.empresa_id = ?", empresaID).
		Count(&count).Error
	
	return count, err
}

func (r *petRepository) GetPetsPorEspecie(empresaID uint) (map[string]int64, error) {
	type Result struct {
		Especie string
		Count   int64
	}
	
	var results []Result
	
	err := r.db.
		Model(&entities.Pet{}).
		Select("especie, COUNT(*) as count").
		Joins("JOIN clientes ON pets.cliente_id = clientes.id").
		Where("clientes.empresa_id = ?", empresaID).
		Group("especie").
		Scan(&results).Error
	
	if err != nil {
		return nil, err
	}
	
	resultMap := make(map[string]int64)
	for _, result := range results {
		resultMap[result.Especie] = result.Count
	}
	
	return resultMap, nil
}

// repositories/cliente_repository.go
package repositories

import (
	"gorm.io/gorm"
	"seu-projeto/models/entities"
	"seu-projeto/models/requests"
)

type clienteRepository struct {
	db *gorm.DB
}

func NewClienteRepository(db *gorm.DB) ClienteRepository {
	return &clienteRepository{db: db}
}

func (r *clienteRepository) Create(cliente *entities.Cliente) error {
	return r.db.Create(cliente).Error
}

func (r *clienteRepository) GetByID(id uint) (*entities.Cliente, error) {
	var cliente entities.Cliente
	err := r.db.Preload("Pets").First(&cliente, id).Error
	return &cliente, err
}

func (r *clienteRepository) Update(cliente *entities.Cliente) error {
	return r.db.Save(cliente).Error
}

func (r *clienteRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Cliente{}, id).Error
}

func (r *clienteRepository) ListByEmpresa(empresaID uint, filters requests.ClienteFilter) ([]entities.Cliente, error) {
	var clientes []entities.Cliente
	
	query := r.db.Where("empresa_id = ?", empresaID)
	
	if filters.Nome != "" {
		query = query.Where("nome ILIKE ?", "%"+filters.Nome+"%")
	}
	
	if filters.Email != "" {
		query = query.Where("email ILIKE ?", "%"+filters.Email+"%")
	}
	
	if filters.Ativo != nil {
		query = query.Where("ativo = ?", *filters.Ativo)
	}
	
	err := query.Preload("Pets").Order("nome ASC").Find(&clientes).Error
	return clientes, err
}

func (r *clienteRepository) Search(empresaID uint, termo string) ([]entities.Cliente, error) {
	var clientes []entities.Cliente
	
	err := r.db.
		Where("empresa_id = ? AND (nome ILIKE ? OR email ILIKE ? OR telefone ILIKE ?)", 
			empresaID, "%"+termo+"%", "%"+termo+"%", "%"+termo+"%").
		Preload("Pets").
		Order("nome ASC").
		Limit(50).
		Find(&clientes).Error
	
	return clientes, err
}

func (r *clienteRepository) GetWithPets(id uint) (*entities.Cliente, error) {
	var cliente entities.Cliente
	err := r.db.
		Preload("Pets").
		First(&cliente, id).Error
	
	return &cliente, err
}

func (r *clienteRepository) GetTotalClientes(empresaID uint) (int64, error) {
	var count int64
	err := r.db.Model(&entities.Cliente{}).Where("empresa_id = ?", empresaID).Count(&count).Error
	return count, err
}

func (r *clienteRepository) GetClientesNovos(empresaID uint, mes int, ano int) ([]entities.Cliente, error) {
	var clientes []entities.Cliente
	
	err := r.db.
		Where("empresa_id = ? AND EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", 
			empresaID, mes, ano).
		Order("created_at DESC").
		Find(&clientes).Error
	
	return clientes, err
}