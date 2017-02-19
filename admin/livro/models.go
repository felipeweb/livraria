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
	_, err = db.Exec(`INSERT INTO livro 
	(isbn, titulo, descricao, data_publicacao, preco)
	VALUES 
	($1, $2, $3, $4, $5)`,
		l.ISBN, l.Titulo, l.Descricao, l.DataDePublicacao, l.Preco)
	return
}
