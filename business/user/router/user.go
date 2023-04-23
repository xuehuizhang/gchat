package router

import (
	"com.gchat.business/user/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(g *gin.RouterGroup) {
	u := g.Group("/user")
	{
		u.POST("/register", api.Register)
		u.POST("/login", api.Login)
		u.POST("/auth", api.Auth)
	}
}
