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

	// game.Entities = append(game.Entities, lib.NewEntity(0, &game, lib.Vector2{400, 300}, []lib.Component{
	// 	lib.NewPlaceholderSprite(lib.Vector2{X: 20, Y: 20}, rl.Brown),
	// 	lib.NewMove(6),
	// }))

	//chão
	game.Entities = append(game.Entities, lib.NewEntity(0, &game, lib.Vector2{X: 0, Y: 500}, []lib.Component{
		lib.NewPlaceholderSprite(lib.Vector2{X: 800, Y: 100}, rl.Blue),
		lib.NewPhysicalBody(lib.Vector2{X: 800, Y: 100}),
	}))

	//personagem
	game.Entities = append(game.Entities, lib.NewEntity(1, &game, lib.Vector2{X: 400, Y: 300}, []lib.Component{
		lib.NewSprite(rl.LoadTexture("src/shrek.jpg")),
		//lib.NewPlaceholderSprite(lib.Vector2{X: 20, Y: 20}, rl.Beige),
		lib.NewPhysicalBody(lib.Vector2{X: 256, Y: 256}),
		//lib.NewGravity(lib.Vector2{X: 0, Y: 2}),
		lib.NewMove(5),
	}))

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
