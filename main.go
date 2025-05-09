package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1280, 720, "raylib [core] example - basic window")
	rl.SetWindowMonitor(0)

	texture := rl.LoadTexture("resources/erika_texture.png")
	//texture := rl.LoadTexture("resources/erika.png")
	card := rl.LoadModel("resources/Card.obj")
	rl.SetMaterialTexture(&card.GetMaterials()[0], rl.MapDiffuse, texture)
	defer rl.CloseWindow()
	defer rl.UnloadModel(card)
	defer rl.UnloadTexture(texture)

	camera := rl.Camera3D{}
	camera.Position = rl.Vector3{X: 0.0, Y: 1.0, Z: 4.0}
	camera.Target = rl.Vector3{X: 0.0, Y: 0.5, Z: 0.0}
	camera.Up = rl.Vector3{X: 0.0, Y: 1, Z: 0.0}
	camera.Fovy = 50
	camera.Projection = rl.CameraPerspective

	//rl.DisableCursor()
	rl.SetTargetFPS(60)

	dragging := false
	returning := false
	initialDraggingPos := rl.Vector2{}
	currentDraggingPos := rl.Vector2{}
	yaw := float32(0)
	roll := float32(0)

	for !rl.WindowShouldClose() {
		//rl.UpdateCamera(&camera, rl.CameraFree)

		// if rl.IsKeyPressed('Z') {
		// 	camera.Target = rl.Vector3{X: 0, Y: 0, Z: 0}
		// }
		if rl.IsMouseButtonDown(rl.MouseLeftButton) && !returning {
			if !dragging {
				initialDraggingPos = rl.GetMousePosition()
				dragging = true
				returning = false
			} else {
				currentDraggingPos = rl.GetMousePosition()

				// This ugly code is to avoid flickering when the card is returning to 0,0,0 caused by float decimals
				yaw = float32(int32((-(currentDraggingPos.X - initialDraggingPos.X)) / 10))
				roll = float32(int32((currentDraggingPos.Y - initialDraggingPos.Y) / 10))

			}
		} else {
			if yaw > 0.0 {
				yaw -= 1
			} else if yaw < 0.0 {
				yaw += 1
			}
			if roll > 0.0 {
				roll -= 1
			} else if roll < 0.0 {
				roll += 1
			}
			if yaw == 0.0 && roll == 0.0 {
				returning = false
			}
		}
		if rl.IsMouseButtonReleased(rl.MouseLeftButton) && dragging {
			dragging = false
			returning = true
			initialDraggingPos = rl.Vector2{}
			currentDraggingPos = rl.Vector2{}
		}

		card.Transform = rl.MatrixRotateXYZ(rl.Vector3{Y: rl.Deg2rad * yaw, Z: rl.Deg2rad * roll})

		rl.BeginDrawing()
		rl.ClearBackground(rl.LightGray)
		rl.BeginMode3D(camera)
		// To rotate model I have to change rotation axis and rotation angle
		rl.DrawModelEx(card, rl.Vector3{
			X: 0,
			Y: 0,
			Z: 0,
		}, rl.Vector3{
			X: 0,
			Y: 1,
			Z: 0,
		}, -90, rl.Vector3{
			X: 2,
			Y: 2,
			Z: 2,
		}, rl.White)

		//rl.DrawGrid(10, 1.0)

		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		mouseDragging := fmt.Sprintf("Mouse Dragging (%0.2f, %0.2f)",
			currentDraggingPos.X-initialDraggingPos.X,
			currentDraggingPos.Y-initialDraggingPos.Y)
		rl.DrawText(mouseDragging, 10, 40, 20, rl.DarkGreen)
		rl.EndDrawing()
	}
}
