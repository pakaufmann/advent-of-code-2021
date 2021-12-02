package day2

import (
	"advent-of-code-2021/input"
	"fmt"
	"strconv"
	"strings"
)

type Submarine interface {
	up(num int64)
	down(num int64)
	forward(num int64)
	result() int64
}

type Submarine1 struct {
	x int64
	y int64
}

func (sub *Submarine1) up(num int64) {
	sub.y -= num
}

func (sub *Submarine1) down(num int64) {
	sub.y += num
}

func (sub *Submarine1) forward(num int64) {
	sub.x += num
}

func (sub *Submarine1) result() int64 {
	return sub.x * sub.y
}

type Submarine2 struct {
	x   int64
	y   int64
	aim int64
}

func (sub *Submarine2) up(num int64) {
	sub.aim -= num
}

func (sub *Submarine2) down(num int64) {
	sub.aim += num
}

func (sub *Submarine2) forward(num int64) {
	sub.x += num
	sub.y += num * sub.aim
}

func (sub *Submarine2) result() int64 {
	return sub.x * sub.y
}

func Day2Dedup() {
	commands, err := input.ReadLinesString("inputs/day2.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		fmt.Println(run(commands, &Submarine1{0, 0}))
		fmt.Print(run(commands, &Submarine2{0, 0, 0}))
	}
}

func run(commands []string, submarine Submarine) int64 {
	for _, line := range commands {
		command := strings.Split(line, " ")
		num, _ := strconv.ParseInt(command[1], 10, 64)

		switch command[0] {
		case "forward":
			submarine.forward(num)
		case "down":
			submarine.down(num)
		case "up":
			submarine.up(num)
		}
	}

	return submarine.result()
}
