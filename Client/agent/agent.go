package agent

import (
	"csmsg"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
)

type Agent struct {
	*network.TCPConn
	network.Processor
}

func NewAgent(conn *network.TCPConn) network.Agent {
	a := new(Agent)
	a.TCPConn = conn
	a.Processor = csmsg.Processor
	return a
}

func (a *Agent) Run() {
	for {
		data, err := a.ReadMsg()
		if err != nil {
			log.Debug("read message: %v", err)
			break
		}
		
		msg, err := a.Unmarshal(data)
		if err != nil {
			log.Debug("unmarshal message error: %v", err)
			break
		}
		a.Processor.Route(msg, a)
	}
}
func (a *Agent) WriteMsgWithMarshal(msg interface{}) {
	byteMsg, _ := a.Marshal(msg)
	a.WriteMsg(byteMsg...)
}

func (a *Agent) OnClose() {}
