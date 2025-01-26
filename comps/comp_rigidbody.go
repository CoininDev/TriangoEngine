package comps

import (
	"triango/utils"
)

// Type: 0 - Static, 1 - Dynamic
type RigidBody struct {
	Active       bool
	Rect         utils.Rect
	Type         int
	Velocity     utils.Vector2
	Acceleration utils.Vector2
}

func (c *RigidBody) GetType() string { return "RigidBody" }
func NewRigidBody(body_type int, w, h float32) *RigidBody {
	return &RigidBody{
		Active: true,
		Rect:   utils.NewRectf(utils.NewVector2f(0, 0), utils.NewVector2f(w, h)),
		Type:   body_type,
	}
}

func (c *RigidBody) AddForce(force utils.Vector2) {
	c.Acceleration = c.Acceleration.Add(force)
}
