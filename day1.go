package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	depths, err := ReadLines("inputs/day1.txt")
	if err != nil {
		fmt.Println("Could not find file")
	} else {
		fmt.Println(part1(depths))
		fmt.Print(part2(depths))
	}
}

func part1(depths []int64) int {
	increased := 0
	var previous int64 = math.MaxInt64

	for _, depth := range depths {
		if depth > previous {
			increased++
		}
		previous = depth
	}

	return increased
}

func part2(depths []int64) int {
	increased := 0
	var lastSum int64 = math.MaxInt64

	for i, depth := range depths[:len(depths)-2] {
		sum := depth + depths[i+1] + depths[i+2]
		if sum > lastSum {
			increased++
		}
		lastSum = sum
	}

	return increased
}

func ReadLines(path string) ([]int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		lines = append(lines, number)
	}
	return lines, scanner.Err()
}
