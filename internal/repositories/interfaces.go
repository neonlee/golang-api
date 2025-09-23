// repositories/interfaces.go
package repositories

import "petApi/internal/models"

// UsuarioRepository interface
type UsuarioRepository interface {
	Create(usuario *models.Usuario, senha string) error
	GetByID(id uint) (*models.Usuario, error)
	GetByEmail(email string) (*models.Usuario, error)
	Update(usuario *models.Usuario) error
	Delete(id uint) error
	UpdateSenha(id uint, novaSenha string) error
	ListByEmpresa(empresaID uint, filters requests.UsuarioFilter) ([]models.Usuario, error)
	GetWithPerfis(id uint) (*responses.UsuarioResponse, error)
	AssignPerfis(usuarioID uint, perfisIDs []uint) error
	GetPermissoes(usuarioID uint) ([]string, error)
	HasPermission(usuarioID uint, modulo string, acao string) bool
	CheckEmpresaAccess(usuarioID uint, empresaID uint) bool
}

// AuthRepository interface
type AuthRepository interface {
	Login(req requests.LoginRequest) (*responses.LoginResponse, error)
	RegistrarEmpresa(req requests.RegistrarEmpresaRequest) error
	GenerateToken(usuario models.Usuario) (string, error)
	ValidateToken(token string) (*models.Usuario, error)
	Logout(token string) error
	RefreshToken(token string) (string, error)
}

// ClienteRepository interface
type ClienteRepository interface {
	Create(cliente *models.Cliente) error
	GetByID(id uint) (*models.Cliente, error)
	Update(cliente *models.Cliente) error
	Delete(id uint) error
	ListByEmpresa(empresaID uint, filters requests.ClienteFilter) ([]models.Cliente, error)
	Search(empresaID uint, termo string) ([]models.Cliente, error)
	GetWithPets(id uint) (*models.Cliente, error)
	GetTotalClientes(empresaID uint) (int64, error)
	GetClientesNovos(empresaID uint, mes int, ano int) ([]models.Cliente, error)
}

// PetRepository interface
type PetRepository interface {
	Create(pet *models.Pet) error
	GetByID(id uint) (*models.Pet, error)
	Update(pet *models.Pet) error
	Delete(id uint) error
	GetByCliente(clienteID uint) ([]models.Pet, error)
	GetWithCliente(id uint) (*models.Pet, error)
	ListByEmpresa(empresaID uint, filters requests.PetFilter) ([]models.Pet, error)
	GetTotalPets(empresaID uint) (int64, error)
	GetPetsPorEspecie(empresaID uint) (map[string]int64, error)
}

// ProdutoRepository interface
type ProdutoRepository interface {
	Create(produto *models.Produto) error
	GetByID(id uint) (*models.Produto, error)
	Update(produto *models.Produto) error
	Delete(id uint) error
	ListByEmpresa(empresaID uint, filters requests.ProdutoFilter) ([]models.Produto, error)
	GetByCategoria(categoriaID uint) ([]models.Produto, error)
	Search(empresaID uint, termo string) ([]models.Produto, error)
	GetProdutosBaixoEstoque(empresaID uint) ([]models.Produto, error)
	UpdateEstoque(produtoID uint, quantidade int) error
	GetProdutoComEstoque(id uint) (*models.Produto, error)
}

// VendaRepository interface
type VendaRepository interface {
	Create(venda *models.Venda, itens []models.VendaItem) error
	GetByID(id uint) (*models.Venda, error)
	UpdateStatus(id uint, status string) error
	CancelarVenda(id uint, motivo string) error
	ListByEmpresa(empresaID uint, filters requests.VendaFilter) ([]models.Venda, error)
	GetVendasDoDia(empresaID uint, data string) ([]models.Venda, error)
	GetVendasPorPeriodo(empresaID uint, inicio, fim string) ([]models.Venda, error)
	GetResumoVendas(empresaID uint, periodo string) (*responses.ResumoVendas, error)
	GetProdutosMaisVendidos(empresaID uint, limite int) ([]responses.ProdutoVendas, error)
	GetVendasPorFormaPagamento(empresaID uint, inicio, fim string) (map[string]float64, error)
}

// EstoqueRepository interface
type EstoqueRepository interface {
	MovimentarEstoque(movimentacao *models.MovimentacaoEstoque) error
	GetHistoricoEstoque(produtoID uint) ([]models.MovimentacaoEstoque, error)
	GetSaldoAtual(produtoID uint) (int, error)
	AjustarEstoque(produtoID uint, novaQuantidade int, motivo string, usuarioID uint) error
	TransferirEstoque(origemID, destinoID uint, quantidade int, usuarioID uint) error
	GetRelatorioEstoque(empresaID uint) ([]responses.RelatorioEstoque, error)
	GetMovimentacoesPorPeriodo(empresaID uint, inicio, fim string) ([]models.MovimentacaoEstoque, error)
}

// ServicoRepository interface
type ServicoRepository interface {
	CreateTipoServico(servico *models.TipoServico) error
	GetTipoServicoByID(id uint) (*models.TipoServico, error)
	UpdateTipoServico(servico *models.TipoServico) error
	DeleteTipoServico(id uint) error
	ListTiposServico(empresaID uint, categoria string) ([]models.TipoServico, error)
	GetServicosMaisUtilizados(empresaID uint, limite int) ([]responses.ServicoUtilizado, error)
}

