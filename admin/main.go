package main

import (
	"net/http"

	"github.com/felipeweb/livraria/admin/livro"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá mundo Go!"))
	})
	registerAllRoutes(router)
	http.ListenAndServe(":8081", router)
}

func registerAllRoutes(router *mux.Router) {
	livro.RegisterAllRoutes(router)
}
