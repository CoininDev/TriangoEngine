package systems

import (
	"triango/comps"
	"triango/game"
	"triango/utils"
)

type PhysicsSystem struct {
	Gravity utils.Vector2
	notable []*game.Entity
}

func (s *PhysicsSystem) Start(g *game.Game) {
	s.Gravity = utils.NewVector2f(0, 98)
}

func (s *PhysicsSystem) Refresh(g *game.Game) {
	for _, e := range g.Entities {
		if i := e.GetComponent("RigidBody"); i != -1 {
			s.notable = append(s.notable, &e)
		}
	}
}

func (s *PhysicsSystem) End(g *game.Game) {}
func (s *PhysicsSystem) Update(g *game.Game, delta float64) {
	for _, e := range s.notable {
		b := e.Components[e.GetComponent("RigidBody")].(*comps.RigidBody)
		t := e.Components[e.GetComponent("Transform")].(*comps.Transform)
		b.Rect.Position = t.Position

		// CHECKING ACTIVITY
		if !b.Active {
			continue
		}

		// APPLYING GRAVITY
		if b.Type == 1 {
			b.Velocity = b.Velocity.Add(s.Gravity.Mulf(float32(delta)))
		}

		// CHECKING COLLISIONS
		for _, otherE := range g.Entities {
			if otherE.ID == e.ID {
				continue
			}

			if j := otherE.GetComponent("RigidBody"); j != -1 {
				other := otherE.Components[j].(*comps.RigidBody)
				intersection := b.Rect.GetIntersectionWith(other.Rect)
				if intersection.PenetrationDepth > 0 && b.Type == 1 {
					b.AddForce(intersection.NormalVector)
				}
			}
		}

		// APPLYING VELOCITY
		b.Velocity = b.Velocity.Add(b.Acceleration.Mulf(float32(delta)))
		t.Position = t.Position.Add(b.Velocity)
	}
}
