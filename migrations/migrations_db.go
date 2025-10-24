package migrations

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

// import (
// 	"gorm.io/gorm"
// )

type MigrationsDB struct {
	connection *gorm.DB
}

func NewMigrationsDB(db *gorm.DB) *MigrationsDB {
	return &MigrationsDB{db}
}

func (r *MigrationsDB) RunMigrations() error {
	return r.connection.AutoMigrate(
		&models.Empresas{},
		&models.Usuarios{},
		&models.Clientes{},
		&models.Pets{},
		&models.Produtos{},
		&models.CategoriasProdutos{},
		&models.Fornecedores{},
		models.Funcionarios{},
		&models.Vendas{},
		&models.VendaItem{},
		&models.TiposServicos{},
		&models.Agendamento{},
		&models.ContaReceber{},
		&models.ContaPagar{},
		&models.CategoriaDespesa{},
		&models.Prontuarios{},
		&models.Vacinas{},
		&models.Compras{},
		&models.CompraItens{},
		&models.MovimentacaoEstoques{},

		// models COM soft delete (se necessário)
		// &models.LogSistema{}, // Se quiser soft delete para logs
	)
}
