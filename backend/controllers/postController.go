package controllers

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/arthurbrit0/blog-backend/database"
	"github.com/arthurbrit0/blog-backend/models"
	"github.com/arthurbrit0/blog-backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CriarPost(context *fiber.Ctx) error {
	var postBlog models.Post // criando variavel postBlog para armazenar os dados da requisicao em formato de post

	if err := context.BodyParser(&postBlog); err != nil { // dando parse no corpo da requiiscao para o postBlog
		fmt.Println("Não foi possível dar parse no corpo da requisição")
		context.Status(400)
		return context.JSON(fiber.Map{
			"mensagem": "Payload inválido",
		})
	}

	usuarioID, err := utils.ParseJWT(context.Cookies("jwt")) // pegando o id do usuario autenticado a partir do cookie jwt
	if err != nil {
		context.Status(400)
		return context.JSON(fiber.Map{
			"mensagem": "Não foi possível pegar o id do usuario",
		})
	}

	usuarioIDInt, err := strconv.Atoi(usuarioID) // captura tanto o valor quanto o erro
	if err != nil {
		context.Status(400)
		return context.JSON(fiber.Map{
			"mensagem": "ID do usuário inválido",
		})
	}

	postBlog.UsuarioId = uint(usuarioIDInt) // passando o id do usuario autenticado para o campo UsuarioId do postBlog

	if err := database.DB.Create(&postBlog).Error; err != nil { // criando postBlog no banco de dados
		context.Status(400)
		return context.JSON(fiber.Map{
			"mensagem": "Payload inválido",
		})
	}

	return context.JSON(fiber.Map{ // caso nao tenha erro, retornar mensagem de sucesso
		"mensagem": "Parabéns, seu post foi feito com sucesso!",
	})
}

func GetTodosPosts(context *fiber.Ctx) error { // função para pegar todo os posts com paginação
	pagina, _ := strconv.Atoi(context.Query("pagina", "1")) // pegando o valor da pagina a partir da query pagina, presente na url
	limite := 5                                             // por exemplo: localhost:3000/posts?pagina=1, pagina=1 é a query, e o valor da pag é 1
	offset := (pagina - 1) * limite                         // adicionamos um limite de 5 posts por pagina, e calculamos o offset (a partir de qual post a busca vai começar)

	var total int64            // variável para armazenar o total de posts no banco de dados
	var postBlog []models.Post // variavel para armazenar os posts

	database.DB.Preload("Usuario").Order("created_at DESC").Offset(offset).Limit(limite).Find(&postBlog) // populando o campo usuario de cada post, com offset e limite
	database.DB.Model(&models.Post{}).Count(&total)                                                      // contando o total de posts no banco de dados

	return context.JSON(fiber.Map{
		"data": postBlog, // retornamos os posts do blog, já populados com as informacoes de cada autor
		"meta": fiber.Map{
			"total":         total,                                       // total de posts
			"pagina":        pagina,                                      // pagina atual
			"ultima_pagina": math.Ceil(float64(total) / float64(limite)), // ultima pagina
		},
	})
}

func GetDetalhesPost(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id")) // pegando o id do post a partir dos parametros da url

	if err != nil {
		return context.JSON(fiber.Map{
			"erro": "ID do post inválido!",
		})
	}

	var postBlog models.Post // variável para armazenar os dados do post específico

	database.DB.Where("id = ?", id).Preload("Usuario").First(&postBlog) // buscando o post com o id passado no url no banco de dados, populando o campo usuario

	return context.JSON(fiber.Map{
		"mensagem": "Post encontrado com sucesso!",
		"data":     postBlog, // retornando os dados do post se ele tiver sido encontrado no banco de dados
	})

}

func EditarPost(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id")) // pegando o id do post que sera atualizado a partir dos parametros da url

	if err != nil {
		return context.JSON(fiber.Map{
			"erro": "ID do post inválido!", // retornando um erro caso o id do post seja invalido
		})
	}

	var postBlog models.Post // variável para armazenar os dados do post específico

	if err := context.BodyParser(&postBlog); err != nil { // dando parse no corpo da requiscao para o postBlog
		fmt.Println("Não foi possível dar parse no corpo da requisição")
	}

	database.DB.Model(&postBlog).Where("id = ?", id).Updates(&postBlog) // atualizando os dados do post com o id passado na url

	return context.JSON(fiber.Map{
		"mensagem": "Post atualizado com sucesso!",
	})
}

func GetMeusPosts(context *fiber.Ctx) error {
	cookie := context.Cookies("jwt") // pegando o cookie jwt

	id, err := utils.ParseJWT(cookie) // parseando o jwt para pegar o id do usuario
	if err != nil {
		return context.JSON(fiber.Map{
			"erro": "Não foi possível pegar o id do usuario",
		})
	}

	var posts = []models.Post{}                                                                      // criando variavel para armazenar os posts do usuario autenticado
	database.DB.Where("usuario_id = ?", id).Order("created_at DESC").Preload("Usuario").Find(&posts) // buscando os posts do usuario autenticado
	return context.JSON(posts)

}

func DeletarPost(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id")) // pegando o id a partir dos parametros da url

	if err != nil {
		return context.JSON(fiber.Map{ // se não conseguirmos converter o id dos parametros, retornamos um erro
			"erro": "ID do post inválido",
		})
	}

	var post models.Post // criando uma variavel para armazenar os dados do post com o id especifico

	if err := database.DB.First(&post, id).Error; err != nil {
		context.Status(404)
		return context.JSON(fiber.Map{
			"erro": "Post não encontrado",
		})
	}

	deleteQuery := database.DB.Delete(&post) // deletando esse post que criamos do banco de dados

	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		context.Status(404)
		return context.JSON(fiber.Map{
			"erro": "Post não encontrado",
		})
	}

	return context.JSON(fiber.Map{
		"mensagem": "Post deletado com sucesso",
	})

}
