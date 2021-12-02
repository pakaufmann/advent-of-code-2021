package input

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLinesInt(path string) ([]int64, error) {
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

func ReadLinesString(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
