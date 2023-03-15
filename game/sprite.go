package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite interface {
	Update()
	Draw(screen *ebiten.Image)
	Frame() Frame
}

type Frame struct {
	x float64
	y float64
	w float64
	h float64
}

type Position struct {
	x float64
	y float64
}

type Size struct {
	w float64
	h float64
}

type RectSprite struct {
	position Position
	size     Size
	color    color.Color
}

func (s *RectSprite) Update() {
}

func (s *RectSprite) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, s.position.x, s.position.y, s.size.w, s.size.h, s.color)
}

func (s *RectSprite) Frame() Frame {
	return Frame{s.position.x, s.position.y, s.size.w, s.size.h}
}

type CircleSprite struct {
	position Position
	r        float64
	color    color.Color
}

func (s *CircleSprite) Update() {
}

func (s *CircleSprite) Draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, s.position.x+s.r, s.position.y+s.r, s.r, s.color)
}

func (s *CircleSprite) Frame() Frame {
	return Frame{s.position.x, s.position.y, s.r * 2, s.r * 2}
}
