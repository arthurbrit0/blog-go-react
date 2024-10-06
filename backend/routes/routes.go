package routes

import (
	"github.com/arthurbrit0/blog-backend/controllers"
	"github.com/arthurbrit0/blog-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetarRotas(app *fiber.App) {
	app.Post("/api/registrar", controllers.Registrar)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsUsuarioAutenticado) // protegendo as rotas que precisam de autenticacao

	app.Post("/api/post", controllers.CriarPost)
}
