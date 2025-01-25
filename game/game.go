package game

import "time"

type Game struct {
	Entities       []Entity
	Systems        []System
	LastUpdateTime time.Time
}

func (g *Game) StartGame() {
	g.LastUpdateTime = time.Now()
	for idx := range g.Entities {
		g.Entities[idx].ID = int32(idx)
	}
	for _, system := range g.Systems {
		system.Start(g)
	}
}
func (g *Game) Run() {
	now := time.Now()
	delta := now.Sub(g.LastUpdateTime).Seconds()
	g.LastUpdateTime = now
	for _, system := range g.Systems {
		system.Update(g, delta)
	}
}
func (g *Game) End() {
	for _, system := range g.Systems {
		system.End(g)
	}
}
