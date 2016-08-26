package chess

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var files string = "ABCDEFGH"

func rankFileToXY(file string, rank string) (byte, byte) {
	x := strings.Index(files, file)
	y, _ := strconv.ParseInt(rank, 10, 8)
	y = int64(math.Abs(float64(y - 8)))
	if x > 7 || x < 0 || y > 7 || y < 0 {
		panic(fmt.Sprintf("Invalid conversion happened! %v, %v to %v, %v", file, rank, x, y))
	}
	return byte(x), byte(y)
}

func xyToRankFile(x int, y int) (string, string) {
	if x > 7 || x < 0 || y > 7 || y < 0 {
		panic(fmt.Sprintf("Invalid conversion expected! %v, %v", x, y))
	}

	var r = strconv.Itoa(y + 1)
	var f = string(files[x])

	return f, r
}
