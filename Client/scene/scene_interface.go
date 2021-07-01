package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/name5566/leaf/log"
)

type Scene interface {
	OnEnter()
	OnQuit()
	Update()
	Draw(dst *ebiten.Image)
}

type SceneManager struct {
	currentScene Scene
	sceneList    map[string]Scene
}

var SceneMgr = NewSceneManager()

func NewSceneManager() *SceneManager {
	sceneMgr := &SceneManager{
		sceneList: make(map[string]Scene),
	}
	return sceneMgr
}

func (mgr *SceneManager) regScene(sceneName string, s Scene) {
	_, exists := SceneMgr.sceneList[sceneName]
	if exists {
		panic("scene [" + sceneName + "] exists before")
	}
	SceneMgr.sceneList[sceneName] = s
}
func (mgr *SceneManager) changeScene(sceneName string) {
	scene, exists := SceneMgr.sceneList[sceneName]
	if !exists {
		log.Fatal("not exists [%v]")
		return
	}
	if mgr.currentScene != nil{
		mgr.currentScene.OnQuit()
	}
	mgr.currentScene = scene
	mgr.currentScene.OnEnter()
}

func (mgr *SceneManager) Update() {
	mgr.currentScene.Update()
}

func (mgr *SceneManager) Draw(dst *ebiten.Image) {
	mgr.currentScene.Draw(dst)
}

func RegScene(sceneName string, s Scene) {
	SceneMgr.regScene(sceneName, s)
}

func ChangeScene(sceneName string) {
	SceneMgr.changeScene(sceneName)
}

func Update() {
	SceneMgr.Update()
}

func Draw(dst *ebiten.Image) {
	SceneMgr.Draw(dst)
}