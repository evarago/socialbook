package modelos

import "time"

// Publicacao representa uma publicação feita por um usuário
type Publicacao struct {
	ID        uint64    `json:"ID,omitempty"`
	Titulo    string    `json:"Titulo,omitempty"`
	Conteudo  string    `json:"Conteudo,omitempty"`
	AutorID   uint64    `json:"AutorID,omitempty"`
	AutorNick string    `json:"AutorNick,omitempty"`
	Curtidas  uint64    `json:"Curtidas"`
	Criacao   time.Time `json:"Criacao"`
}
