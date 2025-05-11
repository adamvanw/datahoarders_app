package main

import (
	"database/sql"
	rl "github.com/gen2brain/raylib-go/raylib"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
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
	SearchResultsPage
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
	var timer int32 = 1000
	var notification Notification
	resolutions := []rl.Vector2{rl.Vector2{1280, 720}, rl.Vector2{1920, 1080}, rl.Vector2{2560, 1440}, rl.Vector2{3840, 2160}}
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(int32(resolutions[0].X), int32(resolutions[0].Y), "Absolute Football Matches - Data Hoarders")
	rl.SetTargetFPS(120)
	camera := rl.NewCamera2D(rl.Vector2{0, 0}, rl.Vector2{0, 0}, 0, 1.0)
	var prevPage PageType = LoginPage

	var renderSize = rl.Vector2{float32(1280), float32(720)}

	textBoxes, buttons = InitializeLoginPage(renderSize)
	for !rl.WindowShouldClose() {
		if prevPage != pageType {
			timer = 0
			prevPage = pageType
		}
		prevPage = pageType
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
				texts = []Text{}
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
		case GamePage:
			for i := 0; i < len(scrollTables); i++ {
				elementId, elementPage = scrollTables[i].DetectInput(timer)
				if elementId != -1 {
					pageType = PageType(elementPage)
					if elementPage == int(PlayerPage) {
						scrollTables, texts = InitializePlayerPage(db, elementId)
					}
				}
			}
			break
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
					scrollTables, texts = InitializeResultsPage(db, textBoxes[0].text)
					textBoxes = []TextBox{}
					buttons = []Button{}
					pageType = SearchResultsPage
					break
				}
			}
			break
		case SearchPage:
			elementId, elementPage = scrollTables[0].DetectInput(timer)
			break
		case PlayersPage:
			elementId, elementPage = scrollTables[0].DetectInput(timer)
			if elementId != -1 {
				pageType = PageType(elementPage)
				scrollTables, texts = InitializePlayerPage(db, elementId)
			}
			break
		case TeamsPage:
			elementId, elementPage = scrollTables[0].DetectInput(timer)
			if elementId != -1 {
				pageType = PageType(elementPage)
				scrollTables, texts = InitializeTeamPage(db, elementId)
			}
			break
		case CoachesPage:
			elementId, elementPage = scrollTables[0].DetectInput(timer)
			if elementId != -1 {
				pageType = PageType(elementPage)
				scrollTables, texts = InitializeCoachPage(db, elementId)
			}
			break
		case GamesPage:
			elementId, elementPage = scrollTables[0].DetectInput(timer)
			if elementId != -1 {
				pageType = PageType(elementPage)
				scrollTables, texts = InitializeGamePage(db, elementId)
			}
			break
		case PlayerPage:
			for i := 0; i < len(scrollTables); i++ {
				elementId, elementPage = scrollTables[i].DetectInput(timer)
				if elementId != -1 || elementPage != -1 {
					pageType = PageType(elementPage)
					switch pageType {
					case CoachPage:
						scrollTables, texts = InitializeCoachPage(db, elementId)
						break
					case TeamPage:
						scrollTables, texts = InitializeTeamPage(db, elementId)
						break
					}
					break
				}
			}
			break
		case SearchResultsPage:
			if len(scrollTables) == 1 {
				elementId, elementPage = scrollTables[0].DetectInput(timer)
				if elementId != -1 {
					pageType = PageType(elementPage)
					switch pageType {
					case CoachPage:
						scrollTables, texts = InitializeCoachPage(db, elementId)
						break
					case TeamPage:
						scrollTables, texts = InitializeTeamPage(db, elementId)
						break
					case PlayerPage:
						scrollTables, texts = InitializePlayerPage(db, elementId)
						break
					case GamePage:
						scrollTables, texts = InitializeGamePage(db, elementId)
						break
					}
					break
				}
			}
			break
		case CoachPage:
			for i := 0; i < len(scrollTables); i++ {
				elementId, elementPage = scrollTables[i].DetectInput(timer)
				if elementId != -1 || elementPage != -1 {
					pageType = PageType(elementPage)
					scrollTables, texts = InitializeTeamPage(db, elementId)
				}
			}
			break
		case TeamPage:
			for i := 0; i < len(scrollTables); i++ {
				elementId, elementPage = scrollTables[i].DetectInput(timer)
				if elementId != -1 || elementPage != -1 {
					pageType = PageType(elementPage)
					switch pageType {
					case CoachPage:
						scrollTables, texts = InitializeCoachPage(db, elementId)
						break
					case GamePage:
						scrollTables, texts = InitializeGamePage(db, elementId)
						break
					}
					break
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
		timer += int32(rl.GetFrameTime() * 1000)
		if timer < 250 {
			var opacity = float32(1 - float32(timer)/250)
			rl.DrawRectangleRec(rl.Rectangle{0, 0, 1280, 720 - 40}, rl.ColorAlpha(rl.RayWhite, opacity))
		}
		notification.Draw()
		rl.DrawText(strconv.Itoa(int(pageType)), 0, 0, 20, rl.Green)
		rl.EndMode2D()
		rl.EndDrawing()
	}
	if err != nil {
		panic(err)
	}
}
