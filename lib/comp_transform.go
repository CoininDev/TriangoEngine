package lib

type Transform struct {
	Active   bool
	Position Vector2f
	Scale    Vector2f
	Rotation float32
	e        *Entity
}

func NewTransform(position, scale Vector2f, rotation float32) *Transform {
	return &Transform{
		Active:   true,
		Position: position,
		Scale:    scale,
		Rotation: rotation,
	}
}

func (t *Transform) IsActive() bool     { return t.Active }
func (t *Transform) SetActive(new bool) { t.Active = new }
func (t *Transform) Start(e *Entity)    { t.e = e }
func (t *Transform) Tick()              {}
func (t *Transform) End()               {}
func (t *Transform) GetType() string    { return "Transform" }
