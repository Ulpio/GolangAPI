package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	connStr := "host=localhost port=5432 user=root password=root dbname=root sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Falha ao conectar ao banco", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Banco tá fora do alcance.\n", err)
	} else {
		fmt.Println("Banco conectado")
	}
}
