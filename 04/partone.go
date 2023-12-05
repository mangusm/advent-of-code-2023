package main

import (
	"math"
	"strings"

	helpers "github.com/mangusm/advent-of-code-2023/v2"
)

func Contains(value string, l []string) bool {
	for _, n := range l {
		if value == n {
			return true
		}
	}
	return false
}

func Overlap(l1 []string, l2 []string) int {
	matches := 0
	for _, n := range l1 {
		if Contains(n, l2) {
			matches++
		}
	}
	return matches
}

func PartOne() int {
	lines, err := helpers.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0

	for _, line := range lines {
		cardWinningNumbers := strings.Split(line[10:], "|")[0]
		cardNumbersYouHave := strings.Split(line[10:], "|")[1]
		winningNumbers := strings.Fields(cardWinningNumbers)
		numbersYouHave := strings.Fields(cardNumbersYouHave)
		numWinningNumbers := Overlap(winningNumbers, numbersYouHave)
		if numWinningNumbers > 0 {
			sum += int(math.Pow(2, float64(numWinningNumbers-1)))
		}
	}
	return sum
}
