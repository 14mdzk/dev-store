package controller

import (
	"net/http"
	"strconv"

	"github.com/14mdzk/dev-store/internal/app/schema"
	"github.com/14mdzk/dev-store/internal/app/service"
	"github.com/gin-gonic/gin"
)

type UserAddressController struct {
	service service.IUserAddressService
}

func NewUserAddressController(service service.IUserAddressService) *UserAddressController {
	return &UserAddressController{service: service}
}

func (cc *UserAddressController) BrowseUserAddress(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	resp, err := cc.service.BrowseAll(userId)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": gin.H{"message": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
	return
}

func (cc *UserAddressController) GetByIdUserAddress(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("addressId"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	resp, err := cc.service.GetById(id, userId)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
	return
}

func (cc *UserAddressController) CreateUserAddress(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	var req schema.CreateUserAddressReq
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	err = cc.service.Create(userId, req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success create user address"})
	return
}

func (cc *UserAddressController) UpdateUserAddress(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("addressId"))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	var req schema.CreateUserAddressReq
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	err = cc.service.Update(id, userId, req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success update user address"})
	return
}

func (cc *UserAddressController) DeleteUserAddress(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("addressId"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	err = cc.service.Delete(id, userId)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success delete user address"})
	return
}
