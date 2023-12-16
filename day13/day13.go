package day13

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
)

type Day13 struct {}
type mirror struct {
    rows []int
    cols []int
}

func (d Day13) Name() string {
    return "Day 13 - Point of Incidence"
}

func (d Day13) PartOne(ch chan string) {
    defer close(ch)

    if sum, err := d.ProcessPartOne("./input/day13_example.txt"); err == nil {
        ch <- fmt.Sprintf("Example output: %d\n", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one example: %v\n", err)
    }

    if sum, err := d.ProcessPartOne("./input/day13.txt"); err == nil {
        ch <- fmt.Sprintf("Output: %d", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one: %v", err)
    }
}

func (d Day13) PartTwo(ch chan string) {
    defer close(ch)

    if sum, err := d.ProcessPartTwo("./input/day13_example.txt"); err == nil {
        ch <- fmt.Sprintf("Example output: %d\n", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one example: %v\n", err)
    }

    if sum, err := d.ProcessPartTwo("./input/day13.txt"); err == nil {
        ch <- fmt.Sprintf("Output: %d", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one: %v", err)
    }
}

func (d Day13) ProcessPartOne(name string) (int, error) {
    sum := 0

    mirrors, err := processFile(name)
    if err != nil {
        return 0, err
    }

    for _, mirror := range mirrors {
        cols := detectMirrors(mirror.cols)
        rows := detectMirrors(mirror.rows)

        log.Printf("(%d, %d)", rows, cols)

        sum += cols 
        sum += 100 * rows
    }

    return sum, nil
}

func (d Day13) ProcessPartTwo(name string) (int, error) {
    sum := 0

    mirrors, err := processFile(name)
    if err != nil {
        return 0, err
    }

    for i, mirror := range mirrors {
        prevCols := detectMirrors(mirror.cols)
        prevRows := detectMirrors(mirror.rows)
        cols := detectSmudgedMirrors(mirror.cols, prevCols)
        rows := detectSmudgedMirrors(mirror.rows, prevRows)

        if rows == 0 && cols == 0 {
            log.Printf("index: %d", i)
            printMirror(mirror)
            log.Println("")
        }

            sum += cols 
            sum += 100 * rows
    }

    return sum, nil
}

func processFile(name string) ([]mirror, error) {
    file, err := os.Open(name)
    if err != nil {
        return nil, fmt.Errorf("error opening file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    mirrors, err := parseMirrors(scanner)
    if err != nil {
        return nil, err
    }

    return mirrors, nil
}

func parseMirrors(scanner *bufio.Scanner) ([]mirror, error) {
    var mirrors []mirror
    current := mirror{}
    row := 0

    for scanner.Scan() {
        if scanner.Text() == "" {
            row = 0 
            mirrors = append(mirrors, current)
            current = mirror{}
        } else {
            value := 0

            if current.cols == nil {
                current.cols = make([]int, len(scanner.Text()))
            }

            for i, ch := range scanner.Text() {
                if ch == '#' {
                    value |= (1 << i) 
                    current.cols[i] |= (1 << row)
                }
            }

            current.rows = append(current.rows, value)
            row++
        }
    }

    mirrors = append(mirrors, current)

    return mirrors, nil
}

func detectMirrors(values []int) int {
    var visited []int
    top := -1
    result := 0

    for i, value := range values {
        if top == len(visited) - 1 {
            if top == -1 || value != visited[top] {
                top++ 
            } else {
                top-- 
                result = i
            }
        } else {
            if top < 0 {
                return result 
            }

            if visited[top] == value {
                top--
            } else {
                top = len(visited)
                result = 0
            }
        }
        visited = append(visited, value) 
    }

    return result 
}

func detectSmudgedMirrors(values []int, prev int) int {
    var visited []int
    checkMirror := false
    smudge := false
    top := -1 
    result := 0

    for i := 0; i < len(values); i++ { 
        if checkMirror {
            if top < 0 {
                return result
            } else {
                if values[i] == visited[top] {
                    top--
                } else if !smudge && possibleSmudge(values[i], visited[top]) {
                    top--
                    smudge = true
                } else {
                    i, top = result, result
                    checkMirror = false
                    result = 0
                }
            }
        } else {
            if top == -1 || i == prev {
                top++
            } else if equal := values[i] == visited[top]; equal || possibleSmudge(values[i], visited[top]) {
                smudge = !equal && possibleSmudge(values[i], visited[top])
                top--
                result = i
                checkMirror = true
            } else {
                top++
            }
        }

        if len(visited) == i {
            visited = append(visited, values[i])
        }
    }

    return result 
}

func possibleSmudge(a, b int) bool {
    adjusted := a ^ b
    logrithm := math.Log2(float64(adjusted)) 

    return logrithm == math.Floor(logrithm)
}

func printMirror(mirror mirror) {
    for _, row := range mirror.rows {
        str := "" 

        for i := 0; i < len(mirror.cols); i++ {
            if ((1 << i) & row) != 0 {
                str += "#"
            } else {
                str += "."
            }
        }

        log.Printf("%s", str)
    }
}
