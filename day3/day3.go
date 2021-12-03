package day3

import (
	"advent-of-code-2021/input"
	"fmt"
	"math"
)

func Day3() {
	positions, err := input.ReadLinesString("inputs/day3.txt")

	if err != nil {
		fmt.Println("Could not open file")
	} else {
		fmt.Println(part1(positions))
		fmt.Println(part2(positions))
	}
}

func part1(positions []string) int {
	posLength := len(positions[0])
	var counts = make([]int, posLength)

	for _, position := range positions {
		for i, digit := range position {
			if digit == '1' {
				counts[i] += 1
			}
		}
	}

	gamma := 0
	epsilon := 0
	half := len(positions) / 2
	for i, count := range counts {
		if count > half {
			gamma |= 1 << (posLength - i - 1)
		} else {
			epsilon |= 1 << (posLength - i - 1)
		}
	}

	return gamma * epsilon
}

func keepMostCommon(ones int, half int) uint8 {
	if ones >= half {
		return '1'
	}
	return '0'
}

func keepLeastCommon(ones int, half int) uint8 {
	if ones < half {
		return '1'
	}
	return '0'
}

func part2(positions []string) int {
	remainingOxygen := positions
	remainingScrubber := positions
	posLength := len(positions[0])

	for i := 0; i < posLength; i++ {
		remainingOxygen = keep(remainingOxygen, i, keepMostCommon)
		remainingScrubber = keep(remainingScrubber, i, keepLeastCommon)
	}

	return createNum(remainingOxygen[0]) * createNum(remainingScrubber[0])
}

func createNum(position string) int {
	num := 0
	posLength := len(position)

	for i, pos := range position {
		if pos == '1' {
			num |= 1 << (posLength - i - 1)
		}
	}

	return num
}

func keep(positions []string, index int, bitToKeep func(int, int) uint8) []string {
	if len(positions) == 1 {
		return positions
	}

	keep := bitToKeep(
		countOnes(positions, index),
		int(math.Ceil(float64(len(positions))/2.0)),
	)

	remaining := make([]string, 0)

	for _, position := range positions {
		if position[index] == keep {
			remaining = append(remaining, position)
		}
	}

	return remaining
}

func countOnes(positions []string, index int) int {
	ones := 0
	for _, position := range positions {
		if position[index] == '1' {
			ones += 1
		}
	}

	return ones
}
