package schemas

import (
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
