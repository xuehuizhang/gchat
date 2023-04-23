package api

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	g := r.Group("/api")
	{
		g.POST("/send", SendMsg)
	}
}
