package chess

import (
	"errors"
	"math"
	"strconv"
)

type side rune
const (
	White side = 'w'
	Black side = 'b'
)

var fileMappings map[string]byte = map[string]byte{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
	"F": 5,
	"G": 6,
	"H": 7,
}

func rankFileToXY(rank string, file string) (byte, byte) {
	x := fileMappings[file]
	y, _ := strconv.ParseInt(rank, 10, 8)
	y = int64(math.Abs(float64(y - 8)))
	return x, byte(y)
}

type Board [8][8]*Piece

type Game struct {
	Pieces []*Piece
	Side   side
}


func NewGame() *Game {
	game := Game{make([]*Piece, 0), White}

	game.Pieces = append(game.Pieces, &Piece{0, 0, false, Black, 'r'})
	game.Pieces = append(game.Pieces, &Piece{1, 0, false, Black, 'n'})
	game.Pieces = append(game.Pieces, &Piece{2, 0, false, Black, 'b'})
	game.Pieces = append(game.Pieces, &Piece{3, 0, false, Black, 'q'})
	game.Pieces = append(game.Pieces, &Piece{4, 0, false, Black, 'k'})
	game.Pieces = append(game.Pieces, &Piece{5, 0, false, Black, 'b'})
	game.Pieces = append(game.Pieces, &Piece{6, 0, false, Black, 'n'})
	game.Pieces = append(game.Pieces, &Piece{7, 0, false, Black, 'r'})

	game.Pieces = append(game.Pieces, &Piece{0, 1, false, Black, 'p'})
	game.Pieces = append(game.Pieces, &Piece{1, 1, false, Black, 'p'})
	game.Pieces = append(game.Pieces, &Piece{2, 1, false, Black, 'p'})
	game.Pieces = append(game.Pieces, &Piece{3, 1, false, Black, 'p'})
	game.Pieces = append(game.Pieces, &Piece{4, 1, false, Black, 'p'})
	game.Pieces = append(game.Pieces, &Piece{5, 1, false, Black, 'p'})
	game.Pieces = append(game.Pieces, &Piece{6, 1, false, Black, 'p'})
	game.Pieces = append(game.Pieces, &Piece{7, 1, false, Black, 'p'})

	game.Pieces = append(game.Pieces, &Piece{0, 6, false, White, 'p'})
	game.Pieces = append(game.Pieces, &Piece{1, 6, false, White, 'p'})
	game.Pieces = append(game.Pieces, &Piece{2, 6, false, White, 'p'})
	game.Pieces = append(game.Pieces, &Piece{3, 6, false, White, 'p'})
	game.Pieces = append(game.Pieces, &Piece{4, 6, false, White, 'p'})
	game.Pieces = append(game.Pieces, &Piece{5, 6, false, White, 'p'})
	game.Pieces = append(game.Pieces, &Piece{6, 6, false, White, 'p'})
	game.Pieces = append(game.Pieces, &Piece{7, 6, false, White, 'p'})

	game.Pieces = append(game.Pieces, &Piece{0, 7, false, White, 'r'})
	game.Pieces = append(game.Pieces, &Piece{1, 7, false, White, 'n'})
	game.Pieces = append(game.Pieces, &Piece{2, 7, false, White, 'b'})
	game.Pieces = append(game.Pieces, &Piece{3, 7, false, White, 'k'})
	game.Pieces = append(game.Pieces, &Piece{4, 7, false, White, 'q'})
	game.Pieces = append(game.Pieces, &Piece{5, 7, false, White, 'b'})
	game.Pieces = append(game.Pieces, &Piece{6, 7, false, White, 'n'})
	game.Pieces = append(game.Pieces, &Piece{7, 7, false, White, 'r'})

	return &game
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

func (g *Game) Piece(rank string, file string) *Piece {
	x, y := rankFileToXY(rank, file)

	for _, piece := range g.Pieces {
		if piece.x == x && piece.y == byte(y) {
			return piece
		}
	}

	return nil
}

func (g *Game) Move(srcRank string, srcFile string, dstRank string, dstFile string) error {
	moving := g.Piece(srcRank, srcFile)
	if moving == nil {
		return errors.New("No Piece Found there!")
	}

	if srcRank == dstRank && srcFile == dstFile {
		return errors.New("Invalid Move, same space!")
	}

	if moving.side != g.Side {
		return errors.New("Invalid Move, it's not your turn!")
	}

	target := g.Piece(dstRank, dstFile)
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

