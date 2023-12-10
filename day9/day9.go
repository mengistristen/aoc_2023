package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day9 struct {}

func (d Day9) Name() string {
    return "Day 9 - Mirage Maintenance"
}

func (d Day9) PartOne(ch chan string) {
    defer close(ch)

	if sum, err := d.ProcessPartOne("./input/day9_example.txt"); err == nil {
		ch <- fmt.Sprintf("Example 1 output: %d\n", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one example: %v\n", err)
	}

	if sum, err := d.ProcessPartOne("./input/day9.txt"); err == nil {
		ch <- fmt.Sprintf("Output: %d", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one: %v", err)
	}
}

func (d Day9) PartTwo(ch chan string) {
    defer close(ch)

	if sum, err := d.ProcessPartTwo("./input/day9_example.txt"); err == nil {
		ch <- fmt.Sprintf("Example output: %d\n", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one example: %v\n", err)
	}

	if sum, err := d.ProcessPartTwo("./input/day9.txt"); err == nil {
		ch <- fmt.Sprintf("Output: %d", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one: %v", err)
	}
}

func (d Day9) ProcessPartOne(name string) (int, error) {
    sum := 0

    histories, err := processFile(name)
    if err != nil {
        return 0, err
    }

    for _, history := range histories {
        var next []int
        complete := false
        current := history

        for !complete { 
            complete = true

            for i := 0; i < len(current) - 1; i++ {
                difference := current[i + 1] - current[i]

                next = append(next, difference)

                if difference != 0 {
                    complete = false
                }

                if (i + 1) == len(current) - 1 {
                    sum += current[i + 1]
                }
            }

            current = next
            next = nil
        }
    }

    return sum, nil
}

func (d Day9) ProcessPartTwo(name string) (int, error) {
    sum := 0

    histories, err := processFile(name)
    if err != nil {
        return 0, err
    }

    for _, history := range histories {
        var next []int
        var stack []int
        complete := false
        current := history

        for !complete { 
            complete = true

            stack = append(stack, current[0])

            for i := 0; i < len(current) - 1; i++ {
                difference := current[i + 1] - current[i]

                next = append(next, difference)

                if difference != 0 {
                    complete = false
                }
            }

            current = next
            next = nil
        }

        difference := 0
        for i := len(stack) - 1; i >= 0; i-- {
            minuend := stack[i]
            difference = minuend - difference
        }

        sum += difference
    }

    return sum, nil
}

func processFile(name string) ([][]int, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	histories, err := parseHistories(scanner)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

func parseHistories(scanner *bufio.Scanner) ([][]int, error) {
    var histories [][]int

    for scanner.Scan() {
        var history []int
        historyStr := strings.Fields(scanner.Text())

        for _, valueStr := range historyStr {
            value, err := strconv.Atoi(valueStr)
            if err != nil {
                return nil, fmt.Errorf("error parsing value: %v", err)
            }

            history = append(history, value)
        }
        
        histories = append(histories, history)
    }

    return histories, nil
}
