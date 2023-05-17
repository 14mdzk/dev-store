package controller

import (
	"net/http"
	"strconv"

	"github.com/14mdzk/dev-store/internal/app/schema"
	"github.com/14mdzk/dev-store/internal/pkg/handler"
	"github.com/14mdzk/dev-store/internal/pkg/reason"
	"github.com/gin-gonic/gin"
)

type SessionService interface {
	Login(req *schema.LoginReq) (schema.LoginResp, error)
	Logout(userID int) error
	Refresh(req *schema.RefreshTokenReq) (schema.RefreshTokenResp, error)
}

type RefreshTokenVerifier interface {
	VerifyRefreshToken(tokenString string) (string, error)
}

type SessionController struct {
	service    SessionService
	tokenMaker RefreshTokenVerifier
}

func NewSessionController(service SessionService, tokenMaker RefreshTokenVerifier) *SessionController {
	return &SessionController{
		service:    service,
		tokenMaker: tokenMaker,
	}
}

func (c *SessionController) Login(ctx *gin.Context) {
	req := &schema.LoginReq{}
	if !handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := c.service.Login(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success login", resp)
}

func (c *SessionController) Refresh(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("refresh_token")
	sub, err := c.tokenMaker.VerifyRefreshToken(refreshToken)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, reason.RefreshTokenFailed)
		return
	}

	intSub, _ := strconv.Atoi(sub)
	req := &schema.RefreshTokenReq{}
	req.RefreshToken = refreshToken
	req.UserID = intSub

	resp, err := c.service.Refresh(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.RefreshTokenFailed)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success refresh token", resp)
}

func (c *SessionController) Logout(ctx *gin.Context) {

}
