package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	leafLog "github.com/name5566/leaf/log"
	"log"
	"server/conf"
	"server/game"
	"server/gate"
)

func init() {
	gLogger, _ := leafLog.New("debug", "./server_log/", log.LstdFlags)
	leafLog.Export(gLogger)
}
func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath
	
	leaf.Run(
		game.Module,
		gate.Module,
	)
}
