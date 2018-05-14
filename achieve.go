package communication

import (
	"encoding/json"
	// "fmt"
	so "socket_asp"
)

type Achieve struct {
	*Event
}

// 初始化连接执行
func (obj *Achieve) InitConn(param *so.NetIO) {
	obj.NetIO = param // 把连接对象赋值给当前对象
	if obj.ConnEvent != nil {
		obj.ConnEvent.Connection()
	}
}

// 断开连接时执行
func (obj *Achieve) Disconn() {
	if obj.ConnEvent != nil {
		obj.ConnEvent.Disconnect()
	}
}

// 读取到消息时执行
func (obj *Achieve) ProcessRead(buf []byte) bool {
	if obj.ConnEvent != nil {
		obj.ConnEvent.Reading(buf)
	}
	read := &ReqMessage{}
	json.Unmarshal(buf, read)

	flag := true
	func() {
		defer func() {
			if err := recover(); err != nil {
				flag = false
				obj.Close() // 断开连接
			}
		}()
		obj.Request <- read
	}()
	return flag
}

// 发送消息时执行
func (obj *Achieve) ProcessWrite(buf []byte) ([]byte, bool) {
	if obj.ConnEvent != nil {
		obj.ConnEvent.Writing(buf)
	}
	return buf, true
}
