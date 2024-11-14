package database

import (
	"github.com/Ulpio/gin-api/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connStr := "host=localhost port=5432 user=root password=root dbname=root sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar ao banco")
	}
	db.AutoMigrate(&models.Livro{})
	DB = db
}
