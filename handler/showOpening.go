package handler

import (
	"fmt"
	"net/http"

	"github.com/JoaoPaulo-creator/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	openings := schemas.Opening{}

	if err := db.First(&openings, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	sendSuccess(ctx, http.StatusOK, "find-opening", openings)

}
