package routes

import (
	"github.com/arthurbrit0/blog-backend/controllers"
	"github.com/arthurbrit0/blog-backend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetarRotas(app *fiber.App) {

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // permitindo apenas requisições do localhost:5173
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// rotas de autenticação

	app.Post("/api/registrar", controllers.Registrar)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/me", controllers.GetMe)

	app.Use(middleware.IsUsuarioAutenticado) // protegendo as rotas que precisam de autenticacao

	// rotas de posts

	app.Post("/api/post", controllers.CriarPost)
	app.Get("/api/posts", controllers.GetTodosPosts)
	app.Get("/api/post/:id", controllers.GetDetalhesPost)
	app.Put("/api/post/:id", controllers.EditarPost)
	app.Get("/api/meusposts", controllers.GetMeusPosts)
	app.Delete("/api/post/:id", controllers.DeletarPost)

	// rotas de imagem

	app.Post("/api/upload", controllers.Upload)
	app.Static("/api/uploads", "./uploads")
}
