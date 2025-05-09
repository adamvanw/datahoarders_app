package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	text                 string
	active               bool
	selected             bool
	insideColor          rl.Color
	outlineColor         rl.Color
	outlineThickness     float32
	selectedOutlineColor rl.Color
	activeOutlineColor   rl.Color
	textColor            rl.Color
	textSize             int32
	rect                 rl.Rectangle
}

func NewButton(text string, rect rl.Rectangle) Button {
	return Button{
		text:                 text,
		rect:                 rect,
		outlineThickness:     4,
		active:               false,
		selected:             false,
		insideColor:          rl.Green,
		outlineColor:         rl.Lime,
		selectedOutlineColor: rl.SkyBlue,
		activeOutlineColor:   rl.DarkGreen,
		textColor:            rl.RayWhite,
		textSize:             int32(rect.Height - 4),
	}
}

func (b Button) Draw() {
	var color = b.outlineColor
	if b.active {
		color = b.activeOutlineColor
	} else if b.selected {
		color = b.selectedOutlineColor
	}
	rl.DrawRectangle(int32(b.rect.X), int32(b.rect.Y), int32(b.rect.Width), int32(b.rect.Height), b.insideColor)
	rl.DrawRectangleLinesEx(b.rect, b.outlineThickness, color)
	rl.DrawText(b.text, int32(b.rect.X+b.outlineThickness+b.rect.Width/2-float32(rl.MeasureText(b.text, b.textSize)/2)), int32(b.rect.Y-(float32(b.textSize)/2)+(b.rect.Height/2)), b.textSize, b.textColor)
}

func (b *Button) DetectActivation(mousePos rl.Vector2) {
	if rl.CheckCollisionPointRec(mousePos, b.rect) {
		b.active = true
	} else {
		b.active = false
	}
}
