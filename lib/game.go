package lib

import "time"

type Game struct {
	Entities       []Entity
	Servers        []Server
	LastUpdateTime time.Time
}

func (g *Game) StartGame() {
	g.LastUpdateTime = time.Now()
	for _, server := range g.Servers {
		server.Start(g)
	}
	for _, entity := range g.Entities {
		entity.Start()
	}
}
func (g *Game) Run() {
	now := time.Now()
	delta := now.Sub(g.LastUpdateTime).Seconds()
	g.LastUpdateTime = now
	for _, server := range g.Servers {
		server.Update(g, delta)
	}
	for _, entity := range g.Entities {
		entity.Update(delta)
	}
}
func (g *Game) End() {
	for _, server := range g.Servers {
		server.End(g)
	}
	for _, entity := range g.Entities {
		entity.Die()
	}
}
