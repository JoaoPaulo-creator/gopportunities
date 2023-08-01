package handler

import (
	"net/http"

	"github.com/JoaoPaulo-creator/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show a list of openings
// @Description shows a list of existing openings
// @Tags Openings
// @Accept json
// @Procedure json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error linsting openings")
		return
	}

	sendSuccess(ctx, http.StatusOK, "list-openings", openings)

}
