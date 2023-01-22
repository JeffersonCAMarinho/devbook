package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

// criar um repositorio de Usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// criar insere um usuario no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into usuarios(nome, nick, email, senha) value (?,?,?,?)",)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {return 0, erro}

	return uint64(ultimoIDInserido), nil

	}
