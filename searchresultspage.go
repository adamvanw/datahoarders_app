package main

import (
	"database/sql"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"strings"
)

func InitializeResultsPage(db *sql.DB, search string) ([]ScrollTable, []Text) {
	elements := []Element{}

	rows, err := db.Query("SELECT m.first_name, m.last_name, p.player_id FROM players p JOIN members m ON m.member_id = p.member_id WHERE LOWER(m.first_name) LIKE CONCAT('%',?,'%') OR LOWER(m.last_name) LIKE CONCAT('%',?,'%')", strings.ToLower(search), strings.ToLower(search))
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		playerFirstName := ""
		playerLastName := ""
		playerId := 0
		rows.Scan(&playerFirstName, &playerLastName, &playerId)
		elements = append(elements, *NewElement(playerFirstName+" "+playerLastName, playerId, "Player", "Player"))
		fmt.Printf("Added player\n")
	}
	rows, err = db.Query("SELECT m.first_name, m.last_name, coach_id FROM coaches JOIN members m ON m.member_id = coaches.member_id WHERE LOWER(m.first_name) LIKE CONCAT('%',?,'%') OR LOWER(m.last_name) LIKE CONCAT('%',?,'%')", strings.ToLower(search), strings.ToLower(search))
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		coachFirstName := ""
		coachLastName := ""
		coachId := 0
		rows.Scan(&coachFirstName, &coachLastName, &coachId)
		elements = append(elements, *NewElement(coachFirstName+" "+coachLastName, coachId, "Coach", "Coach"))
		fmt.Printf("Added coach\n")
	}
	rows.Close()

	rows, err = db.Query("SELECT game_id, game_date, (SELECT team_name FROM teams WHERE teams.team_id = games.away_id), (SELECT team_name FROM teams WHERE teams.team_id = games.home_id) FROM games WHERE CONVERT(game_date, CHAR) LIKE CONCAT('%', ?, '%')", strings.ToLower(search))
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		gameId := 0
		date := ""
		awayTeam := ""
		homeTeam := ""
		rows.Scan(&gameId, &date, &awayTeam, &homeTeam)
		elements = append(elements, *NewElement(awayTeam+" "+homeTeam, gameId, "Game: "+date, "Game"))
	}
	rows.Close()

	rows, err = db.Query("SELECT team_id, team_name FROM teams WHERE team_name LIKE CONCAT('%', ?, '%')", strings.ToLower(search))
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		teamId := 0
		teamName := ""
		rows.Scan(&teamId, &teamName)
		elements = append(elements, *NewElement(teamName, teamId, string("Team"), "Team"))
	}
	rows.Close()

	height := 600
	if len(elements) < 12 {
		height = len(elements) * 50
	}
	if len(elements) == 0 {
		return []ScrollTable{}, []Text{*NewText("Could not find search results for \""+search+"\"", rl.Vector2{50, 50}, 50, rl.Red)}
	}
	return []ScrollTable{*NewScrollTableDescript(elements, rl.NewRectangle(1280/2-1000/2, 50, 1000, float32(height)), "Search Results for \""+search+"\"")}, []Text{}
}
