package database

import (
	"database/sql"
	// go postgres driver
	_ "github.com/lib/pq"
)

var db *sql.DB

// Connect retorna uma conexão com o banco de dados e um erro caso não consiga conectar
func Connect() (*sql.DB, error) {
	if db == nil {
		var err error
		db, err = sql.Open("postgres", "postgres://postgres:@localhost/livraria?sslmode=disable")
		if err != nil {
			return nil, err
		}
		err = db.Ping()
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
