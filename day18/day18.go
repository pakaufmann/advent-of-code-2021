package day18

import (
	"advent-of-code-2021/input"
	"fmt"
	"strconv"
)

type Number struct {
	previous *Number
	next     *Number
	number   int
	level    int
}

func Day18() {
	lines, err := input.ReadLinesString("inputs/day18.txt")

	if err != nil {
		fmt.Print("Could not find file")
	} else {
		fmt.Println(part1(readNumbers(lines)).magnitude())
		fmt.Println(part2(readNumbers(lines)))
	}
}

func part2(numbers []*Number) int {
	maxMagnitude := 0

	for _, number1 := range numbers {
		for _, number2 := range numbers {
			if number1 == number2 {
				continue
			}

			result := number1.copy().add(number2.copy()).reduce().magnitude()
			if result > maxMagnitude {
				maxMagnitude = result
			}

			result = number2.copy().add(number1.copy()).reduce().magnitude()
			if result > maxMagnitude {
				maxMagnitude = result
			}
		}
	}

	return maxMagnitude
}

func (number *Number) copy() *Number {
	first := &Number{nil, nil, number.number, number.level}
	previous := first
	current := number.next

	for current != nil {
		copied := &Number{previous, nil, current.number, current.level}
		previous.next = copied
		previous = copied
		current = current.next
	}

	return first
}

func part1(numbers []*Number) *Number {
	sum := numbers[0]

	for i, number := range numbers {
		if i == 0 {
			continue
		}
		sum = sum.add(number).reduce()
	}
	return sum
}

func (number *Number) magnitude() int {
	start := number
	current := number

	for current.next != nil {
		next := current.next

		if next != nil && next.level == current.level {
			sum := current.number*3 + next.number*2

			newNumber := &Number{current.previous, next, sum, current.level - 1}
			newNumber.previous = current.previous
			if current.previous != nil {
				current.previous.next = newNumber
			}

			newNumber.next = next.next
			if next.next != nil {
				next.next.previous = newNumber
			}

			if current.previous == nil {
				start = newNumber
			}
			current = start
		} else {
			current = current.next
		}
	}

	return current.number
}

func (number *Number) add(second *Number) *Number {
	last := number

	for last.next != nil {
		last = last.next
	}

	last.next = second
	second.previous = last

	number.increaseLevel()

	return number
}

func (number *Number) increaseLevel() {
	cur := number
	for cur != nil {
		cur.level += 1
		cur = cur.next
	}
}

func (number *Number) reduce() *Number {
	first := number

	hasExploded := true
	hasSplit := true

	for hasExploded || hasSplit {
		hasExploded = false
		hasSplit = false

		first, hasExploded = first.explode()
		if !hasExploded {
			first, hasSplit = first.split()
		}
	}

	return first
}

func (number *Number) explode() (*Number, bool) {
	current := number

	for current != nil {
		if current.level >= 4 {
			start := number
			second := current.next

			newNum := &Number{current.previous, second.next, 0, current.level - 1}

			if current.previous != nil {
				current.previous.number += current.number
				current.previous.next = newNum
			} else {
				start = newNum
			}

			if second.next != nil {
				second.next.number += second.number
				second.next.previous = newNum
			}

			return start, true
		} else {
			current = current.next
		}
	}
	return number, false
}

func (number *Number) split() (*Number, bool) {
	current := number

	for current != nil {
		if current.number > 9 {
			start := number
			leftNum := &Number{nil, nil, current.number / 2, current.level + 1}
			rightNum := &Number{nil, nil, (current.number + 1) / 2, current.level + 1}

			leftNum.next = rightNum
			leftNum.previous = current.previous
			if current.previous != nil {
				current.previous.next = leftNum
			} else {
				start = leftNum
			}

			rightNum.previous = leftNum
			rightNum.next = current.next
			if current.next != nil {
				current.next.previous = rightNum
			}
			return start, true
		} else {
			current = current.next
		}
	}

	return number, false
}

func readNumbers(lines []string) []*Number {
	pairs := make([]*Number, 0)

	for _, line := range lines {
		pairs = append(pairs, readNumber(line))
	}

	return pairs
}

func readNumber(line string) *Number {
	var current *Number
	var first *Number
	level := -1

	for _, char := range line {
		if char == '[' {
			level++
		} else if char == ']' {
			level--
		} else if char == ',' {
			// do nothing
		} else {
			number, _ := strconv.ParseInt(string(char), 10, 64)

			newNum := &Number{current, nil, -1, level}

			if current != nil {
				current.next = newNum
			} else {
				first = newNum
			}
			current = newNum

			current.number = int(number)
		}
	}

	return first
}
