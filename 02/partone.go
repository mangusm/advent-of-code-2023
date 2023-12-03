package main

import (
	"regexp"
	"strconv"

	helpers "github.com/mangusm/advent-of-code-2023/v2"
)

func PartOne() int {
	lines, err := helpers.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	reGame := regexp.MustCompile(`^Game (\d+)`)
	reRed := regexp.MustCompile(`(1[3-9]|[2-9]\d) red`)
	reGreen := regexp.MustCompile(`(1[4-9]|[2-9]\d) green`)
	reBlue := regexp.MustCompile(`(1[5-9]|[2-9]\d) blue`)

	for _, line := range lines {
		matches := reGame.FindStringSubmatch(line)
		gameId, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}

		if reRed.MatchString(line) || reGreen.MatchString(line) || reBlue.MatchString(line) {
			continue
		}
		sum += gameId
	}
	return sum
}
