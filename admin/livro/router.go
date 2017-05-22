package livro

import (
	"github.com/gorilla/mux"
)

const (
	path = "/livro"
)

func routes(router *mux.Router) {
	router.Path("/insert").HandlerFunc(InsertForm).Methods("GET")
	router.Path("/insert").HandlerFunc(SaveBook).Methods("POST")
	router.Path("/list").HandlerFunc(ListBooks).Methods("GET")
}

// RegisterAllRoutes registro todas as rotas de livro
func RegisterAllRoutes(router *mux.Router) {
	livroRouter := router.PathPrefix(path).Subrouter()
	routes(livroRouter)
}
