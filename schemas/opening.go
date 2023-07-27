package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model // serve para fazer o unwrap do gorm, sendo possivel acessar as propriedades do gorm

	role     string
	company  string
	location string
	isRemote bool
	roleLink string
	salary   int64
}
