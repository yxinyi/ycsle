package gate

import (
	"csmsg"
	"server/game"
)

func init() {
	// 这里指定消息 Hello 路由到 game 模块
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	csmsg.Processor.SetRouter(&csmsg.Hello{}, game.ChanRPC)
	csmsg.Processor.SetRouter(&csmsg.C2S_EnterSquare{}, game.ChanRPC)
}