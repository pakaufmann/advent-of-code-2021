package day7

import (
	"advent-of-code-2021/input"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day7() {
	lines, err := input.ReadLinesString("inputs/day7.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		fmt.Println(part1(createPositions(lines[0])))
		fmt.Println(part2(createPositions(lines[0])))
	}
}

func part1(positions []int) int {
	sort.Ints(positions)

	posToAlignTo := positions[len(positions)/2]
	fuelToUse := 0

	for _, position := range positions {
		fuelToUse += abs(position - posToAlignTo)
	}

	return fuelToUse
}

func createPositions(s string) []int {
	positions := make([]int, 0)
	for _, position := range strings.Split(s, ",") {
		pos, _ := strconv.ParseInt(position, 10, 64)
		positions = append(positions, int(pos))
	}
	return positions
}

func part2(positions []int) int {
	sum := 0
	for _, position := range positions {
		sum += position
	}

	lower := int(math.Floor(float64(sum) / float64(len(positions))))

	fuelToUseLower := 0
	fuelToUseHigher := 0

	for _, position := range positions {
		toMoveLower := abs(position - lower)
		toMoveHigher := abs(position - (lower + 1))
		fuelToUseLower += (toMoveLower * (toMoveLower + 1)) / 2
		fuelToUseHigher += (toMoveHigher * (toMoveHigher + 1)) / 2
	}

	if fuelToUseLower < fuelToUseHigher {
		return fuelToUseLower
	}
	return fuelToUseHigher
}

func abs(x int) (o int) {
	o = x
	if x < 0 {
		o = -x
	}
	return
}
