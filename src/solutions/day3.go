package solutions

import (
	"fmt"
	"karlc/aoc2020/src/util"
	"strings"
)

type Cell int

const (
	TREE Cell = iota
	SPACE
	ABOVE
	BELOW
)

type mapGrid struct {
	grid [][]string
	h    int
	w    int
	x    int
	y    int
}

func (m mapGrid) getCell() Cell {
	symbol := m.grid[m.x][m.y]

	if symbol == "#" {
		return TREE
	}

	return SPACE
}

func (m *mapGrid) stepHorizontal(steps int) Cell {
	m.x = (m.x + steps) % m.w

	if m.x >= m.h-1 {
		return BELOW
	}

	if m.x < 0 {
		return ABOVE
	}

	return m.getCell()
}

func (m *mapGrid) stepVertical(steps int) Cell {
	m.y = (m.y + steps) % m.w

	if m.x >= m.h-1 {
		return BELOW
	}

	if m.x < 0 {
		return ABOVE
	}

	return m.getCell()
}

func newMapGrid(lines []string) mapGrid {
	ret := make([][]string, 0, 0)
	for _, line := range lines {
		in := strings.Split(line, "")
		ret = append(ret, in)
	}

	return mapGrid{
		grid: ret,
		h:    len(lines),
		w:    len(lines[0]),
	}
}

func Day3Part1() {
	numTrees := 0
	mg := newMapGrid(util.GetLinesAsStrings(3))

outer:
	for {
		fmt.Printf("x: %d, y: %d\n", mg.x, mg.y)
		mg.stepHorizontal(3)
		res := mg.stepVertical(1)
		switch res {
		case TREE:
			numTrees++
		case BELOW:
			break outer
		}
	}

	fmt.Println(numTrees)
}
