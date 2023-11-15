package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var BUTTON_SIZE = rl.Vector2{X: 100, Y: 50}

type Button struct {
	Pos    rl.Vector2
	Size   rl.Vector2
	V      rl.Vector2
	Hidden bool
}

func NewButton(pos rl.Vector2) *Button {
	return &Button{
		Pos:    pos,
		Size:   BUTTON_SIZE,
		V:      RandV(pos),
		Hidden: false,
	}
}

func (b *Button) Draw() {
	if b.Hidden {
		return
	}
	rl.DrawRectangleV(b.Pos, b.Size, rl.Pink)
	text := "Click me"
	rl.MeasureText(text, 20)
	rl.DrawText(text, int32(b.Pos.X)+10, int32(b.Pos.Y)+10, 20, rl.White)
}

func (b *Button) MouseClicked(mouse rl.Vector2) bool {
	if b.Pos.X <= mouse.X && mouse.X <= b.Pos.X+b.Size.X &&
		b.Pos.Y <= mouse.Y && mouse.Y <= b.Pos.Y+b.Size.Y {
		return true
	}
	return false
}

func (b *Button) Update(mouse *rl.Vector2) {
	if b.Hidden {
		return
	}

	if mouse != nil {
		if b.MouseClicked(*mouse) {
			b.Hidden = true
			SCORE += 10
			EXPLOSIONS = append(EXPLOSIONS,
				NewExplosionEffect(*mouse))
			return
		}
	}

	b.Pos = rl.Vector2Add(b.Pos, b.V)

	if b.Pos.X+b.Size.X < 0 ||
		b.Pos.Y+b.Size.Y < 0 ||
		b.Pos.X > SCREEN_WIDTH ||
		b.Pos.Y > SCREEN_HEIGHT {
		b.Hidden = true
		SCORE -= 10
	}
}

type SpawnPoint int

const (
	Left   SpawnPoint = 0
	Right  SpawnPoint = 1
	Top    SpawnPoint = 2
	Bottom SpawnPoint = 3
)

func RandomSpawnPoint() (p rl.Vector2) {
	spawnPoint := SpawnPoint(rand.Intn(4))

	switch spawnPoint {
	case Left:
		p.X = 0 - BUTTON_SIZE.X
		p.Y = RandFloat(0, SCREEN_HEIGHT)
	case Right:
		p.X = SCREEN_WIDTH
		p.Y = RandFloat(0, SCREEN_HEIGHT)
	case Top:
		p.Y = 0 - BUTTON_SIZE.Y
		p.X = RandFloat(0, SCREEN_WIDTH)
	case Bottom:
		p.Y = SCREEN_HEIGHT
		p.X = RandFloat(0, SCREEN_WIDTH)
	}

	return
}

const (
	MinSpeed = .5
	MaxSpeed = 2.5
)

func RandV(pos rl.Vector2) (v rl.Vector2) {
	if pos.X < SCREEN_WIDTH/2 {
		v.X = RandFloat(MinSpeed, MaxSpeed)
	} else {
		v.X = RandFloat(-1*MaxSpeed, -1*MinSpeed)
	}

	if pos.Y < SCREEN_HEIGHT/2 {
		v.Y = RandFloat(MinSpeed, MaxSpeed)
	} else {
		v.Y = RandFloat(-1*MaxSpeed, -1*MinSpeed)
	}
	return
}

func RandFloat(min, max float32) float32 {
	return float32(min) + float32(max-min)*rand.Float32()
}
