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
	r.connection.Table("client").AutoMigrate(models.Client{})
	r.connection.Table("pet").AutoMigrate(models.Pet{})
	r.connection.Table("service").AutoMigrate(models.Service{})
}
