package scene

import (
	"client/ui"
	"image"
)

type LoginScene struct {
	*SceneBase
}

func NewLoginScene() *LoginScene {
	s := &LoginScene{}
	s.SceneBase = NewSceneBase()
	loginButton := ui.NewButton()
	loginButton.Rect = image.Rect(16, 16, 144, 48)
	loginButton.Text = "Button 1"
	s.AddUI("loginButton", loginButton)
	return s
}

func init() {
	RegScene("login", NewLoginScene())
}
