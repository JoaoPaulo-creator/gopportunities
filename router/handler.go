package router

import (
	"github.com/JoaoPaulo-creator/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRouter(router *gin.Engine) {

	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ListOpeningsHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
	}
}
