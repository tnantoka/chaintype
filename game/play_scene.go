package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type PlayScene struct {
	screen        Screen
	spriteManager *SpriteManager
	input         string
	isGameOver    bool
}

func (s *PlayScene) Update() {
	s.spriteManager.Update()

	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if inpututil.IsKeyJustPressed(k) {
			str := k.String()
			if len(str) == 1 {
				s.input += str
			}
		}
	}

	for _, sprite := range s.spriteManager.sprites {
		if sprite.Frame().x < 1 {
			s.isGameOver = true
			break
		}
	}
}

func (s *PlayScene) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	text.Draw(screen, "Hello, World!", baseFont, 20, 40, textColor)

	ebitenutil.DebugPrint(screen, s.input)

	s.spriteManager.Draw(screen)
}
