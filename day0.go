package main

import "fmt"

type Day0 struct{}

func (d Day0) Name() string {
	return "Day 0 - Example Day"
}

func (d Day0) PartOne() (string, error) {
	s := fmt.Sprintf("Example day part one\n")

	return s, nil
}

func (d Day0) PartTwo() (string, error) {
	s := fmt.Sprintf("Example day part two\n")

	return s, nil
}

func init() {
	RegisterDay(Day0{})
}
