package day23

import (
	"advent-of-code-2021/input"
	"container/heap"
	"fmt"
)

type Bucket struct {
	remaining int
	unsorted  []rune
	left      int
	right     int
}

func (bucket *Bucket) decreaseRemaining() *Bucket {
	return &Bucket{
		bucket.remaining - 1,
		bucket.unsorted,
		bucket.left,
		bucket.right,
	}
}

func (bucket *Bucket) removeTop() *Bucket {
	newUnsorted := make([]rune, len(bucket.unsorted))
	copy(newUnsorted, bucket.unsorted)

	return &Bucket{
		bucket.remaining,
		newUnsorted[:len(newUnsorted)-1],
		bucket.left,
		bucket.right,
	}
}

type Layout struct {
	parkings [7]rune
	buckets  [4]*Bucket
	cost     int
}

func (layout *Layout) finishedAll() bool {
	for _, bucket := range layout.buckets {
		if bucket.remaining > 0 {
			return false
		}
	}

	return true
}

func (layout *Layout) nextMoves() []*Layout {
	for parkedAt, parkedAmphipod := range layout.parkings {
		if parkedAmphipod == 0 {
			continue
		}

		toBucket := layout.getDestinationBucket(parkedAmphipod)
		if len(toBucket.unsorted) > 0 {
			continue
		}

		if parkedAt == toBucket.right || (parkedAt > toBucket.right && layout.canMove(parkedAt-1, toBucket.right)) {
			return []*Layout{
				layout.fromParkedToBucket(
					parkedAt,
					toBucket,
					toBucket.right-1,
				),
			}
		}
		if parkedAt == toBucket.left || (parkedAt < toBucket.left && layout.canMove(parkedAt+1, toBucket.left)) {
			return []*Layout{
				layout.fromParkedToBucket(
					parkedAt,
					toBucket,
					toBucket.left+1,
				),
			}
		}
	}

	layouts := make([]*Layout, 0)

	for from, fromBucket := range layout.buckets {
		if fromBucket.remaining == 0 || len(fromBucket.unsorted) == 0 {
			continue
		}
		unsortedAmphipod := fromBucket.unsorted[len(fromBucket.unsorted)-1]

		toBucket := layout.buckets[unsortedAmphipod-'A']
		if toBucket.canReceive() {
			if toBucket.right <= fromBucket.left && layout.canMove(fromBucket.left, toBucket.right) {
				return []*Layout{
					layout.moveFromBucketToBucket(
						from,
						unsortedAmphipod,
						cost(unsortedAmphipod, fromBucket.right, toBucket.left),
					)}
			}
			if toBucket.left >= fromBucket.right && layout.canMove(fromBucket.right, toBucket.left) {
				return []*Layout{
					layout.moveFromBucketToBucket(
						from,
						unsortedAmphipod,
						cost(unsortedAmphipod, fromBucket.left, toBucket.right),
					),
				}
			}
		}

		for parkingPosition := 0; parkingPosition <= fromBucket.left; parkingPosition++ {
			if layout.canMove(fromBucket.left, parkingPosition) {
				layouts = append(
					layouts,
					layout.fromBucketToParked(
						from,
						parkingPosition,
						unsortedAmphipod,
						fromBucket.left+1,
					),
				)
			}
		}

		for parkingPosition := fromBucket.right; parkingPosition < len(layout.parkings); parkingPosition++ {
			if layout.canMove(fromBucket.right, parkingPosition) {
				layouts = append(
					layouts,
					layout.fromBucketToParked(
						from,
						parkingPosition,
						unsortedAmphipod,
						fromBucket.right-1,
					),
				)
			}
		}
	}

	return layouts
}

func (layout *Layout) fromBucketToParked(from int, parkingPosition int, amphipod rune, costStart int) *Layout {
	fromBucket := layout.buckets[from]

	newBuckets := layout.buckets
	newBuckets[from] = fromBucket.removeTop()

	newParkings := layout.parkings
	newParkings[parkingPosition] = amphipod

	newLayout := &Layout{
		newParkings,
		newBuckets,
		layout.cost +
			cost(amphipod, costStart, parkingPosition) +
			(fromBucket.remaining-len(fromBucket.unsorted))*times[amphipod],
	}
	return newLayout
}

