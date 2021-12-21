package day21

import (
	"advent-of-code-2021/input"
	"fmt"
	"strconv"
	"strings"
)

type Player struct {
	id       int
	position int
	score    int
}

func (p *Player) createCount() Count {
	if p.id == 0 {
		return Count{1, 0}
	} else {
		return Count{0, 1}
	}
}

type DeterministicDice struct {
	next int
}

type DiracDice struct {
}

func (dice DiracDice) roll() []int {
	return []int{3, 4, 4, 4, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 8, 8, 8, 9}
}

func (dice *DeterministicDice) roll() []int {
	return []int{dice.rollNext() + dice.rollNext() + dice.rollNext()}
}

func (dice *DeterministicDice) rollNext() int {
	next := dice.next
	dice.next += 1
	if dice.next > 100 {
		dice.next = 1
	}
	return next
}

type Dice interface {
	roll() []int
}

type Game struct {
	players [2]Player
	dice    Dice
}

func (game Game) runRound() []Game {
	games := make([]Game, 0)

	toMove := game.players[0]
	other := game.players[1]

	for _, roll := range game.dice.roll() {
		newPos := (toMove.position + roll) % 10
		games = append(
			games,
			Game{
				[2]Player{other, {toMove.id, newPos, toMove.score + newPos + 1}},
				game.dice,
			},
		)
	}

	return games
}

func (game *Game) won(score int) (*Player, int) {
	possibleWon := game.players[1]
	other := game.players[0]

	if possibleWon.score >= score {
		return &possibleWon, other.score
	}

	return nil, 0
}

func Day21() {
	lines, err := input.ReadLinesString("inputs/day21.txt")

	if err != nil {
		fmt.Println("Could not find file")
	} else {
		player1 := readPlayer(lines[0], 0)
		player2 := readPlayer(lines[1], 1)

		fmt.Println(part1(player1, player2))
		fmt.Println(part2(player1, player2))
	}
}

func part1(player1 Player, player2 Player) int {
	dice := &DeterministicDice{1}
	game := Game{[2]Player{player1, player2}, dice}
	round := 1

	for {
		game = game.runRound()[0]
		won, score := game.won(1000)
		if won != nil {
			return score * round * 3
		}
		round++
	}
}

func part2(player1 Player, player2 Player) int {
	result := runRecursive(
		Game{
			[2]Player{
				player1,
				player2,
			},
			DiracDice{},
		},
		make(map[Game]Count, 0),
	)
	return intMax(result.first, result.second)
}

func intMax(first int, second int) int {
	if first > second {
		return first
	}
	return second
}

type Count struct {
	first  int
	second int
}

func runRecursive(game Game, cache map[Game]Count) Count {
	if won, _ := game.won(21); won != nil {
		count := won.createCount()
		cache[game] = count
		return count
	}

	counts := Count{0, 0}
	for _, newGame := range game.runRound() {
		wonCounts, ok := cache[newGame]
		if !ok {
			wonCounts = runRecursive(newGame, cache)
		}
		counts.first += wonCounts.first
		counts.second += wonCounts.second
	}

	cache[game] = counts
	return counts
}

func readPlayer(line string, id int) Player {
	segments := strings.Split(line, ": ")
	position, _ := strconv.ParseInt(segments[1], 10, 64)

	return Player{id, int(position) - 1, 0}
}
