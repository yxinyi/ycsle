package internal

import (
	csmsg "csmsg"
	"github.com/name5566/leaf/gate"
	"reflect"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	skeleton.RegisterChanRPC(reflect.TypeOf(&csmsg.C2S_Login{}), UserLogin)
}

func UserLogin(args []interface{}) {
	// 消息的发送者
	a := args[1].(gate.Agent)
	a.WriteMsg(&csmsg.S2C_Login{})
}
