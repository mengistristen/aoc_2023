package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day6 struct{}

type races struct {
	time     []string
	distance []string
}

func (d Day6) Name() string {
	return "Day 6 - Wait For It"
}

func (d Day6) PartOne(ch chan string) {
	defer close(ch)

	if sum, err := d.ProcessPartOne("./input/day6_example.txt"); err == nil {
		ch <- fmt.Sprintf("Example output: %d\n", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one example: %v\n", err)
	}

	if sum, err := d.ProcessPartOne("./input/day6.txt"); err == nil {
		ch <- fmt.Sprintf("Output: %d", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one: %v", err)
	}
}

func (d Day6) PartTwo(ch chan string) {
	defer close(ch)

	if sum, err := d.ProcessPartTwo("./input/day6_example.txt"); err == nil {
		ch <- fmt.Sprintf("Example output: %d\n", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one example: %v\n", err)
	}

	if sum, err := d.ProcessPartTwo("./input/day6.txt"); err == nil {
		ch <- fmt.Sprintf("Output: %d", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one: %v", err)
	}
}

func (d Day6) ProcessPartOne(name string) (int, error) {
	product := 1

	races, err := processFile(name)
	if err != nil {
		return 0, err
	}

	for race := 0; race < len(races.time); race++ {
		ways := 0
		distance_str := races.distance[race]
		time_str := races.time[race]

		distance, err := strconv.Atoi(distance_str)
		if err != nil {
			return 0, fmt.Errorf("error parsing distance: %v", err)
		}

		time, err := strconv.Atoi(time_str)
		if err != nil {
			return 0, fmt.Errorf("error parsing time: %v", err)
		}

		for held := 1; held < time; held++ {
			if held*(time-held) > distance {
				ways++
			}
		}

		product *= ways
	}

	return product, nil
}

func (d Day6) ProcessPartTwo(name string) (int, error) {
	ways := 0

	races, err := processFile(name)
	if err != nil {
		return 0, nil
	}

	distance_str := strings.Join(races.distance, "")
	time_str := strings.Join(races.time, "")

	distance, err := strconv.Atoi(distance_str)
	if err != nil {
		return 0, fmt.Errorf("error parsing distance: %v", err)
	}

	time, err := strconv.Atoi(time_str)
	if err != nil {
		return 0, fmt.Errorf("error parsing time: %v", err)
	}

	for held := 1; held < time; held++ {
		if held*(time-held) > distance {
			ways++
		}
	}

	return ways, nil
}

func processFile(name string) (*races, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	races, err := parseRaces(scanner)
	if err != nil {
		return nil, err
	}

	return races, nil
}

func parseRaces(scanner *bufio.Scanner) (*races, error) {
	var races races

	scanner.Scan()

	races.time = strings.Fields(strings.Split(scanner.Text(), ":")[1])

	scanner.Scan()

	races.distance = strings.Fields(strings.Split(scanner.Text(), ":")[1])

	return &races, nil
}
