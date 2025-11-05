// repositories/auth_repository.go
package repositories

import (
	"errors"
	"petApi/internal/models"
	"petApi/internal/requests"
	"petApi/internal/responses"
	"petApi/pkg/auth"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(req requests.LoginRequest) (*responses.LoginResponse, error)
	GenerateToken(usuario models.Usuarios) (string, error)
	ValidateToken(token string) (*models.Usuarios, error)
	Logout(token requests.LogoutRequest) error
	RefreshToken(token string) (string, error)
	GetUsuarioByEmailAndEmpresa(email string, empresaID uint) (*models.Usuarios, error)
	checkPassword(hashedPassword, password string) bool
	getUsuarioPermissoes(usuarioID uint) ([]string, error)
	GetUsuarioByID(id uint) (*models.Usuarios, error)
}
type authRepository struct {
	db            *gorm.DB
	jwtSecret     string
	tokenDuration time.Duration
}

func NewAuthRepository(db *gorm.DB, jwtSecret string, tokenDuration time.Duration) AuthRepository {
	return &authRepository{
		db:            db,
		jwtSecret:     jwtSecret,
		tokenDuration: tokenDuration,
	}
}

func (r *authRepository) Login(req requests.LoginRequest) (*responses.LoginResponse, error) {
	// Buscar usuário pelo email
	usuario, err := r.GetUsuarioByEmailAndEmpresa(req.Email, req.EmpresaID)
	if err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	if !r.checkPassword(usuario.SenhaHash, req.Senha) {
		return nil, errors.New("credenciais inválidas")
	}

	// Verificar se usuário está ativo
	if !usuario.Ativo {
		return nil, errors.New("usuário inativo")
	}

	// Gerar token
	token, err := r.GenerateToken(*usuario)
	if err != nil {
		return nil, errors.New("erro ao gerar token")
	}

	// Buscar permissões
	permissoes, err := r.getUsuarioPermissoes(usuario.ID)
	if err != nil {
		return nil, err
	}

	// Buscar empresa
	var empresa models.Empresas
	if err := r.db.First(&empresa, req.EmpresaID).Error; err != nil {
		return nil, errors.New("empresa não encontrada")
	}

	// Atualizar último login
	r.db.Model(usuario).Update("ultimo_login", time.Now())

	return &responses.LoginResponse{
		Usuario:    *usuario,
		Token:      token,
		ExpiraEm:   time.Now().Add(r.tokenDuration).Unix(),
		Empresa:    empresa,
		Permissoes: permissoes,
	}, nil
}

func (r *authRepository) GetUsuarioByEmailAndEmpresa(email string, empresaID uint) (*models.Usuarios, error) {
	var usuario models.Usuarios
	err := r.db.
		Preload("Empresa").
		Where("email = ? AND empresa_id = ?", email, empresaID).
		First(&usuario).Error

	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *authRepository) checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (r *authRepository) GenerateToken(usuario models.Usuarios) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    usuario.ID,
		"empresa_id": usuario.EmpresaID,
		"email":      usuario.Email,
		"exp":        time.Now().Add(r.tokenDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(r.jwtSecret))
}

func (r *authRepository) ValidateToken(tokenString string) (*models.Usuarios, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		return r.GetUsuarioByID(userID)
	}

	return nil, errors.New("token inválido")
}

func (r *authRepository) GetUsuarioByID(id uint) (*models.Usuarios, error) {
	var usuario models.Usuarios
	err := r.db.Preload("Empresa").First(&usuario, id).Error
	return &usuario, err
}

func (r *authRepository) getUsuarioPermissoes(usuarioID uint) ([]string, error) {
	var permissoes []string
	err := r.db.
		Model(&models.Permissoes{}).
		Select("DISTINCT modulos.nome").
		Joins("JOIN usuario_perfis ON permissoes.id = usuario_perfis.perfil_id").
		Joins("JOIN modulos ON permissoes.modulo_id = modulos.id").
		Where("usuario_perfis.usuario_id = ?", usuarioID).
		Where("permissoes.pode_visualizar = ?", true).
		Pluck("modulos.nome", &permissoes).Error

	return permissoes, err
}

func (r *authRepository) Logout(token requests.LogoutRequest) error {
	// Implementar blacklist de tokens se necessário
	return nil
}
func (r *authRepository) RefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.jwtSecret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		empresaID := uint(claims["empresa_id"].(float64))
		email := claims["email"].(string)
		nome := claims["nome"].(string)

		newClaims := auth.NewClaims(userID, empresaID, email, nome)
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
		return newToken.SignedString([]byte(r.jwtSecret))
	}
	return "", errors.New("token inválido")
}
