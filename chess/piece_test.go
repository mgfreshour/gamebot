package chess_test

import (
	. "github.com/mgfreshour/gamebot/chess"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
)

var _ = Describe("Piece", func() {
	var game *Game
	var testee *Piece
	Describe("ValidateMove", func() {
		Context("Black Pawns", func() {
			BeforeEach(func() {
				game = LoadFENGame("8/4p3/8/8/8/8/4P3/8 w KQkq - 0 1")
				fmt.Println(game.DisplayText()) 
			})
			It("Allows first forward movement by 1", func() {
				testee = game.Piece("E", "7")
				Expect(testee.ValidateMove("E", "6", nil)).To(BeNil())
			})
			It("Allows first forward movement by 2", func() {
				testee = game.Piece("E", "7")
				Expect(testee.ValidateMove("E", "5", nil)).To(BeNil())
			})
			It("Disallows first forward movement by 3", func() {
				testee = game.Piece("E", "7")
				Expect(testee.ValidateMove("E", "4", nil)).To(Not(BeNil()))
			})
		})
	})
})
