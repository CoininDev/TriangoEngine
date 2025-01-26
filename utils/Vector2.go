package utils

import "math"

type Vector2 struct {
	X, Y float32
}

func NewVector2f(x, y float32) Vector2 {
	return Vector2{X: x, Y: y}
}

func NewRectf(position, size Vector2) Rect {
	return Rect{Position: position, Size: size}
}

func (v Vector2) Add(other Vector2) Vector2 {
	return NewVector2f(v.X+other.X, v.Y+other.Y)
}

func (v Vector2) Addf(f float32) Vector2 {
	return NewVector2f(v.X+f, v.Y+f)
}

func (v Vector2) Sub(other Vector2) Vector2 {
	return NewVector2f(v.X-other.X, v.Y-other.Y)
}

func (v Vector2) Subf(f float32) Vector2 {
	return NewVector2f(v.X-f, v.Y-f)
}

func (v Vector2) Mul(other Vector2) Vector2 {
	return NewVector2f(v.X*other.X, v.Y*other.Y)
}

func (v Vector2) Mulf(f float32) Vector2 {
	return NewVector2f(v.X*f, v.Y*f)
}

func (v Vector2) Div(other Vector2) Vector2 {
	return NewVector2f(v.X/other.X, v.Y/other.Y)
}

func (v Vector2) Divf(f float32) Vector2 {
	return NewVector2f(v.X/f, v.Y/f)
}

func (v Vector2) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v Vector2) Dot(other Vector2) float32 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vector2) Normalize() Vector2 {
	var magnitude = float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
	if magnitude == 0 {
		return Vector2{0, 0} // Avoid division by zero
	}
	return Vector2{
		X: float32(v.X) / magnitude,
		Y: float32(v.Y) / magnitude,
	}
}

func (v Vector2) Negative() Vector2 {
	return NewVector2f(-v.X, -v.Y)
}

func (v Vector2) DirectionTo(other Vector2) Vector2 {
	return other.Sub(v).Normalize()
}

func (v Vector2) DistanceTo(other Vector2) float32 {
	return float32(math.Sqrt(float64((v.X-other.X)*(v.X-other.X) + (v.Y-other.Y)*(v.Y-other.Y))))
}

func (v Vector2) Roundf() Vector2 {
	return NewVector2f(float32(math.Round(float64(v.X))), float32(math.Round(float64(v.Y))))
}
