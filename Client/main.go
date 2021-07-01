package main

import (
	"client/user"
	leafLog "github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"log"
	"math"
	"time"
)

func init(){
	gLogger, _ := leafLog.New("debug", "./client_log/", log.LstdFlags)
	leafLog.Export(gLogger)
}

func main() {
	tcpClient := &network.TCPClient{
		Addr:            "127.0.0.1:3563",
		ConnNum:         1,
		ConnectInterval: 3 * time.Second,
		PendingWriteNum: 1000,
		LenMsgLen:       2,
		MaxMsgLen:       math.MaxUint32,
		NewAgent:        user.NewAgent,
	}
	
	tcpClient.Start()
	StartGame()
}
