package lib

import "time"

type Game struct {
	Entities       []Entity
	LastUpdateTime time.Time
}

func (g *Game) StartGame() {
	g.LastUpdateTime = time.Now()
	for _, entity := range g.Entities {
		entity.Start()
	}
}
func (g *Game) Run() {
	now := time.Now()
	delta := now.Sub(g.LastUpdateTime).Seconds()
	g.LastUpdateTime = now
	for _, entity := range g.Entities {
		entity.Update(delta)
	}
}
func (g *Game) End() {
	for _, entity := range g.Entities {
		entity.Die()
	}
}
