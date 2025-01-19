package lib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Move struct {
	Active       bool
	e            *Entity
	Velocity     Vector2f
	Speed        float32
	transformIdx int
	spriteIdx    int
	collisionIdx int
}

func NewMove(speed float32) *Move {
	return &Move{true, nil, Vector2f{X: 0, Y: 0}, speed, -1, -1, -1}
}

func (c *Move) IsActive() bool     { return c.Active }
func (c *Move) SetActive(new bool) { c.Active = new }
func (m *Move) Start(e *Entity) {
	m.e = e
	m.transformIdx, _ = e.GetComponent("Transform")
	m.collisionIdx, _ = e.GetComponent("CollisionRect")
	m.spriteIdx, _ = e.GetComponent("Sprite")
	if m.spriteIdx == -1 || m.transformIdx == -1 || m.collisionIdx == -1 {
		panic("Move component requires Transform, CollisionRect and Sprite components")
	}
}
func (m *Move) Tick() {
	t := m.e.components[m.transformIdx].(*Transform)
	sprite := m.e.components[m.spriteIdx].(*Sprite)
	col := m.e.components[m.collisionIdx].(*CollisionRect)

	direction := Vector2f{}
	if rl.IsKeyDown(rl.KeyW) && !col.Top {
		direction.Y -= 1
	}
	if rl.IsKeyDown(rl.KeyS) && !col.Bottom {
		direction.Y += 1
	}
	if rl.IsKeyDown(rl.KeyA) && !col.Left {
		direction.X -= 1
	}
	if rl.IsKeyDown(rl.KeyD) && !col.Right {
		direction.X += 1
	}

	if direction.X != 0 {
		sprite.fliph = int(direction.X)
	}
	m.Velocity.X = direction.Normalize().X * m.Speed
	m.Velocity.Y = direction.Normalize().Y * m.Speed
	t.Position.X += m.Velocity.X
	t.Position.Y += m.Velocity.Y
}
func (m *Move) End()            {}
func (m *Move) GetType() string { return "Move" }
