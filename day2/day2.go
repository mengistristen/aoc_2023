package day2

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Day2 struct {}

type game struct {
    id int
    sessions []session
}

type session struct {
    cubes map[string]int
}

func (d Day2) Name() string {
    return "Day 2 - Cube Conundrum"
}

func (d Day2) PartOne(ch chan string) {
    defer close(ch)

    if sum, err := processPartOne("./input/day2_example.txt"); err == nil {
        ch <- fmt.Sprintf("Example output: %d\n", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one example: %v\n", err)
    }

    if sum, err := processPartOne("./input/day2.txt"); err == nil {
        ch <- fmt.Sprintf("Output: %d", sum)
    } else {
        ch <- fmt.Sprintf("error processing part one: %v", err)
    }
}

func (d Day2) PartTwo(ch chan string) {
    defer close(ch)

    if sum, err := processPartTwo("./input/day2_example.txt"); err == nil {
        ch <- fmt.Sprintf("Example output: %d\n", sum)
    } else {
        ch <- fmt.Sprintf("error processing part two example: %v\n", err)
    }

    if sum, err := processPartTwo("./input/day2.txt"); err == nil {
        ch <- fmt.Sprintf("Example output: %d", sum)
    } else {
        ch <- fmt.Sprintf("error processing part two: %v", err)
    }
}

func processPartOne(name string) (int, error) {
    sum := 0

    games, err := processFile(name)
    if err != nil {
        return 0, err
    }

    cubes := map[string]int{"red": 12, "green": 13, "blue": 14}

    for _, game := range games {
        if validateGame(game, cubes) {
            sum += game.id
        }
    }

    return sum, nil
}

func processPartTwo(name string) (int, error) {
    sum := 0

    games, err := processFile(name)
    if err != nil {
        return 0, err
    }

    for _, game := range games {
        sum += computePower(game)
    }

    return sum, nil
}

func processFile(name string) ([]game, error) {
    file, err := os.Open(name)
    if err != nil {
        return nil, fmt.Errorf("error opening file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    games, err := parseGames(scanner)
    if err != nil {
        return nil, err
    }

    return games, nil
} 

func parseGames(scanner *bufio.Scanner) ([]game, error) {
    var games []game

    for scanner.Scan() {
        line := strings.Split(scanner.Text(), ":")
        game_str := line[0]
        sessions_str := line[1]        

        id, err := strconv.Atoi(strings.Split(game_str, " ")[1])
        if err != nil {
            return nil, fmt.Errorf("failed to parse game id")
        }

        current_game := game {
            id: id,
        }

        for _, sessions := range strings.Split(sessions_str, ";") {
            current_session := session{
                cubes: make(map[string]int),
            }

            for _, cube := range strings.Split(sessions, ",") {
                cube = strings.TrimSpace(cube)
                parsed := strings.Split(cube, " ")

                value, exists := current_session.cubes[parsed[1]]
                if !exists {
                    value = 0
                }

                num_cubes, err := strconv.Atoi(parsed[0])
                if err != nil {
                    return nil, fmt.Errorf("failed to parse cube count")
                }

                current_session.cubes[parsed[1]] = value + num_cubes 
            }

            current_game.sessions = append(current_game.sessions, current_session)
        }

        games = append(games, current_game)
    }

    return games, nil
}

func validateGame(g game, cubes map[string]int) bool {
    for _, session := range g.sessions {
        for key := range session.cubes {
            if value, exists := cubes[key]; exists {
                if session.cubes[key] > value {
                    return false
                }
            } else {
                return false
            }
        }
    }

    return true
}

func computePower(g game) int {
    cube_max := make(map[string]int)

    for _, session := range g.sessions {
        for key := range session.cubes {
            if value, exists := cube_max[key]; exists {
                if session.cubes[key] > value {
                    cube_max[key] = session.cubes[key]
                }
            } else {
                cube_max[key] = session.cubes[key]
            }
        }
    }

    power := 1

    for key := range cube_max {
        power *= cube_max[key]
    }

    return power
}
