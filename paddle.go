package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Paddle struct {
	Rectangle rl.Rectangle
	color rl.Color
	speed float32
}

func NewPaddle(Pos rl.Vector2, height float32, width float32, color rl.Color, speed float32) Paddle {
	rec := rl.NewRectangle(Pos.X, Pos.Y, width, height)
	return Paddle {
		Rectangle: rec,
		color: color,
		speed: speed,
	}
}

func (p *Paddle) Draw() {
	rl.DrawRectangleRec(p.Rectangle, p.color)
}
