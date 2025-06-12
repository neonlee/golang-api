package models

import (
	"time"
)

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleManager  Role = "manager"
	RoleCustomer Role = "customer"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// HasPermission verifica se o usuário tem permissão para acessar um recurso
func (u *User) HasPermission(requiredRole Role) bool {
	// Hierarquia de permissões
	roleHierarchy := map[Role]int{
		RoleAdmin:    3,
		RoleManager:  2,
		RoleCustomer: 1,
	}

	return roleHierarchy[u.Role] >= roleHierarchy[requiredRole]
}
