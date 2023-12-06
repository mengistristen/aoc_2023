package day5

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Day5 struct {}

type almanac struct {
    seeds []int
    head *mapping
}

type mapping struct {
    next *mapping
    items []item
}

type item struct {
    dest int
    source int 
    area int
}

func (d Day5) Name() string {
    return "Day 5 - If You Give A Seed A Fertilizer"
}

func (d Day5) PartOne(ch chan string) {
    defer close(ch)

    if sum, err := processPartOne("./input/day5_example.txt"); err == nil {
        ch <- fmt.Sprintf("Example output: %d\n", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one example: %v\n", err)
    }

    if sum, err := processPartOne("./input/day5.txt"); err == nil {
        ch <- fmt.Sprintf("Output: %d", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one: %v", err)
    }
}

func (d Day5) PartTwo(ch chan string) {
    defer close(ch)
}

func processPartOne(name string) (int, error) {
    lowest := math.MaxInt
    
    almanac, err := processFile(name)
    if err != nil {
        return 0, err
    }

    for _, seed := range almanac.seeds {
        location := processSeed(seed, almanac.head)

        if location < lowest {
            lowest = location
        }
    }

    return lowest, nil 
}

func processFile(name string) (*almanac, error) {
    file, err := os.Open(name)
    if err != nil {
        return nil, fmt.Errorf("error opening file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    almanac, err := parseAlmanac(scanner)
    if err != nil {
        return nil, err
    }

    return almanac, nil
}

func parseAlmanac(scanner *bufio.Scanner) (*almanac, error) {
    var almanac almanac
    var tail *mapping 
    var current *mapping

    scanner.Scan()

    seeds := strings.Fields(strings.Split(scanner.Text(), ":")[1]) 

    for _, str := range seeds {
        seed, err := strconv.Atoi(str)
        if err != nil {
            return nil, fmt.Errorf("error parsing seed: %v", err)
        }

        almanac.seeds = append(almanac.seeds, seed)
    }

    // ignore the blank space
    scanner.Scan()

    for scanner.Scan() {
        text := scanner.Text()
        if text == "" {
            if almanac.head == nil {
                almanac.head = current 
            } else {
                tail.next = current
            }
            tail = current
            current = nil
        } else if strings.Contains(text, ":") {
            continue 
        } else {
            if current == nil {
                current = new(mapping)
            }

            values := strings.Split(text, " ")

            dest, err := strconv.Atoi(values[0])
            if err != nil {
                return nil, fmt.Errorf("error parsing source: %v", err)
            }

            source, err := strconv.Atoi(values[1])
            if err != nil {
                return nil, fmt.Errorf("error parsing destination: %v", err)
            }

            area, err := strconv.Atoi(values[2])
            if err != nil {
                return nil, fmt.Errorf("error parsing area: %v", err)
            }

            current.items = append(current.items, item{
                dest: dest, 
                source: source,
                area: area,
            })
        }
    }

    if current != nil {
        tail.next = current
    }

    return &almanac, nil
}

func processSeed(source int, mapping *mapping) int {
    if mapping == nil {
        return source
    }

    for _, item := range mapping.items {
        value := source - item.source
        
        if value >= 0 && value < item.area {
            return processSeed(value + item.dest, mapping.next)
        }
    }
    
    return processSeed(source, mapping.next) 
}
