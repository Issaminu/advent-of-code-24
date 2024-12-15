package main

import (
	"advent24"
	"strings"
)

func main() {
	input := advent24.ReadInputOfDay(12)
	partOne(input)
}

var plantNeighbourhood = make(map[advent24.Pair[int, int]]map[advent24.Pair[int, int]]struct{}) // map<pair<plant_starting_position_x, plant_starting_position_y>, map<pair<same_plant_type_x, same_plant_type_y>, void>

var numNeighboursInRegion = make(map[advent24.Pair[int, int]]int) // map<pair<plant_starting_position_x, plant_starting_position_y>, same_type_neighbour_plants_count>

var seenPlants = make(map[advent24.Pair[int, int]]struct{}) // map<pair<plant_position_x, plant_position_y>, void>

func partOne(input string) {
	lines := parseInput(input)
	for i := range lines {
		for j := range lines[i] {
			position := advent24.Pair[int, int]{First: i, Second: j}
			if _, seen := seenPlants[position]; seen {
				continue
			}
			plantNeighbourhood[position] = make(map[advent24.Pair[int, int]]struct{})
			currentPath := make(map[advent24.Pair[int, int]]struct{})
			numNeighboursInRegion[position] = 0
			DFS_partOne(lines, currentPath, position, lines[i][j], i, j)
		}
	}

	var answer int
	for startingPlant, plantsInRegion := range plantNeighbourhood {
		numberOfPlantsInRegion := len(plantsInRegion)
		priceOfFencingRegion := (numberOfPlantsInRegion*4 - numNeighboursInRegion[startingPlant]) * numberOfPlantsInRegion
		answer += priceOfFencingRegion
	}

	println(answer)
}

func DFS_partOne(lines [][]byte, currentPath map[advent24.Pair[int, int]]struct{}, startingPosition advent24.Pair[int, int], searchedPlant byte, i int, j int) {
	currentPlant := advent24.SafeGetElement(lines, i, j, 0)
	currentPosition := advent24.Pair[int, int]{First: i, Second: j}
	if currentPlant == 0 { // 0 is used as fail value, since it's impossible to get from input
		return
	}

	if currentPlant != searchedPlant {
		return
	}

	_, seenPlantInCurrentPath := currentPath[currentPosition]
	if seenPlantInCurrentPath {
		return
	}

	currentPath[currentPosition] = struct{}{}
	seenPlants[currentPosition] = struct{}{}
	plantNeighbourhood[startingPosition][currentPosition] = struct{}{}

	neighboursCount := 0

	if advent24.SafeGetElement(lines, i, j-1, 0) == searchedPlant {
		neighboursCount++
		DFS_partOne(lines, currentPath, startingPosition, searchedPlant, i, j-1)
	}
	if advent24.SafeGetElement(lines, i, j+1, 0) == searchedPlant {
		neighboursCount++
		DFS_partOne(lines, currentPath, startingPosition, searchedPlant, i, j+1)
	}
	if advent24.SafeGetElement(lines, i-1, j, 0) == searchedPlant {
		neighboursCount++
		DFS_partOne(lines, currentPath, startingPosition, searchedPlant, i-1, j)
	}
	if advent24.SafeGetElement(lines, i+1, j, 0) == searchedPlant {
		neighboursCount++
		DFS_partOne(lines, currentPath, startingPosition, searchedPlant, i+1, j)
	}

	numNeighboursInRegion[startingPosition] += neighboursCount
}

func parseInput(input string) [][]byte {
	lines := strings.Split(input, "\n")
	linesMap := make([][]byte, len(lines))

	for i := range lines {
		line := make([]byte, len(lines[i]))
		for j := range lines[i] {
			line[j] = lines[i][j]
		}
		linesMap[i] = line
	}

	return linesMap
}
