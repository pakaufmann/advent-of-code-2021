package day22

import (
	"advent-of-code-2021/input"
	"fmt"
	"regexp"
	"strconv"
)

type Range struct {
	from int
	to   int
}

func (r *Range) isOutside(o Range) bool {
	if o.to < r.from || o.from > r.to {
		return true
	}
	return false
}

func (r *Range) clamp(in int) int {
	if in < r.from {
		return r.from
	}
	if in > r.to {
		return r.to
	}
	return in
}

func (r *Range) len() int {
	return intAbs(r.to-r.from) + 1
}

func (r *Range) overlap(o Range) Range {
	return Range{intMax(r.from, o.from), intMin(r.to, o.to)}
}

func intAbs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

type Cube struct {
	x    Range
	y    Range
	z    Range
	sign int
}

func (cube *Cube) overlapping(other Cube) *Cube {
	if cube.x.isOutside(other.x) || cube.y.isOutside(other.y) || cube.z.isOutside(other.z) {
		return nil
	}

	return &Cube{
		cube.x.overlap(other.x),
		cube.y.overlap(other.y),
		cube.z.overlap(other.z),
		cube.sign * -1,
	}
}

func (cube *Cube) signedVolume() int {
	return cube.x.len() * cube.y.len() * cube.z.len() * cube.sign
}

func (cube *Cube) clamp(r Range) *Cube {
	if r.isOutside(cube.x) || r.isOutside(cube.y) || r.isOutside(cube.z) {
		return nil
	}

	return &Cube{
		Range{r.clamp(cube.x.from), r.clamp(cube.x.to)},
		Range{r.clamp(cube.y.from), r.clamp(cube.y.to)},
		Range{r.clamp(cube.z.from), r.clamp(cube.z.to)},
		cube.sign,
	}
}

func intMax(f int, s int) int {
	if f > s {
		return f
	}
	return s
}

func intMin(f int, s int) int {
	if f < s {
		return f
	}
	return s
}

func Day22() {
	lines, err := input.ReadLinesString("inputs/day22.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		instructions := readInstructions(lines)
		cubes := createOverlaps(instructions)

		fmt.Println(part1(cubes))
		fmt.Println(part2(cubes))
	}
}

func part1(cubes []Cube) int {
	return countVolumes(filterByClamping(cubes, Range{-50, 50}))
}

func part2(cubes []Cube) int {
	return countVolumes(cubes)
}

func filterByClamping(cubes []Cube, rangeToCheck Range) []Cube {
	filteredCubes := make([]Cube, 0)

	for _, cube := range cubes {
		clamped := cube.clamp(rangeToCheck)
		if clamped != nil {
			filteredCubes = append(filteredCubes, *clamped)
		}
	}
	return filteredCubes
}

func countVolumes(cubes []Cube) int {
	count := 0
	for _, cube := range cubes {
		count += cube.signedVolume()
	}
	return count
}

func createOverlaps(instructions []Cube) []Cube {
	cubes := make([]Cube, 0)

	for _, instruction := range instructions {
		for _, cube := range cubes {
			overlapping := cube.overlapping(instruction)

			if overlapping != nil {
				cubes = append(cubes, *overlapping)
			}
		}

		if instruction.sign > 0 {
			cubes = append(cubes, instruction)
		}
	}
	return cubes
}

func readInstructions(lines []string) []Cube {
	instructions := make([]Cube, 0)

	r := regexp.MustCompile(`(on|off) x=([-0-9]*)..([-0-9]*),y=([-0-9]*)..([-0-9]*),z=([-0-9]*)..([-0-9]*)`)

	for _, line := range lines {
		res := r.FindAllStringSubmatch(line, -1)[0]

		sign := -1
		if res[1] == "on" {
			sign = 1
		}

		instructions = append(instructions, Cube{
			createRange(res[2:4]),
			createRange(res[4:6]),
			createRange(res[6:8]),
			sign,
		})
	}

	return instructions
}

func createRange(r []string) Range {
	from, _ := strconv.ParseInt(r[0], 10, 64)
	to, _ := strconv.ParseInt(r[1], 10, 64)
	return Range{int(from), int(to)}
}
