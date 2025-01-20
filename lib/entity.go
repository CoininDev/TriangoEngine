package lib

type Entity struct {
	ID         int32
	Game       *Game
	components []Component
}

func NewEntity(id int32, game *Game, comps []Component) Entity {
	return Entity{id, game, comps}
}

func (e *Entity) GetComponent(compType string) int {
	for i, c := range e.components {
		if c.GetType() == compType {
			return i
		}
	}
	return -1
}
