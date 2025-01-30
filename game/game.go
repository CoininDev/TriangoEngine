package game

import "time"

type Game struct {
	Scenes         []Scene
	ActiveScene    int
	GlobalScene    Scene
	Systems        []System
	LastUpdateTime time.Time
}

func (g *Game) StartGame() {
	g.LastUpdateTime = time.Now()
	g.ActiveScene = 0
	for _, system := range g.Systems {
		system.Refresh(g)
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
func (g *Game) ChangeScene(idx int) bool {
	if idx < len(g.Scenes) {
		g.ActiveScene = idx
		return true
	}
	return false
}
func (g *Game) RefreshActiveScene() {

}
