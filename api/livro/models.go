package livro

import (
	"time"
)

// Livro representa um livro
type Livro struct {
	ID               int       `json:"id,omitempty"`
	ISBN             string    `json:"isbn,omitempty"`
	Titulo           string    `json:"titulo,omitempty"`
	Descricao        string    `json:"descricao,omitempty"`
	DataDePublicacao time.Time `json:"publicado,omitempty"`
	Preco            float64   `json:"preco,omitempty"`
}
