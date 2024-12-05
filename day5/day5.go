package main

import (
	"advent24"
	"strconv"
	"strings"
)

func main() {
	input := advent24.ReadInputOfDay(5)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	text := strings.Split(input, "\n")
	var separatorLineIdx int
	for index, line := range text {
		if len(line) == 0 {
			separatorLineIdx = index
			break
		}
	}
	rulesStr := text[:separatorLineIdx]
	updatesStr := text[separatorLineIdx+1:]
	updates := make([][]string, len(updatesStr))
	for i, update := range updatesStr {
		updates[i] = strings.Split(update, ",")
	}

	rules := map[string]struct{}{}
	for _, line := range rulesStr {
		rules[line] = struct{}{}
	}

	answer := 0

	for i := 0; i < len(updates); i++ {
		isCorrect := true
		for j := 0; j < len(updates[i])-1; j++ {
			if _, exists := rules[updates[i][j]+"|"+updates[i][j+1]]; !exists {
				isCorrect = false
				break
			}
		}
		if isCorrect {
			middle, _ := strconv.Atoi(updates[i][len(updates[i])/2])
			answer += middle
		}
	}
	println(answer)
}

func partTwo(input string) {
	text := strings.Split(input, "\n")
	var separatorLineIdx int
	for index, line := range text {
		if len(line) == 0 {
			separatorLineIdx = index
			break
		}
	}
	rulesStr := text[:separatorLineIdx]
	updatesStr := text[separatorLineIdx+1:]
	updates := make([][]string, len(updatesStr))
	for i, update := range updatesStr {
		updates[i] = strings.Split(update, ",")
	}

	rules := map[string]struct{}{}
	for _, line := range rulesStr {
		rules[line] = struct{}{}
	}

	answer := 0

	for i := 0; i < len(updates); i++ {
		hasErrors := false
		hasNotChanged := false
		for !hasNotChanged {
			for j := 0; j < len(updates[i])-1; j++ {
				if _, exists := rules[updates[i][j]+"|"+updates[i][j+1]]; !exists {
					hasErrors = true // Same as setting it once, it's here just to signal that this line has errors so that it counts towards answer

					// Swapping
					temp := updates[i][j]
					updates[i][j] = updates[i][j+1]
					updates[i][j+1] = temp

					hasNotChanged = false
					break
				} else {
					hasNotChanged = true
				}
			}
		}
		if hasErrors {
			middle, _ := strconv.Atoi(updates[i][len(updates[i])/2])
			answer += middle
		}
	}
	println(answer)
}
