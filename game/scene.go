package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
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

	title := TextSprite{
		Position{s.screen.w * 0.5, s.screen.h * 0.4},
		largeFontSize,
		textColor,
		"Chaintype",
		AnchorCenterMiddle,
	}
	title.Draw(screen)

	start := TextSprite{
		Position{s.screen.w * 0.5, s.screen.h * 0.6},
		baseFontSize,
		textColor,
		"Press space to start",
		AnchorCenterMiddle,
	}
	start.Draw(screen)
}

type GameOverScene struct {
	screen Screen
	score  int
}

func (s *GameOverScene) Update() {
}

func (s *GameOverScene) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	gameOver := TextSprite{
		Position{s.screen.w * 0.5, s.screen.h * 0.4},
		largeFontSize,
		textColor,
		"Game over",
		AnchorCenterMiddle,
	}
	gameOver.Draw(screen)

	score := TextSprite{
		Position{s.screen.w * 0.5, s.screen.h * 0.6},
		baseFontSize,
		textColor,
		fmt.Sprintf("Score: %d", s.score),
		AnchorCenterMiddle,
	}
	score.Draw(screen)
}
