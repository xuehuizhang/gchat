package api

import (
	"com.gchat.business/user/api/bo"
	"com.gchat.business/user/api/vo"
	"com.gchat.business/user/dao/userDao"
	"com.gchat.business/user/model"
	"com.gchat.pkg/app"
	"com.gchat.pkg/enum"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary  用户注册
// @Tags User
// @Accept json
// @Produce  json
// @Param   body  body   bo.UserRegisterBo true "body"
// @Success 200 {string}  string
// @Failure 400 {string} string
// @Router /api/user/register  [POST]
func Register(ctx *gin.Context) {
	appG := app.App{ctx}
	registerBo := bo.UserRegisterBo{}
	err := appG.C.BindJSON(&registerBo)
	if err != nil {
		appG.ResponseErr(enum.InvalidParams)
		return
	}

	//校验验证码

	registerVo := vo.UserRegisterVo{}

	//验证用户是否已经注册
	userInfo, err := userDao.GetByField("mobile", registerBo.Mobile)
	if err != nil {
		zap.S().Infof("Register ：查询用户失败 %v", err.Error())
		appG.ResponseErr(enum.InternalError)
		return
	}

	if userInfo != nil {
		registerVo.Nick = userInfo.Nick
		registerVo.Mobile = userInfo.Mobile
		appG.ResponseOk(registerVo)
		return
	}

	//注册
	user := &model.User{
		Base:   model.Base{Status: enum.USER_NORAML},
		Nick:   registerBo.Nick,
		Mobile: registerBo.Mobile,
	}
	err = userDao.Add(user)
	if err != nil {
		zap.S().Infof("用户注册失败：%v", err.Error())
		appG.ResponseErr(enum.UserRegisterError)
		return
	}

	registerVo.Nick = user.Nick
	registerVo.Mobile = user.Mobile

	appG.ResponseOk(registerVo)
}
