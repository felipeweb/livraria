package livro

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/felipeweb/livraria/admin/livro/database"
)

var db *sql.DB

func init() {
	var err error
	db, err = database.Connect()
	if err != nil {
		panic(err)
	}
}

// InsertForm renderiza o formulário
func InsertForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("form.html").ParseFiles("livro/tmpl/form.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SaveBook salva o livro
func SaveBook(w http.ResponseWriter, r *http.Request) {
	livro, err := parseLivroRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = livro.Insert(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/insert", http.StatusFound)
}

func parseLivroRequest(r *http.Request) (livro Livro, err error) {
	err = r.ParseForm()
	if err != nil {
		return
	}
	livro = Livro{
		ISBN:      r.FormValue("isbn"),
		Titulo:    r.FormValue("titulo"),
		Descricao: r.FormValue("descricao"),
	}
	if idInput := r.FormValue("id"); idInput != "" {
		var id int
		id, err = strconv.Atoi(idInput)
		if err != nil {
			return
		}
		livro.ID = id
	}
	if dataInput := r.FormValue("data_publicacao"); dataInput != "" {
		var data time.Time
		data, err = time.Parse("2006-01-02", dataInput)
		if err != nil {
			return
		}
		livro.DataDePublicacao = data
	}
	if precoInput := r.FormValue("preco"); precoInput != "" {
		var preco float64
		preco, err = strconv.ParseFloat(precoInput, 64)
		if err != nil {
			return
		}
		livro.Preco = preco
	}
	return
}
