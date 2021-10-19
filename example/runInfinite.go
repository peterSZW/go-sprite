package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/peterSZW/go-sprite"
)

const (
	windowWidth  = 320 // Width of the window
	windowHeight = 240 // Height of the window
	scale        = 2   // Scale of the window
)

var (
	explosion1, explosion2, explosion3, explosion4, zoom1, rotate1, skew1 *sprite.Sprite
)

// update at every frame
func update(surface *ebiten.Image) error {

	// frame skip
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// draw sprite
	explosion1.Draw(surface)
	explosion2.Draw(surface)
	explosion3.Draw(surface)

	return nil
}

func main() {

	explosionDuration := 500
	// create some explosions
	explosion1 = sprite.NewSprite()
	explosion1.AddAnimation("default", "gfx/explosion1.png", explosionDuration, 5, ebiten.FilterNearest)
	explosion1.Position(10, windowHeight/3*2)
	explosion1.Start()

	explosion2 = sprite.NewSprite()
	explosion2.AddAnimation("default", "gfx/explosion2.png", explosionDuration, 7, ebiten.FilterNearest)
	explosion2.Position(windowWidth/2-24, windowHeight/3*2)
	explosion2.Start()

	explosion3 = sprite.NewSprite()
	explosion3.AddAnimation("default", "gfx/explosion3.png", explosionDuration, 9, ebiten.FilterNearest)
	explosion3.Position(windowWidth-10-48, windowHeight/3*2)
	explosion3.Start()

	// infinite loop
	if err := ebiten.Run(update, windowWidth, windowWidth, scale, "Sprite demo"); err != nil {
		log.Fatal(err)
	}
}
