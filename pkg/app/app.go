package app

import (
	"com.gchat.pkg/enum"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type App struct {
	C *gin.Context
}

type Response struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func (a *App) ResponseOk(data interface{}) {
	a.C.JSON(http.StatusOK, Response{
		Code:      0,
		Msg:       "ok",
		Data:      data,
		Timestamp: time.Now().Unix(),
	})
}

func (a *App) ResponseErr(err string) {
	a.C.JSON(http.StatusOK, Response{
		Code:      1,
		Msg:       err,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	})
}

func (a *App) ResponseCode(code int, err string, data interface{}) {
	a.C.JSON(http.StatusOK, Response{
		Code:      code,
		Msg:       err,
		Data:      data,
		Timestamp: time.Now().Unix(),
	})
}

func (a *App) InvalidToken() {
	a.C.JSON(http.StatusOK, Response{
		Code:      enum.InvalidToken,
		Msg:       enum.InvalidTokenMsg,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	})
}
