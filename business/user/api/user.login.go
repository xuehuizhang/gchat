package api

import (
	"com.gchat.business/user/api/bo"
	"com.gchat.business/user/api/vo"
	"com.gchat.business/user/dao/userDao"
	"com.gchat.pkg/app"
	"com.gchat.pkg/enum"
	"com.gchat.pkg/jwtUt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary  用户登录
// @Tags User
// @Accept json
// @Produce  json
// @Param   body  body   bo.UserLoginBo true "body"
// @Success 200 {string}  string
// @Failure 400 {string} string
// @Router /api/user/login  [POST]
func Login(ctx *gin.Context) {
	appG := app.App{ctx}

	b := bo.UserLoginBo{}

	err := appG.C.BindJSON(&b)
	if err != nil {
		appG.ResponseErr(enum.InvalidParams)
		return
	}

	//判断用户是否存在
	userInfo, err := userDao.GetByField("mobile", b.Mobile)
	if err != nil {
		zap.S().Infof("Login 查询用户失败 %v", err.Error())
		appG.ResponseErr(enum.InternalError)
		return
	}

	if userInfo == nil {
		zap.S().Infof("Login 查询用户失败")
		appG.ResponseErr(enum.UserNoRegisterError)
		return
	}

	token, err := jwtUt.GenerateToken(userInfo.Id, userInfo.Nick, "", enum.MobileLogin)
	if err != nil {
		zap.S().Infof("Login 生成token失败 %v", err.Error())
		appG.ResponseErr(enum.InternalError)
		return
	}

	v := vo.UserLoginVo{
		Nick: userInfo.Nick,
		Toke: token,
	}
	appG.ResponseOk(v)
}
