package day14

import (
	"advent-of-code-2021/input"
	"fmt"
	"math"
	"strings"
)

type Pair struct {
	first  string
	second string
}

type Polymer struct {
	evenPairs map[string]int64
	oddPairs  map[string]int64
	last      rune
}

func Day14() {
	lines, err := input.ReadLinesString("inputs/day14.txt")

	if err != nil {
		fmt.Print("Could not find file")
	} else {
		start, rules := readInput(lines)

		fmt.Println(part1(start, rules))
		fmt.Println(part2(start, rules))
	}
}

func part1(start Polymer, rules map[string]Pair) int64 {
	polymer := createPolymer(start, rules, 10)
	return polymer.readCount()
}

func part2(start Polymer, rules map[string]Pair) int64 {
	polymer := createPolymer(start, rules, 40)
	return polymer.readCount()
}

func createPolymer(polymer Polymer, rules map[string]Pair, count int) Polymer {
	for i := 0; i < count; i++ {
		polymer = polymer.expand(rules)
	}
	return polymer
}

func (polymer *Polymer) readCount() int64 {
	elements := make(map[rune]int64, 0)
	elements[polymer.last] += 1

	for pair, count := range polymer.evenPairs {
		elements[rune(pair[0])] += count
		elements[rune(pair[1])] += count
	}

	var leastCommon int64 = math.MaxInt64
	mostCommon := int64(0)

	for _, count := range elements {
		if count < leastCommon {
			leastCommon = count
		}
		if count > mostCommon {
			mostCommon = count
		}
	}

	return mostCommon - leastCommon
}

func (polymer Polymer) expand(rules map[string]Pair) Polymer {
	newPolymer := Polymer{make(map[string]int64), make(map[string]int64), polymer.last}
	newPolymer.addNewPairs(polymer.evenPairs, rules)
	newPolymer.addNewPairs(polymer.oddPairs, rules)
	return newPolymer
}

func (polymer *Polymer) addNewPairs(pairs map[string]int64, rules map[string]Pair) {
	for pair, count := range pairs {
		insert := rules[pair]

		polymer.evenPairs[insert.first] += count
		polymer.oddPairs[insert.second] += count
	}
}

func readInput(lines []string) (Polymer, map[string]Pair) {
	evenPairs := make(map[string]int64, 0)
	oddPairs := make(map[string]int64, 0)

	startPolymer := lines[0]
	last := ' '
	for i := 0; i < len(startPolymer)-1; i++ {
		if i&2 == 0 {
			evenPairs[startPolymer[i:i+2]] += 1
		} else {
			oddPairs[startPolymer[i:i+2]] += 1
		}
		last = rune(startPolymer[i+1])
	}

	rules := make(map[string]Pair, 0)

	for _, line := range lines[2:] {
		segments := strings.Split(line, " -> ")
		pair := segments[0]
		rules[pair] = Pair{
			string(pair[0]) + segments[1],
			segments[1] + string(pair[1]),
		}
	}

	return Polymer{evenPairs, oddPairs, last}, rules
}
