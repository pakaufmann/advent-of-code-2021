package day1

import (
	"advent-of-code-2021/input"
	"fmt"
	"math"
)

func Day1() {
	depths, err := input.ReadLinesInt("inputs/day1.txt")
	if err != nil {
		fmt.Println("Could not find file")
	} else {
		fmt.Println(part1(depths))
		fmt.Print(part2(depths))
	}
}

func part1(depths []int64) int {
	increased := 0
	var previous int64 = math.MaxInt64

	for _, depth := range depths {
		if depth > previous {
			increased++
		}
		previous = depth
	}

	return increased
}

func part2(depths []int64) int {
	increased := 0
	var lastSum int64 = math.MaxInt64

	for i, depth := range depths[:len(depths)-2] {
		sum := depth + depths[i+1] + depths[i+2]
		if sum > lastSum {
			increased++
		}
		lastSum = sum
	}

	return increased
}
