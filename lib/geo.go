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

func (v Vector2f) Add(other Vector2f) Vector2f {
	return NewVector2f(v.X+other.X, v.Y+other.Y)
}

func (v Vector2f) Addf(f float32) Vector2f {
	return NewVector2f(v.X+f, v.Y+f)
}

func (v Vector2f) Sub(other Vector2f) Vector2f {
	return NewVector2f(v.X-other.X, v.Y-other.Y)
}

func (v Vector2f) Subf(f float32) Vector2f {
	return NewVector2f(v.X-f, v.Y-f)
}

func (v Vector2f) Mul(other Vector2f) Vector2f {
	return NewVector2f(v.X*other.X, v.Y*other.Y)
}

func (v Vector2f) Mulf(f float32) Vector2f {
	return NewVector2f(v.X*f, v.Y*f)
}

func (v Vector2f) Div(other Vector2f) Vector2f {
	return NewVector2f(v.X/other.X, v.Y/other.Y)
}

func (v Vector2f) Divf(f float32) Vector2f {
	return NewVector2f(v.X/f, v.Y/f)
}

func (v Vector2f) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v Vector2f) Dot(other Vector2f) float32 {
	return v.X*other.X + v.Y*other.Y
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

func (v Vector2f) DirectionTo(other Vector2f) Vector2f {
	return other.Sub(v).Normalize()
}

func (v Vector2f) DistanceTo(other Vector2f) float32 {
	return float32(math.Sqrt(float64((v.X-other.X)*(v.X-other.X) + (v.Y-other.Y)*(v.Y-other.Y))))
}

func (v Vector2f) Round() Vector2 {
	return NewVector2(int32(math.Round(float64(v.X))), int32(math.Round(float64(v.Y))))
}

func (v Vector2f) Roundf() Vector2f {
	return NewVector2f(float32(math.Round(float64(v.X))), float32(math.Round(float64(v.Y))))
}

type Rectf struct {
	Position Vector2f
	Size     Vector2f
}

func (r Rectf) GetCenter() Vector2f {
	return Vector2f{r.Position.X - r.Size.X/2, r.Position.Y - r.Size.Y/2}
}

func (r Rectf) Intersects(other Rectf) bool {
	return r.Position.X < other.Position.X+other.Size.X &&
		r.Position.X+r.Size.X > other.Position.X &&
		r.Position.Y < other.Position.Y+other.Size.Y &&
		r.Position.Y+r.Size.Y > other.Position.Y
}
