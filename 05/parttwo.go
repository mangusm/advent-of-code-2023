package main

import (
	"strings"

	helpers "github.com/mangusm/advent-of-code-2023/v2"
)

func parseInput(lines []string) ([]int, [][][]int) {
	seeds := []int{}
	maps := [][][]int{}
	for i, line := range lines {
		if i == 0 {
			seeds = strArrtoIntArr(strings.Fields(line)[1:])
		} else if len(line) == 0 {
			continue
		} else if len(strings.Split(line, "-to-")) > 1 {
			maps = append(maps, [][]int{})
		} else {
			maps[len(maps)-1] = append(maps[len(maps)-1], strArrtoIntArr(strings.Fields(line)))
		}
	}
	return seeds, maps
}

func applyInterval(intervals [][]int, mappings [][]int) [][]int {
	a := [][]int{}
	for _, mapping := range mappings {
		dest := mapping[0]
		src := mapping[1]
		size := mapping[2]
		srcEnd := src + size
		newIntervals := [][]int{}

		for _, interval := range intervals {
			start := interval[0]
			end := interval[1]

			before := []int{start, min(end, src)}
			inter := []int{max(start, src), min(srcEnd, end)}
			after := []int{max(srcEnd, start), end}

			if before[1] > before[0] {
				newIntervals = append(newIntervals, before)
			}
			if inter[1] > inter[0] {
				a = append(a, []int{inter[0] - src + dest, inter[1] - src + dest})
			}
			if after[1] > after[0] {
				newIntervals = append(newIntervals, after)
			}
		}
		intervals = newIntervals
	}
	return append(a, intervals...)
}
func PartTwo() int {
	lines, err := helpers.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	seeds, parts := parseInput(lines)
	p2 := []int{}
	for i := 0; i < len(seeds); i += 2 {
		seedStart := seeds[i]
		seedSize := seeds[i+1]

		intervals := [][]int{{seedStart, seedStart + seedSize}}
		for _, part := range parts {
			intervals = applyInterval(intervals, part)
		}
		smallest := intervals[0][0]
		for _, interval := range intervals {
			if interval[0] < smallest {
				smallest = interval[0]
			}
		}
		p2 = append(p2, smallest)
	}
	answer := p2[0]
	for _, v := range p2 {
		if v < answer {
			answer = v
		}
	}

	return answer
}
