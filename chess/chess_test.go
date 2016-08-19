package chess_test

import (
	. "github.com/mgfreshour/gamebot/chess"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"errors"
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
					Expect(testee.Board()[rank][file].Side).To(Equal('b'))
				}
			}
			// White Ranks
			for rank := 6; rank < 7; rank++ {
				for file := 0; file < 7; file++ {
					Expect(testee.Board()[rank][file].Side).To(Equal('w'))
				}
			}
		})
		It("Set white's turn", func () {
			Expect(testee.Side).To(Equal('w'))
		})
	})

	Describe("Move", func() {
		It("Moves piece", func () {
			testee.Move("2", "A", "4", "A")
			Expect(testee.Piece("2", "A")).To(BeNil())
			Expect(testee.Piece("4", "A").String()).To(Equal("WhitePawn"))
		})
		It("Captures piece", func () {
			sacrifice := testee.Piece("7", "B")
			testee.Move("2", "A", "4", "A")
			testee.Move("7", "B", "5", "B")
			testee.Move("4", "A", "5", "B")
			Expect(sacrifice.String()).To(Equal("BlackPawn(Captured)"))
			Expect(testee.Piece("5", "B").String()).To(Equal("WhitePawn"))
		})
		It("Errors if capture piece is same side", func () {
			ret := testee.Move("1", "C", "2", "B")
			Expect(ret).To(Equal(errors.New("Invalid Move, can't take your own pieces!")))
		})
		It("Errors if no piece in start", func () {
			ret := testee.Move("5", "A", "6", "A")
			Expect(ret).To(Equal(errors.New("No Piece Found there!")))
		})
		It("Errors if source and destination are the same", func () {
			ret := testee.Move("2", "A", "2", "A")
			Expect(ret).To(Equal(errors.New("Invalid Move, same space!")))
		})
		It("Prevents moves on wrong turn", func () {
			ret := testee.Move("7", "A", "5", "A")
			Expect(ret).To(Equal(errors.New("Invalid Move, it's not your turn!")))
			ret = testee.Move("2", "A", "4", "A")
			Expect(ret).To(BeNil())
			ret = testee.Move("2", "B", "4", "B")
			Expect(ret).To(Equal(errors.New("Invalid Move, it's not your turn!")))
		})
		It("Alternates the side at play", func () {
			Expect(testee.Side).To(Equal('w'))
			testee.Move("2", "A", "4", "A")
			Expect(testee.Side).To(Equal('b'))
			testee.Move("7", "B", "5", "B")
			Expect(testee.Side).To(Equal('w'))
			testee.Move("4", "A", "5", "B")
			Expect(testee.Side).To(Equal('b'))
		})
	})
})
