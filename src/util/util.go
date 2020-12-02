package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func GetLinesAsInts(day int) []int {
	p, err := filepath.Abs(fmt.Sprintf("input/%d.txt", day))
	if err != nil {
		panic("Failed to parse path: " + err.Error())
	}
	file, err := os.Open(p)
	if err != nil {
		panic("Failed to open input file: " + err.Error())
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Failed to parse int: " + err.Error())
		}
		lines = append(lines, n)
	}

	if err := scanner.Err(); err != nil {
		panic("Failed to scan input: " + err.Error())
	}
	return lines
}
