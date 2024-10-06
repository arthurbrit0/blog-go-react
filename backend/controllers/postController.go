package controllers

import (
	"fmt"
	"math"
	"strconv"

	"github.com/arthurbrit0/blog-backend/database"
	"github.com/arthurbrit0/blog-backend/models"
	"github.com/gofiber/fiber/v2"
)

// TODO: RELER E COMENTAR CONTROLLER DO POST

func CriarPost(context *fiber.Ctx) error {
	var postBlog models.Post

	if err := context.BodyParser(&postBlog); err != nil {
		fmt.Println("Não foi possível dar parse no corpo da requisição")
	}

	if err := database.DB.Create(&postBlog).Error; err != nil {
		context.Status(400)
		return context.JSON(fiber.Map{
			"mensagem": "Payload inválido",
		})
	}

	return context.JSON(fiber.Map{
		"mensagem": "Parabéns, seu post foi feito com sucesso!",
	})
}

func GetTodosPosts(context *fiber.Ctx) error {
	pagina, _ := strconv.Atoi(context.Query("pagina", "1"))
	limite := 5
	offset := (pagina - 1) & limite

	var total int64
	var postBlog []models.Post

	database.DB.Preload("Usuario").Offset(offset).Limit(limite).Find(&postBlog)
	database.DB.Model(&models.Post{}).Count(&total)

	return context.JSON(fiber.Map{
		"data": postBlog,
		"meta": fiber.Map{
			"total":         total,
			"pagina":        pagina,
			"ultima_pagina": math.Ceil(float64(int(total) / limite)),
		},
	})
}
