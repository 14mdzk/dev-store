package handler

import (
	"net/http"

	"github.com/14mdzk/dev-store/internal/pkg/reason"
	"github.com/14mdzk/dev-store/internal/pkg/validator"
	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Status  string
	Message string
	Data    interface{}
}

func ResponseError(ctx *gin.Context, statusCode int, message string) {
	response := ResponseBody{
		Status:  "error",
		Message: message,
	}

	ctx.JSON(statusCode, response)
}

func ResponseSuccess(ctx *gin.Context, statusCode int, message string, data interface{}) {
	response := ResponseBody{
		Status:  "success",
		Message: message,
		Data:    data,
	}

	ctx.JSON(statusCode, response)
}

func BindAndCheck(ctx *gin.Context, data interface{}) bool {
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return false
	}

	isValid := validator.Check(data)
	if !isValid {
		ResponseError(ctx, http.StatusUnprocessableEntity, reason.RequestFormError)
		return false
	}

	return true
}
