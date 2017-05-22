package livro

import (
	"database/sql"
	"time"
)

// Livro representa um livro
type Livro struct {
	ID               int       `json:"id,omitempty"`
	ISBN             string    `json:"isbn,omitempty"`
	Titulo           string    `json:"titulo,omitempty"`
	Descricao        string    `json:"descricao,omitempty"`
	DataDePublicacao time.Time `json:"data_publicacao,omitempty"`
	Preco            float64   `json:"preco,omitempty"`
}

// Insert um livro no banco de dados
func (l *Livro) Insert(db *sql.DB) (err error) {
	_, err = db.Exec(`INSERT INTO livros 
	(isbn, titulo, descricao, data_publicacao, preco)
	VALUES 
	($1, $2, $3, $4, $5)`,
		l.ISBN, l.Titulo, l.Descricao, l.DataDePublicacao, l.Preco)
	return
}

// List os livros do banco de dados
func (l *Livro) List(db *sql.DB) (livros []Livro, err error) {
	var rows *sql.Rows
	rows, err = db.Query("SELECT * FROM livros")
	if err != nil {
		return
	}
	for rows.Next() {
		err = rows.Scan(
			&l.ID,
			&l.ISBN,
			&l.Titulo,
			&l.Descricao,
			&l.DataDePublicacao,
			&l.Preco,
		)
		if err != nil {
			return
		}
		livros = append(livros, *l)
	}
	return
}
