package main

import (
	"advent24"
	"strings"
	"sync"
)

var yMoves = [4]int{-1, 0, 1, 0}
var xMoves = [4]int{0, 1, 0, -1}

func main() {
	input := advent24.ReadInputOfDay(6)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	inputMapStr := strings.Split(input, "\n")
	inputMap := make([][]byte, len(inputMapStr))

	var guardPosition advent24.Pair[int, int]
	for i := 0; i < len(inputMapStr); i++ {
		newVec := make([]byte, len(inputMapStr[i]))
		for j := 0; j < len(inputMapStr[i]); j++ {
			newVec[j] = inputMapStr[i][j]
			// منها حجة و زيارة
			if inputMapStr[i][j] == '^' {
				guardPosition = advent24.Pair[int, int]{First: i, Second: j}
			}
		}
		inputMap[i] = newVec
	}

	answer := 0
	currentDirection := 0

	visited := make(map[advent24.Pair[int, int]]struct{})

	for {
		if isInBounds(inputMap, guardPosition.First, guardPosition.Second) {
			if isNextPositionAnObstacle(inputMap, currentDirection, guardPosition.First, guardPosition.Second) {
				turnDirection(&currentDirection)
			} else {
				if _, exists := visited[guardPosition]; !exists {
					visited[guardPosition] = struct{}{}
					answer++
				}
				guardPosition.First = guardPosition.First + yMoves[currentDirection]
				guardPosition.Second = guardPosition.Second + xMoves[currentDirection]
			}
		} else {
			break
		}
	}
	println(answer)
}

func partTwo(input string) {
	inputMapStr := strings.Split(input, "\n")
	inputMap := make([][]byte, len(inputMapStr))

	var guardPosition advent24.Pair[int, int]
	for i := 0; i < len(inputMapStr); i++ {
		newVec := make([]byte, len(inputMapStr[i]))
		for j := 0; j < len(inputMapStr[i]); j++ {
			newVec[j] = inputMapStr[i][j]
			if inputMapStr[i][j] == '^' {
				guardPosition = advent24.Pair[int, int]{First: i, Second: j}
			}
		}
		inputMap[i] = newVec
	}

	limit := 10000
	answer := 0
	answerMutex := sync.Mutex{}
	var wg sync.WaitGroup
	sem := make(chan struct{}, 8) // 8 CPU cores

	for i := 0; i < len(inputMap); i++ {
		for j := 0; j < len(inputMap[i]); j++ {

			currentI, currentJ := i, j

			wg.Add(1)
			go func(currentI int, currentJ int) {
				defer wg.Done()

				sem <- struct{}{}
				defer func() { <-sem }()

				if inputMap[currentI][currentJ] == '#' || inputMap[currentI][currentJ] == '^' {
					return
				}

				var tempGuardPosition advent24.Pair[int, int]
				tempGuardPosition.First = guardPosition.First
				tempGuardPosition.Second = guardPosition.Second
				currentDirection := 0
				visited := make(map[advent24.Pair[int, int]]int)
				isLooped := true
				for k := 0; k < limit; k++ {
					if isInBounds(inputMap, tempGuardPosition.First, tempGuardPosition.Second) {
						if isNextPositionAnObstacleOrMyObstacle(inputMap, currentDirection, tempGuardPosition.First, tempGuardPosition.Second, currentI, currentJ) {
							turnDirection(&currentDirection)
						} else {
							if _, exists := visited[tempGuardPosition]; !exists {
								visited[tempGuardPosition] = visited[tempGuardPosition] + 1
							} else if visited[tempGuardPosition] >= 10000 {
								break
							}
							tempGuardPosition.First = tempGuardPosition.First + yMoves[currentDirection]
							tempGuardPosition.Second = tempGuardPosition.Second + xMoves[currentDirection]
						}
					} else {
						isLooped = false
						break
					}
				}
				if isLooped {
					answerMutex.Lock()
					answer++
					answerMutex.Unlock()
				}
			}(currentI, currentJ)
		}
	}
	wg.Wait()
	println(answer)
}

func isInBounds(inputMap [][]byte, i int, j int) bool {
	return i >= 0 && i < len(inputMap) && j >= 0 && j < len(inputMap[i])
}

func isNextPositionAnObstacle(inputMap [][]byte, currentDirection int, i int, j int) bool {
	nextI := i + yMoves[currentDirection]
	nextJ := j + xMoves[currentDirection]
	return isInBounds(inputMap, nextI, nextJ) && inputMap[nextI][nextJ] == '#'
}

func turnDirection(currentDirection *int) {
	*currentDirection = (*currentDirection + 1) % 4
}

func isNextPositionAnObstacleOrMyObstacle(inputMap [][]byte, currentDirection int, i int, j int, myObstacleI int, myObstacleJ int) bool {
	nextI := i + yMoves[currentDirection]
	nextJ := j + xMoves[currentDirection]
	return isInBounds(inputMap, nextI, nextJ) && (inputMap[nextI][nextJ] == '#' || (nextI == myObstacleI && nextJ == myObstacleJ))
}
