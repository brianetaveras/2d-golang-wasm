package gaia

import (
	"fmt"
	component "world/src/gaia/component"
	gaia "world/src/gaia/game_objects"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var Elements []*component.Element = []*component.Element{}

func (game *Game) Update() error {
	for _, element := range Elements {
		err := element.Update()

		if err != nil {
			fmt.Println("Error updating element:", err)
		}

	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	for _, element := range Elements {
		if element.Active {
			err := element.Draw(screen)

			if err != nil {
				fmt.Println("Error updating element:", err)
			}
		}
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func RunGame() error {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gaia")

	game := &Game{}
	
	player := gaia.CreateNewPlayer()
	
	Elements = append(Elements, player)
	
	err := ebiten.RunGame(game)

	if err != nil {
		return err
	}

	return nil

}
