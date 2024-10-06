package controllers

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

var random = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func NomeRandom(n int) string {
	nome := make([]rune, n)
	for i := range nome {
		nome[i] = random[rand.Intn(len(random))]
	}
	return string(nome)
}

func Upload(context *fiber.Ctx) error {
	form, err := context.MultipartForm()
	if err != nil {
		return err
	}

	arquivos := form.File["imagem"]
	nome_arquivo := ""

	for _, arquivo := range arquivos {
		nome_arquivo = NomeRandom(10) + "-" + arquivo.Filename
		if err := context.SaveFile(arquivo, "./uploads/"+nome_arquivo); err != nil {
			return err
		}
	}

	return context.JSON(fiber.Map{
		"url": "http://localhost:3000/uploads/" + nome_arquivo,
	})

}
