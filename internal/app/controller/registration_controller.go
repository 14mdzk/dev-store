package controller

import (
	"net/http"

	"github.com/14mdzk/dev-store/internal/app/schema"
	"github.com/14mdzk/dev-store/internal/app/service"
	"github.com/14mdzk/dev-store/internal/pkg/handler"
	"github.com/14mdzk/dev-store/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type RegistrationController struct {
	service service.IRegistrationService
}

func NewRegistrationController(service service.IRegistrationService) *RegistrationController {
	return &RegistrationController{
		service: service,
	}
}

func (rc *RegistrationController) Register(ctx *gin.Context) {
	var req schema.RegistrationReq
	if !handler.BindAndCheck(ctx, &req) {
		return
	}

	err := rc.service.Register(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, reason.RegistrationSuccess, nil)
}
