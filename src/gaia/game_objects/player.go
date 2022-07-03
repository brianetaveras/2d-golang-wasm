package gaia

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	gaia "world/src/gaia/component"
)

func CreateNewPlayer() *gaia.Element {

	newPlayer := &gaia.Element{}

	newPlayer.Active = true
	newPlayer.Rotation = 0
	newPlayer.Position = gaia.Vector{X: 0, Y: 0}
	newPlayer.Tag = "Player"

	_, b, _, _ := runtime.Caller(2)

	basePath := "/"
	
	if os.Getenv("GOOS") != "js" {
		basePath = filepath.Dir(b)
	}

	fullSpritePath := fmt.Sprintf("%s\\assets\\data\\graphics\\characters\\player.png", basePath)

	opts := gaia.AnimatorSequenceOptions{FramesX: 6, FramesY: 1, FrameSizeX: 36, FrameSizeY: 24, Loop: true, SampleRate: 5}
	walkSequence, _ := gaia.NewAnimationSequence(newPlayer, fullSpritePath, opts)

	animationSequences := map[string]*gaia.AnimatorSequence{
		"walk": walkSequence,
	}

	spriteAnimator := gaia.NewSpriteAnimator(newPlayer, animationSequences, "walk")

	newPlayer.Components = append(newPlayer.Components, spriteAnimator)

	return newPlayer
}
