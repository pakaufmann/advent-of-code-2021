package day25

import (
	"advent-of-code-2021/input"
	"fmt"
)

type Cucumbers = [][]int

func Day25() {
	lines, err := input.ReadLinesString("inputs/day25.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		lastMap := readMap(lines)
		fmt.Println(findStop(lastMap))
	}
}

func findStop(lastMap Cucumbers) int {
	for i := 1; ; i++ {
		newMap := step(lastMap)
		if equals(newMap, lastMap) {
			return i
		}
		lastMap = newMap
	}
}

func readMap(lines []string) Cucumbers {
	cucumbers := make(Cucumbers, len(lines))

	for y, line := range lines {
		row := make([]int, len(line))

		for x, tile := range line {
			if tile == '>' {
				row[x] = 1
			} else if tile == 'v' {
				row[x] = 2
			}
		}

		cucumbers[y] = row
	}
	return cucumbers
}

func step(cucumbers Cucumbers) Cucumbers {
	newMap := make([][]int, len(cucumbers))

	for y, row := range cucumbers {
		width := len(row)
		newMap[y] = make([]int, width)

		for x, tile := range row {
			next := (x + 1) % width
			if tile == 1 {
				if cucumbers[y][next] == 0 {
					newMap[y][next] = 1
				} else {
					newMap[y][x] = 1
				}
			}
		}
	}

	height := len(cucumbers)

	for y, row := range cucumbers {
		for x, tile := range row {
			next := (y + 1) % height
			if tile == 2 {
				if newMap[next][x] == 0 && cucumbers[next][x] != 2 {
					newMap[next][x] = 2
				} else {
					newMap[y][x] = 2
				}
			}
		}
	}

	return newMap
}

func equals(f Cucumbers, s Cucumbers) bool {
	for y, row := range f {
		for x, tile := range row {
			if s[y][x] != tile {
				return false
			}
		}
	}

	return true
}
