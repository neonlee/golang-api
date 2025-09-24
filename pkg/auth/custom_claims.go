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
