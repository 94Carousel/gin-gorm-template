package api

import (
	"fmt"
	"gin-template/config"

	"github.com/gin-gonic/gin"
)

// API defined API struct
type API struct {
	Config *config.Config
}

// RespondOK Success
func RespondOK(ctx *gin.Context, body interface{}) {
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "",
		"body": body})
}

// RespondError failded
func RespondError(ctx *gin.Context, code int, msg interface{}) {
	ctx.JSON(200, gin.H{"code": code, "msg": msg})
}

// HandleNoRoute return error 404 when routes unused
func HandleNoRoute(ctx *gin.Context) {
	RespondError(ctx, 404, "page not found")
}

// HandleRecovery handle global recovery and return error 500 JSON
func HandleRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				str := fmt.Sprint(err)
				RespondError(ctx, 500, str)
			}
		}()
		ctx.Next()
	}
}
