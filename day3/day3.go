package main

import (
	"advent24"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := advent24.ReadInputOfDay(3)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	multiplications := getMultiplicationsInStr(input)

	result := getResultOfMultiplications(multiplications)
	println(result)
}

func getMultiplicationsInStr(str string) []string {
	mul_pattern := regexp.MustCompile(`mul\([0-9]*,[0-9]*\)`)
	items := mul_pattern.FindAllStringSubmatch(str, -1)
	multiplications := make([]string, len(items))
	for i, item := range items {
		multiplications[i] = item[0]
	}
	return multiplications
}

func getResultOfMultiplications(multiplications []string) int {
	result := 0
	mul_num1_pattern := regexp.MustCompile(`\([0-9]*`)
	mul_num2_pattern := regexp.MustCompile(`,[0-9]*`)
	for _, mult := range multiplications {
		num1_str := strings.Split(mul_num1_pattern.FindString(mult), "(")[1]
		num1, _ := strconv.Atoi(num1_str)
		num2_str := strings.Split(mul_num2_pattern.FindString(mult), ",")[1]
		num2, _ := strconv.Atoi(num2_str)
		result += num1 * num2
	}
	return result
}

func partTwo(input string) {
	// result = edge case multiplications + normal case multiplications
	var result int

	// Edge case: getting all the multiplications at the start (which weren't precedented by a `do()`)
	before_first_dont_pattern := regexp.MustCompile(`[\s\S]*?don't\(\)`)
	first_pattern_matches := before_first_dont_pattern.FindString(input)

	first_multiplications := getMultiplicationsInStr(first_pattern_matches)
	result = getResultOfMultiplications(first_multiplications)

	// Normal case: getting the multiplications that were precedented by a `do()`
	subsequent_items := input[len(first_pattern_matches)+7:] // + 7 to account for "don't()"

	after_first_dont := regexp.MustCompile(`do\(\)[\s\S]*?don't\(\)`) // Normally, I should also handle the case there's a `do()` segment that doesn't end with a `don't()` segment, but this case isn't present in `input.txt` so i'm not gonna bother :p
	second_pattern_matches := after_first_dont.FindAllStringSubmatch(subsequent_items, -1)
	matchesStr := ""
	for _, item := range second_pattern_matches {
		matchesStr += item[0]
	}

	subsequent_multiplications := getMultiplicationsInStr(matchesStr)
	result += getResultOfMultiplications(subsequent_multiplications)

	println(result)
}
