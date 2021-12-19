package day19

import (
	"advent-of-code-2021/input"
	"fmt"
	sort2 "sort"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
	z int
}

func (p *Position) minus(o *Position) *Position {
	return &Position{p.x - o.x, p.y - o.y, p.z - o.z}
}

func (p *Position) plus(o *Position) *Position {
	return &Position{p.x + o.x, p.y + o.y, p.z + o.z}
}

type Beacon struct {
	position  *Position
	distances map[string]*Beacon
	scanner   *Scanner
}

type Scanner struct {
	id      string
	beacons []*Beacon
}

func (scanner *Scanner) buildDistances() {
	for i, beacon := range scanner.beacons {
		for _, otherBeacon := range scanner.beacons[i+1:] {
			distance := beacon.position.minus(otherBeacon.position)
			key := distance.key()
			beacon.distances[key] = otherBeacon
			otherBeacon.distances[key] = beacon
		}
	}
}

func (scanner *Scanner) findOverlaps(otherScanner *Scanner, count int) map[*Beacon]*Beacon {
	for _, beacon := range scanner.beacons {

		for _, other := range otherScanner.beacons {
			overlaps := make(map[*Beacon]*Beacon, 0)
			overlaps[beacon] = other

			for distance, s1 := range beacon.distances {
				if s2, ok := other.distances[distance]; ok {
					overlaps[s1] = s2
				}
			}

			if len(overlaps) >= count {
				return overlaps
			}
		}
	}

	return nil
}

func Day19() {
	lines, err := input.ReadLinesString("inputs/day19.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		scanners := readScanners(lines)

		found, scannerCenters := buildMap(scanners)
		part1(found)
		part2(scannerCenters)
	}
}

func part2(scannerCenters []*Position) {
	maxDistance := 0

	for i, center := range scannerCenters {
		for _, other := range scannerCenters[i:] {
			distance := center.distanceTo(other)
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	fmt.Println(maxDistance)
}

func part1(found *Scanner) {
	seen := make(map[Position]interface{}, 0)
	count := 0

	for _, beacon := range found.beacons {
		if _, ok := seen[*beacon.position]; ok {
			continue
		}
		seen[*beacon.position] = nil
		count++
	}

	fmt.Println(count)
}

func buildMap(scanners []*Scanner) (*Scanner, []*Position) {
	scannerMap := scanners[0]
	remaining := scanners[1:]
	scannerCenters := make([]*Position, 0)

	for i := 0; i < len(remaining); i++ {
		scanner := remaining[i]
		overlaps := scannerMap.findOverlaps(scanner, 12)

		if overlaps != nil {
			rotation, offset := calcRotationAndOffset(overlaps)
			scannerCenters = append(scannerCenters, offset)

			for _, beacon := range scanner.beacons {
				newPos := rotation.rotate(beacon.position).plus(offset)
				beacon.position = newPos
				scannerMap.beacons = append(scannerMap.beacons, beacon)
			}

			remaining = append(remaining[:i], remaining[i+1:]...)
			i = -1
		}
	}
	return scannerMap, scannerCenters
}

func calcRotationAndOffset(overlaps map[*Beacon]*Beacon) (*Rotation, *Position) {
	for origin1, origin2 := range overlaps {
		for dist1, dest1 := range origin1.distances {
			dest2, ok := origin2.distances[dist1]

			if ok {
				diffNew := origin1.position.minus(dest1.position)
				diffOld := origin2.position.minus(dest2.position)

				if absInt(diffNew.x) == absInt(diffNew.y) || absInt(diffNew.y) == absInt(diffNew.z) || absInt(diffNew.x) == absInt(diffNew.z) {
					continue
				}

				rotation := diffOld.getMoves(diffNew)
				newPos := rotation.rotate(origin2.position)

				offset := origin1.position.minus(newPos)
				return rotation, offset
			}
		}
	}

	return nil, nil
}

func (p *Position) getZRotation(zNew int) func(*Position) int {
	if p.x == zNew {
		return func(b *Position) int { return b.x }
	}
	if p.x == zNew*-1 {
		return func(b *Position) int { return b.x * -1 }
	}
	if p.y == zNew {
		return func(b *Position) int { return b.y }
	}
	if p.y == zNew*-1 {
		return func(b *Position) int { return b.y * -1 }
	}
	if p.z == zNew {
		return func(b *Position) int { return b.z }
	}
	if p.z == zNew*-1 {
		return func(b *Position) int { return b.z * -1 }
	}
	return nil
}

func (p *Position) getYRotation(yNew int) func(*Position) int {
	if p.x == yNew {
		return func(b *Position) int { return b.x }
	}
	if p.x == yNew*-1 {
		return func(b *Position) int { return b.x * -1 }
	}
	if p.y == yNew {
		return func(b *Position) int { return b.y }
	}
	if p.y == yNew*-1 {
		return func(b *Position) int { return b.y * -1 }
	}
	if p.z == yNew {
		return func(b *Position) int { return b.z }
	}
	if p.z == yNew*-1 {
		return func(b *Position) int { return b.z * -1 }
	}
	return nil
}

func (p *Position) getXRotation(xNew int) func(*Position) int {
	if p.x == xNew {
		return func(b *Position) int { return b.x }
	}
	if p.x == xNew*-1 {
		return func(b *Position) int { return b.x * -1 }
	}
	if p.y == xNew {
		return func(b *Position) int { return b.y }
	}
	if p.y == xNew*-1 {
		return func(b *Position) int { return b.y * -1 }
	}
	if p.z == xNew {
		return func(b *Position) int { return b.z }
	}
	if p.z == xNew*-1 {
		return func(b *Position) int { return b.z * -1 }
	}
	return nil
}

type Rotation struct {
	x func(*Position) int
	y func(*Position) int
	z func(*Position) int
}

func (m Rotation) rotate(position *Position) *Position {
	return &Position{
		m.x(position),
		m.y(position),
		m.z(position),
	}
}

func (p *Position) getMoves(diffNew *Position) *Rotation {
	return &Rotation{
		p.getXRotation(diffNew.x),
		p.getYRotation(diffNew.y),
		p.getZRotation(diffNew.z),
	}
}

func (p *Position) key() string {
	sort := []int{absInt(p.x), absInt(p.y), absInt(p.z)}
	sort2.Ints(sort)
	key := ""

	for _, s := range sort {
		key += strconv.Itoa(s) + "_"
	}

	return key
}

func (p *Position) distanceTo(other *Position) int {
	distance := p.minus(other)
	return absInt(distance.x) + absInt(distance.y) + absInt(distance.z)
}

func readScanners(lines []string) []*Scanner {
	scanners := make([]*Scanner, 0)
	var currentScanner *Scanner

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, "--- scanner") {
			currentScanner = &Scanner{line, make([]*Beacon, 0)}
			scanners = append(scanners, currentScanner)
			continue
		}

		segments := strings.Split(line, ",")
		x, _ := strconv.ParseInt(segments[0], 10, 64)
		y, _ := strconv.ParseInt(segments[1], 10, 64)
		z, _ := strconv.ParseInt(segments[2], 10, 64)

		sensor := &Beacon{&Position{int(x), int(y), int(z)}, make(map[string]*Beacon, 0), currentScanner}
		currentScanner.beacons = append(currentScanner.beacons, sensor)
	}

	for _, scanner := range scanners {
		scanner.buildDistances()
	}

	return scanners
}

func absInt(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
