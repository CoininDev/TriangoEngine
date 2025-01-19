package lib

type Entity struct {
	ID         int32
	Position   Vector2f
	Game       *Game
	components []Component
}

func NewEntity(id int32, game *Game, pos Vector2f, comps []Component) Entity {
	return Entity{id, pos, game, comps}
}

func (e *Entity) Start() {
	for _, comp := range e.components {
		comp.Start(e)
	}
}

func (e *Entity) Update() {
	for _, c := range e.components {
		if c.IsActive() {
			c.Tick()
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
