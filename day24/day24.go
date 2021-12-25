package day24

import (
	"advent-of-code-2021/input"
	"fmt"
	"strconv"
	"strings"
)

type ALU struct {
	a     []int
	b     []int
	c     []int
	maxZ  []int
	order []int
}

func Day24() {
	lines, err := input.ReadLinesString("inputs/day24.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		a, b, c := readInput(lines)

		maxZ := buildMaxZ(b)

		alu := &ALU{a, b, c, maxZ, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}}
		fmt.Println(part1(alu))
		fmt.Println(part2(alu))
	}
}

func part1(alu *ALU) int {
	_, result := alu.search(0, 0, [14]int{})
	return result
}

func part2(alu *ALU) int {
	alu.order = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	_, result := alu.search(0, 0, [14]int{})
	return result
}

func buildMaxZ(b []int) []int {
	var maxZ [14]int
	maxZ[13] = b[13]

	for i := len(b) - 2; i >= 0; i-- {
		maxZ[i] = maxZ[i+1] * b[i]
	}

	return maxZ[:]
}

func (alu *ALU) search(depth int, z int, solution [14]int) (bool, int) {
	if depth == 14 {
		if z == 0 {
			return true, buildSolution(solution)
		}
		return false, 0
	}

	if z >= alu.maxZ[depth] {
		return false, 0
	}

	for _, i := range alu.order {
		solution[depth] = i
		if found, sol := alu.search(depth+1, alu.stage(depth, i, z), solution); found {
			return true, sol
		}
	}

	return false, 0
}

func buildSolution(solution [14]int) int {
	sol := 0

	for _, digit := range solution {
		sol *= 10
		sol += digit
	}

	return sol
}

func (alu *ALU) stage(n int, w int, z int) int {
	if z%26+alu.a[n] == w {
		return z / alu.b[n]
	} else {
		return 26*(z/alu.b[n]) + w + alu.c[n]
	}
}

func readInput(lines []string) ([]int, []int, []int) {
	a := make([]int, 0)
	b := make([]int, 0)
	c := make([]int, 0)

	for i := 0; i < 14; i++ {
		block := lines[i*18 : 18*i+18]
		a = append(a, int(parseInt(block[5], "add x ")))
		b = append(b, int(parseInt(block[4], "div z ")))
		c = append(c, int(parseInt(block[15], "add y ")))
	}

	return a, b, c
}

func parseInt(block string, pre string) int64 {
	replace := strings.Replace(block, pre, "", -1)
	parsed, _ := strconv.ParseInt(replace, 10, 64)
	return parsed
}
