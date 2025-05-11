package main

import (
	"database/sql"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"strconv"
)

func InitializeTeamPage(db *sql.DB, teamId int) ([]ScrollTable, []Text) {
	rows, err := db.Query("SELECT team_name, team_location FROM teams WHERE team_id = ?", teamId)
	if err != nil {
		log.Fatal(err)
	}
	teamName := ""
	location := ""

	if rows.Next() {
		err = rows.Scan(&teamName, &location)
	}
	rows.Close()

	nameText := NewText(teamName, rl.Vector2{50, 50}, 50, rl.Black)
	locationText := NewText("Location: "+location, rl.Vector2{50, 110}, 30, rl.DarkGray)
	wins := 0
	losses := 0

	gameElements := []Element{}
	rows, err = db.Query("SELECT game_id, (SELECT team_name FROM teams WHERE teams.team_id = games.away_id), (SELECT team_name FROM teams WHERE teams.team_id = games.home_id), g_winner(game_id) FROM games WHERE away_id = ? OR home_id = ?", teamId, teamId)

	for rows.Next() {
		gameId := 0
		awayName := ""
		homeName := ""
		winnerId := 0
		description := ""
		rows.Scan(&gameId, &awayName, &homeName, &winnerId)
		if winnerId == teamId {
			wins++
			description = "Won"
		} else {
			losses++
			description = "Lost"
		}
		gameElements = append(gameElements, *NewElement(awayName+" at "+homeName, gameId, description, "Game"))
	}
	rows.Close()

	var coachSelector ScrollTable
	rows, err = db.Query("SELECT coach_id, m.first_name, m.last_name FROM coaches JOIN members m ON coaches.member_id = m.member_id WHERE team_id = ?", teamId)
	if rows.Next() {
		coachId := 0
		firstName := ""
		lastName := ""
		rows.Scan(&coachId, &firstName, &lastName)
		coachSelector = *NewScrollTable([]Element{*NewElement(firstName+" "+lastName, coachId, "", "Coach")}, rl.NewRectangle(50, 350, 350, 50), "Coach")
	}
	rows.Close()

	scoreText := NewText(strconv.Itoa(wins)+":"+strconv.Itoa(losses), rl.Vector2{50, 150}, 30, rl.DarkGray)
	gameScroll := *NewScrollTable(gameElements, rl.Rectangle{600, 720/2 - 400/2 + 100, 1280 - 650, 400}, "Games Played")
	gameScroll.fontSize = 20

	return []ScrollTable{gameScroll, coachSelector}, []Text{*nameText, *locationText, *scoreText}
}
