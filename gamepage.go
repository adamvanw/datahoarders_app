package main

import (
	"database/sql"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"strconv"
)

func InitializeGamePage(db *sql.DB, gameId int) ([]ScrollTable, []Text) {
	rows, err := db.Query("SELECT game_date, (SELECT team_name FROM teams WHERE teams.team_id = games.away_id), away_id, (SELECT team_name FROM teams WHERE teams.team_id = games.home_id), home_id, away_score, home_score, g_winner(game_id) FROM games WHERE game_id = ?", gameId)
	if err != nil {
		log.Fatal(err)
	}
	date := ""
	awayTeam := ""
	awayId := 0
	homeTeam := ""
	homeId := 0
	awayScore := 0
	homeScore := 0
	winnerId := 0
	if rows.Next() {
		rows.Scan(&date, &awayTeam, &awayId, &homeTeam, &homeId, &awayScore, &homeScore, &winnerId)
	}
	rows.Close()
	gameName := NewText(awayTeam+" at "+homeTeam, rl.Vector2{20, 20}, 50, rl.Black)
	gameDate := NewText(date, rl.Vector2{20, 80}, 30, rl.DarkGray)
	home := NewText("Home: "+homeTeam+". Score: "+strconv.Itoa(homeScore), rl.Vector2{20, 120}, 20, rl.DarkGray)
	away := NewText("Away: "+awayTeam+". Score: "+strconv.Itoa(awayScore), rl.Vector2{20, 145}, 20, rl.DarkGray)
	if winnerId == awayId {
		away.color = rl.Green
	} else {
		home.color = rl.Green
	}
	rows, err = db.Query("SELECT player_id, first_name, last_name FROM players JOIN members ON players.member_id = members.member_id WHERE players.team_id = ?", homeId)
	homeElements := []Element{}
	for rows.Next() {
		first_name := ""
		last_name := ""
		player_id := 0
		rows.Scan(&player_id, &first_name, &last_name)
		element := *NewElement(first_name+" "+last_name, player_id, "", "Player")
		homeElements = append(homeElements, element)
	}
	homeScroll := NewScrollTable(homeElements, rl.NewRectangle(380, 250, 400, 400), homeTeam+" Players")
	rows.Close()

	rows, err = db.Query("SELECT player_id, first_name, last_name FROM players JOIN members ON players.member_id = members.member_id WHERE players.team_id = ?", awayId)
	awayElements := []Element{}
	for rows.Next() {
		player_id := 0
		first_name := ""
		last_name := ""
		rows.Scan(&player_id, &first_name, &last_name)
		element := *NewElement(first_name+" "+last_name, player_id, "", "Player")
		awayElements = append(awayElements, element)
	}
	awayScroll := NewScrollTable(awayElements, rl.NewRectangle(830, 250, 400, 400), awayTeam+" Players")
	awayScroll.fontSize = 20
	homeScroll.fontSize = 20
	rows.Close()

	return []ScrollTable{*awayScroll, *homeScroll}, []Text{*gameName, *gameDate, *home, *away}
}
