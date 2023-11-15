package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	EXPLOSION_TEXTURE_NUM_LINES           = 5
	EXPLOSION_TEXTURE_NUM_FRAMES_PER_LINE = 5
)

type ExplosionEffect struct {
	FrameRect    rl.Rectangle
	Pos          rl.Vector2
	FrameCounter float32
	Active       bool
}

func NewExplosionEffect(pos rl.Vector2) *ExplosionEffect {
	pos.X -= ExplosionFrameWidth / 2
	pos.Y -= ExplosionFrameHeight / 2

	return &ExplosionEffect{
		FrameRect:    rl.Rectangle{X: 0, Y: 0, Width: ExplosionFrameWidth, Height: ExplosionFrameHeight},
		Pos:          pos,
		FrameCounter: 0,
		Active:       true,
	}
}

func (e *ExplosionEffect) Draw() {
	if !e.Active {
		return
	}

	rl.DrawTextureRec(Explosion, e.FrameRect, e.Pos, rl.White)
}

func (e *ExplosionEffect) Update() {
	if !e.Active {
		return
	}

	if e.FrameCounter >= 25 {
		e.Active = false
		return
	}

	currentFrame := int(e.FrameCounter) % EXPLOSION_TEXTURE_NUM_FRAMES_PER_LINE
	currentLine := int(e.FrameCounter) / EXPLOSION_TEXTURE_NUM_FRAMES_PER_LINE

	e.FrameRect.X = ExplosionFrameWidth * float32(currentFrame)
	e.FrameRect.Y = ExplosionFrameHeight * float32(currentLine)
	e.FrameCounter += .5
}
