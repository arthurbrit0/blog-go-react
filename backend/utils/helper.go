package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "secret"

func GerarJWT(issuer string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{ // criando um novo token com o algoritmo HS256
		Issuer:    issuer, // passando as claims, presentes no payload do token, de quem emitiu o token e sua validade
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(SecretKey)) // retornando o token assinado com a Secret Key, presente na Signature do token
}

func ParseJWT(cookie string) (string, error) { // função para validar um token jwt contido em um cookie
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		// na função ParseWithClaims, recebemos o cookie que contem o token jwt, as claims que esperamos que estejam no token,
		// e uma função que retorna a chave secreta que definimos para fazer a verificacao de que o token foi realmente assinado com essa chave
		return []byte(SecretKey), nil // aqui,retornamos a chave secreta para verificacao do token
	})

	if err != nil || !token.Valid { // se houver algum erro no parse do token ou se ele não for mais válido, retornamos erro
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims) // fazendo type assertion para dizermos que os claims do token são do tipo standartclaims, especificamente

	return claims.Issuer, nil // retornamos o nome de quem emitiu o cookie
}
