package lib

type Component interface {
	Start(*Entity)
	Tick()
	End()
	GetType() string
}
