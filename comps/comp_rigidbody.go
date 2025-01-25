package comps

import (
	"triango/utils"
)

// Type: 0 - Static, 1 - Dynamic
type RigidBody struct {
	Active   bool
	Rect     utils.Rectf
	Type     int
	Velocity utils.Vector2f
}

func (c *RigidBody) GetType() string { return "RigidBody" }
func NewRigidBody(body_type int, w, h float32) *RigidBody {
	return &RigidBody{
		Active: true,
		Rect:   utils.NewRectf(utils.NewVector2f(0, 0), utils.NewVector2f(w, h)),
		Type:   body_type,
	}
}
