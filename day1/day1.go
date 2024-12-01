package main

import (
	"advent24"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := advent24.ReadInputOfDay(1)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	lines := strings.Split(string(input), "\n")
	size := len(lines)
	leftSide := make([]int, size)
	rightSide := make([]int, size)
	for i := 0; i < size; i++ {
		numbers := strings.Split(lines[i], "   ")
		num, _ := strconv.Atoi(numbers[0])
		leftSide[i] = num
		num, _ = strconv.Atoi(numbers[1])
		rightSide[i] = num
	}

	sort.Ints(leftSide)
	sort.Ints(rightSide)
	result := 0
	for i := 0; i < size; i++ {
		result += advent24.IntAbs(leftSide[i] - rightSide[i])
	}
	println(result)
}

func partTwo(input string) {
	lines := strings.Split(string(input), "\n")
	size := len(lines)
	leftSide := make([]int, size)
	rightSide := make(map[int]int)

	for i := 0; i < size; i++ {
		numbers := strings.Split(lines[i], "   ")
		num, _ := strconv.Atoi(numbers[0])
		leftSide[i] = num
		num, _ = strconv.Atoi(numbers[1])
		rightSide[num]++
	}
	var result int
	for i := 0; i < size; i++ {
		result += leftSide[i] * rightSide[leftSide[i]]
	}
	println(result)
}
