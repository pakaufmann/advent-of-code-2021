package day13

import (
	"advent-of-code-2021/input"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	horizontal bool
	at         int
}

type Point struct {
	x int
	y int
}

func Day13() {
	lines, err := input.ReadLinesString("inputs/day13.txt")

	if err != nil {
		fmt.Print("Could not read file")
	} else {
		sheet, instructions := readSheetAndInstructions(lines)
		fmt.Println(part1(sheet, instructions))
		fmt.Println(part2(sheet, instructions))
	}
}

type Sheet struct {
	points      []*Point
	bottomRight Point
}

func part1(sheet Sheet, instructions []Instruction) int {
	sheet.fold(instructions[0])
	return sheet.countPoints()
}

func part2(sheet Sheet, instructions []Instruction) string {
	for _, instruction := range instructions {
		sheet.fold(instruction)
	}
	return sheet.showSheet()
}

func (sheet *Sheet) fold(instruction Instruction) {
	for _, point := range sheet.points {
		if instruction.horizontal {
			point.y = move(point.y, instruction.at)
		} else {
			point.x = move(point.x, instruction.at)
		}
	}

	if instruction.horizontal {
		sheet.bottomRight.y = instruction.at
	} else {
		sheet.bottomRight.x = instruction.at
	}
}

func move(pos int, at int) int {
	if pos > at {
		diff := pos - at
		return at - diff
	}
	return pos
}

func (sheet *Sheet) countPoints() int {
	return len(sheet.createPointMap())
}

func (sheet *Sheet) createPointMap() map[Point]interface{} {
	distinct := make(map[Point]interface{}, 0)

	for _, point := range sheet.points {
		distinct[*point] = nil
	}
	return distinct
}

func (sheet *Sheet) showSheet() string {
	pointMap := sheet.createPointMap()

	output := ""
	for y := 0; y < sheet.bottomRight.y; y++ {
		for x := 0; x < sheet.bottomRight.x; x++ {
			if _, ok := pointMap[Point{x, y}]; ok {
				output += "#"
			} else {
				output += "."
			}
		}
		output += "\n"
	}
	return output
}

func readSheetAndInstructions(lines []string) (Sheet, []Instruction) {
	sheet := Sheet{make([]*Point, 0), Point{0, 0}}
	instructions := make([]Instruction, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, "fold along") {
			segments := strings.Split(line, "=")
			horizontal := strings.Contains(segments[0], "y")
			at, _ := strconv.ParseInt(segments[1], 10, 64)

			instructions = append(instructions, Instruction{horizontal, int(at)})
		} else {
			coordinates := strings.Split(line, ",")

			x, _ := strconv.ParseInt(coordinates[0], 10, 64)
			y, _ := strconv.ParseInt(coordinates[1], 10, 64)
			sheet.points = append(sheet.points, &Point{int(x), int(y)})

			if int(x) > sheet.bottomRight.x {
				sheet.bottomRight.x = int(x)
			}
			if int(y) > sheet.bottomRight.y {
				sheet.bottomRight.y = int(y)
			}
		}
	}

	return sheet, instructions
}
