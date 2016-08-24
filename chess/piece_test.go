package chess_test

import (
	. "github.com/mgfreshour/gamebot/chess"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "fmt"
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
		Context("Pawns", func() {
			BeforeEach(func() {
				game = LoadFENGame("8/4p3/8/2pp4/2PP4/8/4P3/8 w KQkq - 0 1")
				// fmt.Println(game.DisplayText())
			})
			Context("First Move", func() {
				// TODO - this actually doesn't test first move as I don't keep that yet.
				allowed := map[string]testMove{
					"Black Allows forward movement by 1": testMove{ "E", "7", "E", "6" },
					"Black Allows forward movement by 2": testMove{ "E", "7", "E", "5" },
					"Black Allows Capturing diagnally": testMove{"D", "5", "C", "4"},

					"White Allows forward movement by 1": testMove{ "E", "2", "E", "3" },
					"White Allows forward movement by 2": testMove{ "E", "2", "E", "4" },
					"White Allows Capturing diagnally": testMove{"C", "4", "D", "5"},
				}
				disallowed := map[string]testMove{
					"Black Disallows forward movement by 3": testMove{ "E", "7", "E", "4" },
					"Black Disallows forward diagnal": testMove{ "E", "7", "F", "6" },
					"Black Disallows backward diagnal": testMove{ "E", "7", "F", "8" },
					"Black Disallows backward movement": testMove{ "E", "7", "E", "8" },
					"Black Disallows capturing forward": testMove{ "C", "5", "C", "4" },

					"White Disallows forward movement by 3": testMove{ "E", "2", "E", "5" },
					"White Disallows forward diagnal": testMove{ "E", "2", "F", "3" },
					"White Disallows backward diagnal": testMove{ "E", "2", "F", "1" },
					"White Disallows backward movement": testMove{ "E", "2", "E", "1" },
					"White Disallows Capturing forward": testMove{"C", "4", "C", "5"},
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
		Context("Kings", func() {
			BeforeEach(func() {
				game = LoadFENGame("8/8/8/4k3/8/8/8/8 w KQkq - 0 1")
				// fmt.Println(game.DisplayText())
			})
			allowed := map[string]testMove{
				"Allows move up": testMove{"E", "5", "E", "6"},
				"Allows move up-right": testMove{"E", "5", "F", "6"},
				"Allows move right": testMove{"E", "5", "F", "5"},
				"Allows move down-right": testMove{"E", "5", "F", "4"},
				"Allows move down": testMove{"E", "5", "E", "4"},
				"Allows move down-left": testMove{"E", "5", "D", "4"},
				"Allows move left": testMove{"E", "5", "D", "5"},
				"Allows move up-left": testMove{"E", "5", "D", "6"},
			}
			disallowed := map[string]testMove{
				"Allows move up by 2": testMove{"E", "5", "E", "7"},
				"Allows move up-right by 2": testMove{"E", "5", "G", "7"},
				"Allows move right by 2": testMove{"E", "5", "G", "5"},
				"Allows move down-right by 2": testMove{"E", "5", "G", "3"},
				"Allows move down by 2": testMove{"E", "5", "E", "3"},
				"Allows move down-left by 2": testMove{"E", "5", "C", "3"},
				"Allows move left by 2": testMove{"E", "5", "C", "5"},
				"Allows move up-left by 2": testMove{"E", "5", "C", "7"},
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
