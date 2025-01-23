package lib

type Transform struct {
	Active   bool
	Position Vector2f
	Scale    Vector2f
	Rotation float32
}

func NewTransform(position, scale Vector2f, rotation float32) *Transform {
	return &Transform{
		Active:   true,
		Position: position,
		Scale:    scale,
		Rotation: rotation,
	}
}
func (t *Transform) GetType() string { return "Transform" }
