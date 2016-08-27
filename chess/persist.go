package chess

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var insertGame *sql.Stmt
var updateGame *sql.Stmt
var loadGame *sql.Stmt

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./chess.db")
	checkErr(err)
	// defer db.Close()

	ddl, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS chess_games (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(64),
		fen VARCHAR(64)
	)`)
	checkErr(err)
	_, err = ddl.Exec()
	checkErr(err)

	insertGame, err = db.Prepare("INSERT INTO chess_games(name, fen) values(?,?)")
	checkErr(err)

	updateGame, err = db.Prepare("UPDATE chess_games SET fen=? WHERE name=?")
	checkErr(err)

	loadGame, err = db.Prepare("SELECT fen FROM chess_games WHERE name=? LIMIT 1")
	checkErr(err)
}

func LoadGame(name string) *Game {
	res, err := loadGame.Query(name)
	defer res.Close()
	checkErr(err)

	var fen string
	if !res.Next() {
		return nil
	}
	err = res.Scan(&fen)
	checkErr(err)

	game := LoadFENGame(fen)
	return game
}

func CreateGame(name string) *Game {
	game := NewGame()
	res, err := insertGame.Exec(name, SaveFENGame(game))
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Printf("Saved Id - %v", id)

	return game
}

func SaveGame(game *Game, name string) {
	_, err := updateGame.Exec(SaveFENGame(game), name)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
