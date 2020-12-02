package solutions

import (
	"fmt"
	"karlc/aoc2020/src/util"
	"sort"
)

func Day1Part1() {

	input := util.GetLinesAsInts(1)
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	var firstIndex int
	var secondIndex int

	for i, e := range input {
		if pos, ok := search(e, i, len(input)-1, 2020, input); ok {
			firstIndex = i
			secondIndex = pos
			break
		}
	}

	fmt.Printf("Result: %d\n", input[firstIndex]*input[secondIndex])
}

func Day1Part2() {

	input := util.GetLinesAsInts(1)
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	var firstIndex int
	var secondIndex int
	var thirdIdnex int

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i == j {
				continue
			}
			offset := input[i] + input[j]
			if pos, ok := search(offset, j, len(input)-1, 2020, input); ok {
				firstIndex = i
				secondIndex = j
				thirdIdnex = pos
				break
			}
		}
	}

	fmt.Printf("Result: %d\n", input[firstIndex]*input[secondIndex]*input[thirdIdnex])
}

func search(i, min, max, target int, list []int) (int, bool) {

	if max == min+1 || min == max {
		return -1, false
	}

	pos := int((max + min) / 2)
	curr := list[pos]
	res := curr + i

	if res < target {
		return search(i, pos, max, target, list)
	}

	if res > target {
		return search(i, min, pos, target, list)
	}

	if res == target {
		return pos, true
	}

	return -1, false
}
