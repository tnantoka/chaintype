package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	chainImageSize = 15
	chainSpeed     = 8
)

type ChainSprite struct {
	position Position
	isDead   bool
	img      *ebiten.Image
}

func NewChainSprite(x float64, y float64) *ChainSprite {
	img, _, err := ebitenutil.NewImageFromFileSystem(imagesFS, "images/chain.png")
	if err != nil {
		log.Fatal(err)
	}

	return &ChainSprite{Position{x, y}, false, img}
}

func (s *ChainSprite) Update() {
	s.position.y += chainSpeed
}

func (s *ChainSprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.position.x, s.position.y)
	screen.DrawImage(s.img, op)
}

func (s *ChainSprite) Frame() Frame {
	return Frame{s.position.x, s.position.y, chainImageSize, chainImageSize}
}
