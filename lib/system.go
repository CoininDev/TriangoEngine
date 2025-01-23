package lib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type System interface {
	Start(*Game)
	Update(*Game, float64)
	End(*Game)
}

type RenderSystem struct{}

func (s *RenderSystem) Start(g *Game) {}
func (s *RenderSystem) End(g *Game) {
	for _, e := range g.Entities {
		if i := e.GetComponent("Sprite"); i != -1 {
			s := e.components[i].(*Sprite)
			rl.UnloadTexture(s.Texture)
		}
	}
}
func (s *RenderSystem) Update(g *Game, delta float64) {
	for _, e := range g.Entities {
		if i := e.GetComponent("Sprite"); i != -1 {
			s := e.components[i].(*Sprite)
			t := e.components[e.GetComponent("Transform")].(*Transform)
			if !s.IsActive() {
				continue
			}
			rl.DrawTexturePro(
				s.Texture,
				rl.NewRectangle(0, 0, float32(s.Texture.Width*int32(s.Fliph)), float32(s.Texture.Height*int32(s.Flipv))),
				rl.NewRectangle(0, 0, float32(s.Texture.Width)*t.Scale.X, float32(s.Texture.Height)*t.Scale.Y),
				rl.Vector2(t.Position.Negative()),
				t.Rotation,
				rl.White,
			)
		}
		if i := e.GetComponent("PlaceholderSprite"); i != -1 {
			s := e.components[i].(*PlaceholderSprite)
			t := e.components[e.GetComponent("Transform")].(*Transform)
			if !s.IsActive() {
				continue
			}
			rl.DrawRectanglePro(
				rl.NewRectangle(0, 0, s.Size.X, s.Size.Y),
				rl.Vector2(t.Position.Negative()),
				t.Rotation,
				s.Color,
			)
		}
	}
}

type PhysicsSystem struct {
	Gravity Vector2f
}

func (s *PhysicsSystem) Start(g *Game) {
	s.Gravity = NewVector2f(0, 98)
}
func (s *PhysicsSystem) End(g *Game) {}
func (s *PhysicsSystem) Update(g *Game, delta float64) {
	for _, e := range g.Entities {
		if i := e.GetComponent("RigidBody"); i != -1 {
			b := e.components[i].(*RigidBody)
			t := e.components[e.GetComponent("Transform")].(*Transform)
			b.Rect.Position = t.Position

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
					other := otherE.components[j].(*RigidBody)
					othert := otherE.components[otherE.GetComponent("Transform")].(*Transform)
					if b.Rect.Intersects(other.Rect) && b.Type == 1 {
						b.Velocity = b.Velocity.Add(t.Position.DirectionTo(othert.Position).Roundf().Negative()) // Square Push back
						//b.Velocity = b.Velocity.Add(t.Position.DirectionTo(othert.Position).Negative()) // Push back
					}
				}
			}

			// APPLYING VELOCITY
			t.Position = t.Position.Add(b.Velocity)
		}
	}
}

type MoveSystem struct{}

func (s *MoveSystem) Start(g *Game) {}
func (s *MoveSystem) End(g *Game)   {}
func (s *MoveSystem) Update(g *Game, delta float64) {
	for _, e := range g.Entities {
		if i := e.GetComponent("Move"); i != -1 {
			m := e.components[i].(*Move)
			rb := e.components[e.GetComponent("RigidBody")].(*RigidBody)
			sprite := e.components[e.GetComponent("Sprite")].(*Sprite)

			direction := Vector2f{}
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
}
