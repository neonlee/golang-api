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
	r.connection.Table("module").AutoMigrate(&models.Modulo{})
	r.connection.Table("type_access_module").AutoMigrate(&models.TipoAcessoModulo{})
	// r.connection.Table("users").AutoMigrate(&models.Users{})
	r.connection.Table("employee").AutoMigrate(&models.Employee{})
	r.connection.Table("client").AutoMigrate(&models.Client{})
	r.connection.Table("pet").AutoMigrate(&models.Pet{})
	r.connection.Table("service").AutoMigrate(&models.Service{})
	r.connection.Table("category").AutoMigrate(&models.Category{})

	//r.connection.Table("tenant").AutoMigrate(models.Tenant{})
}
