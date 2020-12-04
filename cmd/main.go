package main

import (
	"fmt"
	"karlc/aoc2020/src/solutions"
	"time"
)

func main() {
	start := time.Now()
	solutions.Day3Part1()
	fmt.Printf("Elapsed time: %v", time.Since(start))
}
