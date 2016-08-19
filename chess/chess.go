package chess

import (
	"strconv"
	"fmt"
	"errors"
	"math"
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


func rankFileToXY (rank string, file string) (byte, byte) {
	x := fileMappings[file]
	y, _ := strconv.ParseInt(rank, 10, 8)
	y = int64(math.Abs(float64(y - 8)))
	return x, byte(y)
}

type Piece struct {
	X        byte
	Y        byte
	Captured bool
	Side     rune
	Piece    rune
}

func (p *Piece) Move (rank string, file string) {
	p.X, p.Y = rankFileToXY(rank, file)
}

type Board [8][8]*Piece

type Game struct {
	Pieces []*Piece
	Side rune
}

var pieceTrans = map[string]string {
	"wp": "WhitePawn", //"♙",
	"wk": "WhiteKing", //"♕",
	"wq": "WhiteQueen", //"♔",
	"wb": "WhiteBishop", //"♗",
	"wn": "WhiteKnight", //"♘",
	"wr": "WhiteRook", //"♖",
	"bp": "BlackPawn", //"♟",
	"bk": "BlackKing", //"♛",
	"bq": "BlackQueen", //"♚",
	"bb": "BlackBishop", //"♝",
	"bn": "BlackKnight", //"♞",
	"br": "BlackRook", //"♜",
}
var pieceSymbols = map[string]string {
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

func (p *Piece) String () string {
	c := ""
	if p.Captured {
		c = "(Captured)"
	}
	return pieceTrans[fmt.Sprintf("%c%c", p.Side, p.Piece)] + c
}

func (p *Piece) Capture () {
	p.Captured = true
	p.X = 255
	p.Y = 255
}

func NewGame () *Game {
	game := Game{make([]*Piece, 0), 'w'}

	game.Pieces = append(game.Pieces, &Piece{0, 0, false, 'b', 'r'})
	game.Pieces = append(game.Pieces, &Piece{1, 0, false, 'b', 'n'})
	game.Pieces = append(game.Pieces, &Piece{2, 0, false, 'b', 'b'})
	game.Pieces = append(game.Pieces, &Piece{3, 0, false, 'b', 'q'})
	game.Pieces = append(game.Pieces, &Piece{4, 0, false, 'b', 'k'})
	game.Pieces = append(game.Pieces, &Piece{5, 0, false, 'b', 'b'})
	game.Pieces = append(game.Pieces, &Piece{6, 0, false, 'b', 'n'})
	game.Pieces = append(game.Pieces, &Piece{7, 0, false, 'b', 'r'})

	game.Pieces = append(game.Pieces, &Piece{0, 1, false, 'b', 'p'})
	game.Pieces = append(game.Pieces, &Piece{1, 1, false, 'b', 'p'})
	game.Pieces = append(game.Pieces, &Piece{2, 1, false, 'b', 'p'})
	game.Pieces = append(game.Pieces, &Piece{3, 1, false, 'b', 'p'})
	game.Pieces = append(game.Pieces, &Piece{4, 1, false, 'b', 'p'})
	game.Pieces = append(game.Pieces, &Piece{5, 1, false, 'b', 'p'})
	game.Pieces = append(game.Pieces, &Piece{6, 1, false, 'b', 'p'})
	game.Pieces = append(game.Pieces, &Piece{7, 1, false, 'b', 'p'})

	game.Pieces = append(game.Pieces, &Piece{0, 6, false, 'w', 'p'})
	game.Pieces = append(game.Pieces, &Piece{1, 6, false, 'w', 'p'})
	game.Pieces = append(game.Pieces, &Piece{2, 6, false, 'w', 'p'})
	game.Pieces = append(game.Pieces, &Piece{3, 6, false, 'w', 'p'})
	game.Pieces = append(game.Pieces, &Piece{4, 6, false, 'w', 'p'})
	game.Pieces = append(game.Pieces, &Piece{5, 6, false, 'w', 'p'})
	game.Pieces = append(game.Pieces, &Piece{6, 6, false, 'w', 'p'})
	game.Pieces = append(game.Pieces, &Piece{7, 6, false, 'w', 'p'})

	game.Pieces = append(game.Pieces, &Piece{0, 7, false, 'w', 'r'})
	game.Pieces = append(game.Pieces, &Piece{1, 7, false, 'w', 'n'})
	game.Pieces = append(game.Pieces, &Piece{2, 7, false, 'w', 'b'})
	game.Pieces = append(game.Pieces, &Piece{3, 7, false, 'w', 'k'})
	game.Pieces = append(game.Pieces, &Piece{4, 7, false, 'w', 'q'})
	game.Pieces = append(game.Pieces, &Piece{5, 7, false, 'w', 'b'})
	game.Pieces = append(game.Pieces, &Piece{6, 7, false, 'w', 'n'})
	game.Pieces = append(game.Pieces, &Piece{7, 7, false, 'w', 'r'})

	return &game
}

func (g *Game) Board () Board {
	board := Board{}

	for _, piece := range g.Pieces {
		if !piece.Captured {
			board[piece.Y][piece.X] = piece
		}
	}

	return board;
}

func (g *Game) Piece (rank string, file string) *Piece {
	x, y := rankFileToXY(rank, file)

	for _, piece := range g.Pieces {
		if piece.X == x && piece.Y == byte(y) {
			return piece
		}
	}

	return nil
}

func (g *Game) Move (srcRank string, srcFile string, dstRank string, dstFile string) error {
	moving := g.Piece(srcRank, srcFile)
	if moving == nil {
		return errors.New("No Piece Found there!")
	}

	if srcRank == dstRank && srcFile == dstFile {
		return errors.New("Invalid Move, same space!")
	}

	if moving.Side != g.Side {
		return errors.New("Invalid Move, it's not your turn!")
	}

	target := g.Piece(dstRank, dstFile)
	if target != nil {
		if target.Side == g.Side {
			return errors.New("Invalid Move, can't take your own pieces!")
		}
		target.Capture()
	}

	moving.Move(dstRank, dstFile)

	if g.Side == 'w' {
		g.Side = 'b'
	} else {
		g.Side = 'w'
	}

	return nil
}

func (g *Game) DisplaySlack() string {
	ret := "" // TODO - find equivalent in go of ByteBuffer
	file := 8
	for _, row := range g.Board() {
		ret += strconv.Itoa(file)
		rank := 1
		for _, piece := range row {
			var sq string
			color := "w"
			if (rank % 2 == 0 && file % 2 == 0) || (rank % 2 == 1 && file % 2 == 1) {
				color = "b"
			}
			if piece != nil {
				sq = fmt.Sprintf(":chess-%c-%c-%s:", piece.Side, piece.Piece, color)
			} else {
				sq = fmt.Sprintf(":chess-%s:", color)
			}
			ret += sq
			rank++
		}
		ret += "\n"
		file--
	}
	ret += "     A   B   C   D   E    F   G   H"

	return ret
}

func (g *Game) DisplayText() string {
	ret := "\n" // TODO - find equivalent in go of ByteBuffer
	file := 8
	for _, row := range g.Board() {
		ret += strconv.Itoa(file)
		rank := 1
		for _, piece := range row {
			if piece != nil {
				ret += pieceSymbols[fmt.Sprintf("%c%c", piece.Side, piece.Piece)]
			} else {
				if (rank % 2 == 0 && file % 2 == 0) || (rank % 2 == 1 && file % 2 == 1) {
					ret += " "
				} else {
					ret += "#"
				}
			}
			rank++
		}
		ret += "\n"
		file--
	}
	ret += " ABCDEFGH"

	return ret
}