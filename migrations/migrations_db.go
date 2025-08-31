package migrations

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type MigrationsDB struct {
	connection *gorm.DB
}

func NewMigrationsDB(db *gorm.DB) *MigrationsDB {
	return &MigrationsDB{db}
}

func (r *MigrationsDB) RunMigrations() {
	r.connection.Table("type_access").AutoMigrate(&models.TipoAcesso{})
	r.connection.Table("module_permission").AutoMigrate(&models.Modulo{})
	r.connection.Table("type_access_module").AutoMigrate(&models.TipoAcessoModulo{})
	r.connection.Table("users").AutoMigrate(&models.Users{})
	r.connection.Table("employee").AutoMigrate(&models.Employee{})
	r.connection.Table("client").AutoMigrate(&models.Client{})
	r.connection.Table("pet").AutoMigrate(&models.Pet{})
	r.connection.Table("service").AutoMigrate(&models.Service{})
	r.connection.Table("category").AutoMigrate(&models.Category{})
	r.connection.Table("supplier").AutoMigrate(&models.Supplier{})
	r.connection.Table("product").AutoMigrate(&models.Product{})
	// r.connection.Table("service_order").AutoMigrate(&models.ServiceOrder{})
	// r.connection.Table("service_order_item").AutoMigrate(&models.ServiceOrderItem{})
	// r.connection.Table("stock_entry").AutoMigrate(&models.StockEntry{})
	// r.connection.Table("stock_entry_item").AutoMigrate(&models.StockEntryItem{})
	// r.connection.Table("stock_exit").AutoMigrate(&models.StockExit{})
	// r.connection.Table("stock_exit_item").AutoMigrate(&models.StockExitItem{})
	// r.connection.Table("sale").AutoMigrate(&models.Sale{})
	// r.connection.Table("sale_item").AutoMigrate(&models.SaleItem{})
	//r.connection.Table("tenant").AutoMigrate(models.Tenant{})
}
