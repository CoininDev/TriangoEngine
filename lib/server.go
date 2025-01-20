package lib

import rl "github.com/gen2brain/raylib-go/raylib"

// Um server é uma forma de criar uma lógica universal, e não individual como são os componentes
type Server interface {
	Start(*Game)
	Update(*Game, float64)
	End(*Game)
}

type RenderServer struct{}

func (s RenderServer) Start(g *Game) {}
func (s RenderServer) End(g *Game) {
	for _, e := range g.Entities {
		if i := e.GetComponent("Sprite"); i != -1 {
			s := e.components[i].(*Sprite)
			rl.UnloadTexture(s.Texture)
		}
	}
}
func (s RenderServer) Update(g *Game, delta float64) {
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

type PhysicsServer struct{}

func (s PhysicsServer) Start(g *Game) {}
func (s PhysicsServer) End(g *Game)   {}
func (s PhysicsServer) Update(g *Game, delta float64) {
	for _, e := range g.Entities {
		if i := e.GetComponent("CollisionRect"); i != -1 {
			c := e.components[i].(*CollisionRect)
			t := e.components[e.GetComponent("Transform")].(*Transform)
			c.rect.Position = t.Position

			// CHECKING COLLISION OF COLLISIONRECTS
			c.Top = false
			c.Bottom = false
			c.Left = false
			c.Right = false
			for _, otherE := range g.Entities {
				if otherE.ID != e.ID {
					if j := otherE.GetComponent("CollisionRect"); j != -1 {
						other := otherE.components[j].(*CollisionRect)
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

			// MOVING PLAYER
			if j := e.GetComponent("Move"); j != -1 {
				m := e.components[j].(*Move)
				t := e.components[e.GetComponent("Transform")].(*Transform)

				sprite := e.components[e.GetComponent("Sprite")].(*Sprite)
				col := e.components[e.GetComponent("CollisionRect")].(*CollisionRect)

				direction := Vector2f{}
				if rl.IsKeyDown(rl.KeyW) && !col.Top {
					direction.Y -= 1
				}
				if rl.IsKeyDown(rl.KeyS) && !col.Bottom {
					direction.Y += 1
				}
				if rl.IsKeyDown(rl.KeyA) && !col.Left {
					direction.X -= 1
				}
				if rl.IsKeyDown(rl.KeyD) && !col.Right {
					direction.X += 1
				}

				if direction.X != 0 {
					sprite.Fliph = int(direction.X)
				}
				m.Velocity.X = direction.Normalize().X * m.Speed * float32(delta)
				m.Velocity.Y = direction.Normalize().Y * m.Speed * float32(delta)
				t.Position.X += m.Velocity.X
				t.Position.Y += m.Velocity.Y
			}
		}
	}
}
