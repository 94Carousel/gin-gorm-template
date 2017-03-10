package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/yingce/gin-gorm-template/api"
	"github.com/yingce/gin-gorm-template/config"
	"github.com/yingce/gin-gorm-template/models"
	"github.com/yingce/gin-gorm-template/store"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	conf := config.InitEnv()
	// initialize logfile
	f := config.LogFile()
	defer f.Close()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// initialize database ORM
	models.InitDB()
	defer models.DB.Close()

	store.InitRedis()
	// Gin run environment
	gin.SetMode(conf.RunMode)

	// HTTPServer
	api := &api.API{Config: conf}
	server := &http.Server{
		Addr:           conf.Addr,
		Handler:        api.Router(),
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("can't start server: ", err)
	}
}
