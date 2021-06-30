package user

import (
	"csmsg"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
)

type User struct {
	*network.TCPConn
	network.Processor
	LoginState uint32
}

var U *User

func (a *User) Run() {
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
func (a *User) WriteMsgWithMarshal(msg interface{}) {
	byteMsg, _ := a.Marshal(msg)
	a.WriteMsg(byteMsg...)
}

func (a *User) OnClose() {}

func NewAgent(conn *network.TCPConn) network.Agent {
	u := new(User)
	u.TCPConn = conn
	u.Processor = csmsg.Processor
	u.LoginState = LOGIN_NONE

	LoginFunc(u)
	return u
}
