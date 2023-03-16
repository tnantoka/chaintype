package game

import (
	"image/color"
	"math/rand"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

var words = []string{
	"break", "default", "func", "interface", "select",
	"case", "defer", "go", "map", "struct",
	"chan", "else", "goto", "package", "switch",
	"const", "fallthrough", "if", "range", "type",
	"continue", "for", "import", "return", "var",
}

type EnemySprite struct {
	position Position
	fontSize FontSize
	color    color.Color
	text     string
	cursor   int
	isDead   bool
}

func NewEnemySprite() *EnemySprite {
	return &EnemySprite{fontSize: baseFontSize, color: color.RGBA{255, 0, 0, 255}, text: words[rand.Intn(len(words))]}
}

func (s *EnemySprite) Update() {
	s.position.x -= 1
}

func (s *EnemySprite) Draw(screen *ebiten.Image) {
	rectSprite := RectSprite{s.position, Size{s.Frame().w, s.Frame().h}, s.color}
	rectSprite.Draw(screen)

	textSprite := TextSprite{s.position, s.fontSize, color.White, s.text, AnchorLeftTop}
	textSprite.Draw(screen)

	textSprite = TextSprite{s.position, s.fontSize, color.RGBA{200, 200, 200, 255}, s.text[:s.cursor], AnchorLeftTop}
	textSprite.Draw(screen)
}

func (s *EnemySprite) Frame() Frame {
	return Frame{s.position.x, s.position.y, float64(font.MeasureString(s.fontSize.Font(), s.text).Round()), s.fontSize.Raw()}
}

func (s *EnemySprite) Input(str string) {
	if strings.ToUpper(string(s.text[s.cursor])) == str {
		s.cursor++

		if s.cursor >= len(s.text) {
			s.isDead = true
		}
	}
}
