package main

import (
	"triango/lib"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "PintoGames")
	defer rl.CloseWindow()
	//rl.SetTargetFPS(60)

	game := lib.Game{}
	incorporate(&game)

	game.StartGame()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		game.Run()

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
	game.End()
}

func incorporate(game *lib.Game) {
	game.Systems = append(game.Systems,
		&lib.MoveSystem{},
		&lib.PhysicsSystem{},
		&lib.RenderSystem{},
	)

	game.Entities = append(game.Entities,
		// chão
		lib.NewEntity(game, []lib.Component{
			lib.NewTransform(lib.NewVector2f(0, 500), lib.NewVector2f(1, 1), 0),
			lib.NewRigidBody(0, 500, 100),
			lib.NewPlaceholderSprite(500, 100, rl.Blue),
		}),

		// bloco central
		lib.NewEntity(game, []lib.Component{
			lib.NewTransform(lib.NewVector2f(400, 300), lib.NewVector2f(1, 1), 0),
			lib.NewRigidBody(0, 100, 100),
			lib.NewPlaceholderSprite(100, 100, rl.Red),
		}),

		//shrek
		lib.NewEntity(game, []lib.Component{
			lib.NewTransform(lib.NewVector2f(0, 0), lib.NewVector2f(0.5, 0.5), 0),
			lib.NewRigidBody(1, 256/2, 256/2),
			lib.NewSprite(rl.LoadTexture("src/shrek.jpg"), 1, 1),
			lib.NewMove(300),
		}),
	)
}
