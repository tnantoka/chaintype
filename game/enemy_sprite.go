package game

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
)

const (
	enemyImageSize = 50
)

var (
	enemyBgColor    = color.RGBA{100, 100, 100, 255}
	enemyMutedColor = color.RGBA{150, 150, 150, 255}
)

type EnemySprite struct {
	position Position
	fontSize FontSize
	text     string
	cursor   int
	isDead   bool
	img      *ebiten.Image
	speed    float64
}

func NewEnemySprite(screenWidth float64, screenHeight float64) *EnemySprite {
	x := screenWidth
	y := rand.Float64()*(screenHeight-baseFontSize.Raw()-enemyImageSize) + enemyImageSize

	speed := rand.Float64()*1 + 0.5

	img, _, err := ebitenutil.NewImageFromFileSystem(imagesFS, fmt.Sprintf("images/enemy_%d.png", rand.Intn(4)))
	if err != nil {
		log.Fatal(err)
	}
	return &EnemySprite{position: Position{x, y}, fontSize: baseFontSize, text: words[rand.Intn(len(words))], img: img, speed: speed}
}

func (s *EnemySprite) Update() {
	s.position.x -= s.speed
}

func (s *EnemySprite) Draw(screen *ebiten.Image) {
	rectSprite := RectSprite{s.position, Size{s.Frame().w, s.Frame().h}, enemyBgColor}
	rectSprite.Draw(screen)

	textSprite := TextSprite{s.position, s.fontSize, color.White, s.text, AnchorLeftTop}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.position.x+textSprite.Frame().w*0.5-enemyImageSize*0.5, s.position.y-enemyImageSize)
	screen.DrawImage(s.img, op)

	textSprite.Draw(screen)

	textSprite = TextSprite{s.position, s.fontSize, enemyMutedColor, s.text[:s.cursor], AnchorLeftTop}
	textSprite.Draw(screen)
}

func (s *EnemySprite) Frame() Frame {
	return Frame{s.position.x, s.position.y, float64(font.MeasureString(s.fontSize.Font(), s.text).Round()), s.fontSize.Raw()}
}

func (s *EnemySprite) Input(str string) bool {
	if strings.ToUpper(string(s.text[s.cursor])) == str {
		s.cursor++

		if s.cursor >= len(s.text) {
			s.isDead = true
		}

		return true
	}

	return false
}

func (s *EnemySprite) KnockBack() {
	s.isDead = true
}
