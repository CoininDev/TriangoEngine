package lib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Move struct {
	IsActive bool
	e        *Entity
	Velocity Vector2
	Speed    float64
}

func NewMove(speed float64) *Move {
	return &Move{IsActive: true, e: nil, Velocity: Vector2{X: 0, Y: 0}, Speed: speed}
}

func (m *Move) Start(e *Entity) { m.e = e }
func (m *Move) Tick() {
	direction := Vector2f{}
	if rl.IsKeyDown(rl.KeyW) {
		direction.Y -= 1
	}
	if rl.IsKeyDown(rl.KeyS) {
		direction.Y += 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		direction.X -= 1
	}
	if rl.IsKeyDown(rl.KeyD) {
		direction.X += 1
	}
	m.Velocity.X = int32(direction.Normalize().X * m.Speed)
	m.Velocity.Y = int32(direction.Normalize().Y * m.Speed)
	m.e.Position.X += m.Velocity.X
	m.e.Position.Y += m.Velocity.Y
}
func (m *Move) End()            {}
func (m *Move) GetType() string { return "Move" }
