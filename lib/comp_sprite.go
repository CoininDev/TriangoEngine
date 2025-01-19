package lib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	Active  bool
	e       *Entity
	Texture rl.Texture2D
	fliph   int
	flipv   int
	scale   float32
}

func (c *Sprite) IsActive() bool     { return c.Active }
func (c *Sprite) SetActive(new bool) { c.Active = new }

func NewSprite(tex rl.Texture2D, fliph, flipv int, scale float32) *Sprite {
	return &Sprite{true, nil, tex, fliph, flipv, scale}
}
func (s *Sprite) Start(e *Entity) { s.e = e }
func (s *Sprite) Tick() {
	rl.DrawTexturePro(
		s.Texture,
		rl.NewRectangle(0, 0, float32(s.Texture.Width*int32(s.fliph)), float32(s.Texture.Height*int32(s.flipv))),
		rl.NewRectangle(0, 0, float32(s.Texture.Width*int32(s.scale)), float32(s.Texture.Height*int32(s.scale))),
		rl.NewVector2(float32(s.e.Position.X), float32(s.e.Position.Y)),
		0,
		rl.White,
	)

}
func (s *Sprite) End() {
	rl.UnloadTexture(s.Texture)
}
func (s *Sprite) GetType() string { return "Sprite" }

type PlaceholderSprite struct {
	Active bool
	e      *Entity
	Size   Vector2f
	Color  rl.Color
}

func NewPlaceholderSprite(size Vector2f, color rl.Color) *PlaceholderSprite {
	return &PlaceholderSprite{true, nil, size, color}
}

func (c *PlaceholderSprite) IsActive() bool     { return c.Active }
func (c *PlaceholderSprite) SetActive(new bool) { c.Active = new }

func (s *PlaceholderSprite) Start(e *Entity) { s.e = e }
func (s *PlaceholderSprite) Tick() {
	rl.DrawRectanglePro(
		rl.NewRectangle(0, 0, s.Size.X, s.Size.Y),
		rl.NewVector2(s.e.Position.X, s.e.Position.Y),
		0,
		s.Color,
	)
}
func (s *PlaceholderSprite) End()            {}
func (s *PlaceholderSprite) GetType() string { return "PlaceholderSprite" }
