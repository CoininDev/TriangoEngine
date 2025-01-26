package comps

import (
	"triango/utils"
)

type Move struct {
	Active   bool
	Velocity utils.Vector2
	Speed    float32
}

func NewMove(speed float32) *Move {
	return &Move{true, utils.Vector2{X: 0, Y: 0}, speed}
}
func (m *Move) GetType() string { return "Move" }
