package lib

type Component interface {
	Start(*Entity)
	Tick(float64)
	End()
	GetType() string
	IsActive() bool
	SetActive(bool)
}
