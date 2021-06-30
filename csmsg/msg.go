package csmsg

import (
	"github.com/name5566/leaf/network/json"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Hello{})
	Processor.Register(&C2S_EnterSquare{})
	Processor.Register(&S2C_EnterSquare{})
	Processor.Register(&C2S_Login{})
	Processor.Register(&S2C_Login{})
}

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
type Hello struct {
	Name  string
	Count uint32
}

type C2S_EnterSquare struct {
}

type S2C_EnterSquare struct {
	TotalCount uint32
}
