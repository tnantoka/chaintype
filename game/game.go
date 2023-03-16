package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	sceneManager *SceneManager
}

func NewGame() *Game {
	return &Game{
		sceneManager: NewSceneManager(screenWidth, screenHeight),
	}
}

func (g *Game) Update() error {
	g.sceneManager.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
