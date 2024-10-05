package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar as variáveis de ambiente: ", err)
	}
}

func getDSN() string {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("Não foi possível carregar a variável DSN")
	}

	return dsn
}

func connectDB() *gorm.DB {
	dsn := getDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao se conectar com o banco de dados: ", err)
	}

	return db
}

func Connect() {
	loadEnv()
	DB = connectDB()
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso")
}
