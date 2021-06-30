package internal

import (
	csmsg "csmsg"
	"reflect"
)

var (
	SquareCount uint32
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	skeleton.RegisterChanRPC(reflect.TypeOf(&csmsg.C2S_EnterSquare{}), handle_EnterSquare)
	
	//skeleton.RegisterChanRPC("CloseAgent", rpc_EnterSquare)
}

func handle_EnterSquare(args []interface{}) {
	/*m := args[0].(*csmsg.C2S_EnterSquare)
	a := args[1].(gate.Agent)*/
	
	SquareCount++
	for userIt := range UserSet {
		userIt.WriteMsg(&csmsg.S2C_EnterSquare{
			TotalCount: SquareCount,
		})
	}
	
}

func rpc_EnterSquare(args []interface{}) {
	//a := args[0].(gate.Agent)
	SquareCount--
	for userIt := range UserSet {
		userIt.WriteMsg(&csmsg.S2C_EnterSquare{
			TotalCount: SquareCount,
		})
	}
}
