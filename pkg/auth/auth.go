package middleware

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
)

func RoleCheck(repo *repositories.UserRepository, requiredRole models.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Obter ID do usuário do contexto (assumindo que já foi autenticado)
			userID := r.Context().Value("userID").(int)

			user, err := repo.GetByID(userID)
			if err != nil {
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}

			if !user.HasPermission(requiredRole) {
				http.Error(w, "Insufficient permissions", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
