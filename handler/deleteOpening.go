package handler

import (
	"fmt"
	"net/http"

	"github.com/JoaoPaulo-creator/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete an existing opening
// @Description Deletes a job opening
// @Tags Openings
// @Accept json
// @Procedure json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s ", id))
		return
	}

	response := schemas.OpeningResponse{
		CreatedAt: opening.CreatedAt,
		DeletedAt: opening.DeletedAt.Time,
		UpdatedAt: opening.UpdatedAt,
		Role:      opening.Role,
		Company:   opening.Location,
		Location:  opening.Location,
		IsRemote:  opening.IsRemote,
		RoleLink:  opening.RoleLink,
		Salary:    opening.Salary,
	}

	sendSuccess(ctx, http.StatusOK, "delete-opening", response)

}
