package message

type Message struct {
	MsgType int    `json:"msg_type"`
	Data    []byte `json:"data"`
}

type ReplyMessage struct {
	MsgType int    `json:"msg_type"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    []byte `json:"data"`
}

type AuthMessage struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type SyncMessage struct {
	Seq int64 `json:"seq"`
}
