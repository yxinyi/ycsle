package user

import (
	"csmsg"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"queue"
)

type MainUser struct {
	*network.TCPConn
	network.Processor
	User
	msgQueue *queue.SyncQueue
}

type User struct {
	LoginState uint32
}


var Obj *MainUser

func (u *MainUser) Run() {
	for {
		data, err := u.ReadMsg()
		if err != nil {
			log.Debug("read message: %v", err)
			break
		}

		msg, err := u.Unmarshal(data)
		if err != nil {
			log.Debug("unmarshal message error: %v", err)
			break
		}
		
		u.msgQueue.Add(msg)
	}
}


func (u *MainUser)MsgRoute(){
	for {
		if u.msgQueue.Len() == 0 {
			break
		}
		msg :=  u.msgQueue.Pop()
		u.Processor.Route(msg, u)
	}
}

func (u *MainUser) WriteMsgWithMarshal(msg interface{}) {
	byteMsg, _ := u.Marshal(msg)
	u.WriteMsg(byteMsg...)
}

func (u *MainUser) OnClose() {}

func NewAgent(conn *network.TCPConn) network.Agent {
	u := new(MainUser)
	u.TCPConn = conn
	u.Processor = csmsg.Processor
	u.LoginState = LOGIN_NONE
	u.msgQueue = queue.NewSyncQueue()
	LoginFunc(u)
	return u
}
