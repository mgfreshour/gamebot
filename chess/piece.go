package chess

import (
	"errors"
	"fmt"
	"math"
)

type pieceType rune

const (
	Pawn   pieceType = 'p'
	Rook   pieceType = 'r'
	King   pieceType = 'k'
	Queen  pieceType = 'q'
	Knight pieceType = 'n'
	Bishop pieceType = 'b'
)

var pieceNames = map[pieceType]string{
	Pawn:   "Pawn",
	King:   "King",
	Rook:   "Rook",
	Queen:  "Queen",
	Bishop: "Bishop",
	Knight: "Knight",
}

var pieceSymbols = map[string]string{
	"wp": "♙",
	"wk": "♕",
	"wq": "♔",
	"wb": "♗",
	"wn": "♘",
	"wr": "♖",
	"bp": "♟",
	"bk": "♛",
	"bq": "♚",
	"bb": "♝",
	"bn": "♞",
	"br": "♜",
}

type Piece struct {
	x        byte
	y        byte
	captured bool
	side     side
	piece    pieceType
}

func (p *Piece) move(rank string, file string) {
	p.x, p.y = rankFileToXY(rank, file)
}

func (p *Piece) ValidateMove(rank string, file string, target *Piece) error {
	// Get the move vector
	x, y := rankFileToXY(file, rank)
	dx := math.Abs(float64(x) - float64(p.x))
	dy := math.Abs(float64(y) - float64(p.y))

	switch p.piece {
	case Pawn:
		fmt.Printf("MGF - checking %v, %v", dx, dy)
		if dx > 0 || dy > 2 {
			return errors.New("Invalid move")
		}
		// TODO en-passant
	case King:
		// TODO castling
	case Rook:
	case Queen:
	case Bishop:
	case Knight:
	default:
		return errors.New("Unknown piece type " + string(p.piece))
	}
	return nil
}

func (p *Piece) String() string {
	c := ""
	if p.captured {
		c = "(Captured)"
	}
	side := "White"
	if p.side == Black {
		side = "Black"
	}
	return side + pieceNames[p.piece] + c
}

func (p *Piece) Symbol() string {
	return pieceSymbols[fmt.Sprintf("%c%c", p.side, p.piece)]
}

func (p *Piece) Side() side {
	return p.side
}

func (p *Piece) Capture() {
	p.captured = true
}
