package handler

import (
	"fmt"

	"github.com/JoaoPaulo-creator/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, statusCode int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"message":   msg,
		"errorCode": statusCode,
	})
}

func sendSuccess(ctx *gin.Context, statusCode int, operation string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", operation), //Sprintf retorna uma string concatenada
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
