package lib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	Active       bool
	e            *Entity
	Texture      rl.Texture2D
	fliph        int
	flipv        int
	transformIdx int
}

func (c *Sprite) IsActive() bool     { return c.Active }
func (c *Sprite) SetActive(new bool) { c.Active = new }

func NewSprite(tex rl.Texture2D, fliph, flipv int) *Sprite {
	return &Sprite{true, nil, tex, fliph, flipv, -1}
}
func (s *Sprite) Start(e *Entity) {
	s.e = e
	s.transformIdx, _ = s.e.GetComponent("Transform")
}
func (s *Sprite) Tick() {
	t := s.e.components[s.transformIdx].(*Transform)
	rl.DrawTexturePro(
		s.Texture,
		rl.NewRectangle(0, 0, float32(s.Texture.Width*int32(s.fliph)), float32(s.Texture.Height*int32(s.flipv))),
		rl.NewRectangle(0, 0, float32(s.Texture.Width)*t.Scale.X, float32(s.Texture.Height)*t.Scale.Y),
		rl.NewVector2(t.Position.X, t.Position.Y),
		t.Rotation,
		rl.White,
	)
}
func (s *Sprite) End() {
	rl.UnloadTexture(s.Texture)
}
func (s *Sprite) GetType() string { return "Sprite" }

type PlaceholderSprite struct {
	Active       bool
	e            *Entity
	Size         Vector2f
	Color        rl.Color
	transformIdx int
}

func NewPlaceholderSprite(size Vector2f, color rl.Color) *PlaceholderSprite {
	return &PlaceholderSprite{true, nil, size, color, -1}
}

func (c *PlaceholderSprite) IsActive() bool     { return c.Active }
func (c *PlaceholderSprite) SetActive(new bool) { c.Active = new }

func (s *PlaceholderSprite) Start(e *Entity) {
	s.e = e
	s.transformIdx, _ = s.e.GetComponent("Transform")
}
func (s *PlaceholderSprite) Tick() {
	t := s.e.components[s.transformIdx].(*Transform)
	rl.DrawRectanglePro(
		rl.NewRectangle(0, 0, s.Size.X, s.Size.Y),
		rl.NewVector2(t.Position.X, t.Position.Y),
		t.Rotation,
		s.Color,
	)
}
func (s *PlaceholderSprite) End()            {}
func (s *PlaceholderSprite) GetType() string { return "PlaceholderSprite" }
