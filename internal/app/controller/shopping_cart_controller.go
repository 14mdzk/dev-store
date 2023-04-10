package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/14mdzk/dev-store/internal/app/service"
	"github.com/gin-gonic/gin"
)

type ShoppingCartController struct {
	service service.IShoppingCartService
}

func NewShoppingCartController(service service.IShoppingCartService) *ShoppingCartController {
	return &ShoppingCartController{service: service}
}

func (cc *ShoppingCartController) BrowseCart(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {
		log.Print(fmt.Errorf("%w", err))
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": gin.H{"message": err.Error()}})
		return
	}

	resp, err := cc.service.Browse(userId)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": gin.H{"message": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
	return
}

func (cc *ShoppingCartController) CreateCartItem(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("userId"))
	if err != nil {
		log.Print(fmt.Errorf("%w", err))
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": gin.H{"message": err.Error()}})
		return
	}

	resp, err := cc.service.Browse(userId)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": gin.H{"message": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
	return
}
