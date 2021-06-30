package main

import (
	"client/user"
	"github.com/name5566/leaf/network"
	"math"
	"time"
)

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
