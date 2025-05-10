package main

import (
	"database/sql"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
)

func InitializeCoachesPage(db *sql.DB) []ScrollTable {
	rows, err := db.Query("SELECT c.coach_id, m.first_name, m.last_name, t.team_name FROM coaches c JOIN members m ON c.member_id = m.member_id JOIN teams t ON t.team_id = c.team_id ORDER BY m.last_name ASC")
	if err != nil {
		log.Fatal(err)
	}
	elements := []Element{}
	for rows.Next() {
		first_name := ""
		last_name := ""
		team_name := ""
		id := 0
		element_type := "Coach"
		err := rows.Scan(&id, &first_name, &last_name, &team_name)
		if err != nil {
			return nil
			fmt.Printf("Failed scan.")
		}
		elements = append(elements, Element{first_name + " " + last_name, id, team_name, element_type, false, rl.RayWhite})
	}
	rows.Close()
	fmt.Printf("%d elements\n", len(elements))
	return []ScrollTable{*NewScrollTable(elements, rl.Rectangle{320, 100, 600, 550}, "Coaches")}
}
