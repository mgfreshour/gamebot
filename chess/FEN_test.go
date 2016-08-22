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
					Expect(actual).To(Equal(expected[y][x]), "Expected '%c' == '%c' line %d, pos %d", line[x], expected[y][x], y, x)
				}
				y++
				if y >= 8 {
					break
				}
			}
		})
		//It("Loads side")
	})
})
