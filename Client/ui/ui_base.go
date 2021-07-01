package ui

import "github.com/hajimehoshi/ebiten/v2"

type UIInterface interface {
	Draw(dst *ebiten.Image)
	Update()
}

type UIBase struct {
}

func (ui *UIBase)Draw(dst *ebiten.Image){}
func (ui *UIBase)Update(){}