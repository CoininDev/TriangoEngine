package lib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	Active       bool
	e            *Entity
	Texture      rl.Texture2D
	Fliph        int
	Flipv        int
	TransformIdx int
}

func (c *Sprite) IsActive() bool     { return c.Active }
func (c *Sprite) SetActive(new bool) { c.Active = new }

func NewSprite(tex rl.Texture2D, fliph, flipv int) *Sprite {
	return &Sprite{true, nil, tex, fliph, flipv, -1}
}
func (s *Sprite) Start(e *Entity)     {}
func (s *Sprite) Tick(_delta float64) {}
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

func NewPlaceholderSprite(w, h float32, color rl.Color) *PlaceholderSprite {
	return &PlaceholderSprite{true, nil, NewVector2f(w, h), color, -1}
}
func (c *PlaceholderSprite) IsActive() bool      { return c.Active }
func (c *PlaceholderSprite) SetActive(new bool)  { c.Active = new }
func (s *PlaceholderSprite) Start(e *Entity)     {}
func (s *PlaceholderSprite) Tick(_delta float64) {}
func (s *PlaceholderSprite) End()                {}
func (s *PlaceholderSprite) GetType() string     { return "PlaceholderSprite" }
