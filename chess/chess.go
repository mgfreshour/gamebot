package chess

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type side rune

const (
	White side = 'w'
	Black side = 'b'
)

var files string = "ABCDEFGH"

func rankFileToXY(rank string, file string) (byte, byte) {
	x := strings.Index(files, file)
	y, _ := strconv.ParseInt(rank, 10, 8)
	y = int64(math.Abs(float64(y - 8)))
	return byte(x), byte(y)
}

func xyToRankFile(x int, y int) (string, string) {
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
	x, y := rankFileToXY(rank, file)

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
