package day10

import (
	"advent-of-code-2021/input"
	"fmt"
	"sort"
)

var pairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func Day10() {
	lines, err := input.ReadLinesString("inputs/day10.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		corrupt, incompletes := getCorruptAndIncompletes(lines)
		fmt.Println(corrupt)
		fmt.Println(part2(incompletes))
	}
}

func part2(incompletes []stack) int {
	points := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	totalScores := make([]int, 0)

	for _, incomplete := range incompletes {
		totalScore := 0

		for len(incomplete) != 0 {
			var opening rune
			incomplete, opening = incomplete.Pop()
			totalScore = totalScore*5 + points[pairs[opening]]
		}

		totalScores = append(totalScores, totalScore)
	}

	sort.Ints(totalScores)

	return totalScores[len(totalScores)/2]
}

func getCorruptAndIncompletes(lines []string) (int, []stack) {
	incomplete := make([]stack, 0)

	sum := 0

	for _, line := range lines {
		incompleteStack, corruptSum := readLine(line)

		if corruptSum == 0 {
			incomplete = append(incomplete, incompleteStack)
		} else {
			sum += corruptSum
		}
	}

	return sum, incomplete
}

func readLine(line string) (stack, int) {
	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	stack := make(stack, 0)

	for _, paren := range line {
		if _, ok := pairs[paren]; ok {
			stack = stack.Push(paren)
		} else {
			var last rune
			stack, last = stack.Pop()

			if pairs[last] != paren {
				return stack, points[paren]
			}
		}
	}
	return stack, 0
}

type stack []rune

func (s stack) Push(v rune) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, rune) {
	l := len(s)
	return s[:l-1], s[l-1]
}
