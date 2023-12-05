package main

import (
	"strings"

	helpers "github.com/mangusm/advent-of-code-2023/v2"
)

func initMap(len int, value int) map[int]int {
	m := map[int]int{}
	for i := 0; i < len; i++ {
		m[i] = value
	}
	return m
}

func add(m1 map[int]int, m2 map[int]int) map[int]int {
	for k, v := range m1 {
		m2[k] += v
	}
	return m2
}

func processCard(cards map[int]int, cardNumber int) map[int]int {
	counts := initMap(199, 0)

	for i := 1; i <= cards[cardNumber]; i++ {
		counts[cardNumber+i]++
		counts2 := processCard(cards, cardNumber+i)
		counts = add(counts, counts2)
	}
	return counts
}

func PartTwo() int {
	lines, err := helpers.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0

	cards := initMap(199, 0)
	counts := initMap(199, 1)

	for i, line := range lines {
		cardWinningNumbers := strings.Split(line[10:], "|")[0]
		cardNumbersYouHave := strings.Split(line[10:], "|")[1]
		winningNumbers := strings.Fields(cardWinningNumbers)
		numbersYouHave := strings.Fields(cardNumbersYouHave)
		numWinningNumbers := Overlap(winningNumbers, numbersYouHave)
		cards[i] = numWinningNumbers
	}

	// This takes a liiiiiiitle bit of time to run
	for i := 0; i < 199; i++ {
		counts = add(counts, processCard(cards, i))
	}

	for _, v := range counts {
		sum += v
	}

	return sum
}
