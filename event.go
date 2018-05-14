package communication

import (
	"fmt"
	"os"
	so "socket_asp"
	"time"
)

type Event struct {
	*so.NetIO
	Inter      IMaster          // 包外部接口
	MaxOutTime int              // outtime time
	Message    IMessage         // 消息接口
	ConnEvent  IConnEvent       // Socket 接口
	Disconnect chan int         // 断开连接
	Request    chan *ReqMessage // 客户端请求
	Response   chan *ResMessage // 服务端返回
}

// 创建事件对象
func NewEvent(Inter IMaster, MaxOutTime int) *Event {
	event := &Event{
		Inter:      Inter,
		MaxOutTime: MaxOutTime,
		Request:    make(chan *ReqMessage),
		Response:   make(chan *ResMessage),
		Disconnect: make(chan int),
	}

	if event.Inter == nil {
		fmt.Println("外部接口未使用")
		os.Exit(0)
	}
	Inter.InitObj(event)

	return event
}

// 初始化
func (obj *Event) Init() {
	if obj.Message == nil { // 是否存在外部消息类
		obj.Message = &MsgInfo{} // 如果使用者并没有创建message类，则继续使用系统自带的消息类
	}
}

// 等待并读取客户端信息
func (obj *Event) WaitRead() {
	flag, err := obj.NetIO.WaitRead() // 执行读取
	// flag, err := obj.ReadStruct(obj.Request)
	// 判断连接是否己断开
	if !flag && !obj.DisConn {
		// 断开连接
		obj.EventClose()
	}
	if err != nil {
		// 信息处理错误,导致断开连接
		obj.EventClose()
	}
}

// 用户事件
func (obj *Event) Events() {
	defer obj.EventClose()
	for {
		select {
		case data := <-obj.Request:
			go obj.EventRequest(data)
		case data := <-obj.Response:
			go obj.EventResponse(data)
		case <-time.After(time.Duration(obj.MaxOutTime+30) * time.Second):
			go obj.EventOutTime()
		case <-obj.Disconnect:
			return
		}
	}
}

// 处理请求
func (obj *Event) EventRequest(data *ReqMessage) {
	defer func() {
		if err := recover(); err != nil && err != "return" {
			obj.Response <- obj.Message.GenResKeyMsg("unkerr")
		}
	}()
	obj.Inter.ProcessReq(data)
}

// 发送消息
func (obj *Event) EventResponse(data *ResMessage) {
	obj.WriteJson(data)
}

// 超时断开
func (obj *Event) EventOutTime() {
	obj.WriteJson(obj.Message.GenResKeyMsg("outtime"))
	obj.Display()
}

// 关闭连接
func (obj *Event) Display() {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	obj.Disconnect <- 1
}

// 关闭连接
func (obj *Event) EventClose() {
	obj.Close()
	close(obj.Request)
	close(obj.Response)
	close(obj.Disconnect)
	obj.DisConn = true
}
