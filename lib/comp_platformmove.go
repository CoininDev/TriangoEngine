package lib

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlatformMove struct {
	Active    bool
	e         *Entity
	velocity  Vector2
	speed     int32
	jumpForce int32
}

func NewPlatformMove(speed int32, jumpForce int32) *PlatformMove {
	return &PlatformMove{true, nil, Vector2{0, 0}, speed, jumpForce}
}

func (c *PlatformMove) IsActive() bool     { return c.Active }
func (c *PlatformMove) SetActive(new bool) { c.Active = new }
func (c *PlatformMove) GetType() string    { return "PlatformMove" }

func (m *PlatformMove) Start(e *Entity) {
	m.e = e
	if e.GetComponent("Gravity") == nil {
		m.Active = false
		fmt.Println("O componente PlatformMove exige um componente Gravity.")
	}
}

func (m *PlatformMove) Tick() {
	vel := Vector2{0, 0}
	a := 0
	d := 0
	w := 0
	if rl.IsKeyDown(rl.KeyA) {
		a = 1
	}
	if rl.IsKeyDown(rl.KeyD) {
		d = 1
	}
	if rl.IsKeyPressed(rl.KeyW) {
		w = 1
	}
	dir := -a + d

	vel.X = int32(dir * int(m.speed))
	vel.Y = int32(-w * int(m.jumpForce))

	m.e.Position.X += vel.X
	m.e.Position.Y += vel.Y
}

func (m *PlatformMove) End() {}
