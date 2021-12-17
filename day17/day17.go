package day17

import (
	"advent-of-code-2021/input"
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	from int
	to   int
}

type Velocity struct {
	x int
	y int
}

func (velocity *Velocity) maxHeight() int {
	return (velocity.y * (velocity.y + 1)) / 2
}

type Area struct {
	xRange Range
	yRange Range
}

func (area *Area) isHitBy(xVel int, yVel int) bool {
	x := 0
	y := 0

	for x <= area.xRange.to && y >= area.yRange.to {
		x += xVel
		y += yVel
		xVel = maxInt(0, xVel-1)
		yVel--

		if area.in(x, y) {
			return true
		}
	}

	return false
}

func (area *Area) in(x int, y int) bool {
	return area.xRange.from <= x && area.xRange.to >= x &&
		area.yRange.from >= y && area.yRange.to <= y
}

func Day17() {
	lines, err := input.ReadLinesString("inputs/day17.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		targetArea := parseInput(lines)
		velocities := getVelocities(targetArea)
		fmt.Println(part1(velocities))
		fmt.Println(part2(velocities))
	}
}

func part1(velocities []Velocity) int {
	max := 0
	for _, velocity := range velocities {
		max = maxInt(velocity.maxHeight(), max)
	}

	return max
}

func part2(velocities []Velocity) int {
	return len(velocities)
}

func getVelocities(area Area) []Velocity {
	velocities := make([]Velocity, 0)

	maxX := absInt(area.xRange.to)
	maxY := absInt(area.yRange.to)
	for x := 1; x <= maxX; x++ {
		for y := area.yRange.to; y <= maxY; y++ {
			if area.isHitBy(x, y) {
				velocities = append(velocities, Velocity{x, y})
			}
		}
	}

	return velocities
}

func parseInput(lines []string) Area {
	segments := strings.Split(strings.Replace(lines[0], ",", "", -1), " ")
	yRange := parseRange(strings.Replace(segments[3], "y=", "", -1))
	return Area{
		parseRange(strings.Replace(segments[2], "x=", "", -1)),
		Range{yRange.to, yRange.from},
	}
}

func parseRange(r string) Range {
	segments := strings.Split(r, "..")
	from, _ := strconv.ParseInt(segments[0], 10, 64)
	to, _ := strconv.ParseInt(segments[1], 10, 64)

	return Range{int(from), int(to)}
}

func maxInt(f int, s int) int {
	if f > s {
		return f
	}
	return s
}

func absInt(i int) int {
	if i > 0 {
		return i
	}
	return i * -1
}
