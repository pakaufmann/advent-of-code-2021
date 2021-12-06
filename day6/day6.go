package day6

import (
	"advent-of-code-2021/input"
	"fmt"
	"strconv"
	"strings"
)

func Day6() {
	fish, err := readFish()

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		fmt.Println(countFish(runSimulation(fish, 80)))
		fmt.Println(countFish(runSimulation(fish, 256)))
	}
}

func countFish(fish map[int64]int64) int64 {
	var sum int64 = 0
	for _, count := range fish {
		sum += count
	}

	return sum
}

func readFish() (map[int64]int64, error) {
	lines, err := input.ReadLinesString("inputs/day6.txt")
	if err != nil {
		return nil, err
	}

	output := make(map[int64]int64, 0)
	for _, fish := range strings.Split(lines[0], ",") {
		day, _ := strconv.ParseInt(fish, 10, 64)
		output[day] += 1
	}

	return output, nil
}

func runSimulation(fish map[int64]int64, days int) map[int64]int64 {
	for i := 0; i < days; i++ {
		newFish := make(map[int64]int64, 0)

		for days, count := range fish {
			remaining := days - 1
			if remaining == -1 {
				newFish[6] += count
				newFish[8] += count
			} else {
				newFish[remaining] += count
			}
		}

		fish = newFish
	}

	return fish
}
