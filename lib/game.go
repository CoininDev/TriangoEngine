package lib

type Game struct {
	Entities []Entity
}

func (g *Game) StartGame() {
	for _, entity := range g.Entities {
		entity.Start()
	}
}
func (g *Game) Run() {
	for _, entity := range g.Entities {
		entity.Update()
	}
}
func (g *Game) End() {
	for _, entity := range g.Entities {
		entity.Die()
	}
}
