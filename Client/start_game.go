package main

import (
	"client/2048"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/name5566/leaf/log"
)



func StartGame(){
	game, err := twenty48.NewGame()
	if err != nil {
		log.Fatal(err.Error())
	}
	
	game.EventRegister()
	ebiten.SetWindowSize(twenty48.ScreenWidth, twenty48.ScreenHeight)
	ebiten.SetWindowTitle("2048 (Ebiten Demo)")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err.Error())
	}
}