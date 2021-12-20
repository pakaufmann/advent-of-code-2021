package day20

import (
	"advent-of-code-2021/input"
	"fmt"
)

type Position struct {
	x int
	y int
}

type Image struct {
	algorithm     string
	pixels        map[Position]bool
	minX          int
	maxX          int
	minY          int
	maxY          int
	switchEmpties bool
	round         int
}

func (image Image) enhance() Image {
	pixels := make(map[Position]bool, len(image.pixels))

	for x := image.minX - 1; x < image.maxX+1; x++ {
		for y := image.minY - 1; y < image.maxY+1; y++ {
			pos := 0

			for yDelta := -1; yDelta <= 1; yDelta++ {
				for xDelta := -1; xDelta <= 1; xDelta++ {
					pos = image.getPixel(x+xDelta, y+yDelta, pos<<1)
				}
			}

			pixels[Position{x, y}] = rune(image.algorithm[pos]) == '#'
		}
	}

	return Image{
		image.algorithm,
		pixels,
		image.minX - 2,
		image.maxX + 1,
		image.minY - 2,
		image.maxY + 2,
		image.switchEmpties,
		image.round + 1,
	}
}

func (image Image) getPixel(x int, y int, in int) int {
	if lit, ok := image.pixels[Position{x, y}]; ok {
		if lit {
			return in + 1
		} else {
			return in
		}
	}

	if image.switchEmpties && image.round%2 == 1 {
		return in + 1
	}

	return in
}

func (image Image) enhanceTimes(times int) int {
	img := image

	for img.round < times {
		img = img.enhance()
	}

	count := 0
	for _, lit := range img.pixels {
		if lit {
			count++
		}
	}
	return count
}

func Day20() {
	lines, err := input.ReadLinesString("inputs/day20.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		image := readImage(lines)
		fmt.Println(part1(image))
		fmt.Println(part2(image))
	}
}

func part1(image Image) int {
	return image.enhanceTimes(2)
}

func part2(image Image) int {
	return image.enhanceTimes(50)
}

func readImage(lines []string) Image {
	algorithm := lines[0]

	pixels := make(map[Position]bool, 0)

	for y, line := range lines[2:] {
		for x, pixel := range line {
			if rune(pixel) == '#' {
				pixels[Position{x, y}] = true
			} else {
				pixels[Position{x, y}] = false
			}
		}
	}

	return Image{
		algorithm,
		pixels,
		0,
		len(lines[2]),
		0,
		len(lines[2:]),
		lines[0][0] == '#',
		0,
	}
}
