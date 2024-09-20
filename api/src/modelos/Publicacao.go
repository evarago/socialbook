package modelos

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID        uint64    `json:"ID,omitempty"`
	Titulo    string    `json:"Titulo,omitempty"`
	Conteudo  string    `json:"Conteudo,omitempty"`
	AutorID   uint64    `json:"AutorID,omitempty"`
	AutorNick string    `json:"AutorNick,omitempty"`
	Curtidas  uint64    `json:"Curtidas"`
	Criacao   time.Time `json:"Criacao"`
}

// Preparar vai chacar os métodos para validação e formatação de campos
func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O título é obrigatório e não pode estar em branco")
	}
	if publicacao.Conteudo == "" {
		return errors.New("O Conteúdo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
