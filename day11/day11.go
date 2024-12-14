package main

import (
	"advent24"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	input := advent24.ReadInputOfDay(11)
	partOne(input)
	start := time.Now()
	partTwo(input)
	end := time.Now()

	println(end.Sub(start).Milliseconds()) // ~26 ms
}

var cache = make(map[advent24.Pair[int, int]]int) // map<pair<stone_value, blinks_left>, count_of_stones_using_blinks_left>
var stones []int

func partOne(input string) { // nothing to see here, pls just go see part 2 :)
	blinks := 0
	stonesStr := strings.Split(input, " ")
	stones := make([]int, len(stonesStr))
	for i := range stones {
		stones[i], _ = strconv.Atoi(string(stonesStr[i]))
	}

	stonesCpy := make([]int, len(stonesStr))
	copy(stonesCpy, stones)
	for blinks < 25 {
		j := 0
		for i := 0; i < len(stones); i++ {
			stoneVal := strconv.Itoa(stones[i])
			length := len(stoneVal)
			if stones[i] == 0 {
				stonesCpy[j] = 1
			} else if length%2 == 0 {
				leftHalfStr := stoneVal[0 : length/2]
				rightHalfStr := stoneVal[length/2 : length]
				leftHalf, _ := strconv.Atoi(leftHalfStr)
				rightHalf, _ := strconv.Atoi(rightHalfStr)
				stonesCpy[j] = leftHalf
				stonesCpy = slices.Insert(stonesCpy, j+1, rightHalf)
				j++
			} else {
				stonesCpy[j] = stones[i] * 2024
			}

			j++
		}
		stones = make([]int, len(stonesCpy))
		copy(stones, stonesCpy)
		blinks++
	}
	println(len(stonesCpy))
}

func partTwo(input string) {
	stonesStr := strings.Split(input, " ")
	stones = make([]int, len(stonesStr))
	for i := range stones {
		stones[i], _ = strconv.Atoi(string(stonesStr[i]))
	}

	var answer int
	for i := range stones {
		answer += handleStone(stones[i], 75)
	}

	println(answer)
}

func handleStone(stone int, blinksLeft int) int {
	if blinksLeft == 0 {
		return 1
	}

	stoneIdentifier := advent24.Pair[int, int]{First: stone, Second: blinksLeft}
	if resultingStones, exists := cache[stoneIdentifier]; exists {
		return resultingStones
	}

	stoneVal := strconv.Itoa(stone)
	length := len(stoneVal)
	var numStones int
	if stone == 0 {
		numStones = handleStone(1, blinksLeft-1)
		cache[stoneIdentifier] = numStones
		return numStones
	}
	if length%2 == 0 {
		leftHalfStr := stoneVal[0 : length/2]
		rightHalfStr := stoneVal[length/2 : length]
		leftHalf, _ := strconv.Atoi(leftHalfStr)
		rightHalf, _ := strconv.Atoi(rightHalfStr)
		numStones := handleStone(leftHalf, blinksLeft-1) + handleStone(rightHalf, blinksLeft-1)
		cache[stoneIdentifier] = numStones
		return numStones
	}

	numStones = handleStone(stone*2024, blinksLeft-1)
	cache[stoneIdentifier] = numStones

	return numStones
}
