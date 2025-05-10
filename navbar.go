package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type NavBar struct {
	buttons [5]Button
	rect    rl.Rectangle
}

func NewNavBar() *NavBar {
	// rl.MeasureText() doesn't seem to work here, so all buttons have the same width.
	playersLength := float32(80) + 20
	coachesLength := float32(80) + 20
	teamsLength := float32(80) + 20
	gamesLength := float32(80) + 20

	var search = NewButton("Search", rl.Rectangle{0 + 1280/10 - playersLength/2, 720 - 40, playersLength, 40})
	search.textSize = 20
	search.outlineThickness = 0
	search.insideColor = rl.DarkBlue
	
	var players = NewButton("Players", rl.Rectangle{1280/5 + (1280 / 10) - playersLength/2, 720 - 40, playersLength, 40})
	players.textSize = 20
	players.outlineThickness = 0
	players.insideColor = rl.DarkBlue

	var coaches = NewButton("Coaches", rl.Rectangle{2*1280/5 + (1280 / 10) - coachesLength/2, 720 - 40, coachesLength, 40})
	coaches.textSize = 20
	coaches.outlineThickness = 0
	coaches.insideColor = rl.DarkBlue

	var teams = NewButton("Teams", rl.Rectangle{3*1280/5 + (1280 / 10) - teamsLength/2, 720 - 40, teamsLength, 40})
	teams.textSize = 20
	teams.outlineThickness = 0
	teams.insideColor = rl.DarkBlue

	var games = NewButton("Games", rl.Rectangle{4*1280/5 + (1280 / 10) - gamesLength/2, 720 - 40, gamesLength, 40})
	games.textSize = 20
	games.outlineThickness = 0
	games.insideColor = rl.DarkBlue

	return &NavBar{[5]Button{search, players, coaches, teams, games}, rl.Rectangle{0, 720 - 40, 1280, 40}}
}

func (n *NavBar) Draw() {
	rl.DrawRectangleRec(n.rect, rl.DarkGray)
	for i := 0; i < 5; i++ {
		n.buttons[i].Draw()
	}
}

func (n *NavBar) DetectInput() int {
	for i := 0; i < 5; i++ {
		n.buttons[i].DetectActivation(rl.GetMousePosition())
		if n.buttons[i].active {
			return i
		}
	}
	return -1
}
