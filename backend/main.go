package main

import (
	"log"
	"os"

	"github.com/arthurbrit0/blog-backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Não foi possível carregar as variáveis de ambiente")
	}

	port := os.Getenv("PORT")

	app := fiber.New()

	app.Listen(":" + port)
}
