package api

import (
	"com.gchat.business/user/api/bo"
	"com.gchat.business/user/api/vo"
	"com.gchat.pkg/app"
	"com.gchat.pkg/enum"
	"com.gchat.pkg/jwtUt"
	"github.com/gin-gonic/gin"
)

// @Summary  用户鉴权
// @Tags User
// @Accept json
// @Produce  json
// @Param   body  body   bo.UserAuthBo true "body"
// @Success 200 {string}  string
// @Failure 400 {string} string
// @Router /api/user/login  [POST]
func Auth(ctx *gin.Context) {
	appG := app.App{ctx}
	b := bo.UserAuthBo{}
	err := appG.C.BindJSON(&b)
	if err != nil {
		appG.ResponseErr(enum.InvalidParams)
		return
	}
	c, err := jwtUt.ParseToken(b.Token)
	if err != nil {
		appG.ResponseErr(enum.InvalidTokenMsg)
		return
	}

	v := vo.UserAuthVo{UserId: c.Id}

	appG.ResponseOk(v)
}
