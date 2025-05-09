package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Notification struct {
	text     string
	textSize int
	time     int32
	color    rl.Color
}

func NewNotification(text string, time int32, color rl.Color) *Notification {
	return &Notification{text: text, textSize: 20, time: time, color: color}
}

func (n *Notification) Draw() {
	var opacity float32 = 255
	if n.time < 1000 {
		opacity = float32(n.time) / 1000
	}
	rl.DrawText(n.text, 6, 6, int32(n.textSize), rl.ColorAlpha(n.color, opacity))
	n.time -= int32(1000 * rl.GetFrameTime())
}
