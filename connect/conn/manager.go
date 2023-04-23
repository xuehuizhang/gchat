package conn

import "sync"

var ConnsMap = sync.Map{}

func SetConn(userId int64, conn *Conn) {
	ConnsMap.Store(userId, conn)
}

func GetConn(userId int64) *Conn {
	value, ok := ConnsMap.Load(userId)
	if ok {
		return value.(*Conn)
	}
	return nil
}

func RemoveConn(userid int64) {
	conn := GetConn(userid)
	if conn != nil {
		conn.Close()
		ConnsMap.Delete(userid)
	}
}

func RemoveAll() {
	ConnsMap.Range(func(key, value any) bool {
		conn := value.(*Conn)
		conn.Close()
		ConnsMap.Delete(key)
		return true
	})
}
