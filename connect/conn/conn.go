package conn

import (
	"com.gchat.pkg/enum"
	"com.gchat/connect/client/busi_client"
	"com.gchat/connect/message"
	"encoding/json"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"time"
)

type Conn struct {
	UserId   int64           `json:"user_id"`
	ConnType int             `json:"conn_type"`
	WsConn   *websocket.Conn `json:"ws_conn"`
}

func (c *Conn) Close() {
	//用户下线 todo

	//取消订阅 todo

	if c.ConnType == enum.WsConn {
		c.WsConn.Close()
	} else if c.ConnType == enum.TcpConn {
		c.WsConn.Close()
	}
}

func (c *Conn) HandleMessage(bytes []byte) {
	msg := &message.Message{}
	err := json.Unmarshal(bytes, &msg)
	if err != nil {
		zap.S().Errorf("解析用户:%v消息失败:%v", c.UserId, err.Error())
		return
	}

	switch msg.MsgType {
	case enum.AuthMsg:
		c.Auth(msg)
	case enum.SyncMsg:
		c.Sync(msg)
	}
}

func (c *Conn) Auth(msg *message.Message) {
	var authMsg message.AuthMessage
	err := json.Unmarshal(msg.Data, &authMsg)
	if err != nil {
		zap.S().Errorf("鉴权解析用户消息失败:%v", err.Error())
		return
	}

	//发送鉴权请求
	authBo := &busi_client.AuthBo{
		UserId: authMsg.UserId,
		Token:  authMsg.Token,
	}
	authVo := busi_client.Auth(authBo)
	if authVo == nil {
		//鉴权失败，提示用户 todo
		c.Send(msg.MsgType, nil, enum.AuthErrCode)
		return
	}

	//维护长连接
	c.UserId = authVo.UserId
	SetConn(c.UserId, c)
}

func (c *Conn) Sync(msg *message.Message) {

}

func (c *Conn) Send(msgType int, data []byte, errCode int) {
	msg := enum.GetErrorMsg(errCode)

	replyMessage := &message.ReplyMessage{
		MsgType: msgType,
		Code:    errCode,
		Msg:     msg,
		Data:    data,
	}

	if err := c.Write(replyMessage); err != nil {
		zap.S().Errorf("用户%v发送消息失败：%v", c.UserId, err.Error())
		c.Close()
		return
	}
}

func (c *Conn) Write(msg *message.ReplyMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if c.ConnType == enum.WsConn {
		c.WriteToWs(data)
	} else if c.ConnType == enum.TcpConn {

	}
	return nil
}

func (c *Conn) WriteToWs(data []byte) error {
	err := c.WsConn.SetWriteDeadline(time.Now().Add(8 * time.Second))
	if err != nil {
		return err
	}

	return c.WsConn.WriteMessage(websocket.BinaryMessage, data)
}
