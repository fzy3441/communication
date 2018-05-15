# communication

    自定义通讯包

# 使用方法

### 引用包

```go
import (
	"fmt"
	c "communication"
)
```

### 实现 IMaster 主接口(必须)
```go
type Master struct{} // 实现IMaster接口

// 初始化当前对象到使用者连接成功之前操作
func (obj *Master) InitObj(event *c.Event) {
	obj.Event = event
	obj.ConnEvent = &ConnS{} // IConnEvent 实现后，添加该行（否则可省略）
}

// 处理客户端返回消息
func (obj *Master) ProcessReq(data *ReqMessage) {

} // 处理客户端数据信息
func (obj *Master) ProcessDisConn() {

} // 客户端断开连接后访问该方法
```

### IConnEvent 接口，当前连接基本操作响应（可选）

```go
type ConnS struct{} // 实现IConnEvent 接口

func (obj *ConnS) Reading(buf []byte) {} // Socket获取信息时执行
func (obj *ConnS) Writing(buf []byte) {} // Socket发送信息时执行
func (obj *ConnS) Connection() {
	fmt.Println("hello world") // 当用户连接成功时输出'hello world'
}                              // Socket连接成功时执行
func (obj *ConnS) Disconnect() {} // Socket断开连接时执行
```

### 开始连接
```go
co := &Master{}
c.NewDailTcp(co, ":9001")
msg := c.MsgInfo{} //  包自带消息类，可实现消息专用接口 IMessage,修改消息类型
co.EventResponse(msg.GenResKeyMsg("success")) // 当连接成功后发送消息给客户端
```

### 开始监听端口
```go
co := &Master{}
lis := c.Listen("tcp", ":9001")
lis.Accept(co) // 等待连接成功 co 可操作返回数据
```
