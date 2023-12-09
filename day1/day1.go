package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Day1 struct{}

var prefixes = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func (d Day1) Name() string {
	return "Day 1 - Trebuchet"
}

func (d Day1) PartOne(ch chan string) {
	defer close(ch)

	file, err := os.Open("./input/day1_example1.txt")
	if err != nil {
		ch <- fmt.Sprintf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := processPartOne(scanner, ch)

	ch <- fmt.Sprintf("Example output: %d\n", sum)

	file, err = os.Open("./input/day1.txt")
	if err != nil {
		ch <- fmt.Sprintf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)

	sum = processPartOne(scanner, ch)

	ch <- fmt.Sprintf("Output: %d", sum)
}

func (d Day1) PartTwo(ch chan string) {
	defer close(ch)

	file, err := os.Open("./input/day1_example2.txt")
	if err != nil {
		ch <- fmt.Sprintf("Error opening file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := processPartTwo(scanner, ch)

	ch <- fmt.Sprintf("Example output: %d\n", sum)

	file, err = os.Open("./input/day1.txt")
	if err != nil {
		ch <- fmt.Sprintf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)

	sum = processPartTwo(scanner, ch)

	ch <- fmt.Sprintf("Output: %d", sum)
}

func (d Day1) ProcessPartOne(name string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}

func (d Day1) ProcessPartTwo(name string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}

func processPartOne(scanner *bufio.Scanner, ch chan string) int {
	sum := 0

	for scanner.Scan() {
		current := -1

		for _, char := range scanner.Text() {
			if unicode.IsDigit(char) {
				value, err := strconv.Atoi(string(char))
				if err != nil {
					ch <- fmt.Sprintf("Error parsing rune: %v\n", err)
				}

				if current == -1 {
					sum += value * 10
				}

				current = value
			}
		}

		sum += current
	}

	return sum
}

func processPartTwo(scanner *bufio.Scanner, ch chan string) int {
	sum := 0

	for scanner.Scan() {
		current := -1
		line := scanner.Text()

		for index := 0; index < len(line); index++ {
			found := -1

			sub := line[index:]

			if valid, amount := getPrefixValue(sub); valid {
				found = amount
			} else if char := rune(line[index]); unicode.IsDigit(char) {
				value, err := strconv.Atoi(string(char))
				if err != nil {
					ch <- fmt.Sprintf("Error parsing rune: %v\n", err)
				}
				found = value
			}

			if found != -1 {
				if current == -1 {
					sum += found * 10
				}

				current = found
			}
		}

		sum += current
	}

	return sum
}

func getPrefixValue(str string) (bool, int) {
	for key := range prefixes {
		if strings.HasPrefix(str, key) {
			return true, prefixes[key]
		}
	}

	return false, 0
}
