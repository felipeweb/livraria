package livro

import (
	"encoding/json"
	"html/template"
	"net/http"
)

// InsertForm renderiza o formulário
func InsertForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("form.html").ParseFiles("livro/tmpl/form.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SaveBook salva o livro
func SaveBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := make(map[string]interface{})
	for key, value := range r.PostForm {
		form[key] = value[0]
	}
	byt, err := json.Marshal(&form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byt)
}
