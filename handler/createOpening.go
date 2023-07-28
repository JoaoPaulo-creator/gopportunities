package handler

import (
	"net/http"

	"github.com/JoaoPaulo-creator/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{} // as {} inicilizam minha struct

	ctx.BindJSON(&request)
	// %+v expande o que sera logado na hora da request
	logger.Infof("request received: %+v ", request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v ", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Location,
		Location: request.Location,
		IsRemote: *request.IsRemote,
		RoleLink: request.RoleLink,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error on creating opening: %+v ", err.Error())
		sendError(ctx, http.StatusInternalServerError, "an error occured while creating opening on database")
		return
	}

	sendSuccess(ctx, http.StatusCreated, "createOpenin", opening)

}
