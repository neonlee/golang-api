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
		// models SEM soft delete
		&models.Empresa{},
		&models.Usuarios{},
		&models.Clientes{},
		&models.Pets{},
		&models.Produtos{},
		&models.CategoriaProdutos{},
		&models.Fornecedor{},
		&models.Vendas{},
		&models.VendaItem{},
		&models.TipoServico{},
		&models.Agendamento{},
		&models.ContaReceber{},
		&models.ContaPagar{},
		&models.CategoriaDespesa{},
		&models.Prontuario{},
		&models.Vacina{},
		&models.Compra{},
		&models.CompraItem{},

		// models COM soft delete (se necessário)
		// &models.LogSistema{}, // Se quiser soft delete para logs
	)
}
