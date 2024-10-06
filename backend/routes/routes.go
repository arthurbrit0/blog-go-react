package routes

import (
	"github.com/arthurbrit0/blog-backend/controllers"
	"github.com/arthurbrit0/blog-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetarRotas(app *fiber.App) {

	// rotas de autenticação

	app.Post("/api/registrar", controllers.Registrar)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsUsuarioAutenticado) // protegendo as rotas que precisam de autenticacao

	// rotas de posts

	app.Post("/api/post", controllers.CriarPost)
	app.Get("/api/posts", controllers.GetTodosPosts)
	app.Get("/api/post/:id", controllers.GetDetalhesPost)
	app.Put("/api/post/:id", controllers.EditarPost)
	app.Get("/api/meusposts", controllers.GetMeusPosts)
	app.Delete("/api/post/:id", controllers.DeletarPost)
}
