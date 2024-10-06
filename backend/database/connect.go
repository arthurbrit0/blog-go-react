package database

import (
	"fmt"
	"log"
	"os"

	"github.com/arthurbrit0/blog-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func loadEnv() { // função que carrega as variáveis do ambiente usando o godotenv
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar as variáveis de ambiente: ", err)
	}
}

func getDSN() string {
	dsn := os.Getenv("DSN") // função que pega a variavel de ambiente DSN, que sera usada como config para conectar com o banco de dados
	if dsn == "" {
		log.Fatal("Não foi possível carregar a variável DSN")
	}

	return dsn
}

func connectDB() *gorm.DB { // função para conectar com o banco de dados usando o gorm, que retornará um ponteiro para o bdd
	dsn := getDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // abrindo um banco de dados usando o dsn (obtido na funcao getDSN), e passando as config defautl
	if err != nil {
		log.Panic("Erro ao se conectar com o banco de dados: ", err)
	}

	return db
}

func Connect() {
	loadEnv()                                                            // carregando as variaveis de ambiente
	DB = connectDB()                                                     // obtendo o ponteiro para o banco de dados
	fmt.Println("Conexão com o banco de dados estabelecida com sucesso") // se não houver nenhum erro nas funções chamadas acima, o banco foi conectado com sucesso

	DB.AutoMigrate(&models.Usuario{}, &models.Post{}) // usando a função AutoMigrate do gorm para migrar o modelo de Usuario para uma tabela do banco de dados mysql
}
