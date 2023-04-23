package handler

import (
	"com.gchat.pkg/enum"
	"com.gchat/connect/conn"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
	"time"
)

type WebsocketHandler struct {
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (cs *WebsocketHandler) Handler(resp http.ResponseWriter, req *http.Request) {
	//升级websocket
	wsConn, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		zap.S().Fatalf("websocket 服务升级失败 %v", err.Error())
		return
	}

	c := &conn.Conn{
		ConnType: enum.WsConn,
		WsConn:   wsConn,
	}

	procConn(c)
}

func procConn(conn *conn.Conn) {
	defer func() {
		err := recover()
		if err != nil {
			zap.S().Infof("处理连接异常：%v", err)
		}
	}()

	for {
		err := conn.WsConn.SetReadDeadline(time.Now().Add(8 * time.Minute))
		if err != nil {
			HandleReadErr(conn, err)
			return
		}

		_, data, err := conn.WsConn.ReadMessage()
		if err != nil {
			HandleReadErr(conn, err)
			return
		}

		conn.HandleMessage(data)
	}
}

func HandleReadErr(c *conn.Conn, err error) {
	zap.S().Infof("read conn error userId:%v", c.UserId)
	str := err.Error()
	if strings.HasPrefix(str, "use of closed network connection") {
		return
	}

	loadConn := conn.GetConn(c.UserId)
	if loadConn != nil {
		conn.RemoveConn(c.UserId) //已鉴权
	} else {
		c.Close() //未鉴权，直接关闭
	}

	//客户端关闭连接，或者异常退出
	if err == io.EOF {
		return
	}

	//readDeadline 超时错误
	if strings.HasSuffix(str, "i/o timeout") {
		return
	}
}
