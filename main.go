package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tnantoka/chaiping/game"
)

func main() {
	if err := ebiten.RunGame(&game.Game{}); err != nil {
		log.Fatal(err)
	}
}
