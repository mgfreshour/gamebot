package chess_test

import (
	. "github.com/mgfreshour/gamebot/chess"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testMove struct {
	srcFile string
	srcRank string
	dstFile string 
	dstRank string 
}

var _ = Describe("Piece", func() {
	var game *Game
	// var testee *Piece
	Describe("ValidateMove", func() {
		Context("Black Pawns", func() {
			BeforeEach(func() {
				game = LoadFENGame("8/4p3/8/8/8/8/4P3/8 w KQkq - 0 1")
				// fmt.Println(game.DisplayText()) 
			})
			Context("First Move", func() {
				allowed := map[string]testMove{
					"Allows forward movement by 1": testMove{ "E", "7", "E", "6" },
					"Allows forward movement by 2": testMove{ "E", "7", "E", "5" },
					//"Allows Capturing diagnally":
				}
				disallowed := map[string]testMove{
					"Disallows forward movement by 3": testMove{ "E", "7", "E", "4" },
					"Disallows forward diagnal": testMove{ "E", "7", "F", "6" },
					"Disallows backward diagnal": testMove{ "E", "7", "F", "8" },
					"Disallows backward movement": testMove{ "E", "7", "E", "8" },
					//"Disallows Capturing forward":
				}
				for name, move := range allowed {
					func(move testMove) {
						It(name, func() {
							testee := game.Piece(move.srcFile, move.srcRank)
							target := game.Piece(move.dstFile, move.dstRank)
							Expect(testee.ValidateMove(move.dstFile, move.dstRank, target)).To(BeNil())
						})
					}(move)
				}
				for name, move := range disallowed {
					func(move testMove) {
						It(name, func() {
							testee := game.Piece(move.srcFile, move.srcRank)
							target := game.Piece(move.dstFile, move.dstRank)
							Expect(testee.ValidateMove(move.dstFile, move.dstRank, target)).To(Not(BeNil()))
						})
					}(move)
				}
			})
		})
	})
})
