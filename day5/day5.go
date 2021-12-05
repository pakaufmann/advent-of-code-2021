package day5

import (
	"advent-of-code-2021/input"
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	x int64
	y int64
}

func Day5() {
	rows, err := input.ReadLinesString("inputs/day5.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		fmt.Println(countMultipleCovered(buildUpMap(rows, false)))
		fmt.Println(countMultipleCovered(buildUpMap(rows, true)))
	}
}

func buildUpMap(rows []string, withDiagonal bool) map[Position]int64 {
	covered := make(map[Position]int64)

	for _, row := range rows {
		from, to := readLine(row)

		if from.x == to.x {
			for i := min(from.y, to.y); i <= max(from.y, to.y); i++ {
				covered[Position{from.x, i}] += 1
			}
		} else if from.y == to.y {
			for i := min(from.x, to.x); i <= max(from.x, to.x); i++ {
				covered[Position{i, from.y}] += 1
			}
		} else if withDiagonal {
			left, right := findLeft(from, to)
			x := left.x
			y := left.y

			if left.y > right.y {
				for ; y >= right.y; y-- {
					covered[Position{x, y}] += 1
					x++
				}
			} else {
				for ; y <= right.y; y++ {
					covered[Position{x, y}] += 1
					x++
				}
			}
		}
	}
	return covered
}

func countMultipleCovered(covered map[Position]int64) int {
	sum := 0

	for _, coveredCount := range covered {
		if coveredCount > 1 {
			sum++
		}
	}
	return sum
}

func findLeft(from, to Position) (Position, Position) {
	if from.x < to.x {
		return from, to
	}
	return to, from
}

func readLine(row string) (Position, Position) {
	positions := strings.Split(row, " -> ")
	return readPosition(positions[0]), readPosition(positions[1])
}

func readPosition(position string) Position {
	pos := strings.Split(position, ",")
	x, _ := strconv.ParseInt(pos[0], 10, 64)
	y, _ := strconv.ParseInt(pos[1], 10, 64)
	return Position{x, y}
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
