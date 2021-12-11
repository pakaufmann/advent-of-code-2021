package day11

import (
	"advent-of-code-2021/input"
	"container/list"
	"fmt"
	"strconv"
)

type Octopus struct {
	x       int
	y       int
	energy  int
	flashed bool
}

func Day11() {
	lines, err := input.ReadLinesString("inputs/day11.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		fmt.Println(part1(readGrid(lines)))
		fmt.Println(part2(readGrid(lines)))
	}
}

func readGrid(lines []string) [][]*Octopus {
	grid := make([][]*Octopus, len(lines))

	for y, line := range lines {
		grid[y] = make([]*Octopus, len(line))

		for x, octopus := range line {
			energy, _ := strconv.ParseInt(string(octopus), 10, 64)
			grid[y][x] = &Octopus{x, y, int(energy), false}
		}
	}

	return grid
}

func part1(grid [][]*Octopus) int {
	flashes := 0

	for i := 0; i < 100; i++ {
		flashes += runStep(grid)
	}

	return flashes
}

func part2(grid [][]*Octopus) interface{} {
	for i := 1; ; i++ {
		flashes := runStep(grid)

		if flashes == len(grid)*len(grid[0]) {
			return i
		}
	}
}

func runStep(grid [][]*Octopus) int {
	flashes := 0
	flashed := increaseEnergyLevels(grid)

	for next := flashed.Front(); next != nil; next = next.Next() {
		octopus := next.Value.(*Octopus)
		octopus.energy = 0
		flashes += 1
		flashed.PushBackList(updateNeighbours(grid, octopus))
	}

	return flashes
}

func updateNeighbours(grid [][]*Octopus, octopus *Octopus) *list.List {
	flashedNeighbours := list.New()
	x := octopus.x
	y := octopus.y

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if y+dy < 0 || x+dx < 0 || y+dy >= len(grid) || x+dx >= len(grid[0]) {
				continue
			}

			neighbour := grid[y+dy][x+dx]
			if neighbour.flashed {
				continue
			}

			neighbour.energy++
			if neighbour.energy > 9 {
				neighbour.flashed = true
				flashedNeighbours.PushBack(neighbour)
			}
		}
	}

	return flashedNeighbours
}

func increaseEnergyLevels(grid [][]*Octopus) *list.List {
	flashed := list.New()

	for _, line := range grid {
		for _, octopus := range line {
			octopus.energy += 1
			octopus.flashed = false

			if octopus.energy > 9 {
				octopus.flashed = true
				flashed.PushBack(octopus)
			}
		}
	}

	return flashed
}
