package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	maxEnemies = 4
)

var (
	wallColor = color.RGBA{200, 200, 200, 255}
)

type PlayScene struct {
	screen        Screen
	input         string
	isGameOver    bool
	enemies       []*EnemySprite
	chains        []*ChainSprite
	frameCount    int
	enemyInterval int
	wallCount     int
	score         int
}

func NewPlayScene(s Screen) *PlayScene {
	return &PlayScene{s, "", false, []*EnemySprite{}, []*ChainSprite{}, 0, 40, 3, 0}
}

func (s *PlayScene) Update() {
	s.frameCount++

	if len(s.enemies) <= maxEnemies && s.frameCount%s.enemyInterval == 0 {
		enemy := NewEnemySprite(s.screen.w, s.screen.h)
		s.enemies = append(s.enemies, enemy)
	}

	s.input = ""
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if inpututil.IsKeyJustPressed(k) {
			str := k.String()
			if len(str) == 1 {
				s.input += str

				for _, enemy := range s.enemies {
					if enemy.Input(str) {
						break
					}
				}
			}
		}

		s.enemies = func() []*EnemySprite {
			var result []*EnemySprite
			for _, enemy := range s.enemies {
				if enemy.isDead {
					s.score++

					chain := NewChainSprite(enemy.position.x+enemy.Frame().w*0.5, enemy.position.y)
					s.chains = append(s.chains, chain)
				} else {
					result = append(result, enemy)
				}
			}
			return result
		}()
	}

	for _, enemy := range s.enemies {
		enemy.Update()
		if enemy.Frame().x < (float64(s.wallCount)-1)*20+10 {
			s.wallCount--
			enemy.KnockBack()

			if s.wallCount < 0 {
				s.isGameOver = true
			}
		}
	}

	s.chains = func() []*ChainSprite {
		var result []*ChainSprite
		for _, chain := range s.chains {
			if !chain.isDead {
				result = append(result, chain)
			}
		}
		return result
	}()

	for _, chain := range s.chains {
		chain.Update()

		for _, enemy := range s.enemies {
			if Intersect(chain, enemy) {
				enemy.isDead = true
				chain.isDead = true
			}
		}
	}
}

func (s *PlayScene) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	for _, enemy := range s.enemies {
		enemy.Draw(screen)
	}

	for _, chain := range s.chains {
		chain.Draw(screen)
	}

	for i := 0; i < s.wallCount; i++ {
		wall := RectSprite{Position{float64(i) * 20, 0}, Size{10, s.screen.h}, wallColor}
		wall.Draw(screen)
	}

	score := TextSprite{Position{s.screen.w - 5, 5}, baseFontSize, textColor, fmt.Sprintf("Score: %d", s.score), AnchorRightTop}
	score.Draw(screen)
}
