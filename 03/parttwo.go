package main

import (
	"math"
	"strconv"

	helpers "github.com/mangusm/advent-of-code-2023/v2"
)

func getFullNumber(line string, start int, digit int) int {
	digits := []int{digit}
	direction := 1
	i := start
	lookedBothWays := false
	for {
		i += direction
		if i < 0 || i >= len(line) {
			if lookedBothWays {
				break
			}
			direction *= -1
			i = start
			lookedBothWays = true
			continue
		}
		digit, err := strconv.Atoi(string(line[i]))
		if err != nil {
			if lookedBothWays {
				break
			}
			direction *= -1
			i = start
			lookedBothWays = true
			continue
		}
		if direction < 0 {
			digits = append([]int{digit}, digits...)
			continue
		}
		digits = append(digits, digit)
	}
	sum := 0
	for i, digit := range digits {
		sum += digit * int(math.Pow10(len(digits)-i-1))
	}
	return sum
}

func around(lines []string, lineIndex int, asteriskIndex int) []int {
	gearRatios := []int{}
	for i := -1; i < 2; i++ {
		lastCharWasDigit := false
		for k := -1; k < 2; k++ {
			if lineIndex+i >= 0 && lineIndex+i < len(lines) {
				line := lines[lineIndex+i]
				if asteriskIndex+k > 0 && asteriskIndex+k < len(line) {
					digit, err := strconv.Atoi(string(line[asteriskIndex+k]))
					if err != nil {
						lastCharWasDigit = false
						continue
					}
					if !lastCharWasDigit {
						lastCharWasDigit = true
						gearRatios = append(gearRatios, getFullNumber(line, asteriskIndex+k, digit))
					}
				}
			}
		}
	}
	return gearRatios
}

func PartTwo() int {
	lines, err := helpers.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0

	for i, line := range lines {
		for j, char := range line {
			if string(char) == "*" {
				gearRatios := around(lines, i, j)
				if len(gearRatios) == 2 {
					sum += gearRatios[0] * gearRatios[1]
				}
			}
		}
	}
	return sum
}
