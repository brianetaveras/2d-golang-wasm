package gaia

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SpriteRenderer struct {
	container *Element
	texture   *ebiten.Image
	width     int
	height    int
}

type SpriteDrawOptions struct {
	Width  int
	Height int
	PosX   float64
	PosY   float64
}

func NewSpriteRenderer(container *Element, filePath string, spriteDrawOptions SpriteDrawOptions) *SpriteRenderer {
	spriteRenderer := &SpriteRenderer{}

	img, _, err := ebitenutil.NewImageFromFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	croppedImage := img.SubImage(image.Rectangle{Min: image.Point{X: int(spriteDrawOptions.PosX), Y: int(spriteDrawOptions.PosY)}, Max: image.Point{X: spriteDrawOptions.Width, Y: spriteDrawOptions.Height}})

	ebitenImage := ebiten.NewImageFromImage(croppedImage)
	spriteRenderer.texture = ebitenImage
	spriteRenderer.container = container

	return spriteRenderer
}

func (spriteRenderer *SpriteRenderer) onUpdate() error {
	return nil
}

func (spriteRenderer *SpriteRenderer) onDraw(renderer *ebiten.Image) error {
	renderer.DrawImage(spriteRenderer.texture, nil)
	return nil
}
