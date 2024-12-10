package main

import (
	"advent24"
	"strconv"
)

func main() {
	input := advent24.ReadInputOfDay(9)
	partOne(input)
}

func partOne(input string) { // optimized O(n) solution, sadly it doesn't work in part 2 as that requires summing *after* finishing the full run-through
	values := make([]int, len(input))
	for i := range input {
		values[i], _ = strconv.Atoi(string(input[i]))
	}

	left := 0
	var right int
	if len(values)%2 == 0 {
		right = len(values) - 2 // skip last index if it is odd
	} else {
		right = len(values) - 1
	}
	offset := 0
	var answer uint64
	for left < len(values) {
		leftVal := values[left]
		if left%2 == 0 {
			answer += uint64((left / 2) * ((offset+leftVal)*(offset+leftVal-1)/2 - (offset-1)*(offset)/2)) // quick mafs
			offset += leftVal
			left++
		} else if left < right {
			currLeft := leftVal
			rightVal := values[right]
			for rightVal > 0 && currLeft > 0 {
				answer += uint64((right / 2) * offset)
				rightVal--
				currLeft--
				offset++
			}
			values[right] = rightVal
			values[left] = currLeft
			if values[right] == 0 {
				right -= 2
			}
			if currLeft == 0 {
				left++
			}
		} else {
			left++
		}
	}
	println(answer)
}
