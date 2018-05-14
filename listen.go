package communication

import (
	"fmt"
	"os"
	so "socket_asp"
)

type Listener struct {
	*so.Listener
}

func Listen(network, address string) *Listener {
	listener, err := so.Listen(network, address)
	if err != nil {
		// log.Errorf("打开服务器%s端口: %s 失败", network, address)
		fmt.Printf("打开服务器%s端口: %s 失败", network, address)
		os.Exit(0)
	}

	return &Listener{listener}
}

// 等待客户连接
func (obj *Listener) Accept(inter IMaster) {
	event := NewEvent(inter, 30) // 创建事件对象
	// inter.InitObj(obj)           // 使用者初始化
	achieve := &Achieve{event}   // 创建NetIO主接口实现对象
	obj.Listener.Accept(achieve) // 等待连接
	event.Init()                 // 初始化事件对象

	go event.WaitRead()
	go event.Events()
}
