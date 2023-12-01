package main

type Day0 struct{}

func (d Day0) Name() string {
	return "Day 0 - Example Day"
}

func (d Day0) PartOne(ch chan string) {
	defer close(ch)

	ch <- "Part one example output.\n"
	ch <- "Here is some other output.\n"
	ch <- "Should there be more?"
}

func (d Day0) PartTwo(ch chan string) {
	defer close(ch)

	ch <- "Example day part two. This is a really long line of text that may or may not fit withing the boundaries that are specified by the application.\n"
}

func init() {
	RegisterDay(Day0{})
}
