package middleware

import (
	"github.com/arthurbrit0/blog-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func IsUsuarioAutenticado(context *fiber.Ctx) error { // função para ver se o usuário está autenticado, ou seja, se tem um token jwt válido para, então, acessar as rotas da aplicação
	cookie := context.Cookies("jwt") // pegando o cookie que contém o token jwt a partir da requisição de um usuário

	if _, err := utils.ParseJWT(cookie); err != nil { // verificando se há erro na hora de validar o cookie com a função do diretorio utils ParseJWT
		context.Status(fiber.StatusUnauthorized)
		return context.JSON(fiber.Map{
			"erro": "Você não está autorizado a acessar essa rota!", // se houver erro, o usuário não será autorizado a acessar aquela rota específica
		})
	}

	return context.Next()
}

// FUNÇÃO PARSEJWT PARA COMPARAÇÃO:

/*

	func ParseJWT(cookie string) (string, error) {
		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) { ----> Pega o cookie que enviamos,
			return []byte(SecretKey), nil                                                                          ----> as claims padrões do pacote jwt
		})                                                                                                         ----> e uma função para retornar a chave secreta

		if err != nil || !token.Valid {                                                                            ----> Se o token for inválido, retornamos um erro
			return "", err
		}

		claims := token.Claims.(*jwt.StandardClaims)                                                               -----> Se não, retornamos o issuer do token

		return claims.Issuer

*/
