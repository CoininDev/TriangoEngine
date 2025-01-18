package lib

type Component interface {
	Start(*Entity)
	Tick()
	End()
	GetType() string
	IsActive() bool
	SetActive(bool)
}
