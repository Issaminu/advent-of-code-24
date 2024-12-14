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

func SafeGetElement[T any](matrix [][]T, i int, j int, nilValue T) T {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) {
		return nilValue
	}
	return matrix[i][j]
}
