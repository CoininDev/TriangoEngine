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
}
func (g *Game) Run() {
	now := time.Now()
	delta := now.Sub(g.LastUpdateTime).Seconds()
	g.LastUpdateTime = now
	for _, server := range g.Servers {
		server.Update(g, delta)
	}
}
func (g *Game) End() {
	for _, server := range g.Servers {
		server.End(g)
	}
}
