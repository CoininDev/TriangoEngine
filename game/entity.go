package game

type Entity struct {
	ID         int32
	Game       *Game
	Components []Component
}

func NewEntity(game *Game, comps []Component) Entity {
	return Entity{-1, game, comps}
}

func (e *Entity) GetComponent(compType string) int {
	for i, c := range e.Components {
		if c.GetType() == compType {
			return i
		}
	}
	return -1
}
