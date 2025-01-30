package game

import (
	"triango/utils"

	"github.com/kelindar/column"
)

type Scene struct {
	EntityCollection *column.Collection
	game             *Game
}

func NewScene(game *Game, col *column.Collection) *Scene {
	return &Scene{col, game}
}

func (s *Scene) Spawn(entity Entity) {
	idx, _ := s.EntityCollection.Insert(func(r column.Row) error {
		r.SetRecord("Transform", &utils.Vector2f{X: 0, Y: 0})
		return nil
	})
	s.game.RefreshActiveScene()
}
