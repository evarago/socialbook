package modelos

// DadosAutenticacao contém o token e o id do usuário autenticado
type DadosAutenticacao struct {
	ID    string `json:"ID"`
	Token string `json:"Token"`
}
