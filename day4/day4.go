package day4

import (
	"advent-of-code-2021/input"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Board struct {
	rows    [5]map[int64]bool
	columns [5]map[int64]bool
}

func (board *Board) removeNumber(number int64) {
	for _, row := range board.rows {
		delete(row, number)
	}
	for _, column := range board.columns {
		delete(column, number)
	}
}

func (board *Board) hasWon() bool {
	return hasWon(board.rows) || hasWon(board.columns)
}

func hasWon(fields [5]map[int64]bool) bool {
	for _, field := range fields {
		if len(field) == 0 {
			return true
		}
	}
	return false
}

func (board *Board) sum() int64 {
	var sum int64 = 0

	for _, row := range board.rows {
		for num := range row {
			sum += num
		}
	}

	return sum
}

func Day4() {
	rows, err := input.ReadLinesString("inputs/day4.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		numbers := strings.Split(rows[0], ",")
		boards := buildBoards(rows[2:])

		fmt.Println(part1(numbers, boards))
		fmt.Println(part2(numbers, boards))
	}
}

func part1(numbers []string, boards []Board) (int64, error) {
	for _, num := range numbers {
		number, _ := strconv.ParseInt(num, 10, 64)

		for _, board := range boards {
			board.removeNumber(number)

			if board.hasWon() {
				return board.sum() * number, nil
			}
		}
	}

	return 0, errors.New("could not find winning board")
}

func part2(numbers []string, boards []Board) (int64, error) {
	for _, num := range numbers {
		number, _ := strconv.ParseInt(num, 10, 64)
		remainingBoards := make([]Board, 0)

		for _, board := range boards {
			board.removeNumber(number)

			if !board.hasWon() {
				remainingBoards = append(remainingBoards, board)
			}
		}

		if len(remainingBoards) == 0 {
			return boards[0].sum() * number, nil
		}

		boards = remainingBoards
	}

	return 0, errors.New("could not loosing board")
}

func buildBoards(rows []string) []Board {
	boards := make([]Board, 0)

	for i := 0; i < len(rows); i += 6 {
		boards = append(boards, buildBoard(rows[i:i+5]))
	}

	return boards
}

func buildBoard(input []string) Board {
	re := regexp.MustCompile(" +")

	board := Board{createArray(), createArray()}

	for y, row := range input {
		for x, field := range re.Split(strings.Trim(row, " "), -1) {
			number, _ := strconv.ParseInt(field, 10, 64)
			board.rows[y][number] = true
			board.columns[x][number] = true
		}
	}
	return board
}

func createArray() [5]map[int64]bool {
	var arr [5]map[int64]bool
	for i := range arr {
		arr[i] = make(map[int64]bool)
	}
	return arr
}
