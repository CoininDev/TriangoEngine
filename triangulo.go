package main

import (
	"triangulo/lib"

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
		rl.ClearBackground(rl.LightGray)

		game.Run()

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
	game.End()
}

func incorporate(game *lib.Game) {
	//chão
	game.Entities = append(game.Entities, lib.NewEntity(0, game, lib.Vector2{0, 500}, []lib.Component{
		lib.NewPlaceholderSprite(lib.Vector2{X: 800, Y: 100}, rl.Blue),
		lib.NewPhysicalBody(true, lib.Vector2f{X: 800, Y: 100}),
	}))

	//personagem
	game.Entities = append(game.Entities, lib.NewEntity(1, game, lib.Vector2{400, 100}, []lib.Component{
		lib.NewPlaceholderSprite(lib.Vector2{X: 20, Y: 20}, rl.Brown),
		lib.NewPhysicalBody(true, lib.Vector2f{X: 20, Y: 20}),
		lib.NewGravity(lib.Vector2{X: 0, Y: 2}),
		lib.NewPlatformMove(5, 20),
	}))
}
