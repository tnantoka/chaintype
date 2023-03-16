package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	maxEnemies = 10
)

type PlayScene struct {
	screen        Screen
	input         string
	isGameOver    bool
	enemies       []*EnemySprite
	frameCount    int
	enemyInterval int
}

func NewPlayScene(s Screen) *PlayScene {
	return &PlayScene{s, "", false, []*EnemySprite{}, 0, 100}
}

func (s *PlayScene) Update() {
	s.frameCount++

	if len(s.enemies) <= maxEnemies && s.frameCount%s.enemyInterval == 0 {
		enemy := NewEnemySprite()
		enemy.position.x = s.screen.w
		enemy.position.y = rand.Float64() * (s.screen.h - enemy.Frame().h)
		s.enemies = append(s.enemies, enemy)
	}

	s.input = ""
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if inpututil.IsKeyJustPressed(k) {
			str := k.String()
			if len(str) == 1 {
				s.input += str

				for _, enemy := range s.enemies {
					enemy.Input(str)
				}
			}
		}

		s.enemies = func() []*EnemySprite {
			var result []*EnemySprite
			for _, enemy := range s.enemies {
				if !enemy.isDead {
					result = append(result, enemy)
				}
			}
			return result
		}()
	}

	for _, enemy := range s.enemies {
		enemy.Update()
		if enemy.Frame().x < 1 {
			s.isGameOver = true
		}
	}
}

func (s *PlayScene) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	text.Draw(screen, "Hello, World!", baseFont, 20, 40, textColor)

	ebitenutil.DebugPrint(screen, s.input)

	for _, enemy := range s.enemies {
		enemy.Draw(screen)
	}
}
