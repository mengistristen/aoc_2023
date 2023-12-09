package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/mengistristen/aoc_2023/day0"
	"github.com/mengistristen/aoc_2023/day1"
	"github.com/mengistristen/aoc_2023/day2"
	"github.com/mengistristen/aoc_2023/day3"
	"github.com/mengistristen/aoc_2023/day4"
	"github.com/mengistristen/aoc_2023/day5"
	"github.com/mengistristen/aoc_2023/day6"
	"github.com/mengistristen/aoc_2023/day7"
	"github.com/mengistristen/aoc_2023/day8"
	"github.com/spf13/cobra"
)

func registerChallenges() {
	RegisterDay(day0.Day0{})
	RegisterDay(day1.Day1{})
	RegisterDay(day2.Day2{})
	RegisterDay(day3.Day3{})
	RegisterDay(day4.Day4{})
	RegisterDay(day5.Day5{})
	RegisterDay(day6.Day6{})
	RegisterDay(day7.Day7{})
	RegisterDay(day8.Day8{})
}

func setup() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "aoc_2023",
	}

	cmdUi := &cobra.Command{
		Use:   "ui",
		Short: "Show a TUI for running AOC challenges",
		Run:   runUi,
	}

	cmdRun := &cobra.Command{
		Use:   "run <day> <part> <input>",
		Short: "Run a single part of an AOC challenge",
		Run:   runChallenge,
	}

	rootCmd.AddCommand(cmdUi, cmdRun)

	return rootCmd
}

func runUi(cmd *cobra.Command, args []string) {
	RunApp()
}

func runChallenge(cmd *cobra.Command, args []string) {
	if len(args) < 3 {
		cmd.Help()
		os.Exit(1)
	}

	challenge, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("error parsing challenge: %v\n", err)
	}

	if challenge < 0 || challenge > len(days) {
		log.Fatalf("invalid day; available days: 0-%d\n", len(days))
	}

	part := args[1]

	if part != "1" && part != "2" {
		log.Fatalf("invalid part; choose part 1 or 2\n")
	}

	keys := make([]string, 0, len(days))
	for k := range days {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	day := days[keys[challenge]]

	if part == "1" {
		start := time.Now()
		if result, err := day.ProcessPartOne(args[2]); err != nil {
			fmt.Printf("error running day \"%s\": %v\n", day.Name(), err)
		} else {
			fmt.Printf("result: %d\n", result)
		}
		fmt.Printf("duration: %v\n", time.Since(start))
	} else if part == "2" {
		start := time.Now()
		if result, err := day.ProcessPartTwo(args[2]); err != nil {
			fmt.Printf("error running day \"%s\": %v\n", day.Name(), err)
		} else {
			fmt.Printf("result: %d\n", result)
		}
		fmt.Printf("duration: %v\n", time.Since(start))
	}
}

func main() {
	registerChallenges()

	command := setup()

	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
