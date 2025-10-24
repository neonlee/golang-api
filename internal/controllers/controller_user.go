package controllers

import (
	"fmt"
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"petApi/internal/requests"

	"github.com/gin-gonic/gin"
)

type ControllersUser struct {
	repository repositories.UsuarioRepository
}

func NewUsuarioController(connection repositories.UsuarioRepository) *ControllersUser {
	return &ControllersUser{repository: connection}
}

func (u *ControllersUser) CreateUser(ctx *gin.Context) {
	var usuario models.Usuarios
	if err := ctx.BindJSON(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	err := u.repository.Create(&usuario, usuario.SenhaHash)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, &usuario)
}

func (u *ControllersUser) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	usuario, err := u.repository.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar usuário"})
		return
	}
	if usuario == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, usuario)
}

func (u *ControllersUser) GetByEmail(ctx *gin.Context) {
	email := ctx.Query("email")
	usuario, err := u.repository.GetByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar usuário"})
		return
	}
	if usuario == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, usuario)
}

func (u *ControllersUser) Update(ctx *gin.Context) {
	var usuario models.Usuarios
	if err := ctx.BindJSON(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})

		return
	}
	err := u.repository.Update(&usuario)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao atualizar usuário"})
		return
	}
	ctx.JSON(http.StatusOK, &usuario)
}
func (u *ControllersUser) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	err = u.repository.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao deletar usuário"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"mensagem": "Usuário deletado com sucesso"})
}

func (u *ControllersUser) UpdateSenha(ctx *gin.Context) {
	var req requests.UpdateSenhaRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	err := u.repository.UpdateSenha(req.UsuarioID, req.NovaSenha)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao atualizar senha"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"mensagem": "Senha atualizada com sucesso"})
}
func (u *ControllersUser) ListByEmpresa(ctx *gin.Context) {
	empresaIDParam := ctx.Param("empresa_id")
	var empresaID uint
	_, err := fmt.Sscan(empresaIDParam, &empresaID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID da empresa inválido"})
		return
	}
	var filters requests.UsuarioFilter
	if err := ctx.BindQuery(&filters); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Filtros inválidos"})
		return
	}
	usuarios, err := u.repository.ListByEmpresa(empresaID, filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao listar usuários"})
		return
	}
	ctx.JSON(http.StatusOK, usuarios)
}

func (u *ControllersUser) GetWithPerfis(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	usuario, err := u.repository.GetWithPerfis(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar usuário"})
		return
	}
	if usuario == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, usuario)
}

func (u *ControllersUser) AssignPerfis(ctx *gin.Context) {
	var req requests.AssignPerfilRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	err := u.repository.AssignPerfis(req.UsuarioID, req.Perfil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao atribuir perfis"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"mensagem": "Perfis atribuídos com sucesso"})
}
func (u *ControllersUser) GetPermissoes(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	permissoes, err := u.repository.GetPermissoes(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar permissões"})
		return
	}
	ctx.JSON(http.StatusOK, permissoes)
}
func (u *ControllersUser) HasPermission(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	permission := ctx.Query("permission")
	acao := ctx.Query("acao")

	has := u.repository.HasPermission(uint(id), permission, acao)

	ctx.JSON(http.StatusOK, gin.H{"has_permission": has})
}
func (u *ControllersUser) CheckEmpresaAccess(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	_, err := fmt.Sscan(idParam, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	empresaIDParam := ctx.Query("empresa_id")
	var empresaID uint
	_, err = fmt.Sscan(empresaIDParam, &empresaID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID da empresa inválido"})
		return
	}
	hasAccess := u.repository.CheckEmpresaAccess(id, empresaID)
	ctx.JSON(http.StatusOK, gin.H{"has_access": hasAccess})
}
