package repositories

import (
	"petApi/internal/models"
	"petApi/internal/responses"

	"gorm.io/gorm"
)

type ServicoRepository interface {
	CreateTipoServico(servico *models.TiposServicos) error
	GetTipoServicoByID(id uint) (*models.TiposServicos, error)
	UpdateTipoServico(servico *models.TiposServicos) error
	DeleteTipoServico(id uint) error
	ListTiposServico(empresaID uint, categoria string) ([]models.TiposServicos, error)
	GetServicosMaisUtilizados(empresaID uint, limite int) ([]responses.ServicoUtilizado, error)
}
type servicoRepository struct {
	db *gorm.DB
}

func NewServicoRepository(db *gorm.DB) ServicoRepository {
	return &servicoRepository{db: db}
}

func (r *servicoRepository) CreateTipoServico(servico *models.TiposServicos) error {
	return r.db.Create(servico).Error
}

func (r *servicoRepository) GetTipoServicoByID(id uint) (*models.TiposServicos, error) {
	var servico models.TiposServicos
	err := r.db.First(&servico, id).Error
	return &servico, err
}

func (r *servicoRepository) UpdateTipoServico(servico *models.TiposServicos) error {
	return r.db.Save(servico).Error
}

func (r *servicoRepository) DeleteTipoServico(id uint) error {
	return r.db.Delete(&models.TiposServicos{}, id).Error
}

func (r *servicoRepository) ListTiposServico(empresaID uint, categoria string) ([]models.TiposServicos, error) {
	var servicos []models.TiposServicos

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
