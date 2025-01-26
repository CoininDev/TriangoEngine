package comps

import (
	"triango/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	Active       bool
	Texture      rl.Texture2D
	Fliph        int
	Flipv        int
	TransformIdx int
}

func NewSprite(tex rl.Texture2D, fliph, flipv int) *Sprite {
	return &Sprite{true, tex, fliph, flipv, -1}
}

func (s *Sprite) GetType() string { return "Sprite" }

type PlaceholderSprite struct {
	Active       bool
	Size         utils.Vector2
	Color        rl.Color
	transformIdx int
}

func NewPlaceholderSprite(w, h float32, color rl.Color) *PlaceholderSprite {
	return &PlaceholderSprite{true, utils.NewVector2f(w, h), color, -1}
}
func (s *PlaceholderSprite) GetType() string { return "PlaceholderSprite" }
