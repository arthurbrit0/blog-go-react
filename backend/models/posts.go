package models

import "time"

type Post struct {
	Id        uint      `gorm:"primaryKey;autoIncrement" json:"id"` // modelo para o post, usando como chave estrangeira o id do usuario
	Titulo    string    `json:"titulo"`
	Descricao string    `json:"descricao"`
	Imagem    string    `json:"imagem"`
	CreatedAt time.Time `json:"created_at"`
	UsuarioId uint      `json:"usuarioid"`
	Usuario   Usuario   `json:"usuario" gorm:"foreignkey:UsuarioId;references:Id"`
}
