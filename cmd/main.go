package main

import (
	"fmt"
	"log"

	"github.com/14mdzk/dev-store/internal/app/controller"
	"github.com/14mdzk/dev-store/internal/app/repository"
	"github.com/14mdzk/dev-store/internal/app/service"
	"github.com/14mdzk/dev-store/internal/pkg/config"
	"github.com/14mdzk/dev-store/internal/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var (
	cfg    config.Config
	DBConn *sqlx.DB
)

func init() {
	configLoad, err := config.LoadConfig(".")
	if err != nil {
		log.Panic(err.Error())
	}
	cfg = configLoad

	db, err := db.ConnectDB(cfg.DBDriver, cfg.DBConnection)
	if err != nil {
		log.Panic(err.Error())
	}

	DBConn = db
}

func main() {
	r := gin.Default()

	categoryRepository := repository.NewCategoryRepository(DBConn)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	r.GET("/categories", categoryController.BrowseCategory)
	r.POST("/categories", categoryController.CreateCategory)
	r.GET("/categories/:id", categoryController.GetByIdCategory)
	r.DELETE("/categories/:id", categoryController.DeleteCategory)
	r.PATCH("/categories/:id", categoryController.UpdateCategory)

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	r.Run(appPort)
}
