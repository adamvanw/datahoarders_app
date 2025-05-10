package main

import (
	"database/sql"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"strconv"
)

func InitializePlayerPage(db *sql.DB, playerId int) ([]ScrollTable, []Text) {
	rows, err := db.Query("SELECT m.first_name, m.last_name, m.experience FROM players p JOIN members m ON m.member_id = p.member_id WHERE p.player_id = ?", playerId)
	if err != nil {
		log.Fatal(err)
	}
	firstName := ""
	lastName := ""
	experience := 0

	if rows.Next() {
		err = rows.Scan(&firstName, &lastName, &experience)
	}
	rows.Close()

	playerName := NewText(firstName+" "+lastName, rl.Vector2{50, 285}, 50, rl.Black)
	experienceText := NewText("Experience: "+strconv.Itoa(experience)+" Years", rl.Vector2{50, 285 + 60}, 30, rl.DarkGray)

	rows, err = db.Query("SELECT team_id, team_name FROM teams WHERE team_id = (SELECT team_id FROM players WHERE player_id = ?)", playerId)
	teamId := 0
	teamName := ""
	rows.Next()
	err = rows.Scan(&teamId, &teamName)
	rows.Close()

	rows, err = db.Query("SELECT coach_id, m.first_name, m.last_name FROM coaches JOIN members m ON m.member_id = coaches.member_id WHERE coaches.team_id = ?", teamId)
	coachId := 0
	coachFirstName := ""
	coachLastName := ""
	rows.Next()
	err = rows.Scan(&coachId, &coachFirstName, &coachLastName)
	fmt.Printf("%s %s\n", teamName, coachFirstName)
	rows.Close()

	teamSelector := NewScrollTable([]Element{*NewElement(teamName, teamId, "", "Team")}, rl.NewRectangle(1280-450, 250, 400, 50), "Team")
	coachSelector := NewScrollTable([]Element{*NewElement(coachFirstName+" "+coachLastName, coachId, "", "Coach")}, rl.NewRectangle(1280-450, 400, 400, 50), "Coach")

	return []ScrollTable{*teamSelector, *coachSelector}, []Text{*playerName, *experienceText}
}
