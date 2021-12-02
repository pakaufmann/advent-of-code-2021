package day2

import (
	"advent-of-code-2021/input"
	"fmt"
	"strconv"
	"strings"
)

func Day2() {
	commands, err := input.ReadLinesString("inputs/day2.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		fmt.Println(part1(commands))
		fmt.Print(part2(commands))
	}
}

func part1(commands []string) int64 {
	var x int64 = 0
	var y int64 = 0

	for _, line := range commands {
		command := strings.Split(line, " ")
		num, _ := strconv.ParseInt(command[1], 10, 64)

		switch command[0] {
		case "forward":
			x += num
		case "down":
			y += num
		case "up":
			y -= num
		}
	}

	return x * y
}

func part2(commands []string) int64 {
	var x int64 = 0
	var y int64 = 0
	var aim int64 = 0

	for _, line := range commands {
		command := strings.Split(line, " ")
		num, _ := strconv.ParseInt(command[1], 10, 64)

		switch command[0] {
		case "forward":
			x += num
			y += num * aim
		case "down":
			aim += num
		case "up":
			aim -= num
		}
	}

	return x * y
}
