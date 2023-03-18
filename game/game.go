package game

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

//go:embed images/*
var imagesFS embed.FS

type Game struct {
	currentScene Scene
}

func NewGame() *Game {
	return &Game{&TitleScene{Screen{screenWidth, screenHeight}}}
}

func (g *Game) Update() error {
	g.currentScene.Update()

	switch s := g.currentScene.(type) {
	case *TitleScene:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.currentScene = NewPlayScene(Screen{screenWidth, screenHeight})
		}
	case *PlayScene:
		if s.isGameOver {
			g.currentScene = &GameOverScene{Screen{screenWidth, screenHeight}, s.score}
		}
	case *GameOverScene:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.currentScene = &TitleScene{Screen{screenWidth, screenHeight}}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentScene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
