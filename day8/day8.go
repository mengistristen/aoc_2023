package day8

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Day8 struct{}

type graph struct {
	instructions []int8
	nodes        map[string]int
	edges        [][2]int
}

func (d Day8) Name() string {
	return "Day 8 - Haunted Wasteland"
}

func (d Day8) PartOne(ch chan string) {
	defer close(ch)

	if sum, err := processPartOne("./input/day8_example1.txt"); err == nil {
		ch <- fmt.Sprintf("Example 1 output: %d\n", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one example: %v\n", err)
	}

	if sum, err := processPartOne("./input/day8_example2.txt"); err == nil {
		ch <- fmt.Sprintf("Example 2 output: %d\n", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one example: %v\n", err)
	}

	if sum, err := processPartOne("./input/day8.txt"); err == nil {
		ch <- fmt.Sprintf("Output: %d", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one: %v", err)
	}
}

func (d Day8) PartTwo(ch chan string) {
	defer close(ch)

	if sum, err := processPartTwo("./input/day8_example3.txt"); err == nil {
		ch <- fmt.Sprintf("Example output: %d\n", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one example: %v\n", err)
	}

	if sum, err := processPartTwo("./input/day8.txt"); err == nil {
		ch <- fmt.Sprintf("Output: %d", sum)
	} else {
		ch <- fmt.Sprintf("error processing part one: %v", err)
	}
}

func processPartOne(name string) (int, error) {
	steps := 0

	graph, err := processFile(name)
	if err != nil {
		return 0, err
	}

	current := graph.nodes["AAA"]

	index := 0

	for current != graph.nodes["ZZZ"] {
		current = graph.edges[current][graph.instructions[index]]

		steps++
		index++

		if index >= len(graph.instructions) {
			index = 0
		}
	}

	return calculateSteps(graph, graph.nodes["AAA"], func(end int) bool {
		return end == graph.nodes["ZZZ"]
	}), nil
}

func processPartTwo(name string) (int, error) {
	graph, err := processFile(name)
	if err != nil {
		return 0, err
	}

	var startNodes []int
	var endNodes []int
	var results []int

	for key := range graph.nodes {
		if key[2] == 'A' {
			startNodes = append(startNodes, graph.nodes[key])
		}
	}

	for key := range graph.nodes {
		if key[2] == 'Z' {
			endNodes = append(endNodes, graph.nodes[key])
		}
	}

	for _, node := range startNodes {
		results = append(results, calculateSteps(graph, node, func(end int) bool {
			return slices.Contains(endNodes, end)
		}))
	}

	return lcm(results), nil
}

func processFile(name string) (*graph, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	games, err := parseGraph(scanner)
	if err != nil {
		return nil, err
	}

	return games, nil
}

func parseGraph(scanner *bufio.Scanner) (*graph, error) {
	graph := graph{
		nodes: make(map[string]int),
	}

	scanner.Scan()

	for _, char := range scanner.Text() {
		if char == 'L' {
			graph.instructions = append(graph.instructions, 0)
		} else {
			graph.instructions = append(graph.instructions, 1)
		}
	}

	scanner.Scan()

	var edges [][2]string
	index := 0

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "=")
		node := strings.TrimSpace(split[0])

		graph.nodes[node] = index

		edges = append(edges, [2]string(strings.Split(strings.Trim(split[1], " ()"), ",")))
		index++
	}

	for _, edge := range edges {
		graph.edges = append(graph.edges, [2]int{graph.nodes[strings.TrimSpace(edge[0])], graph.nodes[strings.TrimSpace(edge[1])]})
	}

	return &graph, nil
}

func calculateSteps(g *graph, start int, end func(int) bool) int {
	steps := 0
	current := start
	index := 0

	for !end(current) {
		current = g.edges[current][g.instructions[index]]

		steps++
		index++

		if index >= len(g.instructions) {
			index = 0
		}
	}

	return steps
}

func lcm(n []int) int {
	current := n[0]

	for i := 1; i < len(n); i++ {
		current = current / gcd(current, n[i]) * n[i]
	}

	return current
}

func gcd(a, b int) int {
	result := min(a, b)

	for result > 1 {
		if (a%result) == 0 && (b%result) == 0 {
			break
		}
		result--
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
