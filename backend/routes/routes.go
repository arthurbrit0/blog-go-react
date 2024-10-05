package routes

import (
	"github.com/arthurbrit0/blog-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetarRotas(app *fiber.App) {
	app.Post("/api/registrar", controllers.Registrar)
}
