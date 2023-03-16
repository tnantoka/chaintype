package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	bgColor   = color.RGBA{50, 50, 50, 255}
	textColor = color.White
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Screen struct {
	w float64
	h float64
}

type TitleScene struct {
	screen Screen
}

func (s *TitleScene) Update() {
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	grid(screen, s.screen)

	title := TextSprite{
		Position{s.screen.w * 0.5, s.screen.h * 0.4},
		largeFontSize,
		textColor,
		"Chaintype",
	}
	title.Draw(screen)

	start := TextSprite{
		Position{s.screen.w * 0.5, s.screen.h * 0.6},
		baseFontSize,
		textColor,
		"Press space to start",
	}
	start.Draw(screen)
}

type GameOverScene struct {
	screen Screen
}

func (s *GameOverScene) Update() {
}

func (s *GameOverScene) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	grid(screen, s.screen)

	gameOver := TextSprite{
		Position{s.screen.w * 0.5, s.screen.h * 0.5},
		largeFontSize,
		textColor,
		"Game over",
	}
	gameOver.Draw(screen)
}

func grid(screen *ebiten.Image, s Screen) {
	ebitenutil.DrawLine(screen, 0, s.h*0.5, s.w, s.h*0.5, textColor)
	ebitenutil.DrawLine(screen, s.w*0.5, 0, s.w*0.5, s.h, textColor)
}
