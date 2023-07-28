package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model // serve para fazer o unwrap do gorm, sendo possivel acessar as propriedades do gorm
	// Aparentemente, as propriedades tem que obrigatoriamente possuir a primeira letra maiuscula
	// Pelo que vi, se for minusculo, nao é gerado uma base com todos os campos
	Role     string
	Company  string
	Location string
	IsRemote bool
	RoleLink string
	Salary   int64
}

// Objeto a ser retornado nas chamadas http

// A tag `json:"algumacoisa"` dita como deve ser o nome e ortografia de cada prop no body da request
type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"` // se o campo for falsy/vazio, etc, o campo sera omitido
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	IsRemote  bool      `json:"isRemote"`
	RoleLink  string    `json:"roleLink"`
	Salary    int64     `json:"salary"`
}
