package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func GetLinesAsStrings(day int) []string {
	p, err := filepath.Abs(fmt.Sprintf("input/%d.txt", day))
	if err != nil {
		panic("Failed to parse path: " + err.Error())
	}
	file, err := os.Open(p)
	if err != nil {
		panic("Failed to open input file: " + err.Error())
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic("Failed to scan input: " + err.Error())
	}
	return lines
}

func GetLinesAsInts(day int) []int {
	lines := GetLinesAsStrings(day)
	ints := make([]int, len(lines), len(lines))

	for i, e := range lines {
		v, err := strconv.Atoi(e)
		if err != nil {
			panic("Failed to convert line input to int: " + err.Error())
		}
		ints[i] = v
	}
	return ints
}
