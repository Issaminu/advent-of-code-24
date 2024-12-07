package main

import (
	"advent24"
	"math/rand"
	"strconv"
	"strings"
	"sync"
)

func main() {
	input := advent24.ReadInputOfDay(7)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	expressions := parseExpressions(input)

	var answer uint64
	answerMutex := sync.Mutex{}
	sym := make(chan (struct{}), 8) // 8 CPU cores
	var wg sync.WaitGroup

	limit := 10000

	for _, exp := range expressions {
		wg.Add(1)
		go func() {
			defer wg.Done()

			sym <- struct{}{}
			defer func() { <-sym }()

			for k := 0; k < limit; k++ {
				curr := uint64(exp.Second[0])
				for i := 1; i < len(exp.Second); i++ {
					randomOperation := rand.Intn(2) // returns 0 or 1
					if randomOperation == 0 {
						curr += uint64(exp.Second[i])
					} else {
						curr *= uint64(exp.Second[i])
					}
					if curr > exp.First {
						break
					}
				}
				if curr == exp.First {
					answerMutex.Lock()
					answer += uint64(exp.First)
					answerMutex.Unlock()
					return
				}
			}
		}()
	}
	wg.Wait()
	println(answer)
}

func partTwo(input string) {
	expressions := parseExpressions(input)

	var answer uint64
	answerMutex := sync.Mutex{}
	sym := make(chan (struct{}), 8) // 8 CPU cores
	var wg sync.WaitGroup

	limit := 1000000

	operations := [3]string{"+", "*", "||"}

	for _, exp := range expressions {
		wg.Add(1)
		go func() {
			defer wg.Done()

			sym <- struct{}{}
			defer func() { <-sym }()

			for k := 0; k < limit; k++ {
				curr := uint64(exp.Second[0])
				for i := 1; i < len(exp.Second); i++ {
					randomOperation := rand.Intn(3) // returns 0 or 1 or 2
					if operations[randomOperation] == "+" {
						curr += uint64(exp.Second[i])
					} else if operations[randomOperation] == "*" {
						curr *= uint64(exp.Second[i])
					} else if operations[randomOperation] == "||" {
						newCurr, _ := strconv.Atoi(strconv.FormatUint(curr, 10) + strconv.Itoa(exp.Second[i]))
						curr = uint64(newCurr)
					}
					if curr > exp.First {
						break
					}
				}

				if curr == exp.First {
					answerMutex.Lock()
					answer += uint64(exp.First)
					answerMutex.Unlock()
					return
				}
			}
		}()
	}
	wg.Wait()
	println(answer)
}

func parseExpressions(input string) []advent24.Pair[uint64, []int] {
	lines := strings.Split(input, "\n")
	expressions := make([]advent24.Pair[uint64, []int], len(lines))

	for i, line := range lines {
		lineSplit := strings.Split(line, ":")
		val, _ := strconv.Atoi(lineSplit[0])
		testValue := uint64(val)
		numbersStr := strings.Split(lineSplit[1][1:], " ")
		numbers := make([]int, len(numbersStr))
		for j := range numbersStr {
			numbers[j], _ = strconv.Atoi(numbersStr[j])
		}

		expressions[i] = advent24.Pair[uint64, []int]{First: testValue, Second: numbers}
	}

	return expressions
}
