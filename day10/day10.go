package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		ch <- fmt.Sprintf("Example output: %d\n", sum)
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

	if sum, err := d.ProcessPartTwo("./input/day10.txt"); err == nil {
		ch <- fmt.Sprintf("Output: %d", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one: %v", err)
	}
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

	// find the first segment of the pipeline from the left
	for i, t := range typeMap {
		if t == 1 {
			start = i
			break
		}
	}

	current := current{
		index: start,
		from:  south,
	}

	vector := east

	for processing := true; processing; processing = (current.index != start) {
		switch pipeline[current.index] {
		case pipes['L']:
			gradients[current.index] |= vector
			if current.from == north {
				vector = rotateCounterClockwise(vector)
			} else {
				vector = rotateClockwise(vector)
			}
		case pipes['J']:
			gradients[current.index] |= vector
			if current.from == north {
				vector = rotateClockwise(vector)
			} else {
				vector = rotateCounterClockwise(vector)
			}
		case pipes['7']:
			gradients[current.index] |= vector
			if current.from == south {
				vector = rotateCounterClockwise(vector)
			} else {
				vector = rotateClockwise(vector)
			}
		case pipes['F']:
			gradients[current.index] |= vector
			if current.from == south {
				vector = rotateClockwise(vector)
			} else {
				vector = rotateCounterClockwise(vector)
			}
		}

		gradients[current.index] |= vector
		current = moveNext(pipeline, current.index, cols, current.from)
	}

	rows := len(pipeline) / cols
	for row, col := 0, 0; row < rows/2 || col < cols/2; {
		// Check top left corner
		if index := coordToIndex(col, row, cols); typeMap[index] == 0 {
			if col == 0 && row == 0 {
				typeMap[index] = -1
			} else if ((gradients[coordToIndex(col, row-1, cols)] & south) != 0) && ((gradients[coordToIndex(col-1, row, cols)] & east) != 0) {
				typeMap[index] = 2
				gradients[index] = 0b1111
				contained++
			} else {
				typeMap[index] = -1
			}
		}

		// Check top right corner
		if index := coordToIndex(cols-col-1, row, cols); typeMap[index] == 0 {
			if col == 0 && row == 0 {
				typeMap[index] = -1
			} else if ((gradients[coordToIndex(cols-col-1, row-1, cols)] & south) != 0) && ((gradients[coordToIndex(cols-col, row, cols)] & west) != 0) {
				typeMap[index] = 2
				gradients[index] = 0b1111
				contained++
			} else {
				typeMap[index] = -1
			}
		}

		// Check bottom left corner
		if index := coordToIndex(col, rows-row-1, cols); typeMap[index] == 0 {
			if col == 0 && row == 0 {
				typeMap[index] = -1
			} else if ((gradients[coordToIndex(col, rows-row, cols)] & north) != 0) && ((gradients[coordToIndex(col-1, rows-row-1, cols)] & east) != 0) {
				typeMap[index] = 2
				gradients[index] = 0b1111
				contained++
			} else {
				typeMap[index] = -1
			}
		}

		// Check bottom right corner
		if index := coordToIndex(cols-col-1, rows-row-1, cols); typeMap[index] == 0 {
			if col == 0 && row == 0 {
				typeMap[index] = -1
			} else if ((gradients[coordToIndex(cols-col-1, rows-row, cols)] & north) != 0) && ((gradients[coordToIndex(cols-col, rows-row-1, cols)] & west) != 0) {
				typeMap[index] = 2
				gradients[index] = 0b1111
				contained++
			} else {
				typeMap[index] = -1
			}
		}

		// Check the top row
		for x := col + 1; x < cols-col-1; x++ {
			if index := coordToIndex(x, row, cols); typeMap[index] == 0 {
				if row == 0 {
					typeMap[index] = -1
				} else if (gradients[coordToIndex(x, row-1, cols)] & south) != 0 {
					typeMap[index] = 2
					gradients[index] = 0b1111
					contained++
				} else {
					typeMap[index] = -1
				}
			}
		}

		// Check the sides
		for y := row + 1; y < rows-row-1; y++ {
			// Check the left side
			if index := coordToIndex(col, y, cols); typeMap[index] == 0 {
				if col == 0 {
					typeMap[index] = -1
				} else if (gradients[coordToIndex(col-1, y, cols)] & east) != 0 {
					typeMap[index] = 2
					gradients[index] = 0b1111
					contained++
				} else {
					typeMap[index] = -1
				}
			}

			// Check the right side
			if index := coordToIndex(cols-col-1, y, cols); typeMap[index] == 0 {
				if col == 0 {
					typeMap[index] = -1
				} else if (gradients[coordToIndex(cols-col, y, cols)] & west) != 0 {
					typeMap[index] = 2
					gradients[index] = 0b1111
					contained++
				} else {
					typeMap[index] = -1
				}
			}
		}

		// Check the bottom row
		for x := col + 1; x < cols-col-1; x++ {
			if index := coordToIndex(x, rows-row-1, cols); typeMap[index] == 0 {
				if row == 0 {
					typeMap[index] = -1
				} else if (gradients[coordToIndex(x, rows-row, cols)] & north) != 0 {
					typeMap[index] = 2
					gradients[index] = 0b1111
					contained++
				} else {
					typeMap[index] = -1
				}
			}
		}

		if row < rows/2 {
			row++
		}

		if col < cols/2 {
			col++
		}
	}

	for row := 0; row < len(pipeline)/cols; row++ {
		str := ""

		for col := 0; col < cols; col++ {
			index := coordToIndex(col, row, cols)
			str += fmt.Sprintf(" %2d ", typeMap[index])
		}

		log.Print(str)
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

func rotateClockwise(dir int8) int8 {
	if dir == west {
		return north
	}
	return dir << 1
}

func rotateCounterClockwise(dir int8) int8 {
	if dir == north {
		return west
	}
	return dir >> 1
}
