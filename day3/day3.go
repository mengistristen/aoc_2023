package day3

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "unicode"
)

type Day3 struct{}

type part struct{
    value int
    area [2]int
}

type symbol struct{
    c string
    x int 
    y int
}

func (d Day3) Name() string {
    return "Day 3 - Gear Ratios"
}

func (d Day3) PartOne(ch chan string) {
    defer close(ch)

    if sum, err := processPartOne("./input/day3_example.txt"); err == nil {
        ch <- fmt.Sprintf("Example output: %d\n", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one example: %v\n", err)
    }

    if sum, err := processPartOne("./input/day3.txt"); err == nil {
        ch <- fmt.Sprintf("Output: %d", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one: %v", err)
    }
}

func (d Day3) PartTwo(ch chan string) {
    defer close(ch)
    
    if sum, err := processPartTwo("./input/day3_example.txt"); err == nil {
        ch <- fmt.Sprintf("Example output: %d\n", sum)
    } else {
        ch <- fmt.Sprintf("error processing part two example: %v\n", err)
    }

    if sum, err := processPartTwo("./input/day3.txt"); err == nil {
        ch <- fmt.Sprintf("Example output: %d", sum)
    } else {
        ch <- fmt.Sprintf("error processing part two: %v", err)
    }
}

func processPartOne(name string) (int, error) {
    sum := 0

    schematic, symbols, err := processFile(name)
    if err != nil {
        return 0, err
    }

    for _, symbol := range symbols {
        if symbol.y != 0 {
            sum += processRow(&schematic[symbol.y - 1], symbol.x)
        }

        sum += processRow(&schematic[symbol.y], symbol.x)

        if symbol.y != len(schematic) - 1 {
            sum += processRow(&schematic[symbol.y + 1], symbol.x)
        }
    }

    return sum, nil
}

func processPartTwo(name string) (int, error) {
    sum := 0

    schematic, symbols, err := processFile(name)
    if err != nil {
        return 0, err
    }

    for _, symbol := range symbols {
        if symbol.c == "*" {
            var adjacent []int 

            if symbol.y != 0 {
                adjacent = append(adjacent, processGear(schematic[symbol.y - 1], symbol.x)...)
            }

            adjacent = append(adjacent, processGear(schematic[symbol.y], symbol.x)...)

            if symbol.y != len(schematic) - 1 {
                adjacent = append(adjacent, processGear(schematic[symbol.y + 1], symbol.x)...)
            }

            if len(adjacent) == 2 {
                sum += adjacent[0] * adjacent[1]
            }
        }
    }

    return sum, nil
}

func processFile(name string) ([][]part, []symbol, error) {
    file, err := os.Open(name)
    if err != nil {
        return nil, nil, fmt.Errorf("error opening file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    schematic, symbols, err := parseSchematic(scanner)
    if err != nil {
        return nil, nil, err
    }

    return schematic, symbols, nil
} 

func parseSchematic(scanner *bufio.Scanner) ([][]part, []symbol, error) {
    var schematic [][]part
    var symbols []symbol

    for scanner.Scan() {
        var parts []part
        index := 0
        current := ""

        for _, char := range scanner.Text() {
            if unicode.IsDigit(char) {
                current += string(char)
            } else if current != "" {
                value, err := strconv.Atoi(current)
                if err != nil {
                    return nil, nil, fmt.Errorf("failed to parse part number: %v", err)
                }

                parts = append(parts, part{
                    value: value,
                    area: [2]int{index - len(current), index - 1}, 
                }) 

                current = ""
            } 

            if !unicode.IsDigit(char) && char != '.' {
                symbols = append(symbols, symbol{
                    c: string(char),
                    x: index,
                    y: len(schematic),
                })
            }

            index++
        }

        if current != "" {
            value, err := strconv.Atoi(current)
            if err != nil {
                return nil, nil, fmt.Errorf("failed to parse part number: %v", err)
            }

            parts = append(parts, part{
                value: value,
                area: [2]int{index - len(current), index - 1}, 
            }) 
        }

        schematic = append(schematic, parts)
    }

    return schematic, symbols, nil
}

func processRow(row *[]part, x int) int {
    sum := 0

    for _, part := range *row {
        if x >= (part.area[0] - 1) && x <= (part.area[1] + 1) {
            sum += part.value
        }
    }

    return sum
}

func processGear(row []part, x int) []int {
    var adjacent []int

    for _, part := range row {
        if x >= (part.area[0] - 1) && x <= (part.area[1] + 1) {
            adjacent = append(adjacent, part.value)
        }
    }
    
    return adjacent
}
