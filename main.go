package main

import (
	"database/sql"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type PageType int

const (
	LoginPage = iota
	LandingPage
	PlayerPage
	TeamPage
	CoachPage
)

func main() {
	var pageType PageType = LoginPage
	var textBoxes []TextBox
	var buttons []Button
	var db *sql.DB
	var err error
	var notification Notification
	resolutions := []rl.Vector2{rl.Vector2{1280, 720}, rl.Vector2{1920, 1080}, rl.Vector2{2560, 1440}, rl.Vector2{3840, 2160}}
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(int32(resolutions[0].X), int32(resolutions[0].Y), "Absolute Football Matches - Data Hoarders")
	rl.SetTargetFPS(120)
	camera := rl.NewCamera2D(rl.Vector2{0, 0}, rl.Vector2{0, 0}, 0, 1.0)

	var renderSize = rl.Vector2{float32(rl.GetRenderWidth()), float32(rl.GetRenderHeight())}

	textBoxes, buttons = InitializeLoginPage(renderSize)
	fmt.Printf("%d, %d", rl.GetRenderWidth(), rl.GetRenderHeight())
	for !rl.WindowShouldClose() {
		for i := 0; i < len(textBoxes); i++ {
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				textBoxes[i].DetectActivation(rl.GetMousePosition())
			}
		}
		if rl.IsKeyPressed(rl.KeyTab) || rl.IsKeyPressed(rl.KeyEnter) {
			for i := 0; i < len(textBoxes); i++ {
				if textBoxes[i].active {
					textBoxes[i].active = false
					textBoxes[(i+1)%len(textBoxes)].active = true
					break
				}
			}
		}
		for i := 0; i < len(textBoxes); i++ {
			if textBoxes[i].active {
				textBoxes[i].DetectInput()
			}
		}
		switch pageType {
		case LoginPage:
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				buttons[0].DetectActivation(rl.GetMousePosition())
				if buttons[0].active {
					fmt.Printf("Huh\n")
					buttons[0].active = false
					db, err = sql.Open("mysql", textBoxes[2].text+":"+textBoxes[3].text+"@tcp("+textBoxes[0].text+":"+textBoxes[1].text+")/football_db")
					if err != nil {
						notification = *NewNotification("Error: Incorrect credentials. Try again.", 5000, rl.Red)
					} else {
						notification = *NewNotification("Success: Connection Successful!", 3000, rl.Green)
						db.SetConnMaxLifetime(time.Minute * 3)
						db.SetMaxOpenConns(10)
						db.SetMaxIdleConns(10)

						members, err := db.Query("SELECT member_id, first_name, last_name, experience FROM members")
						if err != nil {
							notification = *NewNotification("Error: Incorrect credentials. Try again.", 5000, rl.Red)
						}
						defer members.Close()

						if notification.color == rl.Red {

						} else {
							pageType = LandingPage
							// InitializeLandingPage()
						}
					}

				}
			}
		case LandingPage:

		case PlayerPage:

		case TeamPage:

		case CoachPage:

		}
		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		rl.ClearBackground(rl.RayWhite)
		for i := 0; i < len(textBoxes); i++ {
			textBoxes[i].Draw()
		}
		for i := 0; i < len(buttons); i++ {
			buttons[i].Draw()
		}
		notification.Draw()
		rl.EndMode2D()
		rl.EndDrawing()
	}
	if err != nil {
		panic(err)
	}
}
