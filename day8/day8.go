package main

import (
	"advent24"
	"strings"
	"unicode"
)

func main() {
	input := advent24.ReadInputOfDay(8)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	inputMap := strings.Split(input, "\n")
	answer := make(map[advent24.Pair[int, int]]struct{})
	for i := 0; i < len(inputMap); i++ {
		for j := 0; j < len(inputMap[i]); j++ {
			charToLookFor := inputMap[i][j]

			if !unicode.IsLetter(rune(charToLookFor)) && !unicode.IsDigit(rune(charToLookFor)) {
				continue
			}
			for a := 0; a < len(inputMap); a++ {
				for b := 0; b < len(inputMap[a]); b++ {

					if a == i && b == j {
						continue
					}

					if inputMap[a][b] == charToLookFor {
						antinodePosition := advent24.Pair[int, int]{First: i + (i - a), Second: j + (j - b)}
						_, exists := answer[antinodePosition]
						if isInBounds(inputMap, antinodePosition.First, antinodePosition.Second) && !exists {
							answer[antinodePosition] = struct{}{}
						}
					}
				}
			}
		}
	}
	println(len(answer))
}

func partTwo(input string) {
	inputMap := strings.Split(input, "\n")
	answer := make(map[advent24.Pair[int, int]]struct{})
	for i := 0; i < len(inputMap); i++ {
		for j := 0; j < len(inputMap[i]); j++ {
			charToLookFor := inputMap[i][j]

			if !unicode.IsLetter(rune(charToLookFor)) && !unicode.IsDigit(rune(charToLookFor)) {
				continue
			}
			for a := 0; a < len(inputMap); a++ {
				for b := 0; b < len(inputMap[a]); b++ {

					if a == i && b == j {
						continue
					}

					if inputMap[a][b] != charToLookFor {
						continue
					}

					answer[advent24.Pair[int, int]{First: a, Second: b}] = struct{}{}

					counter := 1
					for {
						newA, newB := i+counter*(i-a), j+counter*(j-b)
						if !isInBounds(inputMap, newA, newB) {
							break
						}

						antinodePosition := advent24.Pair[int, int]{First: newA, Second: newB}
						_, exists := answer[antinodePosition]
						if !exists {
							answer[antinodePosition] = struct{}{}
						}

						counter++
					}
				}
			}
		}
	}
	println(len(answer))
}

func isInBounds(inputMap []string, i int, j int) bool {
	return i >= 0 && i < len(inputMap) && j >= 0 && j < len(inputMap[i])
}
