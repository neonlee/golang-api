package controllers

import (
	"net/http"
	"petApi/internal/repositories"
	"petApi/internal/requests"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	repositorio repositories.AuthRepository
}

func NewAuthController(repo repositories.AuthRepository) *AuthController {
	return &AuthController{repositorio: repo}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var login requests.LoginRequest
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	categorias, err := c.repositorio.Login(login)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, categorias)
}

func (c *AuthController) Logout(ctx *gin.Context) {
	var login requests.LogoutRequest
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.repositorio.Logout(login)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (c *AuthController) RefreshToken(ctx *gin.Context) {
	var req requests.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newToken, err := c.repositorio.RefreshToken(req.Token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": newToken})
}
