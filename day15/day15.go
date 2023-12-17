package day15

import (
    "bufio"
    "fmt"
    "os"
    "slices"
    "strconv"
    "strings"
)

type Day15 struct {}

type lens struct {
    label string
    focalLength int
}

func (d Day15) Name() string {
    return "Day 15 - Lens Library"
}

func (d Day15) PartOne(ch chan string) {
    defer close(ch)

    if sum, err := d.ProcessPartOne("./input/day15_example.txt"); err == nil {
        ch <- fmt.Sprintf("Example output: %d\n", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one example: %v\n", err)
    }

    if sum, err := d.ProcessPartOne("./input/day15.txt"); err == nil {
        ch <- fmt.Sprintf("Output: %d", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one: %v", err)
    }
}

func (d Day15) PartTwo(ch chan string) {
    defer close(ch)

    if sum, err := d.ProcessPartTwo("./input/day15_example.txt"); err == nil {
        ch <- fmt.Sprintf("Example output: %d\n", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one example: %v\n", err)
    }

    if sum, err := d.ProcessPartTwo("./input/day15.txt"); err == nil {
        ch <- fmt.Sprintf("Output: %d", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one: %v", err)
    }
}

func (d Day15) ProcessPartOne(name string) (int, error) {
    sum := 0

    steps, err := processFile(name)
    if err != nil {
        return 0, err
    }

    for _, step := range steps {
        sum += calculateHash(step)
    }

    return sum, nil
}

func (d Day15) ProcessPartTwo(name string) (int, error) {
    focusingPower := 0

    steps, err := processFile(name)
    if err != nil {
        return 0, err
    }

    boxes := make([][]lens, 256)

    for _, step := range steps {
        var label string
        split := strings.Split(step, "=")

        if len(split) == 1 {
            label = split[0][:len(split[0])-1]
        } else {
            label = split[0]
        }

        hash := calculateHash(label)

        if len(split) == 1 {
            boxes[hash] = slices.DeleteFunc(boxes[hash], func(l lens) bool {
                return l.label == label
            })
        } else {
            focalLength, err := strconv.Atoi(split[1])
            if err != nil {
                return 0, fmt.Errorf("Error parsing focal length: %v", err)
            }

            if index := slices.IndexFunc(boxes[hash], func(l lens) bool {
                return l.label == label
            }); index != -1 {
                boxes[hash][index].focalLength = focalLength 
            } else {
                boxes[hash] = append(boxes[hash], lens{
                    label: label,
                    focalLength: focalLength,
                })
            }
        }
    }

    for i, box := range boxes {
        for j, lens := range box {
            focusingPower += (i + 1) * (j + 1) * lens.focalLength
        }
    }

    return focusingPower, nil
}

func processFile(name string) ([]string, error) {
    file, err := os.Open(name)
    if err != nil {
        return nil, fmt.Errorf("error opening file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    steps, err := parseSteps(scanner)
    if err != nil {
        return nil, err
    }

    return steps, nil
}

func parseSteps(scanner *bufio.Scanner) ([]string, error) {
    var steps []string

    scanner.Scan()

    steps = strings.Split(scanner.Text(), ",")

    return steps, nil
}

func calculateHash(step string) int {
    current := 0

    for _, char := range step {
        current += int(char)
        current *= 17
        current %= 256
    }

    return current
}
