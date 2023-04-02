package controller

import (
	"net/http"
	"strconv"

	"github.com/14mdzk/dev-store/internal/app/schema"
	"github.com/14mdzk/dev-store/internal/app/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.IUserService
}

func NewUserController(service service.IUserService) *UserController {
	return &UserController{service: service}
}

func (cc *UserController) BrowseUser(ctx *gin.Context) {
	resp, err := cc.service.BrowseAll()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": gin.H{"message": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
	return
}

func (cc *UserController) GetByIdUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	resp, err := cc.service.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
	return
}

func (cc *UserController) CreateUser(ctx *gin.Context) {
	var req schema.CreateUserReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	err = cc.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success create User"})
	return
}

func (cc *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	var req schema.CreateUserReq
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	err = cc.service.Update(id, req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success update User"})
	return
}

func (cc *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	err = cc.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success delete User"})
	return
}
