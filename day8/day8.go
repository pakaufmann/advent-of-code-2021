package day8

import (
	"advent-of-code-2021/input"
	"fmt"
	"sort"
	"strings"
)

type Display struct {
	patterns []string
	digits   []string
}

func (display *Display) removeNumberConstrainedBy(length int, segments ...string) string {
	for i, pattern := range display.patterns {
		if len(pattern) == length && patternContainsAllSegments(pattern, segments) {
			display.patterns = append(display.patterns[:i], display.patterns[i+1:]...)
			return pattern
		}
	}

	return ""
}

func patternContainsAllSegments(pattern string, segments []string) bool {
	for _, segment := range segments {
		if !strings.Contains(pattern, segment) {
			return false
		}
	}
	return true
}

func Day8() {
	lines, err := input.ReadLinesString("inputs/day8.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		fmt.Println(part1(readSegments(lines)))
		fmt.Println(part2(readSegments(lines)))
	}
}

func readSegments(lines []string) []Display {
	result := make([]Display, 0)

	for _, line := range lines {
		segments := strings.Split(line, " | ")

		result = append(
			result,
			Display{strings.Split(segments[0], " "), strings.Split(segments[1], " ")},
		)
	}

	return result
}

func part1(displays []Display) int {
	sum := 0

	for _, segment := range displays {
		for _, digit := range segment.digits {
			length := len(digit)
			if length == 2 || length == 4 || length == 3 || length == 7 {
				sum += 1
			}
		}
	}

	return sum
}

func part2(displays []Display) int {
	sum := 0

	for _, display := range displays {
		one := display.removeNumberConstrainedBy(2)
		four := display.removeNumberConstrainedBy(4)
		seven := display.removeNumberConstrainedBy(3)
		eight := display.removeNumberConstrainedBy(7)
		three := display.removeNumberConstrainedBy(5, strings.Split(seven, "")...)
		nine := display.removeNumberConstrainedBy(6, strings.Split(three, "")...)
		zero := display.removeNumberConstrainedBy(6, strings.Split(seven, "")...)
		six := display.removeNumberConstrainedBy(6)
		five := display.removeNumberConstrainedBy(5, diffBetweenNumbers(one, six)...)
		two := display.removeNumberConstrainedBy(5)

		numbers := map[string]int{
			sortString(zero):  0,
			sortString(one):   1,
			sortString(two):   2,
			sortString(three): 3,
			sortString(four):  4,
			sortString(five):  5,
			sortString(six):   6,
			sortString(seven): 7,
			sortString(eight): 8,
			sortString(nine):  9,
		}

		number := 0

		for _, digit := range display.digits {
			number *= 10
			number += numbers[sortString(digit)]
		}

		sum += number
	}

	return sum
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func diffBetweenNumbers(first string, second string) (same []string) {
	same = make([]string, 0)

	for _, firstDigit := range first {
		if strings.ContainsRune(second, firstDigit) {
			same = append(same, string(firstDigit))
		}
	}
	return
}
