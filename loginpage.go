package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func InitializeLoginPage(windowSize rl.Vector2) ([]TextBox, []Button) {
	var yRatio = 720 / float32(windowSize.Y)
	var xRatio = 1280 / float32(windowSize.X)
	var layout []rl.Rectangle
	for i := 0; i < 4; i++ {
		layout = append(layout, rl.Rectangle{float32(40 * xRatio), float32(float32(i)*windowSize.Y/5 + 60*yRatio), float32(windowSize.X - 80*xRatio), float32(windowSize.Y/5 - 100*yRatio)})
	}
	layout = append(layout, rl.Rectangle{windowSize.X/2 - (windowSize.X/4 - 80*xRatio/2), 4*windowSize.Y/5 + 60*yRatio, windowSize.X/2 - 160*xRatio, windowSize.Y/5 - 80*yRatio})
	var address = NewTextBox(layout[0], "IP Address (defaults to localhost)")
	address.text = "localhost"
	var port = NewTextBox(layout[1], "Port (defaults to 3306)")
	port.text = "3306"
	var user = NewTextBox(layout[2], "User (defaults to root)")
	user.text = "root"
	var password = NewTextBox(layout[3], "Password (defaults to football11)")
	password.text = "football11"
	submit := NewButton("Submit", layout[4])

	return []TextBox{address, port, user, password}, []Button{submit}
}
