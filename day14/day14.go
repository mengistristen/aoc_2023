package day14

import (
	"bufio"
	"fmt"
	"os"
)

type Day14 struct {}
type rock struct {
    x int
    y int
}

func (d Day14) Name() string {
    return "Day 14 - Parabolic Reflector Disk" 
}

func (d Day14) PartOne(ch chan string) {
    defer close(ch)
}

func (d Day14) PartTwo(ch chan string) {
    defer close(ch)
}

func (d Day14) ProcessPartOne(name string) (int, error) {
    load := 0

    round, cube, err := processFile(name)
    if err != nil {
        return 0, err
    }

    for _, rock := range round {
        row := cube[rock.y][rock.x]

        for next := row + 1; next < len(cube); next++ {
            if cube[next][rock.x] != row {
                break
            }

            cube[next][rock.x] = row + 1
        }

        load += len(cube) - row
    }

    return load, nil
}

func (d Day14) ProcessPartTwo(name string) (int, error) {
    return 0, nil
}

func processFile(name string) ([]rock, [][]int, error) {
    file, err := os.Open(name)
    if err != nil {
        return nil, nil, fmt.Errorf("error opening file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    round, cube, err := parseRocks(scanner)
    if err != nil {
        return nil, nil, err
    }

    return round, cube, nil
}

func parseRocks(scanner *bufio.Scanner) ([]rock, [][]int, error) {
    var round []rock
    var cube [][]int
    var latest []int

    for row := 0; scanner.Scan(); row++ {
        if latest == nil {
            latest = make([]int, len(scanner.Text()))
        }

        for col, ch := range scanner.Text() {
            if ch == 'O' {
                round = append(round, rock{
                    x: col,
                    y: row,
                })
            } else if ch == '#' {
                latest[col] = row + 1
            }
        }

        cube = append(cube, make([]int, len(latest)))
        
        copy(cube[row], latest)
    }

    return round, cube, nil
}
