package main

import (
	"strconv"
	"strings"

	helpers "github.com/mangusm/advent-of-code-2023/v2"
)

func strArrtoIntArr(arr []string) []int {
	out := []int{}
	for _, el := range arr {
		n, err := strconv.Atoi(el)
		if err != nil {
			panic(err)
		}
		out = append(out, n)
	}
	return out
}

func inRangeItems(items []int, ranges []int) map[int]int {
	out := map[int]int{}
	upperBound := ranges[1] + ranges[2]
	for _, item := range items {
		if item >= ranges[1] && item < upperBound {
			delta := item - ranges[1]
			out[item] = ranges[0] + delta
		}
	}
	return out
}

func removeFromArr(arr []int, value int) []int {
	out := []int{}
	for _, el := range arr {
		if el != value {
			out = append(out, el)
		}
	}
	return out
}

func arrMin(arr []int) int {
	least := arr[0]
	for _, n := range arr {
		if n < least {

			least = n
		}
	}
	return least
}

func PartOne() int {
	lines, err := helpers.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	seedStrings := strings.Fields(lines[0])[1:]
	items := strArrtoIntArr(seedStrings)
	tmpItems := []int{}
	startCreatingMap := false
	for i, line := range lines[1:] {
		if len(strings.Split(line, "-to-")) > 1 {
			startCreatingMap = true
			continue
		}
		if len(line) > 0 && startCreatingMap {
			vals := strArrtoIntArr(strings.Split(line, " "))
			itemsToTmpItems := inRangeItems(items, vals)
			for k, v := range itemsToTmpItems {
				tmpItems = append(tmpItems, v)
				items = removeFromArr(items, k)
			}
		}
		if len(line) == 0 || i == len(lines)-2 {
			if startCreatingMap {
				tmpItems = append(tmpItems, items...)
				items = tmpItems
				tmpItems = []int{}
			}
			startCreatingMap = false
			continue
		}
	}
	return arrMin(items)
}
