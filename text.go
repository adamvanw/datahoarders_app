package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Text struct {
	text     string
	pos      rl.Vector2
	fontSize int32
	color    rl.Color
}

func NewText(text string, pos rl.Vector2, fontSize int32, color rl.Color) *Text {
	return &Text{text: text, pos: pos, fontSize: fontSize, color: color}
}

func (text *Text) Draw() {
	rl.DrawText(text.text, int32(text.pos.X), int32(text.pos.Y), text.fontSize, text.color)
}
