package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Day7 struct{}
type game struct {
	hand string
	bid  int
}

var cardValues = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

func (d Day7) Name() string {
	return "Day 7 - Camel Cards"
}

func (d Day7) PartOne(ch chan string) {
	defer close(ch)

	if sum, err := d.ProcessPartOne("./input/day7_example.txt"); err == nil {
		ch <- fmt.Sprintf("Example output: %d\n", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one example: %v\n", err)
	}

	if sum, err := d.ProcessPartOne("./input/day7.txt"); err == nil {
		ch <- fmt.Sprintf("Output: %d", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one: %v", err)
	}
}

func (d Day7) PartTwo(ch chan string) {
	defer close(ch)

	if sum, err := d.ProcessPartTwo("./input/day7_example.txt"); err == nil {
		ch <- fmt.Sprintf("Example output: %d\n", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one example: %v\n", err)
	}

	if sum, err := d.ProcessPartTwo("./input/day7.txt"); err == nil {
		ch <- fmt.Sprintf("Output: %d", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one: %v", err)
	}
}

func (d Day7) ProcessPartOne(name string) (int, error) {
	winnings := 0

	games, err := processFile(name)
	if err != nil {
		return 0, err
	}

	slices.SortFunc(games, func(a, b game) int {
		scoreA := scoreCard(a.hand, false)
		scoreB := scoreCard(b.hand, false)

		for i := 0; scoreA == scoreB && i < 5; i++ {
			scoreA = cardValues[rune(a.hand[i])]
			scoreB = cardValues[rune(b.hand[i])]
		}

		if scoreA > scoreB {
			return 1
		} else {
			return -1
		}
	})

	for rank := 1; rank <= len(games); rank++ {
		winnings += rank * games[rank-1].bid
	}

	return winnings, nil
}

func (d Day7) ProcessPartTwo(name string) (int, error) {
	winnings := 0

	cardValues['J'] = 0

	games, err := processFile(name)
	if err != nil {
		return 0, err
	}

	slices.SortFunc(games, func(a, b game) int {
		scoreA := scoreCard(a.hand, true)
		scoreB := scoreCard(b.hand, true)

		for i := 0; scoreA == scoreB && i < 5; i++ {
			scoreA = cardValues[rune(a.hand[i])]
			scoreB = cardValues[rune(b.hand[i])]
		}

		if scoreA > scoreB {
			return 1
		} else {
			return -1
		}
	})

	for _, game := range games {
		log.Printf("%s -> %d", game.hand, scoreCard(game.hand, true))
	}

	for rank := 1; rank <= len(games); rank++ {
		winnings += rank * games[rank-1].bid
	}

	return winnings, nil
}

func processFile(name string) ([]game, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	games, err := parseGames(scanner)
	if err != nil {
		return nil, err
	}

	return games, nil
}

func parseGames(scanner *bufio.Scanner) ([]game, error) {
	var games []game

	for scanner.Scan() {
		split := strings.Fields(scanner.Text())

		bid, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing bid: %v", err)
		}

		games = append(games, game{
			hand: split[0],
			bid:  bid,
		})
	}

	return games, nil
}

func scoreCard(hand string, jokers bool) int {
	counts := countCards(hand)
	numJokers := 0
	types := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
	}

	if jokers {
		var highestKey rune
		numJokers = counts['J']
		highest := 0

		for key := range counts {
			if key != 'J' && counts[key] > highest {
				highest = counts[key]
				highestKey = key
			}
		}

		if highest != 0 {
			counts[highestKey] += numJokers
			counts['J'] = 0
		}
	}

	for key := range counts {
		count := counts[key]

		types[count] = types[count] + 1
	}

	if types[5] == 1 {
		return 7
	} else if types[4] == 1 {
		return 6
	} else if types[3] == 1 && types[2] == 1 {
		return 5
	} else if types[3] == 1 {
		return 4
	} else if types[2] == 2 {
		return 3
	} else if types[2] == 1 {
		return 2
	}
	return 1
}

func countCards(hand string) map[rune]int {
	cardCounts := make(map[rune]int)

	for _, card := range hand {
		if value, exists := cardCounts[card]; !exists {
			cardCounts[card] = 1
		} else {
			cardCounts[card] = value + 1
		}
	}

	return cardCounts
}
