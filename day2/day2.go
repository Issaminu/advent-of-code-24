package main

import (
	"advent24"
	"strconv"
	"strings"
)

func main() {
	input := advent24.ReadInputOfDay(2)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	lines := strings.Split(string(input), "\n")
	size := len(lines)
	levels := make([][]int, size)

	for i := 0; i < size; i++ {
		numbersStr := strings.Split(lines[i], " ")
		var level []int
		for i := 0; i < len(numbersStr); i++ {
			num, _ := strconv.Atoi(numbersStr[i])
			level = append(level, num)

		}
		levels[i] = level
	}
	answer := 0
	for i := 0; i < len(levels); i++ {
		isAscending := false
		level := levels[i]
		for j := 1; j < len(level); j++ {
			if level[0] < level[1] {
				isAscending = true
			}

			delta := advent24.IntAbs(level[j] - level[j-1])

			if (isAscending && level[j] < level[j-1]) || (!isAscending && level[j] > level[j-1]) || (delta < 1 || delta > 3) {
				break
			}
			if j == len(level)-1 {
				answer++
			}
		}
	}
	println(answer)
}

func partTwo(input string) {
	lines := strings.Split(string(input), "\n")
	size := len(lines)
	levels := make([][]int, size)

	for i := 0; i < size; i++ {
		numbersStr := strings.Split(lines[i], " ")
		var level []int
		for i := 0; i < len(numbersStr); i++ {
			num, _ := strconv.Atoi(numbersStr[i])
			level = append(level, num)

		}
		levels[i] = level
	}
	answer := 0
	for i := 0; i < size; i++ {
		level := levels[i]
		if len(level) < 2 {
			answer++
			continue
		}
		isSafe := false
		for skipIdx := 0; skipIdx < len(level); skipIdx++ {
			isAscending := false
			isValidSequence := true
			if skipIdx == 0 {
				isAscending = level[1] < level[2]
			} else if skipIdx == 1 {
				isAscending = level[0] < level[2]
			} else {
				isAscending = level[0] < level[1]
			}

			for k := 0; k < len(level)-1; k++ {
				if skipIdx == k {
					continue
				}

				next := k + 1
				if next == skipIdx {
					if next == len(level)-1 {
						break
					} else {
						next = k + 2
					}
				}

				delta := advent24.IntAbs(level[k] - level[next])

				if (isAscending && level[k] > level[next]) || (!isAscending && level[k] < level[next]) || (delta < 1 || delta > 3) {
					isValidSequence = false
					break
				}
			}
			if isValidSequence {
				isSafe = true
				break
			}
		}
		if isSafe {
			answer++
		}
	}
	println(answer)
}
