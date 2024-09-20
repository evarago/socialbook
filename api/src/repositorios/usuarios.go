package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

// Cria um repositorio de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {

	statement, erro := repositorio.db.Prepare(
		"insert into Usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

// Buscar usuários pelo nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNik%
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criacao from Usuarios Where nome like ? or nick like ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Criacao,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// Buscar usuário pelo ID
func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criacao from Usuarios Where id = ?",
		ID,
	)

	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Criacao,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar os dados de um usuário no banco de dados
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {

	statement, erro := repositorio.db.Prepare(
		"Update Usuarios SET nome = ?, nick = ?, email = ? Where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID)
	if erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui um usuário no banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {

	statement, erro := repositorio.db.Prepare(
		"Delete From Usuarios Where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(ID)
	if erro != nil {
		return erro
	}

	return nil
}

// Buscar usuários pelo email e retorna o seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query(
		"select id, senha from Usuarios Where email = ?",
		email,
	)

	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Senha,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Seguir permite que um usuário siga o outro
func (repositorio Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"insert ignore into Seguidores (usuario_id, seguidor_id) values (?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

// Parar de seguir um usário
func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from Seguidores where usuario_id = ? and seguidor_id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

// Buscar seguidores de um usuário
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criacao from Usuarios u 
		inner join Seguidores s on u.id = s.seguidor_id where s.usuario_id = ? `,
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Criacao,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// Buscar quem um usuário está seguindo
func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criacao from Usuarios u 
		inner join Seguidores s on u.id = s.usuario_id where s.seguidor_id = ? `,
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Criacao,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// Buscar senha do usuário pelo ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {

	linhas, erro := repositorio.db.Query(
		"select senha from Usuarios Where id = ?",
		usuarioID,
	)

	if erro != nil {
		return "", erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

// Atualizar a senha do usuário
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare(
		"update Usuarios SET senha = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil
}
