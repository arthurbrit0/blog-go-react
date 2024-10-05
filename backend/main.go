package main

import (
	"log"
	"os"

	"github.com/arthurbrit0/blog-backend/database"
	"github.com/arthurbrit0/blog-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()     // conectando ao banco de dados
	err := godotenv.Load() // carregando as variaveis de ambiente para obter a porta que a api rodara

	if err != nil {
		log.Fatal("Não foi possível carregar as variáveis de ambiente")
	}

	port := os.Getenv("PORT")

	app := fiber.New() // inicializando o app fiber

	routes.SetarRotas(app) // setando as rotas da aplicacao

	app.Listen(":" + port) // passando uma porta para o app ouvir (no caso, 3000)
}
