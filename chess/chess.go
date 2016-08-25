package chess

import (
	"errors"
	"fmt"
	"math"
)

type side rune

const (
	White side = 'w'
	Black side = 'b'
)

type Board [8][8]*Piece

type castlingStatus struct {
	BlackKing  bool
	BlackQueen bool
	WhiteKing  bool
	WhiteQueen bool
}
type Game struct {
	Pieces        []*Piece
	Side          side
	FullMoveClock int
	HalfMoveClock int
	Castling      castlingStatus
	EnPassantFile string
	EnPassantRank string
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

	err := g.ValidateMove(srcFile, srcRank, dstFile, dstRank)
	if err != nil {
		return err
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

	g.FullMoveClock++
	if moving.piece == Pawn || target != nil {
		g.HalfMoveClock = 0
	} else {
		g.HalfMoveClock++
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
	if p.side == Black {
		up = 1
	}
	inv := fmt.Sprintf(" vec(%v,%v) %vX%v", dx, dy*up, p, target)

	switch p.piece {
	case Pawn:
		// TODO en-passant
		if target != nil {
			if adx != 1 || dy*up != 1 {
				return errors.New("Invalid capture!" + inv)
			}
		} else if dx > 0 || dy*up > 2 {
			return errors.New("Invalid move, going too far!" + inv)
		} else if dy*up <= 0 {
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
		if !((adx == 0 && ady >= 1) ||
			(adx >= 1 && ady == 0) ||
			(adx == ady)) {
			return errors.New("Invalid move, something ain't right!" + inv)
		}
	case Bishop:
		if adx != ady {
			return errors.New("Invalid move, not diagnal!" + inv)
		}
	case Knight:
		if !((adx == 2 && ady == 1) || (adx == 1 && ady == 2)) {
			return errors.New("Invalid move, not knightly!" + inv)
		}
	default:
		return errors.New("Unknown piece type " + string(p.piece))
	}
	return nil
}
