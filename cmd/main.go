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

	userRepository := repository.NewUserRepository(DBConn)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userAddressRepository := repository.NewUserAddressRepository(DBConn)
	userAddressService := service.NewUserAddressService(userAddressRepository, userRepository)
	userAddressController := controller.NewUserAddressController(userAddressService)

	productRepository := repository.NewProductRepository(DBConn)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	r.GET("/categories", categoryController.BrowseCategory)
	r.POST("/categories", categoryController.CreateCategory)
	r.GET("/categories/:id", categoryController.GetByIdCategory)
	r.DELETE("/categories/:id", categoryController.DeleteCategory)
	r.PATCH("/categories/:id", categoryController.UpdateCategory)

	r.GET("/users", userController.BrowseUser)
	r.POST("/users", userController.CreateUser)
	r.GET("/users/:id", userController.GetByIdUser)
	r.DELETE("/users/:id", userController.DeleteUser)
	r.PATCH("/users/:id", userController.UpdateUser)

	r.GET("/users/:id/addresses", userAddressController.BrowseUserAddress)
	r.POST("/users/:id/addresses", userAddressController.CreateUserAddress)
	r.GET("/users/:id/addresses/:addressId", userAddressController.GetByIdUserAddress)
	r.PATCH("/users/:id/addresses/:addressId", userAddressController.UpdateUserAddress)
	r.DELETE("/users/:id/addresses/:addressId", userAddressController.DeleteUserAddress)

	r.GET("/products", productController.BrowseProduct)
	r.POST("/products", productController.CreateProduct)
	r.GET("/products/:id", productController.GetByIdProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)
	r.PATCH("/products/:id", productController.UpdateProduct)

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	r.Run(appPort)
}
