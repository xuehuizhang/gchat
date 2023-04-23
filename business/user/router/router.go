package router

import (
	_ "com.gchat.business/user/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(r *gin.Engine) {
	g := r.Group("/api")

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	InitHealthRouter(g)

	InitUserRouter(g)
}
