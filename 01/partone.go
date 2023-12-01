package main

import (
	helpers "github.com/mangusm/advent-of-code-2023/v2"
)

func PartOne() int {
	lines, err := helpers.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, line := range lines {
		lineSum := 0
		firstNumberFound := false
		lastNumber := 0
		for _, charCode := range line {
			if charCode >= 48 && charCode <= 57 {
				if !firstNumberFound {
					lineSum += 10 * (int(charCode) - 48)
				}
				lastNumber = int(charCode) - 48
				firstNumberFound = true
			}
		}
		lineSum += lastNumber
		sum += lineSum
	}
	return sum
}
