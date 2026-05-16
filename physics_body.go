package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PhysicsBody struct {
	Pos    rl.Vector2
	Vel    rl.Vector2
	Radius float32
	speed float32
	ignoreCollisions bool
}

func NewPhysicsBody(newPos rl.Vector2, newVel rl.Vector2, newRadius float32, speed float32) PhysicsBody {
	pb := PhysicsBody{Pos: newPos, Vel: newVel, Radius: newRadius, speed: speed}
	pb.ignoreCollisions = false
	return pb
}

func (pb *PhysicsBody) Bounce() {
	if pb.Pos.X < 5 || pb.Pos.X > 795 {
		pb.Vel.X = pb.Vel.X * -1
	}
	if pb.Pos.Y < 5 {
		pb.Vel.Y = pb.Vel.Y * -1
	}
}

func (pb PhysicsBody) DrawBoundary() {
	rl.DrawCircleLines(int32(pb.Pos.X), int32(pb.Pos.Y), pb.Radius, rl.Lime)
}

func (pb *PhysicsBody) SetIgnoreCollisions(ignore bool) {
	pb.ignoreCollisions = ignore
}

func (pb *PhysicsBody) PhysicsUpdate() {
	pb.VelocityTick()
}

func (pb *PhysicsBody) VelocityTick() {
	adjustedVel := rl.Vector2Scale(pb.Vel, rl.GetFrameTime())
	pb.Pos = rl.Vector2Add(pb.Pos, adjustedVel)
}

func (pb *PhysicsBody) Launch() {
	if rl.IsKeyDown(rl.KeyA) {
		pb.Vel = rl.NewVector2(-pb.speed, -pb.speed)
	} else if rl.IsKeyDown(rl.KeyD) {
		pb.Vel = rl.NewVector2(pb.speed, -pb.speed)
	} else {
		pb.Vel = rl.NewVector2(0, -pb.speed)
	}
}

func (pb *PhysicsBody) HitPaddle(paddle Paddle) {
	hit := rl.CheckCollisionCircleRec(pb.Pos, pb.Radius, paddle.Rectangle)
	if hit {
		paddleCenter := paddle.Rectangle.X + paddle.Rectangle.Width/2
		distance := pb.Pos.X - paddleCenter
		ratio := distance / (paddle.Rectangle.Width / 2)
		pb.Vel.Y = pb.speed * -1
		pb.Vel.X = pb.speed * ratio
	}
}

// func (pb *PhysicsBody) GravityTick() {
// 	pb.Vel = rl.Vector2Add(pb.Vel, rl.Vector2Scale(rl.NewVector2(0, pb.Gravity), rl.GetFrameTime()))
// }
