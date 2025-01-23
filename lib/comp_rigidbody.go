package lib

// Type: 0 - Static, 1 - Dynamic
type RigidBody struct {
	Active   bool
	Rect     Rectf
	Type     int
	Velocity Vector2f
}

func (c *RigidBody) GetType() string { return "RigidBody" }
func NewRigidBody(body_type int, w, h float32) *RigidBody {
	return &RigidBody{
		Active: true,
		Rect:   NewRectf(NewVector2f(0, 0), NewVector2f(w, h)),
		Type:   body_type,
	}
}
