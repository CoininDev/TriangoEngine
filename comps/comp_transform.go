package comps

import "triango/utils"

type Transform struct {
	Active   bool
	Position utils.Vector2f
	Scale    utils.Vector2f
	Rotation float32
}

func NewTransform(position, scale utils.Vector2f, rotation float32) *Transform {
	return &Transform{
		Active:   true,
		Position: position,
		Scale:    scale,
		Rotation: rotation,
	}
}
func (t *Transform) GetType() string { return "Transform" }
