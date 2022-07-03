package gaia

import (
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SpriteAnimator struct {
	container       *Element
	sequences       map[string]*AnimatorSequence
	current         string
	lastFrameChange time.Time
	finished        bool
}

func NewSpriteAnimator(container *Element, sequences map[string]*AnimatorSequence, defaultSequence string) *SpriteAnimator {

	animator := &SpriteAnimator{}

	animator.container = container
	animator.sequences = sequences
	animator.current = defaultSequence
	animator.lastFrameChange = time.Now()

	return animator

}

func (animator *SpriteAnimator) onUpdate() error {
	sequence := animator.sequences[animator.current]

	frameInterval := float64(time.Second) / sequence.SampleRate

	if time.Since(animator.lastFrameChange) >= time.Duration(frameInterval) {
		animator.finished = sequence.NextFrame()
		animator.lastFrameChange = time.Now()
	}

	return nil
}

func (animator *SpriteAnimator) onDraw(renderer *ebiten.Image) error {
	img := animator.sequences[animator.current].Texture()

	drawImageOptions := &ebiten.DrawImageOptions{}
	drawImageOptions.GeoM.Translate(animator.container.Position.X, animator.container.Position.Y)
	renderer.DrawImage(img, nil)

	return nil
}

func (animator *SpriteAnimator) SetSequence(name string) {
	animator.current = name
	animator.lastFrameChange = time.Now()
}

type AnimatorSequence struct {
	Textures   []*ebiten.Image
	Frame      int
	SampleRate float64
	Loop       bool
}

type AnimatorSequenceOptions struct {
	SampleRate float64
	Loop       bool
	FramesX    int
	FramesY    int
	FrameSizeX int
	FrameSizeY int
}

func NewAnimationSequence(container *Element, filePath string, options AnimatorSequenceOptions) (*AnimatorSequence, error) {

	sequence := &AnimatorSequence{}

	img, _, err := ebitenutil.NewImageFromFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	for y := 0; y <= options.FramesY-1; y++ {
		for x := 0; x <= options.FramesX-1; x++ {
			frameRect := image.Rectangle{Min: image.Point{X: options.FramesX * x, Y: options.FrameSizeY * y}, Max: image.Point{X: x + options.FrameSizeX, Y: y + options.FrameSizeY}}
			frameImage := img.SubImage(frameRect)

			ebitenImage := ebiten.NewImageFromImage(frameImage)

			sequence.Textures = append(sequence.Textures, ebitenImage)

		}
	}

	sequence.Loop = options.Loop
	sequence.SampleRate = options.SampleRate

	return sequence, nil

}
func (sequence *AnimatorSequence) Texture() *ebiten.Image {
	return sequence.Textures[sequence.Frame]
}

func (sequence *AnimatorSequence) NextFrame() bool {
	if sequence.Frame == len(sequence.Textures)-1 {
		if sequence.Loop {
			sequence.Frame = 0
		} else {
			return true
		}
	} else {
		sequence.Frame++
	}
	return false
}
