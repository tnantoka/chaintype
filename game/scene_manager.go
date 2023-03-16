package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SceneManager struct {
	screenWidth  float64
	screenHeight float64
	currentScene Scene
}

func NewSceneManager(screenWidth float64, screenHeight float64) *SceneManager {
	return &SceneManager{screenWidth, screenHeight, &TitleScene{Screen{screenWidth, screenHeight}}}
}

func (sm *SceneManager) Update() {
	sm.currentScene.Update()

	switch s := sm.currentScene.(type) {
	case *TitleScene:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			sm.currentScene = &PlayScene{Screen{screenWidth, screenHeight}, NewSpriteManager(sm.screenWidth, sm.screenHeight), "", false}
		}
	case *PlayScene:
		if s.isGameOver {
			sm.currentScene = &GameOverScene{Screen{screenWidth, screenHeight}}
		}
	case *GameOverScene:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			sm.currentScene = &TitleScene{Screen{screenWidth, screenHeight}}
		}
	}
}

func (sm *SceneManager) Draw(screen *ebiten.Image) {
	sm.currentScene.Draw(screen)
}
