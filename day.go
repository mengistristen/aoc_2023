package main

import "fmt"

type Day interface {
	Name() string
	PartOne() (string, error)
	PartTwo() (string, error)
}

var days = make(map[string]Day)

func RegisterDay(d Day) {
	days[d.Name()] = d
}

func ExecutePartOne(name string) string {
	if output, err := days[name].PartOne(); err != nil {
		return fmt.Sprintf("%v", err)
	} else {
		return output
	}
}

func ExecutePartTwo(name string) string {
	if output, err := days[name].PartTwo(); err != nil {
		return fmt.Sprintf("%v", err)
	} else {
		return output
	}
}
