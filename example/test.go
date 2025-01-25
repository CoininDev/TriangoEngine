package main

import (
	"triango/comps"
	"triango/game"
	"triango/systems"
	"triango/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "PintoGames")
	defer rl.CloseWindow()
	//rl.SetTargetFPS(60)

	g := game.Game{}
	incorporate(&g)

	g.StartGame()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		g.Run()

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
	g.End()
}

func incorporate(g *game.Game) {
	g.Systems = append(g.Systems,
		&systems.MoveSystem{},
		&systems.PhysicsSystem{},
		&systems.RenderSystem{},
	)

	g.Entities = append(g.Entities,
		// chão
		game.NewEntity(g, []game.Component{
			comps.NewTransform(utils.NewVector2f(0, 500), utils.NewVector2f(1, 1), 0),
			comps.NewRigidBody(0, 500, 100),
			comps.NewPlaceholderSprite(500, 100, rl.Blue),
		}),

		// bloco central
		game.NewEntity(g, []game.Component{
			comps.NewTransform(utils.NewVector2f(400, 300), utils.NewVector2f(1, 1), 0),
			comps.NewRigidBody(0, 100, 100),
			comps.NewPlaceholderSprite(100, 100, rl.Red),
		}),

		//shrek
		game.NewEntity(g, []game.Component{
			comps.NewTransform(utils.NewVector2f(0, 0), utils.NewVector2f(0.5, 0.5), 0),
			comps.NewRigidBody(1, 256/2, 256/2),
			comps.NewSprite(rl.LoadTexture("src/shrek.jpg"), 1, 1),
			comps.NewMove(300),
		}),
	)
}
