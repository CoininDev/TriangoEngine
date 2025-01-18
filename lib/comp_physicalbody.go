package lib

import (
	"fmt"

	"github.com/solarlune/resolv"
)

type PhysicalBody struct {
	Active  bool
	e       *Entity
	shape   resolv.IShape
	onFloor bool
}

func NewPhysicalBody(active bool, size Vector2f) *PhysicalBody {
	return &PhysicalBody{active, nil, resolv.NewRectangle(0, 0, size.X, size.Y), false}
}

func (c *PhysicalBody) IsActive() bool     { return c.Active }
func (c *PhysicalBody) SetActive(new bool) { c.Active = new }
func (b *PhysicalBody) Start(e *Entity)    { b.e = e }
func (b *PhysicalBody) GetType() string    { return "PhysicalBody" }
func (b *PhysicalBody) Tick() {
	b.shape.SetPosition(float64(b.e.Position.X), float64(b.e.Position.Y))
	fmt.Println(b.shape.Position())
	b.onFloor = false
	for _, ent := range b.e.Game.Entities {
		// Ignorar colisão com a própria entidade
		if &ent == b.e {
			continue
		}
		body := ent.GetComponent("PhysicalBody").(*PhysicalBody)
		if body != nil {
			intersection := b.shape.Intersection(body.shape)
			if !intersection.IsEmpty() {
				b.onFloor = true
			}
		}
	}
}
func (b *PhysicalBody) End() {}
