package chess

import (
	"strconv"
	"unicode"
)

func LoadFENGame(s string) *Game {
	game := &Game{make([]*Piece, 0), White}

	n := loadBoardPositions(s, game)
	n = loadSide(s, n, game)
	n = loadCastling(s, n, game)
	n = loadEnPassant(s, n, game)
	n = loadHalfMove(s, n, game)
	n = loadFullMove(s, n, game)

	return game
}

func loadBoardPositions(s string, game *Game) int {
	x, y, i := 0, 0, 0
	for i = 0; i < len(s); i++ {
		var c rune = rune(s[i])
		if unicode.IsDigit(c) {
			offset, _ := strconv.ParseInt(string(c), 10, 8)
			x += int(offset)
		} else if c == '/' {
			y++
			x = 0
		} else if c == ' ' {
			i++
			// we're done with this portion
			break
		} else {
			side := White
			// Must be a piece!
			if unicode.IsLower(c) {
				side = Black
			}
			p := Piece{byte(x), byte(y), false, side, pieceType(unicode.ToLower(c))}
			game.Pieces = append(game.Pieces, &p)
			x++
		}
	}
	return i
}

func loadSide(s string, n int, game *Game) int {
	return n + 1
}
func loadCastling(s string, n int, game *Game) int {
	return n + 1

}
func loadEnPassant(s string, n int, game *Game) int {
	return n + 1

}
func loadHalfMove(s string, n int, game *Game) int {
	return n + 1

}
func loadFullMove(s string, n int, game *Game) int {
	return n + 1

}
