package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day12 struct{}
type row struct {
	springs string
	counts  []int
}

func (d Day12) Name() string {
	return "Day 12 - Hot Springs"
}

func (d Day12) PartOne(ch chan string) {
	defer close(ch)
	/*
		if sum, err := d.ProcessPartOne("./input/day12_example.txt"); err == nil {
			ch <- fmt.Sprintf("Example output: %d\n", sum)
		} else {
			ch <- fmt.Sprintf("error processing part one example: %v\n", err)
		}

		if sum, err := d.ProcessPartOne("./input/day12.txt"); err == nil {
			ch <- fmt.Sprintf("Output: %d", sum)
		} else {
			ch <- fmt.Sprintf("error processing part one: %v", err)
		}
	*/
}

func (d Day12) PartTwo(ch chan string) {
	defer close(ch)

	if sum, err := d.ProcessPartTwo("./input/day12_example.txt"); err == nil {
		ch <- fmt.Sprintf("Example output: %d\n", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one example: %v\n", err)
	}
}

func (d Day12) ProcessPartOne(name string) (int, error) {
	combinations := 0

	hotSprings, err := processFile(name)
	if err != nil {
		return 0, err
	}

	for _, row := range hotSprings {
		processRow(0, row.springs, row.counts, "", &combinations)
	}

	return combinations, nil
}

func (d Day12) ProcessPartTwo(name string) (int, error) {
	combinations := 0

	/*
		hotSprings, err := processFile(name)
		if err != nil {
			return 0, err
		}

		for i, row := range hotSprings {
			comboOne := 0
			comboTwo := 0

			processRow(0, "."+row.springs, row.counts, "", &comboOne)
			processRow(0, "#"+row.springs, row.counts, "", &comboTwo)

			log.Printf("%d - %d,%d", i+1, comboOne, comboTwo)
		}
	*/

	row := row{
		springs: "????.#...#...?????.#...#...?????.#...#...?????.#...#...?????.#...#...",
		counts:  []int{4, 1, 1, 4, 1, 1, 4, 1, 1, 4, 1, 1, 4, 1, 1},
	}
	processRow(0, row.springs, row.counts, "", &combinations)

	return combinations, nil
}

func processFile(name string) ([]row, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	space, err := parseHotSprings(scanner)
	if err != nil {
		return nil, err
	}

	return space, nil
}

func parseHotSprings(scanner *bufio.Scanner) ([]row, error) {
	var rows []row

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")

		row := row{
			springs: split[0],
		}

		for _, count := range strings.Split(split[1], ",") {
			count, err := strconv.Atoi(count)
			if err != nil {
				return nil, fmt.Errorf("error parsing count: %v", err)
			}

			row.counts = append(row.counts, count)
		}

		rows = append(rows, row)
	}

	return rows, nil
}

func processRow(curr int, springs string, counts []int, total string, combinations *int) {
	if springs == "" {
		if (curr == 0 && len(counts) == 0) || (len(counts) == 1 && curr == counts[0]) {
			log.Print(total)
			(*combinations)++
		}
		return
	}

	switch springs[0] {
	case '.':
		if curr == 0 {
			processRow(0, springs[1:], counts, total+".", combinations)
		} else if curr == counts[0] {
			processRow(0, springs[1:], counts[1:], total+".", combinations)
		}
	case '#':
		if len(counts) == 0 {
			return
		} else if curr < counts[0] {
			processRow(curr+1, springs[1:], counts, total+"#", combinations)
		}
	case '?':
		processRow(curr, "#"+springs[1:], counts, total, combinations)
		processRow(curr, "."+springs[1:], counts, total, combinations)
	}
}