// AgendamentoRepository interface
type AgendamentoRepository interface {
	Create(agendamento *models.Agendamento) error
	GetByID(id uint) (*models.Agendamento, error)
	Update(agendamento *models.Agendamento) error
	Cancelar(id uint, motivo string) error
	ListByData(empresaID uint, data string) ([]models.Agendamento, error)
	ListByPeriodo(empresaID uint, inicio, fim string) ([]models.Agendamento, error)
	ListByPet(petID uint) ([]models.Agendamento, error)
	VerificarDisponibilidade(empresaID uint, dataHora string, servicoID uint) (bool, error)
	GetHorariosDisponiveis(empresaID uint, data string, servicoID uint) ([]string, error)
}

// FinanceiroRepository interface
type FinanceiroRepository interface {
	CreateContaReceber(conta *models.ContaReceber) error
	GetContaReceberByID(id uint) (*models.ContaReceber, error)
	BaixarContaReceber(id uint, dataPagamento string, formaPagamento string) error
	CreateContaPagar(conta *models.ContaPagar) error
	GetContaPagarByID(id uint) (*models.ContaPagar, error)
	PagarContaPagar(id uint, dataPagamento string, formaPagamento string) error
	GetFluxoCaixa(empresaID uint, inicio, fim string) (*responses.FluxoCaixa, error)
	GetContasVencidas(empresaID uint) ([]models.ContaReceber, []models.ContaPagar, error)
	GetDemonstrativo(empresaID uint, mes, ano int) (*responses.DemonstrativoFinanceiro, error)
}

// CompraRepository interface
type CompraRepository interface {
	Create(compra *models.Compra, itens []models.CompraItem) error
	GetByID(id uint) (*models.Compra, error)
	Update(compra *models.Compra) error
	Cancelar(id uint, motivo string) error
	ListByEmpresa(empresaID uint, filters requests.CompraFilter) ([]models.Compra, error)
	ListByFornecedor(fornecedorID uint) ([]models.Compra, error)
	GetComprasPorPeriodo(empresaID uint, inicio, fim string) ([]models.Compra, error)
	GetResumoCompras(empresaID uint, periodo string) (*responses.ResumoCompras, error)
}

// ProntuarioRepository interface
type ProntuarioRepository interface {
	Create(prontuario *models.Prontuario) error
	GetByID(id uint) (*models.Prontuario, error)
	Update(prontuario *models.Prontuario) error
	GetByPet(petID uint) ([]models.Prontuario, error)
	GetByVeterinario(veterinarioID uint, inicio, fim string) ([]models.Prontuario, error)
	GetUltimoProntuario(petID uint) (*models.Prontuario, error)
	RegistrarVacina(vacina *models.Vacina) error
	GetVacinasPorPet(petID uint) ([]models.Vacina, error)
	GetVacinasVencidas(empresaID uint) ([]models.Vacina, error)
}

// LogRepository interface
type LogRepository interface {
	Create(log *models.LogSistema) error
	GetByID(id uint) (*models.LogSistema, error)
	GetLogsByUsuario(usuarioID uint, limite int) ([]models.LogSistema, error)
	GetLogsByModulo(empresaID uint, modulo string, inicio, fim string) ([]models.LogSistema, error)
	GetLogsErro(empresaID uint, dias int) ([]models.LogSistema, error)
	GetEstatisticasUso(empresaID uint, periodo string) (*responses.EstatisticasUso, error)
}

// FornecedorRepository interface
type FornecedorRepository interface {
	Create(fornecedor *models.Fornecedor) error
	GetByID(id uint) (*models.Fornecedor, error)
	Update(fornecedor *models.Fornecedor) error
	Delete(id uint) error
	ListByEmpresa(empresaID uint, filters requests.FornecedorFilter) ([]models.Fornecedor, error)
	Search(empresaID uint, termo string) ([]models.Fornecedor, error)
	GetTotalFornecedores(empresaID uint) (int64, error)
}

// CategoriaRepository interface
type CategoriaRepository interface {
	Create(categoria *models.CategoriaProduto) error
	GetByID(id uint) (*models.CategoriaProduto, error)
	Update(categoria *models.CategoriaProduto) error
	Delete(id uint) error
	ListByEmpresa(empresaID uint) ([]models.CategoriaProduto, error)
	GetWithProdutos(id uint) (*models.CategoriaProduto, error)
}

// PlanoRepository interface
type PlanoRepository interface {
	Create(plano *models.Plano) error
	GetByID(id uint) (*models.Plano, error)
	Update(plano *models.Plano) error
	ListAll() ([]models.Plano, error)
	GetByNome(nome string) (*models.Plano, error)
}

// PerfilRepository interface
type PerfilRepository interface {
	Create(perfil *models.Perfil) error
	GetByID(id uint) (*models.Perfil, error)
	Update(perfil *models.Perfil) error
	Delete(id uint) error
	ListByEmpresa(empresaID uint) ([]models.Perfil, error)
	GetWithPermissoes(id uint) (*models.Perfil, error)
	UpdatePermissoes(perfilID uint, permissoes []models.Permissao) error
}

// ModuloRepository interface
type ModuloRepository interface {
	GetByID(id uint) (*models.Modulo, error)
	ListAll() ([]models.Modulo, error)
	ListByCategoria(categoria string) ([]models.Modulo, error)
	GetByNome(nome string) (*models.Modulo, error)
}

// DashboardRepository interface
type DashboardRepository interface {
	GetResumoVendas(empresaID uint, periodo string) (*responses.ResumoVendas, error)
	GetResumoFinanceiro(empresaID uint) (*responses.ResumoFinanceiro, error)
	GetProximosAgendamentos(empresaID uint, dias int) ([]responses.AgendamentoResumo, error)
	GetAlertasEstoque(empresaID uint) ([]responses.AlertaEstoque, error)
	GetMetricasGerais(empresaID uint) (*responses.MetricasGerais, error)
}
