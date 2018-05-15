package communication

// 主接口
type IMaster interface {
	InitObj(event *Event) // 初始化当前对象到使用者
	// Message() *IMessage          // 设计并创建消息类型
	ProcessReq(data *ReqMessage) // 处理客户端数据信息
	ProcessDisConn()             // 客户端断开连接后访问该方法
}

// socket 接口
type IConnEvent interface {
	Reading(buf []byte) // Socket获取信息时执行
	Writing(buf []byte) // Socket发送信息时执行
	Connection()        // Socket连接成功时执行
	Disconnect()        // Socket断开连接时执行
}

// 消息
type IMessage interface {
	GenResCodeMsg(code int, data interface{}) *ResMessage // 根据code生成返回消息
	GenResKeyMsg(key string) *ResMessage                  // 根据key 生成返回消息
}
