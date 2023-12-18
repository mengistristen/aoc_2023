package day16

import (
	"bufio"
	"fmt"
	"os"
)

type Day16 struct{}

type location struct {
	x         int
	y         int
	direction int8
}

func (l location) to(direction int8) location {
	dest := location{
		x:         l.x,
		y:         l.y,
		direction: direction,
	}

	switch direction {
	case up:
		dest.y--
	case right:
		dest.x++
	case down:
		dest.y++
	case left:
		dest.x--
	}

	return dest
}

const (
	up    = int8(1 << 0)
	right = int8(1 << 1)
	down  = int8(1 << 2)
	left  = int8(1 << 3)
)

func (d Day16) Name() string {
	return "Day 16 - The Floor Will Be Lava"
}

func (d Day16) PartOne(ch chan string) {
	defer close(ch)
}

func (d Day16) PartTwo(ch chan string) {
	defer close(ch)
}

func (d Day16) ProcessPartOne(name string) (int, error) {
	contraption, err := processFile(name)
	if err != nil {
		return 0, err
	}

	return processContraption(contraption, location{
		x:         0,
		y:         0,
		direction: right,
	}), nil
}

func (d Day16) ProcessPartTwo(name string) (int, error) {
	return 0, nil
}

func processFile(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	contraption, err := parseContraption(scanner)
	if err != nil {
		return nil, err
	}

	return contraption, nil
}

func parseContraption(scanner *bufio.Scanner) ([]string, error) {
	var contraption []string

	for scanner.Scan() {
		contraption = append(contraption, scanner.Text())
	}

	return contraption, nil
}

func validCoordinates(x, y, rows, cols int) bool {
	return x >= 0 && x < cols && y >= 0 && y < rows
}

func processContraption(contraption []string, start location) int {
	energized := 0
	rows, cols := len(contraption), len(contraption[0])

	lasers := make([][]int8, rows)

	for row := 0; row < rows; row++ {
		lasers[row] = make([]int8, cols)
	}

	var queue []location

	queue = append(queue, start)

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]

		if validCoordinates(curr.x, curr.y, rows, cols) && (lasers[curr.y][curr.x]&curr.direction) == 0 {
			switch contraption[curr.y][curr.x] {
			case '|':
				switch curr.direction {
				case up:
					queue = append(queue, curr.to(up))
				case down:
					queue = append(queue, curr.to(down))
				default:
					queue = append(queue, curr.to(up))
					queue = append(queue, curr.to(down))
				}
			case '-':
				switch curr.direction {
				case left:
					queue = append(queue, curr.to(left))
				case right:
					queue = append(queue, curr.to(right))
				default:
					queue = append(queue, curr.to(left))
					queue = append(queue, curr.to(right))
				}
			case '/':
				switch curr.direction {
				case up:
					queue = append(queue, curr.to(right))
				case right:
					queue = append(queue, curr.to(up))
				case down:
					queue = append(queue, curr.to(left))
				case left:
					queue = append(queue, curr.to(down))
				}
			case '\\':
				switch curr.direction {
				case up:
					queue = append(queue, curr.to(left))
				case right:
					queue = append(queue, curr.to(down))
				case down:
					queue = append(queue, curr.to(right))
				case left:
					queue = append(queue, curr.to(up))
				}
			default:
				queue = append(queue, curr.to(curr.direction))
			}

			if lasers[curr.y][curr.x] == 0 {
				energized++
			}

			lasers[curr.y][curr.x] |= curr.direction
		}
	}

	return energized
}
