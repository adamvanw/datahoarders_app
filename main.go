package main

import (
	"database/sql"
	rl "github.com/gen2brain/raylib-go/raylib"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type PageType int

const (
	LoginPage = iota
	LandingPage
	PlayersPage
	TeamsPage
	CoachesPage
	GamesPage
	SearchPage
	GamePage
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
	var navBar = NewNavBar()
	var scrollTables []ScrollTable
	var texts []Text
	// var timer int32 = 0
	var notification Notification
	resolutions := []rl.Vector2{rl.Vector2{1280, 720}, rl.Vector2{1920, 1080}, rl.Vector2{2560, 1440}, rl.Vector2{3840, 2160}}
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(int32(resolutions[0].X), int32(resolutions[0].Y), "Absolute Football Matches - Data Hoarders")
	rl.SetTargetFPS(120)
	camera := rl.NewCamera2D(rl.Vector2{0, 0}, rl.Vector2{0, 0}, 0, 1.0)

	var renderSize = rl.Vector2{float32(1280), float32(720)}

	textBoxes, buttons = InitializeLoginPage(renderSize)
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
		if pageType != LoginPage {
			var selection = navBar.DetectInput()
			switch selection {
			case 0:
				pageType = LandingPage
				textBoxes, buttons = InitializeLandingPage()
				scrollTables = []ScrollTable{}
			case 1:
				pageType = PlayersPage
				scrollTables = InitializePlayersPage(db)
				buttons = []Button{}
				textBoxes = []TextBox{}
			case 2:
				pageType = CoachesPage
				scrollTables = InitializeCoachesPage(db)
				buttons = []Button{}
				textBoxes = []TextBox{}
			case 3:
				pageType = TeamsPage
				scrollTables = InitializeTeamsPage(db)
				buttons = []Button{}
				textBoxes = []TextBox{}
			case 4:
				pageType = GamesPage
				scrollTables = InitializeGamesPage(db)
				buttons = []Button{}
				textBoxes = []TextBox{}
			}
		}
		var elementId = 0
		var elementPage = 0
		switch pageType {
		case LoginPage:
			buttons[0].DetectActivation(rl.GetMousePosition())
			if buttons[0].active {
				buttons[0].active = false
				db, err = sql.Open("mysql", textBoxes[2].text+":"+textBoxes[3].text+"@tcp("+textBoxes[0].text+":"+textBoxes[1].text+")/football_db")
				if err != nil {
					notification = *NewNotification("Error: Incorrect credentials. Try again.", 5000, rl.Red)
				} else {
					notification = *NewNotification("Success: Connection Successful!", 3000, rl.Green)
					db.SetConnMaxLifetime(time.Minute * 3)
					db.SetMaxOpenConns(10)
					db.SetMaxIdleConns(10)

					// this just tests if we can query the database with our intended attributes.
					rows, err := db.Query("SELECT member_id, first_name, last_name, experience FROM members")
					if err != nil {
						notification = *NewNotification("Error: Incorrect credentials. Try again.", 5000, rl.Red)
					}
					rows.Close()

					if notification.color != rl.Red {
						pageType = LandingPage
						textBoxes, buttons = InitializeLandingPage()
					}
				}
			}
			break
		case LandingPage:
			for i := 0; i < len(buttons); i++ {
				buttons[i].DetectActivation(rl.GetMousePosition())
				if buttons[i].active {

				}
			}
		case SearchPage:
			elementId, elementPage = scrollTables[0].DetectInput(db)
		case PlayersPage:
			elementId, elementPage = scrollTables[0].DetectInput(db)
			if elementId != -1 {
				pageType = PageType(elementPage)
				scrollTables, texts = InitializePlayerPage(db, elementId)
			}
		case TeamsPage:
			elementId, elementPage = scrollTables[0].DetectInput(db)
			if elementId != -1 {
				pageType = PageType(elementPage)
			}
		case CoachesPage:
			elementId, elementPage = scrollTables[0].DetectInput(db)
			if elementId != -1 {
				pageType = PageType(elementPage)
				scrollTables, texts = InitializeCoachPage(db, elementId)
			}
		case GamesPage:
			elementId, elementPage = scrollTables[0].DetectInput(db)
			if elementId != -1 {
				pageType = PageType(elementPage)
				scrollTables, texts = InitializeGamePage(db, elementId)
			}
		case PlayerPage:
			for i := 0; i < len(scrollTables); i++ {
				elementId, elementPage = scrollTables[i].DetectInput(db)
				if elementId != -1 || elementPage != -1 {
					pageType = PageType(elementPage)
					switch pageType {
					case CoachPage:
						scrollTables, texts = InitializeCoachPage(db, elementId)
					}
				}
			}
		case GamePage:
			for i := 0; i < len(scrollTables); i++ {
				elementId, elementPage = scrollTables[i].DetectInput(db)
				if elementId != -1 || elementPage != -1 {
					pageType = PageType(elementPage)
					// it's only likely to be players, so InitializePlayerPage.
					scrollTables, texts = InitializePlayerPage(db, elementId)
				}
			}
		case CoachPage:
			for i := 0; i < len(scrollTables); i++ {
				elementId, elementPage = scrollTables[i].DetectInput(db)
				if elementId != -1 || elementPage != -1 {
					pageType = PageType(elementPage)
					scrollTables, texts = InitializeTeamPage(db, elementId)
				}
			}
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
		if pageType != LoginPage {
			navBar.Draw()
			if pageType != GamesPage && pageType != TeamsPage && pageType != PlayersPage && pageType != CoachesPage {
				for i := 0; i < len(texts); i++ {
					texts[i].Draw()
				}
			}
		}
		for i := 0; i < len(scrollTables); i++ {
			scrollTables[i].Render()
		}
		notification.Draw()
		rl.EndMode2D()
		rl.EndDrawing()
	}
	if err != nil {
		panic(err)
	}
}
