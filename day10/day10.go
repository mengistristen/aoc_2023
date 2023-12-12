package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Day10 struct{}
type current struct {
	index int
	from  int8
}

const (
	north = int8(1 << 0)
	east  = int8(1 << 1)
	south = int8(1 << 2)
	west  = int8(1 << 3)
)

var pipes = map[rune]int8{
	'|': north | south,
	'-': west | east,
	'L': north | east,
	'J': north | west,
	'7': south | west,
	'F': south | east,
}

func (d Day10) Name() string {
	return "Day 10 - Pipe Maze"
}

func (d Day10) PartOne(ch chan string) {
	defer close(ch)

	if sum, err := d.ProcessPartOne("./input/day10_example1.txt"); err == nil {
		ch <- fmt.Sprintf("Example 1 output: %d\n", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one example: %v\n", err)
	}

	if sum, err := d.ProcessPartOne("./input/day10.txt"); err == nil {
		ch <- fmt.Sprintf("Output: %d", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one: %v", err)
	}
}

func (d Day10) PartTwo(ch chan string) {
	defer close(ch)

	if sum, err := d.ProcessPartTwo("./input/day10_example2.txt"); err == nil {
		ch <- fmt.Sprintf("Example output: %d\n", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one example: %v\n", err)
	}

	/*
		if sum, err := d.ProcessPartTwo("./input/day9.txt"); err == nil {
			ch <- fmt.Sprintf("Output: %d", sum)
		} else {
			ch <- fmt.Sprintf("error processing part one: %v", err)
		}
	*/
}

func (d Day10) ProcessPartOne(name string) (int, error) {
	longest := 0

	pipeline, start, cols, err := processFile(name)
	if err != nil {
		return 0, err
	}

	processPipeline(pipeline, start, cols,
		func(_ int) {
			longest++
		}, func(_ int) {})

	return longest, nil
}

func (d Day10) ProcessPartTwo(name string) (int, error) {
	contained := 0

	pipeline, start, cols, err := processFile(name)
	if err != nil {
		return 0, err
	}

	typeMap := make([]int, len(pipeline))
	gradients := make([]int8, len(pipeline))

	typeMap[start] = 1

	processPipeline(pipeline, start, cols,
		func(i int) {
			typeMap[i] = 1
		}, func(i int) {
			typeMap[i] = 1
		})

	processMap(typeMap, pipeline, gradients, cols)

	for row := 0; row < len(pipeline)/cols; row++ {
		str := ""

		for col := 0; col < cols; col++ {
			index := coordToIndex(col, row, cols)

			str += fmt.Sprintf(" %2s:%4b ", strconv.Itoa(typeMap[index]), gradients[index])
		}

		log.Print(str)
	}

	for row := 0; row < len(pipeline)/cols; row++ {
		for col := 0; col < cols; col++ {
			index := coordToIndex(col, row, cols)
			if typeMap[index] == 0 {
				if gradients[coordToIndex(col-1, row, cols)]&east != 0 {
					contained++
					gradients[index] |= east
				}
			}
		}
	}

	return contained, nil
}

func processFile(name string) ([]int8, int, int, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pipeline, start, cols, err := parsePipeline(scanner)
	if err != nil {
		return nil, 0, 0, err
	}

	return pipeline, start, cols, nil
}

func parsePipeline(scanner *bufio.Scanner) ([]int8, int, int, error) {
	var pipeline []int8
	var start int
	cols := 0
	rows := 0

	for scanner.Scan() {
		if cols == 0 {
			cols = len(scanner.Text())
		}

		for i, char := range scanner.Text() {
			if char == '.' {
				pipeline = append(pipeline, 0)
			} else if char == 'S' {
				pipeline = append(pipeline, 0)
				start = coordToIndex(i, rows, cols)
			} else {
				pipeline = append(pipeline, pipes[char])
			}
		}

		rows += 1
	}

	return pipeline, start, cols, nil
}

func coordToIndex(col, row, cols int) int {
	return cols*row + col
}

func indexToCoord(index, cols int) (int, int) {
	return index % cols, index / cols
}

func processPipeline(pipeline []int8, start, cols int, f0 func(int), f1 func(int)) {
	startX, startY := indexToCoord(start, cols)

	if startY-1 >= 0 && (pipeline[coordToIndex(startX, startY-1, cols)]&south) != 0 {
		pipeline[start] |= north
	}

	if startX+1 < cols && (pipeline[coordToIndex(startX+1, startY, cols)]&west) != 0 {
		pipeline[start] |= east
	}

	if startY+1 < len(pipeline)/cols && (pipeline[coordToIndex(startX, startY+1, cols)]&north) != 0 {
		pipeline[start] |= south
	}

	if startX-1 >= 0 && (pipeline[coordToIndex(startX-1, startY, cols)]&east) != 0 {
		pipeline[start] |= west
	}

	var heads []current

	for i := 0; i < 4; i++ {
		if pipeline[start]&(1<<i) != 0 {
			heads = append(heads, moveNext(pipeline, start, cols, pipeline[start]^(1<<i)))
		}
	}

	for heads[0].index != heads[1].index {
		f0(heads[0].index)
		f1(heads[1].index)

		heads[0] = moveNext(pipeline, heads[0].index, cols, heads[0].from)
		heads[1] = moveNext(pipeline, heads[1].index, cols, heads[1].from)
	}

	f0(heads[0].index)
}

func processMap(typeMap []int, pipeline, gradients []int8, cols int) {
	rows := len(typeMap) / cols
	// process each row left to right, right to left
	for row := 0; row < rows; row++ {
		dir := east

		for col := 0; col < cols; col++ {
			index := coordToIndex(col, row, cols)
			prev := coordToIndex(col-1, row, cols)

			processPoint(typeMap, pipeline, gradients, col, index, prev, &dir)
		}

		dir = west

		for col := cols - 1; col >= 0; col-- {
			index := coordToIndex(col, row, cols)
			prev := coordToIndex(col+1, row, cols)

			processPoint(typeMap, pipeline, gradients, cols-(col+1), index, prev, &dir)
		}
	}

	// process each column top to bottom, bottom to top
	for col := 0; col < cols; col++ {
		dir := south

		for row := 0; row < rows; row++ {
			index := coordToIndex(col, row, cols)
			prev := coordToIndex(col, row-1, cols)

			processPoint(typeMap, pipeline, gradients, row, index, prev, &dir)
		}

		dir = north

		for row := rows - 1; row >= 0; row-- {
			index := coordToIndex(col, row, cols)
			prev := coordToIndex(col, row+1, cols)

			processPoint(typeMap, pipeline, gradients, rows-(row+1), index, prev, &dir)
		}
	}
}

func processPoint(typeMap []int, pipeline, gradients []int8, seq, index, prev int, dir *int8) {
	if seq == 0 {
		if typeMap[index] == 0 {
			typeMap[index] = -1
		} else if typeMap[index] == 1 {
			gradients[index] |= *dir
		}
	} else {
		if typeMap[prev] == -1 && typeMap[index] == 0 {
			typeMap[index] = -1
		} else if typeMap[prev] == -1 && typeMap[index] == 1 {
			gradients[index] |= *dir
		} else if typeMap[index] == 1 {
			if (typeMap[prev] == 0) || ((pipeline[index] & *dir) == 0) {
				*dir = opposite(*dir)
			}
			gradients[index] |= *dir
		}
	}
}

func moveNext(pipeline []int8, index, cols int, from int8) current {
	var dest int
	direction := pipeline[index] ^ from

	x, y := indexToCoord(index, cols)

	switch direction {
	case north:
		dest = coordToIndex(x, y-1, cols)
	case east:
		dest = coordToIndex(x+1, y, cols)
	case south:
		dest = coordToIndex(x, y+1, cols)
	case west:
		dest = coordToIndex(x-1, y, cols)
	}

	return current{
		index: dest,
		from:  opposite(direction),
	}
}

func opposite(direction int8) int8 {
	var result int8

	switch direction {
	case north:
		result = south
	case east:
		result = west
	case south:
		result = north
	case west:
		result = east
	}

	return result
}
