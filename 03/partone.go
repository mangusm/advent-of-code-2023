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
	re := regexp.MustCompile(`\d+`)
	re2 := regexp.MustCompile(`[^.\d]`)

	for i, line := range lines {
		indices := re.FindAllStringIndex(line, -1)
		strings := re.FindAllString(line, -1)
		for j, match := range indices {
			partNumber, err := strconv.Atoi(strings[j])
			if err != nil {
				panic(err)
			}
			idx := match[1]
			length := len(strings[j]) + 1

			if i > 0 {
				if re2.MatchString(lines[i-1][max(idx-length, 0):min(idx+1, 140)]) {
					sum += partNumber
					continue
				}
			}

			if re2.MatchString(line[idx-1:idx-1]) || re2.MatchString(line[max(idx-length, 0):min(idx+1, 140)]) {
				sum += partNumber
				continue
			}

			if i < len(lines) {
				if re2.MatchString(lines[i+1][max(idx-length, 0):min(idx+1, 140)]) {
					sum += partNumber
					continue
				}
			}
		}
	}
	return sum
}
