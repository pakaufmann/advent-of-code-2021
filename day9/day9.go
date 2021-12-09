package day9

import (
	"advent-of-code-2021/input"
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
)

type Point struct {
	x      int
	y      int
	height int
}

func Day9() {
	lines, err := input.ReadLinesString("inputs/day9.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		heightMap := createHeightMap(lines)
		lowPoints := lowPoints(heightMap)
		fmt.Println(part1(lowPoints))
		fmt.Println(part2(lowPoints, heightMap))
	}
}

func createHeightMap(lines []string) [][]int {
	heightMap := make([][]int, len(lines))

	for y, line := range lines {
		heightMap[y] = make([]int, len(line))
		for x, point := range line {
			height, _ := strconv.ParseInt(string(point), 10, 64)
			heightMap[y][x] = int(height)
		}
	}

	return heightMap
}

func part1(lowPoints []Point) int {
	sum := 0
	for _, point := range lowPoints {
		sum += 1 + point.height
	}
	return sum
}

func lowPoints(heightMap [][]int) []Point {
	lowPoints := make([]Point, 0)

	for y, line := range heightMap {
		for x, height := range line {
			top := getHeight(heightMap, x, y-1)
			left := getHeight(heightMap, x-1, y)
			bottom := getHeight(heightMap, x, y+1)
			right := getHeight(heightMap, x+1, y)

			if top > height && left > height && right > height && bottom > height {
				lowPoints = append(lowPoints, Point{x, y, height})
			}
		}
	}

	return lowPoints
}

func part2(lowPoints []Point, heightMap [][]int) int {
	basins := make([]int, 0)

	for _, lowPoint := range lowPoints {
		basins = append(basins, findBasinSize(lowPoint, heightMap))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	return basins[0] * basins[1] * basins[2]
}

func findBasinSize(point Point, heightMap [][]int) int {
	remaining := list.New()
	next := remaining.PushFront(point)
	visited := make(map[Point]interface{}, 0)
	size := 0

	for ; next != nil; next = next.Next() {
		nextPoint := next.Value.(Point)
		if _, exists := visited[nextPoint]; exists {
			continue
		}
		visited[nextPoint] = true

		if nextPoint.height < 9 {
			size += 1
			addPoint(getPoint(heightMap, nextPoint.x, nextPoint.y-1), remaining)
			addPoint(getPoint(heightMap, nextPoint.x-1, nextPoint.y), remaining)
			addPoint(getPoint(heightMap, nextPoint.x, nextPoint.y+1), remaining)
			addPoint(getPoint(heightMap, nextPoint.x+1, nextPoint.y), remaining)
		}
	}

	return size
}

func addPoint(point *Point, remaining *list.List) {
	if point != nil {
		remaining.PushBack(*point)
	}
}

func getPoint(heightMap [][]int, x int, y int) *Point {
	height := getHeight(heightMap, x, y)
	if height == math.MaxInt {
		return nil
	}
	return &Point{x, y, height}
}

func getHeight(heightMap [][]int, x int, y int) int {
	if x < 0 || y < 0 {
		return math.MaxInt
	}
	if x >= len(heightMap[0]) || y >= len(heightMap) {
		return math.MaxInt
	}
	return heightMap[y][x]
}
