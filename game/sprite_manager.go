package game

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteManager struct {
	screenWidth  float64
	screenHeight float64
	sprites      []Sprite
}

func NewSpriteManager(screenWidth float64, screenHeight float64) *SpriteManager {
	return &SpriteManager{screenWidth, screenHeight, []Sprite{}}
}

func (sm *SpriteManager) Update() {
	if len(sm.sprites) == 0 {
		rect := &RectSprite{size: Size{w: 20, h: 30}, color: color.RGBA{255, 0, 0, 255}}
		rect.position.x = sm.maxX(rect)
		rect.position.y = rand.Float64() * sm.maxY(rect)
		sm.sprites = append(sm.sprites, rect)

		circle := &CircleSprite{r: 20, color: color.RGBA{0, 255, 0, 255}}
		circle.position.x = sm.maxX(circle)
		circle.position.y = rand.Float64() * sm.maxY(circle)
		sm.sprites = append(sm.sprites, circle)
	}

	for _, sprite := range sm.sprites {
		switch s := sprite.(type) {
		case *RectSprite:
			s.position.x -= 1
			//if s.position.x < sm.minX(s) {
			//	s.position.x = sm.maxX(s)
			//}
		case *CircleSprite:
			s.position.x -= 1
			//if s.position.x < sm.minX(s) {
			//	s.position.x = sm.maxX(s)
			//}
		}
	}
}

func (sm *SpriteManager) Draw(screen *ebiten.Image) {
	for _, s := range sm.sprites {
		s.Draw(screen)
	}
}

func (sm *SpriteManager) minX(s Sprite) float64 {
	return 0
}

func (sm *SpriteManager) maxX(s Sprite) float64 {
	return sm.screenWidth - s.Frame().w
}

func (sm *SpriteManager) minY(s Sprite) float64 {
	return 0
}

func (sm *SpriteManager) maxY(s Sprite) float64 {
	return sm.screenHeight - s.Frame().h
}
