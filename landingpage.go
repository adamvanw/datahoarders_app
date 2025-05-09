package main

import rl "github.com/gen2brain/raylib-go/raylib"

func InitializeLandingPage() ([]TextBox, []Button) {
	var searchBar = NewTextBox(rl.Rectangle{30, 720/2 - 60, float32(rl.GetRenderWidth() - 240), 60}, "Search")
	var searchSubmit = NewButton("Submit", rl.Rectangle{float32(rl.GetRenderWidth() - 30 - 80 - 60 + 5), 720/2 - 50, 130, 40})
	searchSubmit.textSize = 20
	return []TextBox{searchBar}, []Button{searchSubmit}
}
