package lib

type CollisionRect struct {
	Active bool
	Top    bool
	Bottom bool
	Left   bool
	Right  bool
	rect   Rectf
}

func NewCollisionRect(w, h float32) *CollisionRect {
	return &CollisionRect{
		Active: true,
		rect:   NewRectf(NewVector2f(0, 0), NewVector2f(w, h)),
	}
}

func (c *CollisionRect) GetType() string { return "CollisionRect" }
func (r Rectf) Intersects(other Rectf) bool {
	return r.Position.X < other.Position.X+other.Size.X &&
		r.Position.X+r.Size.X > other.Position.X &&
		r.Position.Y < other.Position.Y+other.Size.Y &&
		r.Position.Y+r.Size.Y > other.Position.Y
}
