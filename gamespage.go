package main

import (
	"database/sql"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
)

func InitializeGamesPage(db *sql.DB) []ScrollTable {
	rows, err := db.Query("SELECT game_id, game_date, (SELECT team_name FROM teams WHERE teams.team_id = games.away_id), (SELECT team_name FROM teams WHERE teams.team_id = games.home_id) FROM games;")
	if err != nil {
		log.Fatal(err)
	}
	elements := []Element{}
	for rows.Next() {
		id := 0
		date := ""
		away_name := ""
		home_name := ""
		err := rows.Scan(&id, &date, &away_name, &home_name)
		if err != nil {
			return nil
			fmt.Printf("Failed scan.")
		}
		elements = append(elements, Element{away_name + " at " + home_name, id, date, "Game", false, rl.RayWhite})
	}
	rows.Close()
	fmt.Printf("%d elements\n", len(elements))
	return []ScrollTable{*NewScrollTable(elements, rl.Rectangle{1280/2 - 900/2, 100, 900, 550}, "Games")}
}
