package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	SCREEN_WIDTH  float32 = 800
	SCREEN_HEIGHT float32 = 450

	GAME_SPEED = 5

	Explosion            rl.Texture2D
	ExplosionFrameWidth  float32
	ExplosionFrameHeight float32
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(int32(SCREEN_WIDTH), int32(SCREEN_HEIGHT), "PEB")
	defer rl.CloseWindow()

	Explosion = rl.LoadTexture("resources/explosion.png")
	ExplosionFrameWidth = float32(Explosion.Width / EXPLOSION_TEXTURE_NUM_FRAMES_PER_LINE)
	ExplosionFrameHeight = float32(Explosion.Height / EXPLOSION_TEXTURE_NUM_LINES)

	rl.SetTargetFPS(120)

	InitGame()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		DrawGame()
		UpdateGame()
		rl.EndDrawing()
	}

	rl.UnloadTexture(Explosion)
}

var (
	BUTTONS    = []*Button{}
	EXPLOSIONS = []*ExplosionEffect{}

	SCORE = 0
)

func InitGame() {
	SCORE = 0
	BUTTONS = []*Button{}
	EXPLOSIONS = []*ExplosionEffect{}
}

func DrawGame() {
	scoreText := fmt.Sprintf("Score: %d", SCORE)
	fontSize := int32(24)
	rl.DrawText(scoreText, rl.MeasureText(scoreText, fontSize)/2,
		fontSize/2, fontSize, rl.Gray)

	for _, e := range EXPLOSIONS {
		e.Draw()
	}

	for _, b := range BUTTONS {
		b.Draw()
	}
}

var (
	frame          = 0
	canChangeSpeed = true
)

func UpdateGame() {
	if rl.IsWindowResized() {
		SCREEN_WIDTH = float32(rl.GetScreenWidth())
		SCREEN_HEIGHT = float32(rl.GetScreenHeight())
	}

	frame++
	if frame%(1000/GAME_SPEED) == 0 {
		newButton := NewButton(RandomSpawnPoint())
		BUTTONS = append(BUTTONS, newButton)
	}

	for idx := 0; idx < len(BUTTONS); idx++ {
		b := BUTTONS[idx]
		if b.Hidden {
			BUTTONS = append(BUTTONS[:idx], BUTTONS[idx+1:]...)
			idx--
		}

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			mousePos := rl.GetMousePosition()
			b.Update(&mousePos)
		} else {
			b.Update(nil)
		}
	}

	for idx := 0; idx < len(EXPLOSIONS); idx++ {
		e := EXPLOSIONS[idx]
		e.Update()
		if !e.Active {
			EXPLOSIONS = append(EXPLOSIONS[:idx], EXPLOSIONS[idx+1:]...)
			idx--
		}
	}

	// if player loses points till this happens
	// the speed increases, which is not intended
	// but oh well :shrug:
	if canChangeSpeed && SCORE%500 == 0 {
		GAME_SPEED += 2
		canChangeSpeed = false
	} else if !canChangeSpeed && SCORE%500 != 0 {
		canChangeSpeed = true
	}
}
