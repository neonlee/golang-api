package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type PerfilRepository interface {
	Create(perfil *models.Perfil) error
	GetByID(id uint) (*models.Perfil, error)
	Update(perfil *models.Perfil) error
	Delete(id uint) error
	ListByEmpresa(empresaID uint) ([]models.Perfil, error)
	GetWithPermissoes(id uint) (*models.Perfil, error)
	UpdatePermissoes(perfilID uint, permissoes []models.Permissao) error
}

type perfilRepository struct {
	db *gorm.DB
}

func NewPerfilRepository(db *gorm.DB) PerfilRepository {
	return &perfilRepository{db: db}
}

func (r *perfilRepository) Create(perfil *models.Perfil) error {
	return r.db.Create(perfil).Error
}

func (r *perfilRepository) GetByID(id uint) (*models.Perfil, error) {
	var perfil models.Perfil
	err := r.db.First(&perfil, id).Error
	return &perfil, err
}

func (r *perfilRepository) Update(perfil *models.Perfil) error {
	return r.db.Save(perfil).Error
}

func (r *perfilRepository) Delete(id uint) error {
	return r.db.Delete(&models.Perfil{}, id).Error
}

func (r *perfilRepository) ListByEmpresa(empresaID uint) ([]models.Perfil, error) {
	var perfis []models.Perfil
	err := r.db.Where("empresa_id = ?", empresaID).Find(&perfis).Error
	return perfis, err
}

func (r *perfilRepository) GetWithPermissoes(id uint) (*models.Perfil, error) {
	var perfil models.Perfil
	err := r.db.Preload("Permissoes").First(&perfil, id).Error
	return &perfil, err
}

func (r *perfilRepository) UpdatePermissoes(perfilID uint, permissoes []models.Permissao) error {
	var perfil models.Perfil
	err := r.db.Preload("Permissoes").First(&perfil, perfilID).Error
	if err != nil {
		return err
	}
	return r.db.Model(&perfil).Association("Permissoes").Replace(permissoes)
}
