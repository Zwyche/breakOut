package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Brick struct {
	Rectangle rl.Rectangle
	color rl.Color
}

func NewBrick(Pos rl.Vector2, height float32, width float32, color rl.Color) Brick {
	rec := rl.NewRectangle(Pos.X, Pos.Y, width, height)
	return Brick {
		Rectangle: rec,
		color: color,
	}
}

func (b *Brick) Draw() {
	rl.DrawRectangleRec(b.Rectangle, b.color)
}