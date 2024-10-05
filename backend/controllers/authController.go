package controllers

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/arthurbrit0/blog-backend/database"
	"github.com/arthurbrit0/blog-backend/models"
	"github.com/arthurbrit0/blog-backend/utils"
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

func Login(context *fiber.Ctx) error {
	var dados map[string]string // variavel para armazenar os dados passados pelo usuario na requisicao

	if err := context.BodyParser(&dados); err != nil { // usando o BodyParser no contexto da requisição para transformar o json em uma struct
		fmt.Println("Não foi possível dar parse nos dados da requisição.")
	}

	var dadosUsuario models.Usuario                                     // variavel para armazenar os dados do usuario mandados na requisicao
	database.DB.Where("email = ?", dados["email"]).First(&dadosUsuario) // checando no banco de dados se existe um email cadastrado
	if dadosUsuario.Id == 0 {                                           // se não existir, a função retorna um ID 0
		context.Status(404)
		return context.JSON(fiber.Map{ // retornamos um erro em formato json para o usuario caso o email não tenha sido encontrado
			"mensagem": "Email não existe! Crie uma conta",
		})
	}

	if err := dadosUsuario.CompararSenha(dados["senha"]); err != nil { // comparando a senha enviada na requisicao com a senha hasheada
		context.Status(400) // do usuario que achamos no bdd com o email passado na requisicao
		return context.JSON(fiber.Map{
			"mensagem": "Senha incorreta",
		})
	}

	token, err := utils.GerarJWT(strconv.Itoa(int(dadosUsuario.Id))) // se estiver tudo certo com o usuario e sua senha, geramos um token pra ele com
	if err != nil {                                                  // a funcao GerarJWT feita no diretorio utils.
		context.Status(fiber.StatusInternalServerError) // passamos como issuer desse jwt o id do usuario
		return nil
	}

	cookie := fiber.Cookie{ // criando um cookie com o token que criamos para o usuario
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	context.Cookie(&cookie) // armazenando esse cookie

	return context.JSON(fiber.Map{ // retornando para o usuario, apos todas as verificacoes, a respota de sucesso do login e seus dados
		"mensagem":      "Usuário logado com sucesso!",
		"dados_usuario": dadosUsuario,
	})
}
