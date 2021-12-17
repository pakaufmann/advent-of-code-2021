package day18_broken

import (
	"advent-of-code-2021/input"
	"fmt"
	"strconv"
)

type Number struct {
	number   int
	previous *Number
	next     *Number
	pair     *Pair
}

type Pair struct {
	parent *Pair
	left   interface{}
	right  interface{}
	level  int
}

func Day18_broken() {
	lines, err := input.ReadLinesString("inputs/day18_test.txt")

	if err != nil {
		fmt.Print("Could not find file")
	} else {
		numbers := readNumbers(lines)
		sum := numbers[0]
		sum.reduce()

		for i, number := range numbers {
			if i == 0 {
				continue
			}
			number.reduce()
			sum = sum.add(number)
			sum.reduce()
		}

		fmt.Println(sum)
	}
}

func (pair *Pair) reduce() {
	if pair.level >= 4 {
		leftNum := pair.left.(*Number)
		rightNum := pair.right.(*Number)
		newNum := &Number{0, leftNum.previous, rightNum.next, pair.parent}

		if pair.parent.left == pair {
			pair.parent.left = newNum
		} else {
			pair.parent.right = newNum
		}

		if leftNum.previous != nil {
			leftNum.previous.number += leftNum.number
			leftNum.previous.next = newNum
			leftNum.previous.pair.split()
		}

		if rightNum.next != nil {
			rightNum.next.number += rightNum.number
			rightNum.next.previous = newNum
			//rightNum.next.pair.split()
		}
	} else {
		pair.split()
	}

	switch pair.left.(type) {
	case *Pair:
		pair.left.(*Pair).reduce()
	case *Number:
	}

	switch pair.right.(type) {
	case *Pair:
		pair.right.(*Pair).reduce()
	case *Number:
	}
}

func (pair *Pair) split() {
	switch pair.left.(type) {
	case *Number:
		num := pair.left.(*Number)
		if num.number > 9 {
			leftNum := &Number{num.number / 2, num.previous, num.next, nil}
			rightNum := &Number{(num.number + 1) / 2, leftNum, num.next, nil}
			leftNum.next = rightNum
			newLeft := &Pair{
				pair,
				leftNum,
				rightNum,
				pair.level + 1,
			}
			leftNum.pair = newLeft
			rightNum.pair = newLeft
			if num.previous != nil {
				num.previous.next = leftNum
			}
			if num.next != nil {
				num.next.previous = rightNum
			}
			pair.left = newLeft
			newLeft.reduce()
		}
	}

	switch pair.right.(type) {
	case *Number:
		num := pair.right.(*Number)
		if num.number > 9 {
			leftNum := &Number{num.number / 2, num.previous, num.next, nil}
			rightNum := &Number{(num.number + 1) / 2, leftNum, num.next, nil}
			leftNum.next = rightNum
			newRight := &Pair{
				pair,
				leftNum,
				rightNum,
				pair.level + 1,
			}
			leftNum.pair = newRight
			rightNum.pair = newRight
			if num.previous != nil {
				num.previous.next = leftNum
			}
			if num.next != nil {
				num.next.previous = rightNum
			}
			pair.right = newRight
			newRight.reduce()
		}
	}
}

func (pair *Pair) add(number *Pair) *Pair {
	_, last := pair.increaseLevel()
	first, _ := number.increaseLevel()

	last.next = first
	first.previous = last

	root := &Pair{nil, pair, number, 0}
	pair.parent = root
	number.parent = root
	return root
}

func (pair *Pair) increaseLevel() (*Number, *Number) {
	pair.level += 1
	var first *Number
	var last *Number

	if pair.left != nil {
		switch pair.left.(type) {
		case *Pair:
			first, _ = pair.left.(*Pair).increaseLevel()
		case *Number:
			first = pair.left.(*Number)
		}
	}
	if pair.right != nil {
		switch pair.right.(type) {
		case *Pair:
			_, last = pair.right.(*Pair).increaseLevel()
		case *Number:
			last = pair.right.(*Number)
		}
	}

	return first, last
}

func readNumbers(lines []string) []*Pair {
	pairs := make([]*Pair, 0)

	for _, line := range lines {
		pairs = append(pairs, readNumber(line))
	}

	return pairs
}

func readNumber(line string) *Pair {
	stack := make(stack, 0)
	var current *Pair
	var currentNum *Number
	level := 0

	for _, char := range line {
		if char == '[' {
			pair := &Pair{current, nil, nil, level}
			level++

			if current != nil {
				stack = stack.Push(current)
				if current.left == nil {
					current.left = pair
				} else {
					current.right = pair
				}
			}

			current = pair
		} else if char == ']' {
			if len(stack) != 0 {
				level--
				stack, current = stack.Pop()
			}
		} else if char == ',' {
			// do nothing
		} else {
			number, _ := strconv.ParseInt(string(char), 10, 64)
			num := &Number{int(number), currentNum, nil, current}

			if currentNum != nil {
				currentNum.next = num
			}

			if current.left == nil {
				current.left = num
			} else {
				current.right = num
			}

			currentNum = num
		}
	}

	return current
}

type stack []*Pair

func (s stack) Push(v *Pair) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, *Pair) {
	l := len(s)
	return s[:l-1], s[l-1]
}
