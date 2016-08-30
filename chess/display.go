package chess

import (
	"fmt"
	"strconv"
)

func (g *Game) DisplaySlack() string {
	ret := "" // TODO - find equivalent in go of ByteBuffer
	file := 8
	for _, row := range g.Board() {
		ret += ":chess-" + strconv.Itoa(file) + ":"
		rank := 1
		for _, piece := range row {
			var sq string
			color := "w"
			if (rank%2 == 0 && file%2 == 0) || (rank%2 == 1 && file%2 == 1) {
				color = "b"
			}
			if piece != nil {
				sq = fmt.Sprintf(":chess-%c-%c-%s:", piece.side, piece.piece, color)
			} else {
				sq = fmt.Sprintf(":chess-%s:", color)
			}
			ret += sq
			rank++
		}
		ret += "\n"
		file--
	}
	ret += ":chess-w::chess-ca::chess-cb::chess-cc::chess-cd::chess-ce::chess-cf::chess-cg::chess-ch:\n"

	if g.Side == White {
		ret += "White's move!"
	} else {
		ret += "Black's move!"
	}

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
				ret += piece.Symbol()
			} else {
				if (rank%2 == 0 && file%2 == 0) || (rank%2 == 1 && file%2 == 1) {
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
