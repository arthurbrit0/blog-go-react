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

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			ctx.Set("Access-Control-Allow-Origin", "http://localhost:5173")
			ctx.Set("Access-Control-Allow-Credentials", "true")

			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return ctx.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	}) // inicializando o app fiber

	routes.SetarRotas(app) // setando as rotas da aplicacao

	app.Listen(":" + port) // passando uma porta para o app ouvir (no caso, 3000)
}
