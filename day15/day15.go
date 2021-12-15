package day15

import (
	"advent-of-code-2021/input"
	"container/heap"
	"errors"
	"fmt"
	"math"
	"strconv"
)

type Chiton struct {
	x         int
	y         int
	risk      int
	totalRisk int
}

func Day15() {
	lines, err := input.ReadLinesString("inputs/day15.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		fmt.Println(part1(readCave(lines)))
		fmt.Println(part2(readCave(lines)))
	}
}

func part2(cave [][]*Chiton) (int, error) {
	return findPath(enlargeCave(cave, 5))
}

func enlargeCave(cave [][]*Chiton, repeat int) [][]*Chiton {
	enlargedCave := make([][]*Chiton, len(cave)*repeat)

	height := len(cave)
	length := len(cave[0])

	for y, line := range cave {
		for ry := 0; ry < repeat; ry++ {
			finalY := y + ry*height
			enlargedCave[finalY] = make([]*Chiton, length*repeat)

			for x, chiton := range line {
				for rx := 0; rx < repeat; rx++ {
					finalX := x + rx*length

					newRisk := chiton.risk + rx + ry
					overflows := newRisk / 10
					if newRisk > 9 {
						newRisk = newRisk%10 + overflows%10
					}
					enlargedCave[finalY][finalX] = &Chiton{finalX, finalY, newRisk, chiton.totalRisk}
				}
			}
		}
	}
	return enlargedCave
}

func part1(cave [][]*Chiton) (int, error) {
	return findPath(cave)
}

func findPath(cave [][]*Chiton) (int, error) {
	cave[0][0].totalRisk = 0

	height := len(cave)
	length := len(cave[0])

	chitonHeap := &ChitonHeap{}
	heap.Init(chitonHeap)

	for chiton := cave[0][0]; chiton != nil; chiton = heap.Pop(chitonHeap).(*Chiton) {
		if chiton.x == length-1 && chiton.y == height-1 {
			return chiton.totalRisk, nil
		}

		if chiton.x < length-1 {
			addChiton(chitonHeap, chiton, cave[chiton.y][chiton.x+1])
		}
		if chiton.y < height-1 {
			addChiton(chitonHeap, chiton, cave[chiton.y+1][chiton.x])
		}
		if chiton.x > 0 {
			addChiton(chitonHeap, chiton, cave[chiton.y][chiton.x-1])
		}
		if chiton.y > 0 {
			addChiton(chitonHeap, chiton, cave[chiton.y-1][chiton.x])
		}
	}

	return 0, errors.New("no path found")
}

func addChiton(chitonHeap *ChitonHeap, chiton *Chiton, next *Chiton) {
	if next.totalRisk > next.risk+chiton.totalRisk {
		next.totalRisk = next.risk + chiton.totalRisk
		heap.Push(chitonHeap, next)
	}
}

func readCave(lines []string) [][]*Chiton {
	cave := make([][]*Chiton, len(lines))

	for y, line := range lines {
		cave[y] = make([]*Chiton, len(line))

		for x, pos := range line {
			chiton, _ := strconv.ParseInt(string(pos), 10, 64)
			cave[y][x] = &Chiton{x, y, int(chiton), math.MaxInt}
		}
	}

	return cave
}

type ChitonHeap []*Chiton

func (h ChitonHeap) Len() int {
	return len(h)
}

func (h ChitonHeap) Less(i, j int) bool {
	return h[i].totalRisk < h[j].totalRisk
}

func (h ChitonHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ChitonHeap) Push(x interface{}) {
	*h = append(*h, x.(*Chiton))
}

func (h *ChitonHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
