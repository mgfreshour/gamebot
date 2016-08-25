package chess_test

import (
	. "github.com/mgfreshour/gamebot/chess"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"strings"
)

var _ = Describe("FEN", func() {
	Describe("LoadFENGame", func() {
		It("Loads board positions", func() {
			ret := LoadFENGame("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1").DisplayText()
			expected := []string{"",
				"8♜♞♝♚♛♝♞♜",
				"7♟♟♟♟♟♟♟♟",
				"6# # # # ",
				"5 # # # #",
				"4# # # # ",
				"3 # # # #",
				"2♙♙♙♙♙♙♙♙",
				"1♖♘♗♕♔♗♘♖",
				" ABCDEFGH"}
			y := 0
			for _, line := range strings.Split(ret, "\n") {
				for x := 0; x < len(line); x++ {
					actual := line[x]
					Expect(actual).To(Equal(expected[y][x]), "Expected '%c' == '%c' line %d, pos %d",
						line[x], expected[y][x], y, x)
				}
				y++
				if y >= 8 {
					break
				}
			}
		})
		It("Loads white turn", func() {
			game := LoadFENGame("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
			Expect(game.Side).To(Equal(White))
		})
		It("Loads black turn", func() {
			game := LoadFENGame("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1")
			Expect(game.Side).To(Equal(Black))
		})
		// PIt("Loads available castling", func() {})
		// PIt("Loads available castling", func() {})
		// PIt("Loads En passant available", func() {})
		// PIt("Loads En passant non-available", func() {})
		// PIt("Loads half-move clock", func() {})
		// PIt("Loads full-move clock", func() {})
	})

	Describe("SaveFENGame", func() {
		It("Saves what was loaded", func() {
			tests := []string{
				"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
				// "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
				// "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2",
				// "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2",
			}
			for _, line := range tests {
				game := LoadFENGame(line)
				Expect(SaveFENGame(game)).To(Equal(line))
			}
		})
	})
})
