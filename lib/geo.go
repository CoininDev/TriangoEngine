package lib

import "math"

type Vector2 struct {
	X, Y int32
}

type Vector2f struct {
	X, Y float64
}

func (v Vector2f) Normalize() Vector2f {
	magnitude := math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
	if magnitude == 0 {
		return Vector2f{0, 0} // Avoid division by zero
	}
	return Vector2f{
		X: float64(v.X) / magnitude,
		Y: float64(v.Y) / magnitude,
	}
}

type Rect struct {
	Position Vector2
	Size     Vector2
}

func (r Rect) GetCenter() Vector2 {
	return Vector2{r.Position.X - r.Size.X/2, r.Position.Y - r.Size.Y/2}
}
