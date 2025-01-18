package lib

type Entity struct {
	ID         int32
	Position   Vector2
	Game       *Game
	components []Component
}

func NewEntity(id int32, game *Game, pos Vector2, comps []Component) Entity {
	return Entity{id, pos, game, comps}
}

func (e *Entity) Start() {
	for _, comp := range e.components {
		comp.Start(e)
	}
}

func (e *Entity) Update() {
	for _, c := range e.components {
		c.Tick()
	}
}

func (e *Entity) Die() {
	for _, c := range e.components {
		c.End()
	}
}

func (e *Entity) GetComponent(compType string) Component {
	for i := range e.components {
		if e.components[i].GetType() == compType {
			return e.components[i]
		}
	}
	return nil
}
