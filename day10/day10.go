package day10

import (
	"bufio"
	"fmt"
	"os"
)

type Day10 struct {}
type current struct {
    index int
    from int8
}

const (
    north = (1 << 0)
    east = (1 << 1) 
    south = (1 << 2)
    west = (1 << 3)
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
}

func (d Day10) PartTwo(ch chan string) {
    defer close(ch)
}

func (d Day10) ProcessPartOne(name string) (int, error) {
    longest := 1 

    pipeline, start, cols, err := processFile(name)
    if err != nil {
        return 0, err
    }

    startX, startY := indexToCoord(start, cols)

    if startY - 1 >= 0 && (pipeline[coordToIndex(startX, startY - 1, cols)] & south) != 0 {
        pipeline[start] |= north
    }

    if startX + 1 < cols && (pipeline[coordToIndex(startX + 1, startY, cols)] & west) != 0 {
        pipeline[start] |= east
    }
    
    if startY + 1 < len(pipeline) / cols && (pipeline[coordToIndex(startX, startY + 1, cols)] & north) != 0 {
        pipeline[start] |= south
    }

    if startX - 1 >= 0 && (pipeline[coordToIndex(startX - 1, startY, cols)] & east) != 0 {
        pipeline[start] |= west
    }

    var heads []current

    for i := 0; i < 4; i++ {
        if pipeline[start] & (1 << i) != 0 {
            heads = append(heads, moveNext(pipeline, start, cols, pipeline[start] ^ (1 << i)))                
        }
    }

    for heads[0].index != heads[1].index {
        heads[0] = moveNext(pipeline, heads[0].index, cols, heads[0].from)
        heads[1] = moveNext(pipeline, heads[1].index, cols, heads[1].from)

        longest++
    }

    return longest, nil
}

func (d Day10) ProcessPartTwo(name string) (int, error) {
    return 0, nil
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
    return cols * row + col
}

func indexToCoord(index, cols int) (int, int) {
    return index % cols, index / cols 
}

func moveNext(pipeline []int8, index, cols int, from int8) current {
    var dest int
    direction := pipeline[index] ^ from

    x, y := indexToCoord(index, cols)

    switch direction {
    case north:
        dest = coordToIndex(x, y - 1, cols) 
    case east:
        dest = coordToIndex(x + 1, y, cols)
    case south:
        dest = coordToIndex(x, y + 1, cols)
    case west:
        dest = coordToIndex(x - 1, y, cols)
    }

    return current{
        index: dest,
        from: opposite(direction),
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
