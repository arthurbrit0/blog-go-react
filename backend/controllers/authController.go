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
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func validarEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9.+%+\-]`) // funcao que usa regular expression para validar o email
	return Re.MatchString(email)
}

func Registrar(context *fiber.Ctx) error {
	var dados map[string]interface{} // variavel para armazenar os dados passados pelo usuario na requisicao

	if err := context.BodyParser(&dados); err != nil { // dando parse do corpo na requisicao no map dados
		fmt.Println("Não foi possível dar parse nos dados da requisição.")
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"mensagem": "Dados inválidos!",
		})
	}

	senha, senhaOk := dados["senha"].(string) // validando se a senha existe e é uma string
	if !senhaOk || senha == "" {
		return context.Status(400).JSON(fiber.Map{
			"mensagem": "A senha é obrigatória e deve ser uma string.",
		})
	}

	if len(senha) < 6 { // validando se a senha tem mais de 6 caracteres
		return context.Status(400).JSON(fiber.Map{
			"mensagem": "A senha deve ter mais de 6 caracteres.",
		})
	}

	email, emailOk := dados["email"].(string) // validando se o email existe e é uma string
	if !emailOk || email == "" {
		return context.Status(400).JSON(fiber.Map{
			"mensagem": "O email é obrigatório e deve ser uma string.",
		})
	}

	if !validarEmail(strings.TrimSpace(email)) { // usando a funcao com regex para validar o email passado na requisicao
		return context.Status(400).JSON(fiber.Map{
			"mensagem": "O email é inválido!",
		})
	}

	primeiroNome, primeiroNomeOk := dados["primeiro_nome"].(string) // verificando se o primeiro nome existe e é uma string
	if !primeiroNomeOk || strings.TrimSpace(primeiroNome) == "" {
		return context.Status(400).JSON(fiber.Map{
			"mensagem": "Primeiro nome é obrigatório.",
		})
	}

	ultimoNome, ultimoNomeOk := dados["ultimo_nome"].(string) // verificando se o ultimo nome existe e é uma string
	if !ultimoNomeOk || strings.TrimSpace(ultimoNome) == "" {
		return context.Status(400).JSON(fiber.Map{
			"mensagem": "Último nome é obrigatório.",
		})
	}

	telefone, telefoneOk := dados["telefone"].(string) // verificando se o telefone existe e é uma string
	if !telefoneOk {
		telefone = ""
	}

	var dadosUsuario models.Usuario // criando variavel para armazenar os dados do usuario
	if err := database.DB.Where("email = ?", strings.TrimSpace(email)).First(&dadosUsuario).Error; err == nil {
		return context.Status(400).JSON(fiber.Map{ // achando no banco de dados para ver se ja tem um usuario com esse email registrado
			"mensagem": "O e-mail já está cadastrado!",
		})
	}

	usuario := models.Usuario{ // criando um novo usuario com os dados passados na requisicao
		PrimeiroNome: primeiroNome,
		UltimoNome:   ultimoNome,
		Email:        strings.TrimSpace(email),
		Telefone:     telefone,
	}

	usuario.SetSenha(senha)                                    // usando o metodo do usuario para setar uma senha hasheada a partir da senha enviada na requisicao
	if err := database.DB.Create(&usuario).Error; err != nil { // criando o usuario no banco de dados
		log.Println("Erro ao criar usuário:", err)
		return context.Status(500).JSON(fiber.Map{
			"mensagem": "Erro ao criar usuário",
		})
	}

	return context.Status(201).JSON(fiber.Map{ // passando em todas as verificacoes, os dados do usuario sao retornados na resposta
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

type Claims struct {
	jwt.StandardClaims
}

func GetMe(context *fiber.Ctx) error {
	cookie := context.Cookies("jwt") // pegando o cookie jwt

	id, err := utils.ParseJWT(cookie) // parseando o jwt para pegar o id do usuario
	if err != nil {
		return context.JSON(fiber.Map{
			"erro": "Não foi possível pegar o id do usuario",
		})
	}

	var usuario models.Usuario                      // criando variavel para armazenar os dados do usuario autenticado
	database.DB.Where("id = ?", id).First(&usuario) // buscando o usuario no banco de dados

	return context.JSON(usuario) // retornando os dados do usuario
}
