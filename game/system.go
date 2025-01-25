package game

type System interface {
	Start(*Game)
	Update(*Game, float64)
	End(*Game)
	Refresh(*Game)
}
