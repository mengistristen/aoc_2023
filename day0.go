package main

import "fmt"

type Day0 struct{}

func (d Day0) Name() string {
	return "Day 0 - Example Day"
}

func (d Day0) PartOne() (string, error) {
	s := fmt.Sprintln("Part one example output.")

	s += fmt.Sprintln("Here is some other output.")
	s += fmt.Sprint("Should there be more?")

	return s, nil
}

func (d Day0) PartTwo() (string, error) {
	s := fmt.Sprint("Example day part two")

	return s, nil
}

func init() {
	RegisterDay(Day0{})
}
