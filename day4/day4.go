package main

import (
	"advent24"
	"strings"
)

func main() {
	input := advent24.ReadInputOfDay(4)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	text := strings.Split(input, ";\n")
	answer := 0
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			if text[i][j] == 'X' {
				// Going down
				if safeGetChar(text, i+1, j) == 'M' && safeGetChar(text, i+2, j) == 'A' && safeGetChar(text, i+3, j) == 'S' {
					answer++
				}

				// Going up
				if safeGetChar(text, i-1, j) == 'M' && safeGetChar(text, i-2, j) == 'A' && safeGetChar(text, i-3, j) == 'S' {
					answer++
				}

				// Going right
				if safeGetChar(text, i, j+1) == 'M' && safeGetChar(text, i, j+2) == 'A' && safeGetChar(text, i, j+3) == 'S' {
					answer++
				}

				// Going left
				if safeGetChar(text, i, j-1) == 'M' && safeGetChar(text, i, j-2) == 'A' && safeGetChar(text, i, j-3) == 'S' {
					answer++
				}

				// Going down-right
				if safeGetChar(text, i+1, j+1) == 'M' && safeGetChar(text, i+2, j+2) == 'A' && safeGetChar(text, i+3, j+3) == 'S' {
					answer++
				}

				// Going down-left
				if safeGetChar(text, i+1, j-1) == 'M' && safeGetChar(text, i+2, j-2) == 'A' && safeGetChar(text, i+3, j-3) == 'S' {
					answer++
				}

				// Going up-right
				if safeGetChar(text, i-1, j+1) == 'M' && safeGetChar(text, i-2, j+2) == 'A' && safeGetChar(text, i-3, j+3) == 'S' {
					answer++
				}

				// Going up-left
				if safeGetChar(text, i-1, j-1) == 'M' && safeGetChar(text, i-2, j-2) == 'A' && safeGetChar(text, i-3, j-3) == 'S' {
					answer++
				}
			}
		}
	}
	println(answer)
}

func partTwo(input string) {
	text := strings.Split(input, ";\n")
	answer := 0
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			if text[i][j] == 'S' {
				if safeGetChar(text, i+1, j+1) == 'A' && safeGetChar(text, i+2, j+2) == 'M' {
					if safeGetChar(text, i, j+2) == 'S' && safeGetChar(text, i+2, j) == 'M' {
						answer++
					} else if safeGetChar(text, i, j+2) == 'M' && safeGetChar(text, i+2, j) == 'S' {
						answer++
					}
				}
			} else if text[i][j] == 'M' {
				if safeGetChar(text, i+1, j+1) == 'A' && safeGetChar(text, i+2, j+2) == 'S' {
					if safeGetChar(text, i, j+2) == 'S' && safeGetChar(text, i+2, j) == 'M' {
						answer++
					} else if safeGetChar(text, i, j+2) == 'M' && safeGetChar(text, i+2, j) == 'S' {
						answer++
					}
				}
			}
		}
	}

	println(answer)
}

func safeGetChar(text []string, i, j int) byte {
	if i < 0 || i >= len(text) || j < 0 || j >= len(text[i]) {
		return 'Z'
	}
	return text[i][j]
}
