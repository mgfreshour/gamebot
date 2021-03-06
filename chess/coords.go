package chess

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var files string = "abcdefgh"

func rankFileToXY(rankFile string) (byte, byte) {
	if len(rankFile) != 2 {
		panic(fmt.Sprintf("Invalid coordinates '%v'", rankFile))
	}
	file := strings.ToLower(string(rankFile[0]))
	x := strings.Index(files, file)
	y, _ := strconv.ParseInt(string(rankFile[1]), 10, 8)
	y = int64(math.Abs(float64(y - 8)))
	if x > 7 || x < 0 || y > 7 || y < 0 {
		panic(fmt.Sprintf("Invalid conversion happened! %v to %v, %v", rankFile, x, y))
	}
	return byte(x), byte(y)
}

func xyToRankFile(x int, y int) string {
	if x > 7 || x < 0 || y > 7 || y < 0 {
		panic(fmt.Sprintf("Invalid conversion expected! %v, %v", x, y))
	}

	var r = strconv.Itoa(y + 1)
	var f = string(files[x])

	return f + r
}
