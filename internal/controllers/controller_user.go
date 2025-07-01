package Controllers

import (
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllerUser struct {
	repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *ControllerUser {
	return &ControllerUser{repo: repo}
}

// @create	user statament
func (h *ControllerUser) CreateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	users, err := h.repo.Create(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// GetPets godoc
//
//	@Summary		Lista todos os pets
//	@Description	Retorna todos os pets cadastrados
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.User
//	@Failure		500	{object}	map[string]string
//	@Router			/pets/id [get]
func (h *ControllerUser) GetUser(ctx *gin.Context) {
	id := ctx.Param("productId")

	user, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	users, err := h.repo.GetByID(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
// 	users, err := h.repo.GetAll()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(users)
// }
