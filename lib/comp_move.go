package lib

type Move struct {
	Active   bool
	Velocity Vector2f
	Speed    float32
}

func NewMove(speed float32) *Move {
	return &Move{true, Vector2f{X: 0, Y: 0}, speed}
}
func (m *Move) GetType() string { return "Move" }
