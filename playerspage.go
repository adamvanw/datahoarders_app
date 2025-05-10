package main

import (
	"database/sql"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
)

func InitializePlayersPage(db *sql.DB) []ScrollTable {
	rows, err := db.Query("SELECT p.player_id, m.first_name, m.last_name, t.team_name FROM players p JOIN members m ON p.member_id = m.member_id JOIN teams t ON p.team_id = t.team_id ORDER BY m.first_name ASC")
	if err != nil {
		log.Fatal(err)
	}
	elements := []Element{}
	for rows.Next() {
		id := 0
		first_name := ""
		last_name := ""
		team_name := ""
		err := rows.Scan(&id, &first_name, &last_name, &team_name)
		if err != nil {
			return nil
			fmt.Printf("Failed scan.")
		}
		elements = append(elements, Element{first_name + " " + last_name, id, team_name, "Player", false, rl.RayWhite})
	}
	rows.Close()
	fmt.Printf("%d elements\n", len(elements))
	return []ScrollTable{*NewScrollTable(elements, rl.Rectangle{320, 100, 600, 550}, "Players")}
}
