package chess

import (
	"fmt"
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
	p.x, p.y = rankFileToXY(file, rank)
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
