package communication

import ()

type MsgInfo struct {
}

// 请求信息
type ReqMessage struct {
	Comm string      `json:"comm"`
	Data interface{} `json:"data"`
}

// 返回信息
type ResMessage struct {
	BackCode int         `json:"backcode"`
	BackMsg  string      `json:"backmsg"`
	BackData interface{} `json:"backdata"`
}

var Message = map[int]string{
	10001: "操作正常",
	10002: "连接超时",
	10003: "未知错误",
}

var DefCodes = map[string]int{
	"success": 10001,
	"outtime": 10002,
	"unkerr":  10003,
}

// 根据状态码生成返回消息
func (obj *MsgInfo) GenResCodeMsg(code int, data interface{}) *ResMessage {
	if code == 0 {
		code = DefCodes["success"]
	}
	res := &ResMessage{
		BackCode: code,
		BackMsg:  Message[code],
		BackData: data,
	}
	return res
}

// 根据key生成返回消息
func (obj *MsgInfo) GenResKeyMsg(key string) *ResMessage {
	res := &ResMessage{
		BackCode: DefCodes[key],
		BackMsg:  Message[DefCodes[key]],
		BackData: nil,
	}
	return res
}
