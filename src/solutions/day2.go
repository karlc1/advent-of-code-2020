package solutions

import (
	"fmt"
	"karlc/aoc2020/src/util"
	"strconv"
	"strings"
)

type day2Input struct {
	min      int
	max      int
	letter   string
	password string
}

func newDay2Input(lineIn string) day2Input {
	split := strings.Split(lineIn, " ")
	allowedRange := strings.Split(split[0], "-")

	min, err := strconv.Atoi(allowedRange[0])
	if err != nil {
		panic("Failed to convert string to int" + err.Error())
	}

	max, err := strconv.Atoi(allowedRange[1])
	if err != nil {
		panic("Failed to convert string to int" + err.Error())
	}

	letterStr := strings.Replace(split[1], ":", "", -1)

	return day2Input{
		min:      min,
		max:      max,
		letter:   letterStr,
		password: split[2],
	}
}

func (d day2Input) isValidCount() bool {
	count := strings.Count(d.password, d.letter)

	return count >= d.min && count <= d.max
}

func (d day2Input) isValidPos() bool {
	if d.min >= len(d.password) {
		return false
	}

	pwdList := strings.Split(d.password, "")
	minLetter := pwdList[d.min-1]
	maxLetter := ""

	if d.max <= len(pwdList) {
		maxLetter = pwdList[d.max-1]
	}

	fmt.Println(maxLetter)

	return (minLetter == d.letter && maxLetter != d.letter) || (maxLetter == d.letter && minLetter != d.letter)
}

func Day2Part1() {
	validPasswords := 0
	lines := util.GetLinesAsStrings(2)
	for _, line := range lines {
		if newDay2Input(line).isValidCount() {
			validPasswords++
		}
	}

	fmt.Printf("Result: %d\n", validPasswords)
}

func Day2Part2() {
	validPasswords := 0
	lines := util.GetLinesAsStrings(2)
	for _, line := range lines {
		if newDay2Input(line).isValidPos() {
			validPasswords++
		}
	}

	fmt.Printf("Result: %d\n", validPasswords)
}
