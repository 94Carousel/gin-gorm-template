package main

import (
	"gin-template/config"
	"gin-template/controllers"
	"gin-template/models"
	"io"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// initialize logfile
	f := config.LogFile()
	defer f.Close()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// initialize database
	models.InitDB()
	defer models.DB.Close()

	// initialize gin config
	appSecret := config.Get("app", "secret").Value()
	appName := config.Get("app", "name").Value()
	store := sessions.NewCookieStore([]byte(appSecret))

	r := gin.Default()
	r.Use(sessions.Sessions(appName+"_session", store))

	api := r.Group("/api").Use(controllers.HandleRecovery())
	{
		api.GET("/", func(c *gin.Context) {
			session := sessions.Default(c)
			session.Set("count", "hello")
			session.Save()
			count := session.Get("count")
			controllers.RespondOK(count, c)
		})
	}

	r.NoRoute(controllers.HandleNoRoute)
	port := config.Get("app", "port").Value()
	if port == "" {
		port = "8888"
	}
	r.Run(":" + port)
}
