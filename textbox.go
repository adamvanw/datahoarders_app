package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// yes, I implemented a text box system. what about it.
type TextBox struct {
	rect               rl.Rectangle
	outlineColor       rl.Color
	activeOutlineColor rl.Color
	insideColor        rl.Color
	outlineThickness   float32
	text               string
	textSize           int32
	active             bool
	textPos            int
	title              string
	titleTextSize      int32
	timer              int32
}

func NewTextBox(rect rl.Rectangle, title string) TextBox {
	return TextBox{
		rect:               rect,
		outlineColor:       rl.Gray,
		activeOutlineColor: rl.SkyBlue,
		insideColor:        rl.RayWhite,
		outlineThickness:   4,
		title:              title,
		text:               "",
		timer:              0,
		textPos:            0,
		titleTextSize:      20,
		active:             false,
		textSize:           int32(rect.Height - 4 - 4),
	}
}

func (t *TextBox) DetectActivation(mousePos rl.Vector2) {
	if rl.CheckCollisionPointRec(mousePos, t.rect) {
		t.active = true
		t.textPos = len(t.text)
	} else {
		t.active = false
	}
}

func (t *TextBox) Draw() {
	var color rl.Color
	if t.active {
		color = t.activeOutlineColor
	} else {
		color = t.outlineColor
	}
	rl.DrawRectangle(int32(t.rect.X), int32(t.rect.Y), int32(t.rect.Width), int32(t.rect.Height), t.insideColor)
	rl.DrawRectangleLinesEx(t.rect, t.outlineThickness, color)
	rl.DrawText(t.text, int32(t.rect.X+t.outlineThickness+4), int32(t.rect.Y+t.outlineThickness), t.textSize, color)
	rl.DrawText(t.title, int32(t.rect.X-t.outlineThickness/2), int32(t.rect.Y-t.outlineThickness/2-float32(t.titleTextSize)), t.titleTextSize, color)
	if t.textPos == len(t.text) && t.active {
		if (t.timer/500)%2 == 0 {
			rl.DrawText("_", int32(t.rect.X+2+4+t.outlineThickness+float32(rl.MeasureText(t.text, t.textSize))), int32(t.rect.Y+t.outlineThickness), t.textSize, color)
		}
	} else if t.active && (t.timer/500)%2 == 0 {
		rl.DrawText("|", int32(t.rect.X+6+t.outlineThickness+float32(rl.MeasureText(t.text[:t.textPos], t.textSize))), int32(t.rect.Y+t.outlineThickness), t.textSize, color)
	}
	if t.active {
		t.timer += int32(rl.GetFrameTime() * 1000)
	}
}

func (t *TextBox) DetectInput() {
	if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyV) {
		t.text = t.text[:t.textPos] + rl.GetClipboardText() + t.text[t.textPos:]
		t.textPos = t.textPos + len(rl.GetClipboardText())
	}
	if rl.IsKeyPressed(rl.KeyBackspace) || rl.IsKeyPressedRepeat(rl.KeyBackspace) {
		if t.textPos > 0 {
			t.text = t.text[:t.textPos-1] + t.text[t.textPos:]
		}
		t.textPos = max(t.textPos-1, 0)
	}
	if rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressedRepeat(rl.KeyLeft) {
		t.textPos = max(t.textPos-1, 0)
		fmt.Printf("huh")
	}
	if rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressedRepeat(rl.KeyRight) {
		t.textPos = min(t.textPos+1, len(t.text))
	}
	key := rl.GetCharPressed()
	for key != 0 {
		if key >= 32 && key <= 125 {
			t.text = t.text[:t.textPos] + string(rune(key)) + t.text[t.textPos:]
			t.textPos++
		}
		key = rl.GetCharPressed()
	}
}
