package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConectaBD() *sql.DB {
	conexao := "user=postgres dbname=alura-loja password=349420 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
