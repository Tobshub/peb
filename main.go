package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 450

	GAME_SPEED = 80
)

var (
	Explosion            rl.Texture2D
	ExplosionFrameWidth  float32
	ExplosionFrameHeight float32
)

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "PEB")
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
)

func InitGame() { BUTTONS = []*Button{} }

func DrawGame() {
	for _, e := range EXPLOSIONS {
		e.Draw()
	}

	for _, b := range BUTTONS {
		b.Draw()
	}
}

var frame = 0

func UpdateGame() {
	frame++
	if frame%GAME_SPEED == 0 {
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
}
