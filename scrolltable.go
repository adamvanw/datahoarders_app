package main

import (
	"database/sql"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ScrollTable struct {
	elements            []Element
	rect                rl.Rectangle
	visibleElementCount int
	visiblePosition     int
	title               string
	fontSize            int32
}

func NewScrollTable(elements []Element, rect rl.Rectangle, title string) *ScrollTable {
	visibleElementCount := rect.Height / 50
	return &ScrollTable{elements, rect, int(visibleElementCount), 0, title, 30}
}

func (st *ScrollTable) DetectInput(db *sql.DB) (int, int) {
	// detect whether to scroll table up or down.
	for i := 0; i < st.visibleElementCount; i++ {
		st.elements[i+st.visiblePosition].textColor = rl.RayWhite
	}

	if !rl.CheckCollisionPointRec(rl.GetMousePosition(), st.rect) {
		return -1, -1
	}

	if rl.IsKeyPressed(rl.KeyDown) {
		st.visiblePosition = min(len(st.elements)-st.visibleElementCount, st.visiblePosition+1)
	} else if rl.IsKeyPressedRepeat(rl.KeyDown) {
		st.visiblePosition = min(len(st.elements)-st.visibleElementCount, st.visiblePosition+5)
	} else if rl.IsKeyPressed(rl.KeyUp) {
		st.visiblePosition = max(0, st.visiblePosition-1)
	} else if rl.IsKeyPressedRepeat(rl.KeyUp) {
		st.visiblePosition = max(0, st.visiblePosition-5)
	}

	if rl.GetMouseWheelMove() < 0 {
		st.visiblePosition = min(len(st.elements)-st.visibleElementCount, st.visiblePosition-int(rl.GetMouseWheelMove()))
	} else if rl.GetMouseWheelMove() > 0 {
		st.visiblePosition = max(0, st.visiblePosition-int(rl.GetMouseWheelMove()))
	}

	st.elements[int((rl.GetMousePosition().Y-st.rect.Y)/50)+st.visiblePosition].textColor = rl.SkyBlue
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		element := st.elements[int((rl.GetMousePosition().Y-st.rect.Y)/50)+st.visiblePosition]
		switch element.elementType {
		case "Game":
			return element.id, GamePage
		case "Coach":
			return element.id, CoachPage
		case "Player":
			return element.id, PlayerPage
		case "Team":
			return element.id, TeamPage
		}
	}
	return -1, -1
}

func (st *ScrollTable) Render() {
	rl.DrawText(st.title, int32(st.rect.X), int32(st.rect.Y-20-10), 20, rl.DarkGray)
	rl.DrawRectangleRec(st.rect, rl.Gray)

	if st.visibleElementCount == 1 {
		rl.DrawRectangleRec(rl.Rectangle{st.rect.X, (float32(st.visiblePosition))*50 + st.rect.Y, st.rect.Width, 50}, rl.ColorAlpha(rl.LightGray, 1))
		rl.DrawText(st.elements[0].text, int32(st.rect.X+30), int32((float32(st.visiblePosition))*50+st.rect.Y)+(25-st.fontSize/2), st.fontSize, st.elements[0].textColor)
		return
	}

	for i := st.visiblePosition; i < st.visiblePosition+st.visibleElementCount; i++ {
		if i > len(st.elements) {
			return
		}
		rl.DrawRectangleRec(rl.Rectangle{st.rect.X, (float32(i-st.visiblePosition))*50 + st.rect.Y, st.rect.Width, 50}, rl.ColorAlpha(rl.LightGray, float32(i%2)))
		rl.DrawText(st.elements[i].text, int32(st.rect.X+30), int32((float32(i-st.visiblePosition))*50+st.rect.Y)+(25-st.fontSize/2), st.fontSize, st.elements[i].textColor)
	}
}
