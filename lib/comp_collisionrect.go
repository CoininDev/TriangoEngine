package lib

type CollisionRect struct {
	Active       bool
	e            *Entity
	Top          bool
	Bottom       bool
	Left         bool
	Right        bool
	transformIdx int
	rect         Rectf
}

func NewCollisionRect(w, h float32) *CollisionRect {
	return &CollisionRect{
		Active: true,
		rect:   NewRectf(NewVector2f(0, 0), NewVector2f(w, h)),
	}
}

func (c *CollisionRect) IsActive() bool     { return c.Active }
func (c *CollisionRect) SetActive(new bool) { c.Active = new }

func (c *CollisionRect) Start(e *Entity) {
	c.e = e
	c.transformIdx, _ = e.GetComponent("Transform")
}

func (c *CollisionRect) Tick() {
	t := c.e.components[c.transformIdx].(*Transform)
	c.rect.Position = t.Position

	// Reset collision states
	c.Top = false
	c.Bottom = false
	c.Left = false
	c.Right = false

	// Check for collisions with other entities
	for _, entity := range c.e.Game.Entities {
		if entity.ID != c.e.ID {
			if i, comp := entity.GetComponent("CollisionRect"); i != -1 {
				other := comp.(*CollisionRect)
				if c.rect.Intersects(other.rect) {
					if c.rect.Position.Y < other.rect.Position.Y {
						c.Bottom = true
					}
					if c.rect.Position.Y > other.rect.Position.Y {
						c.Top = true
					}
					if c.rect.Position.X < other.rect.Position.X {
						c.Right = true
					}
					if c.rect.Position.X > other.rect.Position.X {
						c.Left = true
					}
				}
			}
		}
	}
}

func (c *CollisionRect) End()            {}
func (c *CollisionRect) GetType() string { return "CollisionRect" }

func (r Rectf) Intersects(other Rectf) bool {
	return r.Position.X < other.Position.X+other.Size.X &&
		r.Position.X+r.Size.X > other.Position.X &&
		r.Position.Y < other.Position.Y+other.Size.Y &&
		r.Position.Y+r.Size.Y > other.Position.Y
}
