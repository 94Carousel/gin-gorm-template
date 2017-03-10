package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yingce/gin-gorm-template/store"
)

// HomeIndex GET /
func (api *API) HomeIndex(ctx *gin.Context) {
	// session := sessions.Default(ctx)
	// count := session.Get("count")
	// if count == nil {
	// 	session.Set("count", "hello")
	// 	session.Save()
	// }
	// body := gin.H{"count": count, "hello": 100}
	// store.CacheSet("hello", body)
	// a := store.CacheGet("hello")
	id := ctx.Query("id")
	data, ok := store.CacheGet(id)
	fmt.Println(data, ok)
	if !ok {
		data = gin.H{"id": id, "time": time.Now()}
		store.CacheSet(id, data)
	}
	data["cache"] = ok
	RespondOK(ctx, data)
}
