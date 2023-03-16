package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
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

type FontSize float64

const (
	baseFontSize  FontSize = 16
	largeFontSize FontSize = 32
)

func (fs FontSize) Raw() float64 {
	return float64(fs)
}

func (fs FontSize) Font() font.Face {
	switch fs {
	case baseFontSize:
		return baseFont
	case largeFontSize:
		return largeFont
	}
	return nil
}

var (
	baseFont  font.Face
	largeFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72

	baseFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    baseFontSize.Raw(),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	largeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    largeFontSize.Raw(),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Anchor int

const (
	AnchorLeftTop Anchor = iota
	AnchorCenterMiddle
)

type TextSprite struct {
	position Position
	fontSize FontSize
	color    color.Color
	text     string
	anchor   Anchor
}

func (s *TextSprite) Update() {
}

func (s *TextSprite) Draw(screen *ebiten.Image) {
	x := 0
	switch s.anchor {
	case AnchorLeftTop:
		x = int(s.position.x)
	case AnchorCenterMiddle:
		x = int(s.position.x - float64(font.MeasureString(s.fontSize.Font(), s.text).Round())*0.5)
	}
	y := 0
	switch s.anchor {
	case AnchorLeftTop:
		y = int(s.position.y + s.fontSize.Raw())
	case AnchorCenterMiddle:
		y = int(s.position.y + s.fontSize.Raw()*0.5)
	}
	text.Draw(screen, s.text, s.fontSize.Font(), x, y, s.color)
}
