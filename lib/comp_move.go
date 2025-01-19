package lib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Move struct {
	Active   bool
	e        *Entity
	Velocity Vector2f
	Speed    float64
}

func NewMove(speed float64) *Move {
	return &Move{true, nil, Vector2f{X: 0, Y: 0}, speed}
}

func (c *Move) IsActive() bool     { return c.Active }
func (c *Move) SetActive(new bool) { c.Active = new }
func (m *Move) Start(e *Entity)    { m.e = e }
func (m *Move) Tick() {
	sprite := NewSprite(rl.Texture2D{}, 1, 1)
	if i, c := m.e.GetComponent("Sprite"); i != -1 {
		sprite = c.(*Sprite)
	}

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

	sprite.fliph = int(direction.X)
	m.Velocity.X = direction.Normalize().X * m.Speed
	m.Velocity.Y = direction.Normalize().Y * m.Speed
	m.e.Position.X += m.Velocity.X
	m.e.Position.Y += m.Velocity.Y
}
func (m *Move) End()            {}
func (m *Move) GetType() string { return "Move" }
