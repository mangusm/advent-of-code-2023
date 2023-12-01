package main

import (
	helpers "github.com/mangusm/advent-of-code-2023/v2"
)

func lookahead(line string, idx int) int {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, word := range words {
		skip := len(word)
		if len(line) >= idx+skip {
			if line[idx:idx+skip] == word {
				return i + 1
			}
		}
	}
	return 0
}

func PartTwo() int {
	lines, err := helpers.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, line := range lines {
		lineSum := 0
		firstNumberFound := false
		lastNumber := 0
		for i := 0; i < len(line); i++ {
			charCode := rune(line[i])
			if charCode >= 48 && charCode <= 57 {
				if !firstNumberFound {
					lineSum += 10 * (int(charCode) - 48)
				}
				firstNumberFound = true
				lastNumber = int(charCode) - 48
			} else {
				n := lookahead(line, i)
				if !firstNumberFound {
					lineSum += 10 * n
				}
				if n != 0 {
					firstNumberFound = true
					lastNumber = n
				}
			}

		}
		lineSum += lastNumber
		sum += lineSum
	}
	return sum
}
