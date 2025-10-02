// auth/jwt_claims.go
package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	UserID    uint   `json:"user_id"`
	EmpresaID uint   `json:"empresa_id"`
	Email     string `json:"email"`
	Nome      string `json:"nome"`
	jwt.RegisteredClaims
}

func NewClaims(userID, empresaID uint, email, nome string) CustomClaims {
	return CustomClaims{
		UserID:    userID,
		EmpresaID: empresaID,
		Email:     email,
		Nome:      nome,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 horas
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "pet-system-api",
		},
	}
}

func (c *CustomClaims) ToMap() jwt.MapClaims {
	return jwt.MapClaims{
		"user_id":    c.UserID,
		"empresa_id": c.EmpresaID,
		"email":      c.Email,
		"nome":       c.Nome,
		"exp":        c.ExpiresAt.Unix(),
		"iat":        c.IssuedAt.Unix(),
		"nbf":        c.NotBefore.Unix(),
		"iss":        c.Issuer,
	}
}

func ClaimsFromMap(claims jwt.MapClaims) (*CustomClaims, error) {
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return nil, jwt.ErrInvalidKeyType
	}
	empresaID, ok := claims["empresa_id"].(float64)
	if !ok {
		return nil, jwt.ErrInvalidKeyType
	}
	email, ok := claims["email"].(string)
	if !ok {
		return nil, jwt.ErrInvalidKeyType
	}
	nome, ok := claims["nome"].(string)
	if !ok {
		return nil, jwt.ErrInvalidKeyType
	}
	return &CustomClaims{
		UserID:    uint(userID),
		EmpresaID: uint(empresaID),
		Email:     email,
		Nome:      nome,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(int64(claims["exp"].(float64)), 0)),
			IssuedAt:  jwt.NewNumericDate(time.Unix(int64(claims["iat"].(float64)), 0)),
			NotBefore: jwt.NewNumericDate(time.Unix(int64(claims["nbf"].(float64)), 0)),
			Issuer:    claims["iss"].(string),
		},
	}, nil
}

func GenerateToken(claims CustomClaims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims.ToMap())
	return token.SignedString([]byte(secretKey))
}
