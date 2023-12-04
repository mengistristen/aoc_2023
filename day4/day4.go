package day4

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Day4 struct{}

type scratchcard struct{
    id int
    winners [2]int64
    numbers []int
}

func (d Day4) Name() string {
    return "Day 4 - Scratchcards"
}

func (d Day4) PartOne(ch chan string) {
    defer close(ch)

    if sum, err := processPartOne("./input/day4_example.txt"); err == nil {
        ch <- fmt.Sprintf("Example output: %d\n", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one example: %v\n", err)
    }

    if sum, err := processPartOne("./input/day4.txt"); err == nil {
        ch <- fmt.Sprintf("Output: %d", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one: %v", err)
    }
}

func (d Day4) PartTwo(ch chan string) {
    defer close(ch)
    
    if sum, err := processPartTwo("./input/day4_example.txt"); err == nil {
        ch <- fmt.Sprintf("Example output: %d\n", sum)
    } else {
        ch <- fmt.Sprintf("error processing part two example: %v\n", err)
    }

    if sum, err := processPartTwo("./input/day4.txt"); err == nil {
        ch <- fmt.Sprintf("Example output: %d", sum)
    } else {
        ch <- fmt.Sprintf("error processing part two: %v", err)
    }
}

func processPartOne(name string) (int, error) {
    score := 0

    scratchcards, err := processFile(name)
    if err != nil {
        return 0, err
    }

    for _, card := range scratchcards {
        score += scoreCard(card)
    }

    return score, nil
}

func processPartTwo(name string) (int, error) {
    scoredCards := make(map[int]int)

    scratchcards, err := processFile(name)
    if err != nil {
        return 0, err
    }

    for _, card := range scratchcards {
        scoredCards[card.id] = 1 
    }

    totalCards := 0

    for _, card := range scratchcards {
        count := countCardWins(card) 

        for id := card.id + 1; id <= card.id + count; id++ {
            scoredCards[id] += scoredCards[card.id] 
        }

        totalCards += scoredCards[card.id]
    }

    return totalCards, nil
}

func processFile(name string) ([]scratchcard, error) {
    file, err := os.Open(name)
    if err != nil {
        return nil, fmt.Errorf("error opening file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    cards, err := parseCards(scanner)
    if err != nil {
        return nil, err
    }

    return cards, nil
} 

func parseCards(scanner *bufio.Scanner) ([]scratchcard, error) {
    var scratchcards []scratchcard

    for scanner.Scan() {
        line := strings.Split(scanner.Text(), ":")

        id_str := strings.Fields(line[0])[1]
        id, err := strconv.Atoi(id_str)
        if err != nil {
            return nil, fmt.Errorf("error parsing id: %v", err)
        }

        card := scratchcard{
            id: id,
        }

        all := strings.Split(line[1], "|")

        for _, num := range strings.Fields(all[0]) {
            parsed, err := strconv.Atoi(num) 
            if err != nil {
                return nil, fmt.Errorf("error parsing winning number: %v", err)
            }

            if parsed < 64 {
                card.winners[0] |= (1 << parsed)
            } else {
                card.winners[1] |= (1 << (parsed - 64))
            }
        }

        for _, num := range strings.Fields(all[1]) {
            parsed, err := strconv.Atoi(num)
            if err != nil {
                return nil, fmt.Errorf("error parsing card number: %v", err)
            }

            card.numbers = append(card.numbers, parsed)
        }

        scratchcards = append(scratchcards, card) 
    }

    return scratchcards, nil
}

func scoreCard(card scratchcard) int {
    score := 0

    for _, num := range card.numbers {
        found := false

        if num < 64 {
            found = (card.winners[0] & (1 << num) != 0)
        } else {
            found = (card.winners[1] & (1 << (num - 64)) != 0)
        }

        if found {
            if score == 0 {
                score = 1
            } else {
                score = score * 2
            }
        }
    }

    return score
}

func countCardWins(card scratchcard) int {
    count := 0

    for _, num := range card.numbers {
        found := false

        if num < 64 {
            found = (card.winners[0] & (1 << num) != 0)
        } else {
            found = (card.winners[1] & (1 << (num - 64)) != 0)
        }

        if found {
            count++
        }
    }

    return count
}
