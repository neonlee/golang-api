package Controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllersPet struct {
	connection *sql.DB
}

func NewProductController(connection *sql.DB) ControllersPet {
	return ControllersPet{
		connection: connection,
	}
}

// GetPets godoc
// @Summary Lista todos os pets
// @Description Retorna todos os pets cadastrados
// @Tags pets
// @Accept  json
// @Produce  json
// @Success 200 {array} Pet
// @Failure 500 {object} map[string]string
// @Router /pets [get]
func (p *ControllersPet) GetPets(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, `{"":"teste api"}`)
}
