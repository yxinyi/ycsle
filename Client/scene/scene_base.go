package scene

import (
	"client/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneBase struct {
	uiPool map[string]ui.UIInterface
}

func NewSceneBase()*SceneBase{
	sb := &SceneBase{}
	sb.uiPool = make(map[string]ui.UIInterface)
	return sb
}

func (s *SceneBase) AddUI(name string, ui ui.UIInterface) {
	s.uiPool[name] = ui
}
func (s *SceneBase) Update() {
	for _, it := range s.uiPool {
		it.Update()
	}
}

func (s *SceneBase) Draw(dst *ebiten.Image) {
	for _, it := range s.uiPool {
		it.Draw(dst)
	}
}

func (s *SceneBase) OnEnter() {}
func (s *SceneBase) OnQuit()  {}
