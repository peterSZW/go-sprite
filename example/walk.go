package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"

	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/peterSZW/go-sprite"
)

const (
	windowWidth  = 320 // Width of the window
	windowHeight = 240 // Height of the window
	scale        = 2   // Scale of the window
	girlSpeed    = windowWidth / 160
)

var (
	girl *sprite.Sprite
)

// update at every frame
func update(surface *ebiten.Image) error {

	// manage controle
	binding()

	// reset position if outside of the screen
	if girl.X > windowWidth {
		girl.X = 0 - girl.GetWidth()
	}
	if girl.X+girl.GetWidth() < 0 {
		girl.X = windowWidth
	}
	if girl.Y+girl.GetHeight() < 0 {
		girl.Y = windowHeight + 2*girl.GetHeight()
	}
	if girl.Y-2*girl.GetHeight() > windowHeight {
		girl.Y = 0 - girl.GetHeight()
	}

	// frame skip
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// draw sprite
	girl.Draw(surface)

	// display some informations
	drawFPS(surface)

	return nil
}

func main() {

	// create new sprite and load animations
	girl = sprite.NewSprite()
	girl.AddAnimation("stand-right", "gfx/som_girl_stand_right.png", 0, 1, ebiten.FilterDefault)
	girl.AddAnimation("walk-right", "gfx/som_girl_walk_right.png", 700, 6, ebiten.FilterDefault)
	girl.AddAnimation("stand-left", "gfx/som_girl_stand_left.png", 0, 1, ebiten.FilterDefault)
	girl.AddAnimation("walk-left", "gfx/som_girl_walk_left.png", 700, 6, ebiten.FilterDefault)
	girl.AddAnimation("stand-up", "gfx/som_girl_stand_up.png", 0, 1, ebiten.FilterDefault)
	girl.AddAnimation("walk-up", "gfx/som_girl_walk_up.png", 500, 4, ebiten.FilterDefault)
	girl.AddAnimation("stand-down", "gfx/som_girl_stand_down.png", 0, 1, ebiten.FilterDefault)
	girl.AddAnimation("walk-down", "gfx/som_girl_walk_down.png", 500, 4, ebiten.FilterDefault)

	// set position and first animation
	girl.Position(windowWidth/2, windowHeight/2)
	girl.CurrentAnimation = "stand-right"
	girl.Start()

	// infinite loop
	if err := ebiten.Run(update, windowWidth, windowWidth, scale, "Sprite demo"); err != nil {
		log.Fatal(err)
	}
}

// display some stuff on the screen
func drawFPS(surface *ebiten.Image) {
	ebitenutil.DebugPrint(surface,
		fmt.Sprintf("FPS:%0.1f  X:%d Y:%d %s\nLeft:%v Right:%v Up:%v Down:%v",
			ebiten.CurrentFPS(),
			int(girl.X), int(girl.Y),
			girl.CurrentAnimation,
			ebiten.IsKeyPressed(ebiten.KeyLeft),
			ebiten.IsKeyPressed(ebiten.KeyRight),
			ebiten.IsKeyPressed(ebiten.KeyUp),
			ebiten.IsKeyPressed(ebiten.KeyDown),
		))
}

func binding() {

	//////////////////////////// GO THE RIGHT
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {

		if ebiten.IsKeyPressed(ebiten.KeyUp) { // Right+Up
			girl.Direction = 45
			girl.Speed = girlSpeed + 1
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) { // Right+Down
			girl.Direction = -45
			girl.Speed = girlSpeed + 1
		} else { // Right
			girl.Direction = 0
			girl.Speed = girlSpeed
		}
		girl.CurrentAnimation = "walk-right"
		girl.Start() // Show, Reset, Resume
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyRight) {
		girl.Speed = 0
		girl.CurrentAnimation = "stand-right"
	}

	//////////////////////////// GO THE LEFT
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {

		if ebiten.IsKeyPressed(ebiten.KeyUp) { // Left+Up
			girl.Direction = 135
			girl.Speed = girlSpeed + 1
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) { // Left+Down
			girl.Direction = 225
			girl.Speed = girlSpeed + 1
		} else { // Left
			girl.Speed = girlSpeed
			girl.Direction = 180
		}

		girl.CurrentAnimation = "walk-left"
		girl.Start() // Show, Reset, Resume
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		girl.Speed = 0
		girl.CurrentAnimation = "stand-left"
	}

	//////////////////////////// GO THE TOP
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {

		if ebiten.IsKeyPressed(ebiten.KeyRight) { // Up+Right
			girl.Direction = 45
			girl.Speed = girlSpeed + 1
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) { // Up+Left
			girl.Direction = 135
			girl.Speed = girlSpeed + 1
		} else { // Up
			girl.Direction = 90
			girl.Speed = girlSpeed
		}

		girl.CurrentAnimation = "walk-up"
		girl.Start() // Show, Reset, Resume
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
		girl.Speed = 0
		girl.CurrentAnimation = "stand-up"
	}

	//////////////////////////// GO THE BOTTOM
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {

		if ebiten.IsKeyPressed(ebiten.KeyRight) { // Down+Right
			girl.Direction = -45
			girl.Speed = girlSpeed + 1
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) { // Down+Left
			girl.Direction = 225
			girl.Speed = girlSpeed + 1
		} else { // Down
			girl.Speed = girlSpeed
			girl.Direction = 270
		}

		girl.CurrentAnimation = "walk-down"
		girl.Start() // Show, Reset, Resume
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		girl.Speed = 0
		girl.CurrentAnimation = "stand-down"
	}
}
