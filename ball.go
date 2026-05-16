package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Ball struct {
	PhysicsBody
	color rl.Color
}

func NewBall(Pos rl.Vector2, radius float32, color rl.Color, vel rl.Vector2, speed float32) Ball {
	pb := NewPhysicsBody(Pos, vel, radius, speed)
	return Ball{
		PhysicsBody: pb,
		color: color,
	}
}

func (b *Ball) Draw() {
	rl.DrawCircle(int32(b.Pos.X), int32(b.Pos.Y), b.Radius, b.color)
}