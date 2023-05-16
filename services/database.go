package services

import (
	"ApiRest/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB    *gorm.DB
	errDb error
)

func InitDB() {
	stringConn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"127.0.0.1", "postgres", "123456", "postgres", "5432")
	fmt.Println(stringConn)
	DB, errDb = gorm.Open(postgres.Open(stringConn), &gorm.Config{})
	if errDb != nil {
		log.Fatal(errDb)
	} else {
		log.Println("Database iniciado com sucesso!")
		err := DB.AutoMigrate(&models.User{})
		if err != nil {
			log.Fatal("Erro ao criar tabela de usuários")
		}
	}
}
