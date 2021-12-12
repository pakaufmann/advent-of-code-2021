package day12

import (
	"advent-of-code-2021/input"
	"container/list"
	"fmt"
	"strings"
)

type Connections = map[string][]string

type Path struct {
	at             string
	visitedSmalls  map[string]interface{}
	multiVisit     string
	remainingCount int
	visited        string
}

func Day12() {
	lines, err := input.ReadLinesString("inputs/day12.txt")

	if err != nil {
		fmt.Println("Could not load file")
	} else {
		connections := readConnections(lines)
		fmt.Println(part1(connections))
		fmt.Println(part2(connections))
	}
}

func part1(connections Connections) int {
	return findPaths(connections, 1)
}

func part2(connections Connections) int {
	return findPaths(connections, 2)
}

func findPaths(connections Connections, count int) int {
	toCheck := list.New()
	toCheck.PushBackList(findStarts(connections, count))
	next := toCheck.Front()

	foundPaths := make(map[string]interface{}, 0)

	for ; next != nil; next = next.Next() {
		path := next.Value.(Path)

		possibilities := connections[path.at]

		if path.at == "end" {
			foundPaths[path.visited] = nil
			continue
		}

		for _, nextPosition := range possibilities {
			_, alreadyVisited := path.visitedSmalls[nextPosition]
			if nextPosition != path.multiVisit && (alreadyVisited || nextPosition == "start") {
				continue
			}

			remainingCount := path.remainingCount

			if path.multiVisit == nextPosition {
				if path.remainingCount == 0 {
					continue
				}
				remainingCount--
			}

			toCheck.PushBack(Path{
				nextPosition,
				appendTo(path.visitedSmalls, path.at),
				path.multiVisit,
				remainingCount,
				path.visited + path.at,
			})
		}
	}

	return len(foundPaths)
}

func appendTo(smalls map[string]interface{}, at string) map[string]interface{} {
	if strings.ToLower(at) == at {
		newSmalls := make(map[string]interface{}, 0)

		for k, v := range smalls {
			newSmalls[k] = v
		}

		newSmalls[at] = nil
		return newSmalls
	}

	return smalls
}

func findStarts(connections Connections, count int) *list.List {
	starts := list.New()

	for from := range connections {
		if from == "start" {
			for multi := range connections {
				if multi != "start" && multi != "end" && strings.ToLower(multi) == multi {
					starts.PushBack(Path{from, make(map[string]interface{}, 0), multi, count, ""})
				}
			}
		}
	}

	return starts
}

func readConnections(lines []string) Connections {
	connections := make(map[string][]string, 0)

	for _, line := range lines {
		segments := strings.Split(line, "-")
		from := segments[0]
		to := segments[1]

		addConnection(&connections, from, to)
		addConnection(&connections, to, from)
	}

	return connections
}

func addConnection(connections *map[string][]string, from string, to string) {
	tos, ok := (*connections)[from]

	if !ok {
		tos = make([]string, 0)
	}

	tos = append(tos, to)
	(*connections)[from] = tos
}
