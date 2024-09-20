package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuario utilizando a rede social
type Usuario struct {
	ID      uint64    `json:"ID,omitempty"`
	Nome    string    `json:"Nome,omitempty"`
	Nick    string    `json:"Nick,omitempty"`
	Email   string    `json:"Email,omitempty"`
	Senha   string    `json:"Senha,omitempty"`
	Criacao time.Time `json:"Criacao,omitempty"`
}

// Preparar modelo
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

// Validar campos
func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O Nome é obrigatório.")
	}

	if usuario.Nick == "" {
		return errors.New("O Nick é obrigatório.")
	}

	if usuario.Email == "" {
		return errors.New("O Email é obrigatório.")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O Email inserido é inválido.")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatória.")
	}

	return nil
}

// Formatar campos
func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}
