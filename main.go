package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	chosen     bool
	selected   int
	challenges []string
	output     string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "ctrl+c" {
			return m, tea.Quit
		}
	}

	if m.chosen {
		return updateChosen(msg, m)
	}

	return updateChoices(msg, m)
}

func updateChosen(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.chosen = false
		}
	}

	return m, nil
}

func updateChoices(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.selected > 0 {
				m.selected--
			}
		case "down", "j":
			if m.selected < len(m.challenges)-1 {
				m.selected++
			}
		case "enter":
			s := "Part One:\n\n"
			s += ExecutePartOne(m.challenges[m.selected])

			s += "\nPart Two:\n\n"
			s += ExecutePartTwo(m.challenges[m.selected])

			m.output = s
			m.chosen = true
		}
	}

	return m, nil
}

func (m model) View() string {
	s := ""

	if m.chosen {
		s += viewChosen(m)
	} else {
		s += viewChoices(m)
	}

	s += "\nPress q to quit.\n"

	return s
}

func viewChosen(m model) string {
	return m.output
}

func viewChoices(m model) string {
	s := ""

	for i, choice := range m.challenges {
		cursor := " "
		if m.selected == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s [%s]\n", cursor, choice)
	}

	return s
}

func initialModel() model {
	keys := make([]string, 0, len(days))
	for k := range days {
		keys = append(keys, k)
	}

	return model{
		chosen:     false,
		selected:   0,
		challenges: keys,
	}
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("error running program: %v", err)
		os.Exit(1)
	}
}
