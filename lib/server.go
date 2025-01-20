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
func (s RenderServer) End(g *Game)   {}
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
