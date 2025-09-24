package requests

type RegistrarEmpresaRequest struct {
	NomeEmpresa    string `json:"nome_empresa" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	NomeUsuario    string `json:"nome_usuario" binding:"required"`
	Senha          string `json:"senha" binding:"required,min=6"`
	PlanoID        uint   `json:"plano_id" binding:"required"`
	FormaPagamento string `json:"forma_pagamento" binding:"required,oneof=cartao_boleto"`
	TokenRecaptcha string `json:"token_recaptcha" binding:"required"`
	TermosAceitos  bool   `json:"termos_aceitos" binding:"required,eq=true"`
}
