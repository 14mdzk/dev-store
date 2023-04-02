package controller

import (
	"net/http"
	"strconv"

	"github.com/14mdzk/dev-store/internal/app/schema"
	"github.com/14mdzk/dev-store/internal/app/service"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service service.ICategoryService
}

func NewCategoryController(service service.ICategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (cc *CategoryController) BrowseCategory(ctx *gin.Context) {
	resp, err := cc.service.BrowseAll()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": gin.H{"message": err.Error()}})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

func (cc *CategoryController) GetByIdCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
	}

	resp, err := cc.service.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

func (cc *CategoryController) CreateCategory(ctx *gin.Context) {
	var req schema.CreateCategoryReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
	}

	err = cc.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success create category"})
}

func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
	}

	var req schema.CreateCategoryReq
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
	}

	err = cc.service.Update(id, req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success update category"})
}

func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
	}

	err = cc.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success delete category"})
}
