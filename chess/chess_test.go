package chess_test

import (
	. "github.com/mgfreshour/gamebot/chess"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"
	"testing"
)

func TestChess(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chess Suite")
}

type testMove struct {
	srcFile string
	srcRank string
	dstFile string
	dstRank string
}

var _ = Describe("Chess", func() {
	var testee *Game

	BeforeEach(func() {
		testee = NewGame()
	})
	AfterEach(func() {
		//fmt.Println(testee.DisplayText())
	})

	Describe("Game Creation", func() {
		It("creates a board", func() {
			// Black Ranks
			for rank := 0; rank < 2; rank++ {
				for file := 0; file < 7; file++ {
					Expect(testee.Board()[rank][file].Side()).To(Equal(Black))
				}
			}
			// White Ranks
			for rank := 6; rank < 7; rank++ {
				for file := 0; file < 7; file++ {
					Expect(testee.Board()[rank][file].Side()).To(Equal(White))
				}
			}
		})
		It("Set white's turn", func() {
			Expect(testee.Side).To(Equal(White))
		})
	})

	Describe("Move", func() {
		It("Moves piece", func() {
			testee.Move("A", "2", "A", "4")
			Expect(testee.Piece("A", "2")).To(BeNil())
			Expect(testee.Piece("A", "4").String()).To(Equal("WhitePawn"))
		})
		It("Captures piece", func() {
			sacrifice := testee.Piece("B", "7")
			testee.Move("A", "2", "A", "4")
			testee.Move("B", "7", "B", "5")
			testee.Move("A", "4", "B", "5")
			Expect(sacrifice.String()).To(Equal("BlackPawn(Captured)"))
			Expect(testee.Piece("B", "5").String()).To(Equal("WhitePawn"))
		})
		It("Errors if capture piece is same side", func() {
			ret := testee.Move("C", "1", "B", "2")
			Expect(ret).To(Equal(errors.New("Invalid Move, can't take your own pieces!")))
		})
		It("Errors if no piece in start", func() {
			ret := testee.Move("A", "5", "A", "6")
			Expect(ret).To(Equal(errors.New("No Piece Found there!")))
		})
		It("Errors if source and destination are the same", func() {
			ret := testee.Move("A", "2", "A", "2")
			Expect(ret).To(Equal(errors.New("Invalid Move, same space!")))
		})
		It("Prevents moves on wrong turn", func() {
			ret := testee.Move("A", "7", "A", "5")
			Expect(ret).To(Equal(errors.New("Invalid Move, it's not your turn!")))
			ret = testee.Move("A", "2", "A", "4")
			Expect(ret).To(BeNil())
			ret = testee.Move("B", "2", "B", "4")
			Expect(ret).To(Equal(errors.New("Invalid Move, it's not your turn!")))
		})
		It("Alternates the side at play", func() {
			Expect(testee.Side).To(Equal(White))
			testee.Move("A", "2", "A", "4")
			Expect(testee.Side).To(Equal(Black))
			testee.Move("B", "7", "B", "5")
			Expect(testee.Side).To(Equal(White))
			testee.Move("A", "4", "B", "5")
			Expect(testee.Side).To(Equal(Black))
		})
	})

	Describe("ValidateMove", func() {
		Context("Pawns", func() {
			BeforeEach(func() {
				testee = LoadFENGame("8/4p3/8/2pp4/2PP4/8/4P3/8 w KQkq - 0 1")
				// fmt.Println(game.DisplayText())
			})
			// TODO - Test First move vs other moves for pawn charge
			allowed := map[string]testMove{
				"Black Allows forward movement by 1": testMove{"E", "7", "E", "6"},
				"Black Allows forward movement by 2": testMove{"E", "7", "E", "5"},
				"Black Allows Capturing diagnally":   testMove{"D", "5", "C", "4"},

				"White Allows forward movement by 1": testMove{"E", "2", "E", "3"},
				"White Allows forward movement by 2": testMove{"E", "2", "E", "4"},
				"White Allows Capturing diagnally":   testMove{"C", "4", "D", "5"},
			}
			disallowed := map[string]testMove{
				"Black Disallows forward movement by 3": testMove{"E", "7", "E", "4"},
				"Black Disallows forward diagnal":       testMove{"E", "7", "F", "6"},
				"Black Disallows backward diagnal":      testMove{"E", "7", "F", "8"},
				"Black Disallows backward movement":     testMove{"E", "7", "E", "8"},
				"Black Disallows capturing forward":     testMove{"C", "5", "C", "4"},

				"White Disallows forward movement by 3": testMove{"E", "2", "E", "5"},
				"White Disallows forward diagnal":       testMove{"E", "2", "F", "3"},
				"White Disallows backward diagnal":      testMove{"E", "2", "F", "1"},
				"White Disallows backward movement":     testMove{"E", "2", "E", "1"},
				"White Disallows Capturing forward":     testMove{"C", "4", "C", "5"},
			}
			for name, move := range allowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcFile, move.srcRank, move.dstFile, move.dstRank)).To(BeNil())
					})
				}(move)
			}
			for name, move := range disallowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcFile, move.srcRank, move.dstFile, move.dstRank)).To(Not(BeNil()))
					})
				}(move)
			}
		})
		Context("Kings", func() {
			BeforeEach(func() {
				testee = LoadFENGame("8/8/8/4k3/8/8/8/8 w KQkq - 0 1")
				// fmt.Println(game.DisplayText())
			})
			allowed := map[string]testMove{
				"Allows move up":         testMove{"E", "5", "E", "6"},
				"Allows move up-right":   testMove{"E", "5", "F", "6"},
				"Allows move right":      testMove{"E", "5", "F", "5"},
				"Allows move down-right": testMove{"E", "5", "F", "4"},
				"Allows move down":       testMove{"E", "5", "E", "4"},
				"Allows move down-left":  testMove{"E", "5", "D", "4"},
				"Allows move left":       testMove{"E", "5", "D", "5"},
				"Allows move up-left":    testMove{"E", "5", "D", "6"},
			}
			disallowed := map[string]testMove{
				"Disallows move up by 2":         testMove{"E", "5", "E", "7"},
				"Disallows move up-right by 2":   testMove{"E", "5", "G", "7"},
				"Disallows move right by 2":      testMove{"E", "5", "G", "5"},
				"Disallows move down-right by 2": testMove{"E", "5", "G", "3"},
				"Disallows move down by 2":       testMove{"E", "5", "E", "3"},
				"Disallows move down-left by 2":  testMove{"E", "5", "C", "3"},
				"Disallows move left by 2":       testMove{"E", "5", "C", "5"},
				"Disallows move up-left by 2":    testMove{"E", "5", "C", "7"},
			}
			for name, move := range allowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcFile, move.srcRank, move.dstFile, move.dstRank)).To(BeNil())
					})
				}(move)
			}
			for name, move := range disallowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcFile, move.srcRank, move.dstFile, move.dstRank)).To(Not(BeNil()))
					})
				}(move)
			}
		})
		Context("Rooks", func() {
			BeforeEach(func() {
				testee = LoadFENGame("8/8/8/4r3/8/8/8/8 w KQkq - 0 1")
				// fmt.Println(game.DisplayText())
			})
			allowed := map[string]testMove{
				"Allows move up":         testMove{"E", "5", "E", "6"},
				"Allows move up by 2":    testMove{"E", "5", "E", "7"},
				"Allows move right":      testMove{"E", "5", "F", "5"},
				"Allows move right by 2": testMove{"E", "5", "G", "5"},
				"Allows move down":       testMove{"E", "5", "E", "4"},
				"Allows move down by 2":  testMove{"E", "5", "E", "3"},
				"Allows move left":       testMove{"E", "5", "D", "5"},
				"Allows move left by 2":  testMove{"E", "5", "C", "5"},
			}
			disallowed := map[string]testMove{
				"Disallows move up-right by 2":   testMove{"E", "5", "G", "7"},
				"Disallows move down-right by 2": testMove{"E", "5", "G", "3"},
				"Disallows move down-left by 2":  testMove{"E", "5", "C", "3"},
				"Disallows move up-left by 2":    testMove{"E", "5", "C", "7"},
				"Disallows move down-right":      testMove{"E", "5", "F", "4"},
				"Disallows move down-left":       testMove{"E", "5", "D", "4"},
				"Disallows move up-left":         testMove{"E", "5", "D", "6"},
			}
			for name, move := range allowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcFile, move.srcRank, move.dstFile, move.dstRank)).To(BeNil())
					})
				}(move)
			}
			for name, move := range disallowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcFile, move.srcRank, move.dstFile, move.dstRank)).To(Not(BeNil()))
					})
				}(move)
			}
		})
		Context("Queens", func() {
			BeforeEach(func() {
				testee = LoadFENGame("8/8/8/4q3/8/8/8/8 w KQkq - 0 1")
				// fmt.Println(game.DisplayText())
			})
			allowed := map[string]testMove{
				"Allows move up":              testMove{"E", "5", "E", "6"},
				"Allows move up by 2":         testMove{"E", "5", "E", "7"},
				"Allows move right":           testMove{"E", "5", "F", "5"},
				"Allows move right by 2":      testMove{"E", "5", "G", "5"},
				"Allows move down":            testMove{"E", "5", "E", "4"},
				"Allows move down by 2":       testMove{"E", "5", "E", "3"},
				"Allows move left":            testMove{"E", "5", "D", "5"},
				"Allows move left by 2":       testMove{"E", "5", "C", "5"},
				"Allows move up-right by 2":   testMove{"E", "5", "G", "7"},
				"Allows move down-right by 2": testMove{"E", "5", "G", "3"},
				"Allows move down-left by 2":  testMove{"E", "5", "C", "3"},
				"Allows move up-left by 2":    testMove{"E", "5", "C", "7"},
				"Allows move down-right":      testMove{"E", "5", "F", "4"},
				"Allows move down-left":       testMove{"E", "5", "D", "4"},
				"Allows move up-left":         testMove{"E", "5", "D", "6"},
			}
			disallowed := map[string]testMove{
			// TODO negative case, knight like?
			}
			for name, move := range allowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcFile, move.srcRank, move.dstFile, move.dstRank)).To(BeNil())
					})
				}(move)
			}
			for name, move := range disallowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcFile, move.srcRank, move.dstFile, move.dstRank)).To(Not(BeNil()))
					})
				}(move)
			}
		})
		Context("Bishops", func() {
			BeforeEach(func() {
				testee = LoadFENGame("8/8/8/4b3/8/8/8/8 w KQkq - 0 1")
				// fmt.Println(game.DisplayText())
			})
			allowed := map[string]testMove{
				"Allows move up-right by 2":   testMove{"E", "5", "G", "7"},
				"Allows move down-right by 2": testMove{"E", "5", "G", "3"},
				"Allows move down-left by 2":  testMove{"E", "5", "C", "3"},
				"Allows move up-left by 2":    testMove{"E", "5", "C", "7"},
				"Allows move down-right":      testMove{"E", "5", "F", "4"},
				"Allows move down-left":       testMove{"E", "5", "D", "4"},
				"Allows move up-left":         testMove{"E", "5", "D", "6"},
			}
			disallowed := map[string]testMove{
				"Disallows move up":         testMove{"E", "5", "E", "6"},
				"Disallows move up by 2":    testMove{"E", "5", "E", "7"},
				"Disallows move right":      testMove{"E", "5", "F", "5"},
				"Disallows move right by 2": testMove{"E", "5", "G", "5"},
				"Disallows move down":       testMove{"E", "5", "E", "4"},
				"Disallows move down by 2":  testMove{"E", "5", "E", "3"},
				"Disallows move left":       testMove{"E", "5", "D", "5"},
				"Disallows move left by 2":  testMove{"E", "5", "C", "5"},
			}
			for name, move := range allowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcFile, move.srcRank, move.dstFile, move.dstRank)).To(BeNil())
					})
				}(move)
			}
			for name, move := range disallowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcFile, move.srcRank, move.dstFile, move.dstRank)).To(Not(BeNil()))
					})
				}(move)
			}
		})
		Context("Knights", func() {
			BeforeEach(func() {
				testee = LoadFENGame("8/8/8/4n3/8/8/8/8 w KQkq - 0 1")
				// fmt.Println(game.DisplayText())
			})
			allowed := map[string]testMove{
			// TODO positive case
			}
			disallowed := map[string]testMove{
				"Disallows move up-right by 2":   testMove{"E", "5", "G", "7"},
				"Disallows move down-right by 2": testMove{"E", "5", "G", "3"},
				"Disallows move down-left by 2":  testMove{"E", "5", "C", "3"},
				"Disallows move up-left by 2":    testMove{"E", "5", "C", "7"},
				"Disallows move down-right":      testMove{"E", "5", "F", "4"},
				"Disallows move down-left":       testMove{"E", "5", "D", "4"},
				"Disallows move up-left":         testMove{"E", "5", "D", "6"},
				"Disallows move up":              testMove{"E", "5", "E", "6"},
				"Disallows move up by 2":         testMove{"E", "5", "E", "7"},
				"Disallows move right":           testMove{"E", "5", "F", "5"},
				"Disallows move right by 2":      testMove{"E", "5", "G", "5"},
				"Disallows move down":            testMove{"E", "5", "E", "4"},
				"Disallows move down by 2":       testMove{"E", "5", "E", "3"},
				"Disallows move left":            testMove{"E", "5", "D", "5"},
				"Disallows move left by 2":       testMove{"E", "5", "C", "5"},
			}
			for name, move := range allowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcFile, move.srcRank, move.dstFile, move.dstRank)).To(BeNil())
					})
				}(move)
			}
			for name, move := range disallowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcFile, move.srcRank, move.dstFile, move.dstRank)).To(Not(BeNil()))
					})
				}(move)
			}
		})
	})
})
