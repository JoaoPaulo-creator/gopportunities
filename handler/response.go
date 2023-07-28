package handler

import (
	"fmt"

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
		"message": fmt.Sprintf("operation from handleer: %s successfull", operation), //Sprintf retorna uma string concatenada
		"data":    data,
	})
}
