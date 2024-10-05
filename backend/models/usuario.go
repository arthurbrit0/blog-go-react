package models

import "golang.org/x/crypto/bcrypt"

type Usuario struct {
	Id           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	PrimeiroNome string `json:"primeiro_nome"`
	UltimoNome   string `json:"ultimo_nome"`
	Email        string `json:"email"`
	Senha        []byte `json:"-"`
	Telefone     string `json:"telefone"`
}

func (usuario *Usuario) SetSenha(senha string) {
	senhaHasheada, _ := bcrypt.GenerateFromPassword([]byte(senha), 14) // função para hashear a senha
	usuario.Senha = senhaHasheada                                      // setando a senha que será armazenada no banco de dados como a hasheada
}

func (usuario *Usuario) CompararSenha(senha string) error {
	return bcrypt.CompareHashAndPassword(usuario.Senha, []byte(senha)) // comparando a senha hasheada com a senha passada pelo usuario no login
}
