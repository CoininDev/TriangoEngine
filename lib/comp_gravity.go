package lib

import "fmt"

type Gravity struct {
	Active bool
	e      *Entity
	force  Vector2
}

func NewGravity(force Vector2) *Gravity {
	return &Gravity{true, nil, force}
}

func (c *Gravity) IsActive() bool     { return c.Active }
func (c *Gravity) SetActive(new bool) { c.Active = new }
func (g *Gravity) Start(e *Entity)    { g.e = e }

func (g *Gravity) Tick() {
	velocity := Vector2{g.force.X, g.force.Y}
	body := g.e.GetComponent("PhysicalBody").(*PhysicalBody)
	if body != nil {
		fmt.Printf("%+v\n", body.onFloor)
		if body.onFloor {
			velocity.Y = 0
		}
	}
	g.e.Position.X += velocity.X
	g.e.Position.Y += velocity.Y
}

func (g *Gravity) End() {}

func (g *Gravity) GetType() string        { return "Gravity" }
func (g *Gravity) GetInfo() []interface{} { return []interface{}{} }
