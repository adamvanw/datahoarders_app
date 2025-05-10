package main

import (
	"database/sql"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"strconv"
)

func InitializeCoachPage(db *sql.DB, coachId int) ([]ScrollTable, []Text) {
	rows, err := db.Query("SELECT m.first_name, m.last_name, m.experience FROM coaches c JOIN members m ON m.member_id = c.member_id WHERE c.coach_id = ?", coachId)
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

	coachName := NewText(firstName+" "+lastName, rl.Vector2{50, 285}, 50, rl.Black)
	experienceText := NewText("Experience: "+strconv.Itoa(experience)+" Years", rl.Vector2{50, 285 + 60}, 30, rl.DarkGray)

	rows, err = db.Query("SELECT team_id, team_name FROM teams WHERE team_id = (SELECT team_id FROM coaches WHERE coach_id = ?)", coachId)
	teamId := 0
	teamName := ""
	rows.Next()
	err = rows.Scan(&teamId, &teamName)
	rows.Close()

	teamSelector := NewScrollTable([]Element{*NewElement(teamName, teamId, "", "Team")}, rl.NewRectangle(1280-450, 315, 400, 50), "Team")

	return []ScrollTable{*teamSelector}, []Text{*coachName, *experienceText}
}
