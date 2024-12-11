package main

import (
	"advent24"
	"strconv"
	"strings"
)

func main() {
	input := advent24.ReadInputOfDay(10)
	partOne(input)
	partTwo(input)
}

var visitedHikes = make(map[advent24.Pair[int, int]]map[advent24.Pair[int, int]]struct{}) // map<pair<starting_position_x, starting_position_y>, map<pair<number_nine_x, number_nine_y>, void>

var visitedHikesWithCount = make(map[advent24.Pair[int, int]]int) // map<pair<starting_position_x, starting_position_y>, int>

func partOne(input string) {
	topoMap := parseTopologicalMap(input)
	var answer int
	for i := range topoMap {
		for j := range topoMap[i] {
			if topoMap[i][j] != 0 {
				continue
			}
			startingPosition := advent24.Pair[int, int]{First: i, Second: j}
			visitedHikes[startingPosition] = make(map[advent24.Pair[int, int]]struct{})
			currentHikeVisits := make(map[advent24.Pair[int, int]]struct{})
			sum := DFS_partOne(topoMap, currentHikeVisits, startingPosition, 0, i, j)
			answer += sum
		}
	}
	println(answer)
}

func partTwo(input string) {
	topoMap := parseTopologicalMap(input)

	for i := range topoMap {
		for j := range topoMap[i] {
			if topoMap[i][j] != 0 {
				continue
			}
			startingPosition := advent24.Pair[int, int]{First: i, Second: j}
			currentHikeVisits := make(map[advent24.Pair[int, int]]struct{})
			DFS_partTwo(topoMap, currentHikeVisits, startingPosition, 0, i, j)
		}
	}

	var answer int
	for _, val := range visitedHikesWithCount {
		answer += val
	}
	println(answer)
}

func DFS_partOne(inputMap [][]int, currentHikeVisits map[advent24.Pair[int, int]]struct{}, startingPosition advent24.Pair[int, int], searchedNum int, i int, j int) int {
	if i < 0 || i >= len(inputMap) || j < 0 || j >= len(inputMap[i]) {
		return 0
	}

	currentPosition := advent24.Pair[int, int]{First: i, Second: j}
	_, alreadyVisited := currentHikeVisits[currentPosition]
	if alreadyVisited {
		return 0
	} else {
		currentHikeVisits[currentPosition] = struct{}{}
	}

	if inputMap[i][j] != searchedNum {
		delete(currentHikeVisits, currentPosition)
		return 0
	}

	if inputMap[i][j] == 9 {
		_, nineVisitedFromStartPosition := visitedHikes[startingPosition][currentPosition]
		if !nineVisitedFromStartPosition {
			visitedHikes[startingPosition][currentPosition] = struct{}{}
			delete(currentHikeVisits, currentPosition)
			return 1
		} else {
			delete(currentHikeVisits, currentPosition)
			return 0
		}
	}
	left := DFS_partOne(inputMap, currentHikeVisits, startingPosition, searchedNum+1, i, j-1)
	right := DFS_partOne(inputMap, currentHikeVisits, startingPosition, searchedNum+1, i, j+1)
	up := DFS_partOne(inputMap, currentHikeVisits, startingPosition, searchedNum+1, i-1, j)
	down := DFS_partOne(inputMap, currentHikeVisits, startingPosition, searchedNum+1, i+1, j)
	delete(currentHikeVisits, currentPosition)
	return left + right + up + down
}

func DFS_partTwo(inputMap [][]int, currentHikeVisits map[advent24.Pair[int, int]]struct{}, startingPosition advent24.Pair[int, int], searchedNum int, i int, j int) {
	if i < 0 || i >= len(inputMap) || j < 0 || j >= len(inputMap[i]) {
		return
	}

	currentPosition := advent24.Pair[int, int]{First: i, Second: j}
	_, alreadyVisited := currentHikeVisits[currentPosition]
	if alreadyVisited {
		return
	} else {
		currentHikeVisits[currentPosition] = struct{}{}
	}

	if inputMap[i][j] != searchedNum {
		delete(currentHikeVisits, currentPosition)
		return
	}

	if inputMap[i][j] == 9 {
		visitedHikesWithCount[startingPosition]++
		delete(currentHikeVisits, currentPosition)
		return
	}
	DFS_partTwo(inputMap, currentHikeVisits, startingPosition, searchedNum+1, i, j-1)
	DFS_partTwo(inputMap, currentHikeVisits, startingPosition, searchedNum+1, i, j+1)
	DFS_partTwo(inputMap, currentHikeVisits, startingPosition, searchedNum+1, i-1, j)
	DFS_partTwo(inputMap, currentHikeVisits, startingPosition, searchedNum+1, i+1, j)
	delete(currentHikeVisits, currentPosition)
}

func parseTopologicalMap(input string) [][]int {
	lines := strings.Split(input, "\n")
	topoMap := make([][]int, len(lines))
	for i, line := range lines {
		mapLine := make([]int, len(line))
		for j := range line {
			if line[j] == '.' {
				mapLine[j] = -1
			} else {
				mapLine[j], _ = strconv.Atoi(string(line[j]))
			}
		}
		topoMap[i] = mapLine
	}
	return topoMap
}
