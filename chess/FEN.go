package chess

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
)

func LoadFENGame(s string) *Game {
	game := &Game{make([]*Piece, 0), White, 0, 0, castlingStatus{}, "", ""}

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
	if s[n] == 'w' {
		game.Side = White
	} else {
		game.Side = Black
	}
	return n + 1
}
func loadCastling(s string, n int, game *Game) int {
	if strings.Index("KQkq ", string(s[n])) == -1 {
		return n
	}
	n++
	for ; strings.Index("KQkq ", string(s[n])) != -1; n++ {
		switch s[n] {
		case 'k':
			game.Castling.BlackKing = true
		case 'K':
			game.Castling.WhiteKing = true
		case 'q':
			game.Castling.BlackQueen = true
		case 'Q':
			game.Castling.WhiteQueen = true
		case ' ':
			// Skip it
		default:
			panic("Unknown castling status '" + string(s[n]) + "'")
		}
	}
	return n

}
func loadEnPassant(s string, n int, game *Game) int {
	if s[n] == ' ' {
		n++
	}
	if s[n] != '-' {
		game.EnPassantFile = string(s[n])
		n++
		game.EnPassantRank = string(s[n])
	}
	return n + 1

}
func loadHalfMove(s string, n int, game *Game) int {
	num := ""
	if s[n] == ' ' {
		n++
	}
	for ; s[n] != ' '; n++ {
		num += string(s[n])
	}
	x, _ := strconv.ParseInt(num, 10, 32)
	game.HalfMoveClock = int(x)
	return n

}
func loadFullMove(s string, n int, game *Game) int {
	num := ""

	if s[n] == ' ' {
		n++
	}
	for ; n < len(s) && s[n] != ' '; n++ {
		num += string(s[n])
	}
	x, _ := strconv.ParseInt(num, 10, 32)
	game.FullMoveClock = int(x)
	return n
}

func SaveFENGame(g *Game) string {
	var buf bytes.Buffer

	for y := 0; y < 8; y++ {
		n := 0
		for x := 0; x < 8; x++ {
			f, r := xyToRankFile(x, y)
			p := g.Piece(f, r)
			if p != nil {
				if n > 0 {
					buf.Write([]byte(strconv.Itoa(n)))
					n = 0
				}
				if p.side == White {
					buf.WriteByte(byte(p.piece))
				} else {
					buf.WriteByte(byte(unicode.ToUpper(rune(p.piece))))
				}
			} else {
				n++
			}
		}
		if n > 0 {
			buf.Write([]byte(strconv.Itoa(n)))
			n = 0
		}
		if y < 7 {
			buf.WriteByte('/')
		}
	}
	buf.WriteByte(' ')

	buf.WriteByte(byte(g.Side))

	buf.WriteByte(' ')

	if g.Castling.WhiteKing {
		buf.WriteByte('K')
	}
	if g.Castling.WhiteQueen {
		buf.WriteByte('Q')
	}
	if g.Castling.BlackKing {
		buf.WriteByte('k')
	}
	if g.Castling.BlackQueen {
		buf.WriteByte('q')
	}
	// TODO EnPassant
	buf.WriteByte(' ')
	if g.EnPassantFile != "" {
		buf.Write([]byte(g.EnPassantFile + g.EnPassantRank))
	} else {
		buf.WriteByte('-')
	}
	buf.WriteByte(' ')

	buf.Write([]byte(strconv.Itoa(g.HalfMoveClock)))
	buf.WriteByte(' ')
	buf.Write([]byte(strconv.Itoa(g.FullMoveClock)))

	return buf.String()
}
