// repositories/usuario_repository.go
package repositories

import (
	"errors"
	"petApi/internal/models"
	"petApi/internal/requests"
	"petApi/internal/responses"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UsuarioRepository interface
type UsuarioRepository interface {
	Create(usuario *models.Usuarios, senha string) error
	GetByID(id uint) (*models.Usuarios, error)
	GetByEmail(email string) (*models.Usuarios, error)
	Update(usuario *models.Usuarios) error
	Delete(id uint) error
	UpdateSenha(id uint, novaSenha string) error
	ListByEmpresa(empresaID uint, filters requests.UsuarioFilter) ([]models.Usuarios, error)
	GetWithPerfis(id uint) (*responses.UsuarioResponse, error)
	AssignPerfis(usuarioID uint, perfisIDs []uint) error
	GetPermissoes(usuarioID uint) ([]string, error)
	HasPermission(usuarioID uint, modulo string, acao string) bool
	CheckEmpresaAccess(usuarioID uint, empresaID uint) bool
}

type usuarioRepository struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) UsuarioRepository {
	return &usuarioRepository{db: db}
}

func (r *usuarioRepository) Create(usuario *models.Usuarios, senha string) error {
	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usuario.SenhaHash = string(hashedPassword)

	return r.db.Create(usuario).Error
}

func (r *usuarioRepository) GetByID(id uint) (*models.Usuarios, error) {
	var usuario models.Usuarios
	err := r.db.Preload("Empresa").First(&usuario, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &usuario, nil
}

func (r *usuarioRepository) GetByEmail(email string) (*models.Usuarios, error) {
	var usuario models.Usuarios
	err := r.db.Preload("Empresa").Where("email = ?", email).First(&usuario).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &usuario, nil
}

func (r *usuarioRepository) Update(usuario *models.Usuarios) error {
	return r.db.Save(usuario).Error
}

func (r *usuarioRepository) Delete(id uint) error {
	return r.db.Delete(&models.Usuarios{}, id).Error
}

func (r *usuarioRepository) UpdateSenha(id uint, novaSenha string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(novaSenha), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return r.db.Model(&models.Usuarios{}).Where("id = ?", id).Update("senha_hash", string(hashedPassword)).Error
}

func (r *usuarioRepository) ListByEmpresa(empresaID uint, filters requests.UsuarioFilter) ([]models.Usuarios, error) {
	var usuarios []models.Usuarios

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

	if filters.Cargo != "" {
		query = query.Where("cargo = ?", filters.Cargo)
	}

	err := query.Preload("Empresa").Order("nome ASC").Find(&usuarios).Error
	return usuarios, err
}

func (r *usuarioRepository) GetWithPerfis(id uint) (*responses.UsuarioResponse, error) {
	var usuario models.Usuarios
	err := r.db.
		Preload("Empresa").
		Preload("UsuarioPerfis").
		Preload("UsuarioPerfis.Perfil").
		Preload("UsuarioPerfis.Perfil.Permissoes").
		Preload("UsuarioPerfis.Perfil.Permissoes.Modulo").
		First(&usuario, id).Error

	if err != nil {
		return nil, err
	}

	// Converter para response
	response := &responses.UsuarioResponse{
		Usuario: usuario,
		Perfis:  make([]models.Perfil, 0),
	}

	// Extrair perfis únicos
	perfisMap := make(map[uint]models.Perfil)
	for _, usuarioPerfil := range usuario.UsuarioPerfis {
		perfisMap[usuarioPerfil.PerfilID] = usuarioPerfil.Perfil
	}

	for _, perfil := range perfisMap {
		response.Perfis = append(response.Perfis, perfil)
	}

	return response, nil
}

func (r *usuarioRepository) AssignPerfis(usuarioID uint, perfisIDs []uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Remover perfis atuais
		if err := tx.Where("usuario_id = ?", usuarioID).Delete(&models.UsuarioPerfil{}).Error; err != nil {
			return err
		}

		// Adicionar novos perfis
		for _, perfilID := range perfisIDs {
			usuarioPerfil := models.UsuarioPerfil{
				UsuarioID: usuarioID,
				PerfilID:  perfilID,
				CreatedAt: time.Now(),
			}
			if err := tx.Create(&usuarioPerfil).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *usuarioRepository) GetPermissoes(usuarioID uint) ([]string, error) {
	var permissoes []string

	err := r.db.
		Model(&models.Permissoes{}).
		Joins("JOIN usuario_perfis ON permissoes.perfil_id = usuario_perfis.perfil_id").
		Joins("JOIN modulos ON permissoes.modulo_id = modulos.id").
		Where("usuario_perfis.usuario_id = ?", usuarioID).
		Where("permissoes.pode_visualizar = ?", true).
		Pluck("DISTINCT modulos.nome", &permissoes).Error

	return permissoes, err
}

func (r *usuarioRepository) HasPermission(usuarioID uint, modulo string, acao string) bool {
	var count int64

	query := r.db.
		Model(&models.Permissoes{}).
		Joins("JOIN usuario_perfis ON permissoes.perfil_id = usuario_perfis.perfil_id").
		Joins("JOIN modulos ON permissoes.modulo_id = modulos.id").
		Where("usuario_perfis.usuario_id = ?", usuarioID).
		Where("modulos.nome = ?", modulo)

	switch acao {
	case "visualizar":
		query = query.Where("permissoes.pode_visualizar = ?", true)
	case "editar":
		query = query.Where("permissoes.pode_editar = ?", true)
	case "excluir":
		query = query.Where("permissoes.pode_excluir = ?", true)
	case "relatorio":
		query = query.Where("permissoes.pode_gerar_relatorio = ?", true)
	default:
		return false
	}

	err := query.Count(&count).Error
	return err == nil && count > 0
}

func (r *usuarioRepository) CheckEmpresaAccess(usuarioID uint, empresaID uint) bool {
	var count int64
	err := r.db.Model(&models.Usuarios{}).
		Where("id = ? AND empresa_id = ?", usuarioID, empresaID).
		Count(&count).Error

	return err == nil && count > 0
}
