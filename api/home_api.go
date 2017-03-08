package api

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// HomeIndex GET /
func (api *API) HomeIndex(ctx *gin.Context) {
	session := sessions.Default(ctx)
	count := session.Get("count")
	if count == nil {
		session.Set("count", "hello")
		session.Save()
	}
	body := gin.H{"count": count}
	RespondOK(ctx, body)
}
