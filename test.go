package main

import (
	"triango/lib"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "PintoGames")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

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
	game.Entities = append(game.Entities, lib.NewEntity(0, game, []lib.Component{
		lib.NewTransform(lib.NewVector2f(0, 0), lib.NewVector2f(1, 0.5), 0),
		lib.NewSprite(rl.LoadTexture("src/shrek.jpg"), 1, 1),
		lib.NewMove(7),
	}))
}
