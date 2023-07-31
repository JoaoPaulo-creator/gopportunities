package handler

import (
	"net/http"

	"github.com/JoaoPaulo-creator/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)

	logger.Infof("request received: %+v ", request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.IsRemote != nil {
		opening.IsRemote = *request.IsRemote
	}

	if request.RoleLink != "" {
		opening.RoleLink = request.RoleLink
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error on update operation: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSuccess(ctx, http.StatusOK, "update-opening", opening)
}