func (layout *Layout) fromParkedToBucket(parkedAt int, toBucket *Bucket, costEnd int) *Layout {
	parkedAmphipod := layout.parkings[parkedAt]

	newBuckets := layout.buckets
	newDestination := toBucket.decreaseRemaining()
	newBuckets[parkedAmphipod-'A'] = newDestination

	newParkings := layout.parkings
	newParkings[parkedAt] = 0

	newLayout := &Layout{
		newParkings,
		newBuckets,
		layout.cost +
			cost(parkedAmphipod, parkedAt, costEnd) +
			newDestination.remaining*times[parkedAmphipod],
	}
	return newLayout
}

func (layout *Layout) getDestinationBucket(amphipod rune) *Bucket {
	return layout.buckets[amphipod-'A']
}

func (bucket *Bucket) canReceive() bool {
	return len(bucket.unsorted) == 0
}

func (layout *Layout) moveFromBucketToBucket(from int, toBucket rune, moveCost int) *Layout {
	fromBucket := layout.buckets[from]
	newBuckets := layout.buckets

	destination := layout.buckets[toBucket-'A']
	newDestination := destination.decreaseRemaining()

	newBuckets[toBucket-'A'] = newDestination
	newBuckets[from] = fromBucket.removeTop()

	return &Layout{
		layout.parkings,
		newBuckets,
		layout.cost +
			moveCost +
			newDestination.remaining*times[toBucket] +
			(fromBucket.remaining-len(fromBucket.unsorted))*times[toBucket],
	}
}

var times = map[rune]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

func cost(amphipod rune, from int, to int) int {
	steps := intAbs(from-to) * 2
	if from == 0 || to == 0 || to == 6 || from == 6 {
		steps--
	}

	return steps * times[amphipod]
}

func intAbs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func (layout *Layout) canMove(from int, to int) bool {
	min, max := minMax(from, to)

	for _, parked := range layout.parkings[min : max+1] {
		if parked != 0 {
			return false
		}
	}

	return true
}

func (layout *Layout) hash() string {
	hash := ""

	for _, p := range layout.parkings {
		hash += string(p) + "_"
	}

	for _, bucket := range layout.buckets {
		for _, un := range bucket.unsorted {
			hash += string(un) + "_"
		}
		hash += "__"
	}

	return hash
}

func minMax(f int, s int) (int, int) {
	if f > s {
		return s, f
	}
	return f, s
}

type LayoutHeap []*Layout

func Day23() {
	lines, err := input.ReadLinesString("inputs/day23.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		fmt.Println(part1(readAmphipods(lines)))
		fmt.Println(part2(readAmphipods(expand(lines))))
	}
}

func expand(lines []string) []string {
	expanded := make([]string, 0)
	expanded = append(expanded, lines[0:3]...)

	expanded = append(expanded, "#D#C#B#A#")
	expanded = append(expanded, "#D#B#A#C#")
	expanded = append(expanded, lines[3:]...)

	return expanded
}

func part1(start *Layout) int {
	return findSolution(start)
}

func part2(start *Layout) int {
	return findSolution(start)
}

func findSolution(start *Layout) int {
	amphipodHeap := &LayoutHeap{}
	heap.Init(amphipodHeap)

	seen := make(map[string]interface{})

	for layout := start; layout != nil; layout = heap.Pop(amphipodHeap).(*Layout) {
		hash := layout.hash()
		if _, ok := seen[hash]; ok {
			continue
		}

		seen[hash] = nil

		if layout.finishedAll() {
			return layout.cost
		}

		for _, next := range layout.nextMoves() {
			heap.Push(amphipodHeap, next)
		}
	}

	return 0
}

func readAmphipods(lines []string) *Layout {
	var buckets = [4]*Bucket{
		{0, make([]rune, 0), 1, 2},
		{0, make([]rune, 0), 2, 3},
		{0, make([]rune, 0), 3, 4},
		{0, make([]rune, 0), 4, 5},
	}

	for i := len(lines) - 1; i > 1; i-- {
		bucket := 0
		for _, pos := range lines[i] {
			if pos == '#' || pos == ' ' {
				continue
			}

			b := buckets[bucket]

			if b.remaining != 0 || int(pos-'A') != bucket {
				b.unsorted = append(b.unsorted, pos)
				b.remaining++
			}
			bucket++
		}
	}

	return &Layout{
		[7]rune{},
		buckets,
		0,
	}
}

func (h LayoutHeap) Len() int {
	return len(h)
}

func (h LayoutHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}

func (h LayoutHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *LayoutHeap) Push(x interface{}) {
	*h = append(*h, x.(*Layout))
}

func (h *LayoutHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
