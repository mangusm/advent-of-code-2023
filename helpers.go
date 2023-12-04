package helpers

import (
	"bufio"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func Max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
