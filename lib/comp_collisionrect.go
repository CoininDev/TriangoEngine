package lib

import (
	"fmt"

	"github.com/ByteArena/box2d"
)

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

type Box2dBody struct {
	Active     bool
	e          *Entity
	def        box2d.B2BodyDef
	fixtureDef box2d.B2FixtureDef
	body       box2d.B2Body
	shape      *box2d.B2PolygonShape
	world      box2d.B2World
}

// The body type: static, kinematic, or dynamic.
func NewBox2dBody(w, h float64, body_type uint8) *Box2dBody {
	def := box2d.B2BodyDef{Active: true, Type: body_type}
	shape := box2d.NewB2PolygonShape()
	shape.SetAsBox(w/2, h/2)
	fixtureDef := box2d.B2FixtureDef{Shape: shape, Density: 1.0, Friction: 0.3}

	return &Box2dBody{true, nil, def, fixtureDef, box2d.B2Body{}, shape, box2d.B2World{}}
}
func (c *Box2dBody) GetType() string { return "Box2dBody" }
func (c *Box2dBody) BeginContact(contact box2d.B2ContactInterface) {
	fmt.Printf("%+v", c.e.ID)

}
func (c *Box2dBody) EndContact(contact box2d.B2ContactInterface) {}
func (c *Box2dBody) PreSolve(contact box2d.B2ContactInterface, oldManifold box2d.B2Manifold) {
}
func (c *Box2dBody) PostSolve(contact box2d.B2ContactInterface, impulse *box2d.B2ContactImpulse) {
}
