package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// RespondOK Success
func RespondOK(body interface{}, c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "",
		"body": body})
}

// RespondError failded
func RespondError(code int, msg interface{}, c *gin.Context) {
	c.JSON(200, gin.H{"code": code, "msg": msg})
}

// HandleNoRoute return error 404 when routes unused
func HandleNoRoute(c *gin.Context) {
	RespondError(404, "page not found", c)
}

// HandleRecovery handle global recovery and return error 500 JSON
func HandleRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				str := fmt.Sprint(err)
				RespondError(500, str, c)
			}
		}()
		c.Next()
	}
}
