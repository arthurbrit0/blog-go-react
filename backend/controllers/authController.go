package controllers

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/arthurbrit0/blog-backend/database"
	"github.com/arthurbrit0/blog-backend/models"
	"github.com/gofiber/fiber/v2"
)

func validarEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9.+%+\-]`) // funcao que usa regular expression para validar o email
	return Re.MatchString(email)
}

func Registrar(context *fiber.Ctx) error {
	var dados map[string]interface{} // variavel para armazenar os dados passados pelo usuario na requisicao

	var dadosUsuario models.Usuario

	if err := context.BodyParser(&dados); err != nil { // usando o BodyParser no contexto da requisição para transformar o json em uma struct
		fmt.Println("Não foi possível dar parse nos dados da requisição.")
	}

	if len(dados["senha"].(string)) <= 6 { // validando se a senha tem menos de 6 caracteres
		context.Status(400)
		return context.JSON(fiber.Map{
			"mensagem": "A senha tem que ter mais de 6 caracteres!",
		})
	}

	if !validarEmail(strings.TrimSpace(dados["email"].(string))) { // validando se o email é valido
		context.Status(400)
		return context.JSON(fiber.Map{
			"mensagem": "O email é inválido!",
		})
	}

	// passando o usuario cadastrado no banco de dados que tem o mesmo email do que foi enviado no corpo da requisicao
	database.DB.Where("email=?", strings.TrimSpace(dados["email"].(string))).First(&dadosUsuario)

	if dadosUsuario.Id != 0 { // se não houver esse usuario, o id será 0, mas se houver, retornaremos um erro
		context.Status(400)
		return context.JSON(fiber.Map{
			"mensagem": "O e-mail já está cadastrado!",
		})
	}

	usuario := models.Usuario{ // passando os dados do corpo da requisicao para o model usuario
		PrimeiroNome: dados["primeiro_nome"].(string),
		UltimoNome:   dados["ultimo_nome"].(string),
		Email:        strings.TrimSpace(dados["email"].(string)),
		Telefone:     dados["telefone"].(string),
	}

	usuario.SetSenha(dados["senha"].(string))
	err := database.DB.Create(&usuario)
	if err != nil {
		log.Println(err)
	}

	context.Status(201)

	return context.JSON(fiber.Map{
		"usuario":  usuario,
		"mensagem": "Conta criada com sucesso!",
	})
}
