package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Day11 struct {}
type galaxy struct{
    x int
    y int
}

func (d Day11) Name() string {
    return "Day 11 - Cosmic Expansion"
}

func (d Day11) PartOne(ch chan string) {
    defer close(ch)

	if sum, err := d.ProcessPartOne("./input/day11_example.txt"); err == nil {
		ch <- fmt.Sprintf("Example output: %d\n", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one example: %v\n", err)
	}

	if sum, err := d.ProcessPartOne("./input/day11.txt"); err == nil {
		ch <- fmt.Sprintf("Output: %d", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one: %v", err)
	}
}

func (d Day11) PartTwo(ch chan string) {
    defer close(ch)

	if sum, err := d.ProcessPartTwo("./input/day11_example.txt"); err == nil {
		ch <- fmt.Sprintf("Example output: %d\n", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one example: %v\n", err)
	}

	if sum, err := d.ProcessPartTwo("./input/day11.txt"); err == nil {
		ch <- fmt.Sprintf("Output: %d", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one: %v", err)
	}
}

func (d Day11) ProcessPartOne(name string) (int, error) {
    space, cols, rows, err := processFile(name, 2)
    if err != nil {
        return 0, err
    }

    return processPaths(space, cols, rows), nil
}

func (d Day11) ProcessPartTwo(name string) (int, error) {
    space, cols, rows, err := processFile(name, 1000000)
    if err != nil {
        return 0, err
    }

    return processPaths(space, cols, rows), nil
}

func processFile(name string, expansionFactor int) ([]galaxy, []int, []int, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	space, cols, rows, err := parseSpace(scanner, expansionFactor)
	if err != nil {
		return nil, nil, nil, err
	}

	return space, cols, rows, nil
}

func parseSpace(scanner *bufio.Scanner, expansionFactor int) ([]galaxy, []int, []int, error) {
    var space []galaxy
    var cols, rows, counts []int
    current := 0

    for row := 0; scanner.Scan(); row++ {
        found := false

        if counts == nil {
            counts = make([]int, len(scanner.Text()))
        }

        for col, char := range scanner.Text() {
            if char == '#' {
                counts[col]++
                found = true
                space = append(space, galaxy{
                    x: col,
                    y: row,
                })
            }
        }

        if !found {
            current+=expansionFactor-1
        }

        rows = append(rows, current)
    }

    current = 0

    for _, count := range counts {
        if count == 0 {
            current+=expansionFactor-1
        }
        cols = append(cols, current)
    }

    return space, cols, rows, nil
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func processPaths(space []galaxy, cols, rows []int) int {
    sum := 0

    for i, galaxy := range space {
        for j := i + 1; j < len(space); j++ {
            other := space[j]
            deltaX := abs(galaxy.x - other.x)
            deltaY := abs(galaxy.y - other.y)

            deltaX += abs(cols[galaxy.x] - cols[other.x])
            deltaY += abs(rows[galaxy.y] - rows[other.y])

            log.Printf("%d -> %d = %d", i, j, deltaX + deltaY)
            sum += deltaX + deltaY
        }
    }

    return sum
}
