package systems

import (
	. "triango/comps"
	. "triango/game"
	. "triango/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MoveSystem struct {
	notable []*Entity
}

func (s *MoveSystem) Start(g *Game) {}
func (s *MoveSystem) End(g *Game)   {}
func (s *MoveSystem) Refresh(g *Game) {
	for _, e := range g.Entities {
		if i := e.GetComponent("Move"); i != -1 {
			s.notable = append(s.notable, &e)
		}
	}
}
func (s *MoveSystem) Update(g *Game, delta float64) {
	for _, e := range s.notable {
		m := e.Components[e.GetComponent("Move")].(*Move)
		rb := e.Components[e.GetComponent("RigidBody")].(*RigidBody)
		sprite := e.Components[e.GetComponent("Sprite")].(*Sprite)

		// CHECKING ACTIVITY
		if !m.Active {
			continue
		}

		direction := Vector2{}
		if rl.IsKeyDown(rl.KeyW) {
			direction.Y -= 1
		}
		if rl.IsKeyDown(rl.KeyS) {
			direction.Y += 1
		}
		if rl.IsKeyDown(rl.KeyA) {
			direction.X -= 1
		}
		if rl.IsKeyDown(rl.KeyD) {
			direction.X += 1
		}

		if direction.X != 0 {
			sprite.Fliph = int(direction.X)
		}
		rb.Velocity = direction.Normalize().Mulf(m.Speed * float32(delta))
	}
}
