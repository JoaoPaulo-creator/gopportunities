package router

import "github.com/gin-gonic/gin"

/*

Faz se necessário criar funções onde a primeira letra de seu nome, deve ser maiusculo,
pois assim é indicado de que esssa função será exportada.

Funções com o nome todo em caixa baixa, são funções/variáveis locais
*/

func Initialize() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.Run(":3000")
}
