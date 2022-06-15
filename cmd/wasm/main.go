package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (game *Game) Update() error {
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, wasm!")
	ebitenutil.DebugPrintAt(screen, "Here I go! (get it?)", 0, 20)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, world!")
	err := ebiten.RunGame(&Game{})

	if err != nil {
		log.Fatal(err)
	}

}
