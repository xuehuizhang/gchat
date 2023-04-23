package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitHealthRouter(r *gin.RouterGroup) {
	r.GET("/health", func(ctx *gin.Context) {
		time.Sleep(time.Second * 5)
		ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})
}
