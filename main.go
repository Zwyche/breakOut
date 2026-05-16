package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func reset(paddle *Paddle, ball *Ball, grid *Grid, start *bool, gameOver *bool) {
	paddle.Rectangle.X = 350
	paddle.Rectangle.Y = 390

	ball.PhysicsBody.Pos = rl.NewVector2(paddle.Rectangle.X + paddle.Rectangle.Width/2, paddle.Rectangle.Y - ball.Radius)
	ball.PhysicsBody.Vel = rl.NewVector2(0,0)
	ball.PhysicsBody.speed = 200

	*grid = NewGrid()

	*start = true
	*gameOver = false
}

func main() {
	rl.InitWindow(800, 450, "Break Out!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	paddle := NewPaddle(rl.NewVector2(350, 390), 10, 100, rl.RayWhite, 400)
	ball := NewBall(rl.NewVector2(400, 380), 10, rl.Pink, rl.NewVector2(0, 0), 200)
	start := true
	grid := NewGrid()
	gameOver := false

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		if rl.IsKeyDown(rl.KeyA) && paddle.Rectangle.X > 0 {
			paddle.Rectangle.X -= paddle.speed * rl.GetFrameTime()
			if start {
				ball.Pos.X -= paddle.speed * rl.GetFrameTime()
			}
		}
		if rl.IsKeyDown(rl.KeyD) && paddle.Rectangle.X < 700 {
			paddle.Rectangle.X += paddle.speed * rl.GetFrameTime()
			if start {
				ball.Pos.X += paddle.speed * rl.GetFrameTime()
			}
		}
		if rl.IsKeyPressed(rl.KeySpace) && start {
			start = false
			ball.Launch()
		}
		if !start {
			grid.CheckHit(&ball)
			ball.PhysicsUpdate()
			ball.Bounce()
			ball.HitPaddle(paddle)
		}

		if ball.PhysicsBody.Pos.Y > 460 {
			gameOver = true
		}

		if len(grid.bricks) == 0 {
			gameOver = true
		}

		if gameOver {
			reset(&paddle, &ball, &grid, &start, &gameOver)
		}

		grid.Draw()
		paddle.Draw()
		ball.Draw()

		rl.EndDrawing()
	}
}
