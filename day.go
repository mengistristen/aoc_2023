package main

type Day interface {
	Name() string
	PartOne(chan string)
	PartTwo(chan string)
	ProcessPartOne(string) (int, error)
	ProcessPartTwo(string) (int, error)
}

var daysList []Day
var days = make(map[string]Day)

func RegisterDay(d Day) {
	days[d.Name()] = d
    daysList = append(daysList, d)
}

func ExecutePartOne(name string) string {
	output := ""
	ch := make(chan string)

	go days[name].PartOne(ch)

	for line := range ch {
		output += line
	}

	return output
}

func ExecutePartTwo(name string) string {
	output := ""
	ch := make(chan string)

	go days[name].PartTwo(ch)

	for line := range ch {
		output += line
	}

	return output
}
