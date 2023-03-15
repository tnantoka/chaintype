package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tnantoka/chaiping/game"
)

func main() {
	game := game.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
