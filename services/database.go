package services

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var (
	DB    *sql.DB
	errDb error
)

func InitDB() {
	stringConn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"locahost", 5432, "postgres", "123456", "postgres")
	DB, errDb = sql.Open("postgres", stringConn)
	if errDb != nil {
		log.Fatal(errDb)
	} else {
		log.Println("Database iniciado com sucesso!")
	}
}
