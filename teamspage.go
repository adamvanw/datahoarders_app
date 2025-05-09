package main

import (
	"database/sql"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
)

func InitializeTeamsPage(db *sql.DB) *ScrollTable {
	rows, err := db.Query("SELECT t.team_id, t.team_name, t.team_location FROM teams t")
	if err != nil {
		log.Fatal(err)
	}
	elements := []Element{}
	for rows.Next() {
		id := 0
		name := ""
		location := ""
		err := rows.Scan(&id, &name, &location)
		if err != nil {
			return nil
			fmt.Printf("Failed scan.")
		}
		elements = append(elements, Element{name, id, location, "Team", false, rl.RayWhite})
	}
	fmt.Printf("%d elements\n", len(elements))
	return NewScrollTable(elements, rl.Rectangle{320, 100, 600, 550}, "Teams")
}
