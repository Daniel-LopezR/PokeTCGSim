package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1920, 1080, "raylib [core] example - basic window")
	rl.SetWindowMonitor(0)

	texture := rl.LoadTexture("resources/erika_texture.png")
	//texture := rl.LoadTexture("resources/erika.png")
	card := rl.LoadModel("resources/Card.obj")
	rl.SetMaterialTexture(&card.GetMaterials()[0], rl.MapDiffuse, texture)
	defer rl.CloseWindow()
	defer rl.UnloadModel(card)
	defer rl.UnloadTexture(texture)

	camera := rl.Camera3D{}
	camera.Position = rl.Vector3{X: 0.0, Y: 4.0, Z: 10.0}
	camera.Target = rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0}
	camera.Up = rl.Vector3{X: 0.0, Y: 1.0, Z: 0.0}
	camera.Fovy = 45
	camera.Projection = rl.CameraPerspective

	//rl.DisableCursor()
	rl.SetTargetFPS(240)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFree)

		if rl.IsKeyPressed('Z') {
			camera.Target = rl.Vector3{X: 0, Y: 0, Z: 0}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.LightGray)
		rl.BeginMode3D(camera)
		rl.DrawModelEx(card, rl.Vector3{
			X: 0,
			Y: 0,
			Z: 0,
		}, rl.Vector3{
			X: 0,
			Y: 0,
			Z: 0,
		}, 0, rl.Vector3{
			X: 1,
			Y: 1,
			Z: 1,
		}, rl.White)

		rl.DrawGrid(10, 1.0)

		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
