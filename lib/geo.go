package lib

import "math"

type Vector2 struct {
	X, Y int32
}

type Vector2f struct {
	X, Y float32
}

func NewVector2(x, y int32) Vector2 {
	return Vector2{X: x, Y: y}
}

func NewVector2f(x, y float32) Vector2f {
	return Vector2f{X: x, Y: y}
}

func NewRectf(position, size Vector2f) Rectf {
	return Rectf{Position: position, Size: size}
}

func (v Vector2f) Normalize() Vector2f {
	var magnitude = float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
	if magnitude == 0 {
		return Vector2f{0, 0} // Avoid division by zero
	}
	return Vector2f{
		X: float32(v.X) / magnitude,
		Y: float32(v.Y) / magnitude,
	}
}

func (v Vector2f) Negative() Vector2f {
	return NewVector2f(-v.X, -v.Y)
}

type Rectf struct {
	Position Vector2f
	Size     Vector2f
}

func (r Rectf) GetCenter() Vector2f {
	return Vector2f{r.Position.X - r.Size.X/2, r.Position.Y - r.Size.Y/2}
}
