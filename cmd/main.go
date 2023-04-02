package main

import (
	"log"
	"net/http"

	"github.com/14mdzk/devstore/internal/pkg/config"
	"github.com/14mdzk/devstore/internal/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var cfg config.Config
var DBConn *sqlx.DB

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
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.Run()
}
