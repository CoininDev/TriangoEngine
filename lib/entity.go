package lib

type Entity struct {
	ID         int32
	Game       *Game
	components []Component
}

func NewEntity(id int32, game *Game, comps []Component) Entity {
	return Entity{id, game, comps}
}

func (e *Entity) Start() {
	for _, comp := range e.components {
		comp.Start(e)
	}
}

func (e *Entity) Update(delta float64) {
	for _, c := range e.components {
		if c.IsActive() {
			c.Tick(delta)
		}
	}
}

func (e *Entity) Die() {
	for _, c := range e.components {
		c.End()
	}
}

func (e *Entity) GetComponent(compType string) (int, Component) {
	for i, c := range e.components {
		if c.GetType() == compType {
			return i, c
		}
	}
	return -1, nil
}
