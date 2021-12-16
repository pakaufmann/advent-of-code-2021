package day16

import (
	"advent-of-code-2021/input"
	"fmt"
	"math"
	"strconv"
)

type Expression interface {
	countVersion() int
	calculate() int
}

type Literal struct {
	version int
	id      int
	number  int
}

func (literal Literal) countVersion() int {
	return literal.version
}

func (literal Literal) calculate() int {
	return literal.number
}

type Operator struct {
	version    int
	typeId     int
	subPackets []Expression
}

func (operator Operator) countVersion() int {
	count := operator.version
	for _, subPacket := range operator.subPackets {
		count += subPacket.countVersion()
	}
	return count
}

func (operator *Operator) sum() int {
	sum := 0
	for _, subPacket := range operator.subPackets {
		sum += subPacket.calculate()
	}
	return sum
}

func (operator *Operator) product() int {
	product := 1
	for _, subPacket := range operator.subPackets {
		product *= subPacket.calculate()
	}
	return product
}

func (operator *Operator) min() int {
	min := math.MaxInt
	for _, subPacket := range operator.subPackets {
		if result := subPacket.calculate(); result < min {
			min = result
		}
	}
	return min
}

func (operator *Operator) max() int {
	max := 0
	for _, subPacket := range operator.subPackets {
		if result := subPacket.calculate(); result > max {
			max = result
		}
	}
	return max
}

func (operator *Operator) greaterThan() int {
	if operator.subPackets[0].calculate() > operator.subPackets[1].calculate() {
		return 1
	}
	return 0
}

func (operator *Operator) lessThan() int {
	if operator.subPackets[0].calculate() < operator.subPackets[1].calculate() {
		return 1
	}
	return 0
}

func (operator *Operator) equalTo() int {
	if operator.subPackets[0].calculate() == operator.subPackets[1].calculate() {
		return 1
	}
	return 0
}

func (operator Operator) calculate() int {
	switch operator.typeId {
	case 0:
		return operator.sum()
	case 1:
		return operator.product()
	case 2:
		return operator.min()
	case 3:
		return operator.max()
	case 5:
		return operator.greaterThan()
	case 6:
		return operator.lessThan()
	case 7:
		return operator.equalTo()
	}
	return 0
}

var binaryMapping = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func Day16() {
	lines, err := input.ReadLinesString("inputs/day16.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		expression := parseInput(lines[0])
		fmt.Println(expression.countVersion())
		fmt.Println(expression.calculate())
	}
}

func parseInput(input string) Expression {
	binaryInput := ""

	for _, number := range input {
		binaryInput += binaryMapping[number]
	}

	packet, _ := readPacket(binaryInput)
	return packet
}

func readPacket(input string) (Expression, int) {
	version, _ := strconv.ParseInt(input[0:3], 2, 64)
	typeId, _ := strconv.ParseInt(input[3:6], 2, 64)

	if typeId == 4 {
		number, numLength := parseNumber(input[6:])
		return Literal{int(version), int(typeId), number}, 6 + numLength
	}

	var subPackets []Expression
	var endPosition int

	if input[6] == '0' {
		subPackets, endPosition = readSubPacketsByLength(input[7:])
	} else {
		subPackets, endPosition = readSubPacketsByCount(input[7:])
	}
	return Operator{int(version), int(typeId), subPackets}, 7 + endPosition
}

func readSubPacketsByCount(input string) ([]Expression, int) {
	numSubPackets, _ := strconv.ParseInt(input[0:11], 2, 64)

	i := 11
	subPackets := make([]Expression, 0)
	for num := 0; num < int(numSubPackets); num++ {
		packet, newPos := readPacket(input[i:])
		i += newPos
		subPackets = append(subPackets, packet)
	}
	return subPackets, i
}

func readSubPacketsByLength(input string) ([]Expression, int) {
	length, _ := strconv.ParseInt(input[0:15], 2, 64)

	pos := 0
	subPackets := make([]Expression, 0)
	for pos < int(length) {
		packet, newPos := readPacket(input[15+pos:])
		pos += newPos
		subPackets = append(subPackets, packet)
	}

	return subPackets, 15 + pos
}

func parseNumber(input string) (int, int) {
	binaryNumber := ""
	digitPos := 0

	for {
		binaryNumber += input[digitPos+1 : digitPos+5]
		isLast := input[digitPos]
		digitPos += 5
		if isLast == '0' {
			break
		}
	}

	number, _ := strconv.ParseInt(binaryNumber, 2, 64)
	return int(number), digitPos
}
