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
})
