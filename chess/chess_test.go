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
	srcRankFile string
	dstRankFile string
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
			testee.Move("A2", "A4")
			Expect(testee.Piece("A2")).To(BeNil())
			Expect(testee.Piece("A4").String()).To(Equal("WhitePawn"))
		})
		It("Captures piece", func() {
			sacrifice := testee.Piece("B7")
			testee.Move("A2", "A4")
			testee.Move("B7", "B5")
			testee.Move("A4", "B5")
			Expect(sacrifice.String()).To(Equal("BlackPawn(Captured)"))
			Expect(testee.Piece("B5").String()).To(Equal("WhitePawn"))
		})
		It("Errors if capture piece is same side", func() {
			ret := testee.Move("C1", "B2")
			Expect(ret).To(Equal(errors.New("Invalid Move, can't take your own pieces!")))
		})
		It("Errors if no piece in start", func() {
			ret := testee.Move("A5", "A6")
			Expect(ret).To(Equal(errors.New("No Piece Found there!")))
		})
		It("Errors if source and destination are the same", func() {
			ret := testee.Move("A2", "A2")
			Expect(ret).To(Equal(errors.New("Invalid Move, same space!")))
		})
		It("Prevents moves on wrong turn", func() {
			ret := testee.Move("A7", "A5")
			Expect(ret).To(Equal(errors.New("Invalid Move, it's not your turn!")))
			ret = testee.Move("A2", "A4")
			Expect(ret).To(BeNil())
			ret = testee.Move("B2", "B4")
			Expect(ret).To(Equal(errors.New("Invalid Move, it's not your turn!")))
		})
		It("Alternates the side at play", func() {
			Expect(testee.Side).To(Equal(White))
			testee.Move("A2", "A4")
			Expect(testee.Side).To(Equal(Black))
			testee.Move("B7", "B5")
			Expect(testee.Side).To(Equal(White))
			testee.Move("A4", "B5")
			Expect(testee.Side).To(Equal(Black))
		})
		It("Increments full move clock", func() {
			testee.Move("A2", "A4")
			Expect(testee.FullMoveClock).To(Equal(2))
		})
		It("Increments half move clock", func() {
			testee.Move("B1", "C3")
			Expect(testee.HalfMoveClock).To(Equal(1))
		})
		It("Resets half move clock on pawn move", func() {
			testee.Move("B1", "C3")
			Expect(testee.HalfMoveClock).To(Equal(1))
			testee.Move("B7", "B5")
			Expect(testee.HalfMoveClock).To(Equal(0))
		})
		It("Resets half move clock on capture", func() {
			testee.Move("B2", "B4")
			Expect(testee.HalfMoveClock).To(Equal(0))
			testee.Move("B8", "C6")
			Expect(testee.HalfMoveClock).To(Equal(1))
			testee.Move("B1", "C3")
			Expect(testee.HalfMoveClock).To(Equal(2))
			testee.Move("C6", "B4")
			Expect(testee.HalfMoveClock).To(Equal(0))
		})
		PIt("Sets EnPassant square", func() {})
	})

	Describe("ValidateMove", func() {
		Context("Pawns", func() {
			BeforeEach(func() {
				testee = LoadFENGame("8/4p3/8/2pp4/2PP4/8/4P3/8 w KQkq - 0 1")
				// fmt.Println(game.DisplayText())
			})
			// TODO - Test First move vs other moves for pawn charge
			allowed := map[string]testMove{
				"Black Allows forward movement by 1": testMove{"E7", "E6"},
				"Black Allows forward movement by 2": testMove{"E7", "E5"},
				"Black Allows Capturing diagnally":   testMove{"D5", "C4"},

				"White Allows forward movement by 1": testMove{"E2", "E3"},
				"White Allows forward movement by 2": testMove{"E2", "E4"},
				"White Allows Capturing diagnally":   testMove{"C4", "D5"},
			}
			disallowed := map[string]testMove{
				"Black Disallows forward movement by 3": testMove{"E7", "E4"},
				"Black Disallows forward diagnal":       testMove{"E7", "F6"},
				"Black Disallows backward diagnal":      testMove{"E7", "F8"},
				"Black Disallows backward movement":     testMove{"E7", "E8"},
				"Black Disallows capturing forward":     testMove{"C5", "C4"},

				"White Disallows forward movement by 3": testMove{"E2", "E5"},
				"White Disallows forward diagnal":       testMove{"E2", "F3"},
				"White Disallows backward diagnal":      testMove{"E2", "F1"},
				"White Disallows backward movement":     testMove{"E2", "E1"},
				"White Disallows Capturing forward":     testMove{"C4", "C5"},
			}
			for name, move := range allowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcRankFile, move.dstRankFile)).To(BeNil())
					})
				}(move)
			}
			for name, move := range disallowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcRankFile, move.dstRankFile)).To(Not(BeNil()))
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
				"Allows move up":         testMove{"E5", "E6"},
				"Allows move up-right":   testMove{"E5", "F6"},
				"Allows move right":      testMove{"E5", "F5"},
				"Allows move down-right": testMove{"E5", "F4"},
				"Allows move down":       testMove{"E5", "E4"},
				"Allows move down-left":  testMove{"E5", "D4"},
				"Allows move left":       testMove{"E5", "D5"},
				"Allows move up-left":    testMove{"E5", "D6"},
			}
			disallowed := map[string]testMove{
				"Disallows move up by 2":         testMove{"E5", "E7"},
				"Disallows move up-right by 2":   testMove{"E5", "G7"},
				"Disallows move right by 2":      testMove{"E5", "G5"},
				"Disallows move down-right by 2": testMove{"E5", "G3"},
				"Disallows move down by 2":       testMove{"E5", "E3"},
				"Disallows move down-left by 2":  testMove{"E5", "C3"},
				"Disallows move left by 2":       testMove{"E5", "C5"},
				"Disallows move up-left by 2":    testMove{"E5", "C7"},
			}
			for name, move := range allowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcRankFile, move.dstRankFile)).To(BeNil())
					})
				}(move)
			}
			for name, move := range disallowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcRankFile, move.dstRankFile)).To(Not(BeNil()))
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
				"Allows move up":         testMove{"E5", "E6"},
				"Allows move up by 2":    testMove{"E5", "E7"},
				"Allows move right":      testMove{"E5", "F5"},
				"Allows move right by 2": testMove{"E5", "G5"},
				"Allows move down":       testMove{"E5", "E4"},
				"Allows move down by 2":  testMove{"E5", "E3"},
				"Allows move left":       testMove{"E5", "D5"},
				"Allows move left by 2":  testMove{"E5", "C5"},
			}
			disallowed := map[string]testMove{
				"Disallows move up-right by 2":   testMove{"E5", "G7"},
				"Disallows move down-right by 2": testMove{"E5", "G3"},
				"Disallows move down-left by 2":  testMove{"E5", "C3"},
				"Disallows move up-left by 2":    testMove{"E5", "C7"},
				"Disallows move down-right":      testMove{"E5", "F4"},
				"Disallows move down-left":       testMove{"E5", "D4"},
				"Disallows move up-left":         testMove{"E5", "D6"},
			}
			for name, move := range allowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcRankFile, move.dstRankFile)).To(BeNil())
					})
				}(move)
			}
			for name, move := range disallowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcRankFile, move.dstRankFile)).To(Not(BeNil()))
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
				"Allows move up":              testMove{"E5", "E6"},
				"Allows move up by 2":         testMove{"E5", "E7"},
				"Allows move right":           testMove{"E5", "F5"},
				"Allows move right by 2":      testMove{"E5", "G5"},
				"Allows move down":            testMove{"E5", "E4"},
				"Allows move down by 2":       testMove{"E5", "E3"},
				"Allows move left":            testMove{"E5", "D5"},
				"Allows move left by 2":       testMove{"E5", "C5"},
				"Allows move up-right by 2":   testMove{"E5", "G7"},
				"Allows move down-right by 2": testMove{"E5", "G3"},
				"Allows move down-left by 2":  testMove{"E5", "C3"},
				"Allows move up-left by 2":    testMove{"E5", "C7"},
				"Allows move down-right":      testMove{"E5", "F4"},
				"Allows move down-left":       testMove{"E5", "D4"},
				"Allows move up-left":         testMove{"E5", "D6"},
			}
			disallowed := map[string]testMove{
			// TODO negative case, knight like?
			}
			for name, move := range allowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcRankFile, move.dstRankFile)).To(BeNil())
					})
				}(move)
			}
			for name, move := range disallowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcRankFile, move.dstRankFile)).To(Not(BeNil()))
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
				"Allows move up-right by 2":   testMove{"E5", "G7"},
				"Allows move down-right by 2": testMove{"E5", "G3"},
				"Allows move down-left by 2":  testMove{"E5", "C3"},
				"Allows move up-left by 2":    testMove{"E5", "C7"},
				"Allows move down-right":      testMove{"E5", "F4"},
				"Allows move down-left":       testMove{"E5", "D4"},
				"Allows move up-left":         testMove{"E5", "D6"},
			}
			disallowed := map[string]testMove{
				"Disallows move up":         testMove{"E5", "E6"},
				"Disallows move up by 2":    testMove{"E5", "E7"},
				"Disallows move right":      testMove{"E5", "F5"},
				"Disallows move right by 2": testMove{"E5", "G5"},
				"Disallows move down":       testMove{"E5", "E4"},
				"Disallows move down by 2":  testMove{"E5", "E3"},
				"Disallows move left":       testMove{"E5", "D5"},
				"Disallows move left by 2":  testMove{"E5", "C5"},
			}
			for name, move := range allowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcRankFile, move.dstRankFile)).To(BeNil())
					})
				}(move)
			}
			for name, move := range disallowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcRankFile, move.dstRankFile)).To(Not(BeNil()))
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
				"Disallows move up-right by 2":   testMove{"E5", "G7"},
				"Disallows move down-right by 2": testMove{"E5", "G3"},
				"Disallows move down-left by 2":  testMove{"E5", "C3"},
				"Disallows move up-left by 2":    testMove{"E5", "C7"},
				"Disallows move down-right":      testMove{"E5", "F4"},
				"Disallows move down-left":       testMove{"E5", "D4"},
				"Disallows move up-left":         testMove{"E5", "D6"},
				"Disallows move up":              testMove{"E5", "E6"},
				"Disallows move up by 2":         testMove{"E5", "E7"},
				"Disallows move right":           testMove{"E5", "F5"},
				"Disallows move right by 2":      testMove{"E5", "G5"},
				"Disallows move down":            testMove{"E5", "E4"},
				"Disallows move down by 2":       testMove{"E5", "E3"},
				"Disallows move left":            testMove{"E5", "D5"},
				"Disallows move left by 2":       testMove{"E5", "C5"},
			}
			for name, move := range allowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcRankFile, move.dstRankFile)).To(BeNil())
					})
				}(move)
			}
			for name, move := range disallowed {
				func(move testMove) {
					It(name, func() {
						Expect(testee.ValidateMove(move.srcRankFile, move.dstRankFile)).To(Not(BeNil()))
					})
				}(move)
			}
		})
	})
})
