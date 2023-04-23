package api

import (
	"com.gchat.pkg/app"
	"com.gchat.pkg/enum"
	"com.gchat/connect/conn"
	"com.gchat/connect/message"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type MessageBo struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Data   []byte `json:"data"`
	UserId int64  `json:"user_id"`
}

func SendMsg(ctx *gin.Context) {
	appG := app.App{C: ctx}
	bo := &MessageBo{}
	err := appG.C.BindJSON(&bo)
	if err != nil {
		appG.ResponseErr(enum.InvalidParams)
		return
	}

	c := conn.GetConn(bo.UserId)
	if c == nil {
		appG.ResponseErr(enum.ConnNotFoundError)
		return
	}

	replyMessage := message.ReplyMessage{
		MsgType: enum.BusinessMsg,
		Code:    bo.Code,
		Msg:     bo.Msg,
		Data:    bo.Data,
	}
	data, err := json.Marshal(replyMessage)
	if err != nil {
		appG.ResponseErr(enum.ConnNotFoundError)
		return
	}
	c.Send(enum.BusinessMsg, data, enum.SerializeErrCode)
	appG.ResponseOk("ok")
}
