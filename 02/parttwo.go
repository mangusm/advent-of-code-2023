package main

import (
	"regexp"
	"strconv"

	helpers "github.com/mangusm/advent-of-code-2023/v2"
)

func PartTwo() int {
	lines, err := helpers.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	reRed := regexp.MustCompile(`(\d+) red`)
	reGreen := regexp.MustCompile(`(\d+) green`)
	reBlue := regexp.MustCompile(`(\d+) blue`)

	reColors := []*regexp.Regexp{reRed, reGreen, reBlue}

	for _, line := range lines {
		mins := []int{0, 0, 0}

		for i, reColor := range reColors {
			matches := reColor.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				count, err := strconv.Atoi(match[1])
				if err != nil {
					panic(err)
				}
				if count > mins[i] {
					mins[i] = count
				}
			}
		}
		sum += mins[0] * mins[1] * mins[2]

	}
	return sum
}
