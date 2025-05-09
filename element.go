package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Element struct {
	text        string
	id          int
	description string
	elementType string
	active      bool
	textColor   rl.Color
}

func NewElement(text string, id int, description string, elementType string) *Element {
	return &Element{text, id, description, elementType, false, rl.RayWhite}
}
