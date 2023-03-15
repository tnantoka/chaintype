package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	bgColor   = color.RGBA{200, 200, 200, 255}
	textColor = color.Black
	baseFont  font.Face
)

var spriteManager = NewSpriteManager(screenWidth, screenHeight)

func init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const fontSize = 16
	const dpi = 72
	baseFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	spriteManager.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	text.Draw(screen, "Hello, World!", baseFont, 20, 40, textColor)

	spriteManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
