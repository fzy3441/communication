package communication

import (
	so "socket_asp"
)

type DailTcp struct {
	*so.NetIO
}

func NewDailTcp(inter IMaster, address string) error {
	event := NewEvent(inter, 30) // 创建事件对象
	achieve := &Achieve{event}   // 创建NetIO主接口实现对象
	so.DailTcp(achieve, address) // 主动连接
	event.Init()                 // 初始化事件对象

	go event.WaitRead()
	go event.Events()

	return nil
}
