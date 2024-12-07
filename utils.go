package advent24

import (
	"os"
	"strconv"
)

func ReadInputOfDay(dayNum int) string {
	file, err := os.ReadFile("./day" + strconv.Itoa(dayNum) + "/input.txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Pair[A any, B any] struct {
	First  A
	Second B
}
