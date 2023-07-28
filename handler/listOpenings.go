package handler

import (
	"net/http"

	"github.com/JoaoPaulo-creator/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error linsting openings")
		return
	}

	sendSuccess(ctx, http.StatusOK, "list-openings", openings)

}
