package api

import (
	"github.com/yingce/gin-gorm-template/config"
	"github.com/yingce/gin-gorm-template/middleware"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Router defined api namespace router initialize
func (api *API) Router() *gin.Engine {
	// initialize gin config
	appSecret := config.EnvConfig.SecretKey
	appName := config.EnvConfig.AppName
	sessionStore := sessions.NewCookieStore([]byte(appSecret))

	// defined router
	router := gin.Default()
	router.Use(HandleRecovery())
	router.Use(middleware.CORSMiddleware())
	router.Use(sessions.Sessions("_"+appName+"_session", sessionStore))
	router.GET("/", api.HomeIndex)
	router.NoRoute(HandleNoRoute)
	return router
}
