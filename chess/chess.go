package chess

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"fmt"
)

type side rune

const (
	White side = 'w'
	Black side = 'b'
)

var files string = "ABCDEFGH"

func rankFileToXY(file string, rank string) (byte, byte) {
	x := strings.Index(files, file)
	y, _ := strconv.ParseInt(rank, 10, 8)
	y = int64(math.Abs(float64(y - 8)))
	if x > 7 || x < 0 || y > 7 || y < 0 {
		panic(fmt.Sprintf("Invalid conversion happened! %v, %v to %v, %v", file, rank, x, y))
	}
	return byte(x), byte(y)
}

func xyToRankFile(x int, y int) (string, string) {
	if x > 7 || x < 0 || y > 7 || y < 0 {
		panic(fmt.Sprintf("Invalid conversion expected! %v, %v", x, y))
	}

	var r = strconv.Itoa(y + 1)
	var f = string(files[x])

	return f, r
}

type Board [8][8]*Piece

type Game struct {
	Pieces []*Piece
	Side   side
}

func NewGame() *Game {
	return LoadFENGame("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
}

func (g *Game) Board() Board {
	board := Board{}

	for _, piece := range g.Pieces {
		if !piece.captured {
			board[piece.y][piece.x] = piece
		}
	}

	return board
}

func (g *Game) Piece(file string, rank string) *Piece {
	x, y := rankFileToXY(file, rank)

	for _, piece := range g.Pieces {
		if !piece.captured && piece.x == x && piece.y == byte(y) {
			return piece
		}
	}

	return nil
}

func (g *Game) Move(srcFile string, srcRank string, dstFile string, dstRank string) error {
	moving := g.Piece(srcFile, srcRank)
	if moving == nil {
		return errors.New("No Piece Found there!")
	}

	if srcRank == dstRank && srcFile == dstFile {
		return errors.New("Invalid Move, same space!")
	}

	if moving.side != g.Side {
		return errors.New("Invalid Move, it's not your turn!")
	}

	target := g.Piece(dstFile, dstRank)
	if target != nil {
		if target.side == moving.side {
			return errors.New("Invalid Move, can't take your own pieces!")
		}
		target.Capture()
	}

	moving.move(dstRank, dstFile)

	if g.Side == White {
		g.Side = Black
	} else {
		g.Side = White
	}

	return nil
}


func (g *Game) ValidateMove(srcFile string, srcRank string, file string, rank string) error {
	// Get the move vector
	p := g.Piece(srcFile, srcRank)
	target := g.Piece(file, rank)
	x, y := rankFileToXY(file, rank)
	dx := int(float64(x) - float64(p.x))
	dy := int(float64(y) - float64(p.y))
	adx := int(math.Abs(float64(dx)))
	ady := int(math.Abs(float64(dy)))
	up := -1
	if p.side == Black { up = 1 }
	inv := fmt.Sprintf(" vec(%v,%v) %vX%v", dx, dy * up, p, target)

	switch p.piece {
	case Pawn:
		// TODO en-passant
		if target != nil {
			if adx != 1 || dy * up != 1 {
				return errors.New("Invalid capture!" + inv)
			}
		} else if dx > 0 || dy * up > 2 {
			return errors.New("Invalid move, going too far!" + inv)
		} else if dy * up <= 0 {
			return errors.New("Invalid move, going backwards!" + inv)
		}
	case King:
		// TODO castling
		if adx > 1 || ady > 1 {
			return errors.New("Invalid move, too far!" + inv)
		}
	case Rook:
		if adx >= 1 && ady >= 1 {
			return errors.New("Invalid move, multiple directions!" + inv)
		}
	case Queen:
	case Bishop:
	case Knight:
	default:
		return errors.New("Unknown piece type " + string(p.piece))
	}
	return nil
}
