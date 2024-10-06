package controllers

import (
	"fmt"
	"math"
	"strconv"

	"github.com/arthurbrit0/blog-backend/database"
	"github.com/arthurbrit0/blog-backend/models"
	"github.com/gofiber/fiber/v2"
)

func CriarPost(context *fiber.Ctx) error {
	var postBlog models.Post // criando variavel postBlog para armazenar os dados da requisicao em formato de post

	if err := context.BodyParser(&postBlog); err != nil { // dando parse no corpo da requiiscao para o postBlog
		fmt.Println("Não foi possível dar parse no corpo da requisição")
	}

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

	database.DB.Preload("Usuario").Offset(offset).Limit(limite).Find(&postBlog) // populando o campo usuario de cada post, com offset e limite
	database.DB.Model(&models.Post{}).Count(&total)                             // contando o total de posts no banco de dados

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
