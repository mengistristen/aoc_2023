package day19

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Day19 struct{}

type part struct {
	x int
	m int
	a int
	s int
}

type step struct {
	category rune
	operator rune
	operand  int
	workflow string
}

type stepFunc = func(p part, c string, i int) (string, int)

func (p part) getCategory(c rune) int {
	switch c {
	case 'x':
		return p.x
	case 'm':
		return p.m
	case 'a':
		return p.a
	case 's':
		return p.s
	}
	return 0
}

func (p part) sumCategories() int {
	return p.x + p.m + p.a + p.s
}

func (p *part) setCategory(c rune, v int) {
	switch c {
	case 'x':
		p.x = v
	case 'm':
		p.m = v
	case 'a':
		p.a = v
	case 's':
		p.s = v
	}
}

func (d Day19) Name() string {
	return "Day 19 - Aplenty"
}

func (d Day19) PartOne(ch chan string) {
	defer close(ch)
}

func (d Day19) PartTwo(ch chan string) {
	defer close(ch)
}

func (d Day19) ProcessPartOne(name string) (int, error) {
	sum := 0
	workflows, parts, err := processFile(name)
	if err != nil {
		return 0, err
	}

	for _, part := range parts {
		workflow, step := "in", 0

		for workflow != "A" && workflow != "R" {
			workflow, step = workflows[workflow][step](part, workflow, step)
		}

		if workflow == "A" {
			sum += part.sumCategories()
		}
	}

	log.Printf("workflows:\n%v\nparts:\n%v\n", workflows, parts)

	return sum, nil
}

func (d Day19) ProcessPartTwo(name string) (int, error) {
	return 0, nil
}

func processFile(name string) (map[string][]stepFunc, []part, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	workflows, parts, err := parseSystem(scanner)
	if err != nil {
		return nil, nil, err
	}

	return workflows, parts, nil
}

func parseSystem(scanner *bufio.Scanner) (map[string][]stepFunc, []part, error) {
	var parts []part
	workflows := make(map[string][]stepFunc)
	processWorkflows := true

	for scanner.Scan() {
		if scanner.Text() == "" {
			processWorkflows = false
		} else {
			if processWorkflows {
				split := strings.Split(scanner.Text(), "{")
				workflow := split[0]
				steps := strings.Split(split[1][:len(split[1])-1], ",")

				for _, step := range steps {
					if !strings.Contains(step, ":") {
						workflows[workflow] = append(workflows[workflow], unit(step))
					} else {
						step, err := parseStep(step)
						if err != nil {
							return nil, nil, err
						}
						workflows[workflow] = append(workflows[workflow], createStepFunc(step))
					}
				}
			} else {
				var part part
				categories := strings.Split(scanner.Text()[1:len(scanner.Text())-1], ",")

				for _, category := range categories {
					split := strings.Split(category, "=")
					category := rune(split[0][0])

					value, err := strconv.Atoi(split[1])
					if err != nil {
						return nil, nil, fmt.Errorf("error parsing category value: %v", err)
					}

					part.setCategory(category, value)
				}

				parts = append(parts, part)
			}
		}
	}

	return workflows, parts, nil
}

func parseStep(s string) (*step, error) {
	var step step

	step.category = rune(s[0])
	step.operator = rune(s[1])

	split := strings.Split(s[2:], ":")

	operand, err := strconv.Atoi(split[0])

	if err != nil {
		return nil, fmt.Errorf("error parsing operand: %v", err)
	}

	step.operand = operand
	step.workflow = split[1]

	return &step, nil
}

func withOperator(operator rune, operand int) func(int) bool {
	switch operator {
	case '<':
		return func(x int) bool {
			return x < operand
		}
	case '>':
		return func(x int) bool {
			return x > operand
		}
	}

	return nil
}

func createStepFunc(step *step) stepFunc {
	operatorFunc := withOperator(step.operator, step.operand)

	return func(p part, c string, i int) (string, int) {
		if operatorFunc(p.getCategory(step.category)) {
			return step.workflow, 0
		}
		return c, i + 1
	}
}

func unit(workflow string) stepFunc {
	return func(p part, c string, i int) (string, int) {
		return workflow, 0
	}
}
