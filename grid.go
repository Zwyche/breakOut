package main

import (
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	bricks []Brick
}

func NewGrid() Grid {
	bricks := make([]Brick, 0, 50)

	var startX float32 = 55
	var startY float32 = 55
	var width float32 = 60
	var height float32 = 20
	var gap float32 = 10

	for x := 0; x < 10; x++ {
		for y := 0; y < 5; y++ {
			bricks = append(bricks, NewBrick(rl.NewVector2(startX+(width*(float32(x)))+(gap*(float32(x))), startY+(height*(float32(y)))+(gap*(float32(y)))), height, width, rl.Purple))
		}
	}

	return Grid{
		bricks: bricks,
	}
}

func (g *Grid) Draw() {
	for x := range g.bricks {
		g.bricks[x].Draw()
	}
}

func (g *Grid) CheckHit(b *Ball) {
	for x := range g.bricks {
		if rl.CheckCollisionCircleRec(b.Pos, b.Radius, g.bricks[x].Rectangle) {
			g.bricks = slices.Delete(g.bricks, x, x+1)
			b.PhysicsBody.Vel.Y = b.PhysicsBody.Vel.Y * -1
			b.PhysicsBody.speed += 5
			break
		}
	}
}
