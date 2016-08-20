package chess

import "fmt"

type pieceType rune

var pieceTrans = map[pieceType]string{
	'p': "Pawn",
	'k': "King",
	'q': "Queen",
	'b': "Bishop",
	'n': "Knight",
	'r': "Rook",
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

func (p *Piece) String() string {
	c := ""
	if p.captured {
		c = "(Captured)"
	}
	side := "White"
	if p.side == Black {
		side = "Black"
	}
	return side + pieceTrans[p.piece] + c
}

func (p *Piece) Symbol() string {
	return pieceSymbols[fmt.Sprintf("%c%c", p.side, p.piece)]
}

func (p *Piece) Side() side {
	return p.side
}

func (p *Piece) Capture() {
	p.captured = true
	p.x = 255
	p.y = 255
}
