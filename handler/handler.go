package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "createing opening",
	})
}

func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "createing opening",
	})
}
func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "createing opening",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "createing opening",
	})
}

func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "createing opening",
	})
}
