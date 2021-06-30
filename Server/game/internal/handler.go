package internal

import (
	csmsg "csmsg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	skeleton.RegisterChanRPC(reflect.TypeOf(&csmsg.Hello{}), handleHello)
}

func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*csmsg.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)
	
	// 输出收到的消息的内容
	log.Debug("count %v", m.Count)
	
	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&csmsg.Hello{
		Count: m.Count + 1,
	})
}
