package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mengistristen/aoc_2023/day0"
	"github.com/mengistristen/aoc_2023/day1"
	"github.com/mengistristen/aoc_2023/day2"
	"github.com/mengistristen/aoc_2023/day3"
	"github.com/mengistristen/aoc_2023/day4"
	"github.com/mengistristen/aoc_2023/day5"
)

type keyMap struct {
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
	Quit   key.Binding
	Back   key.Binding
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Back, k.Quit},
	}
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Select, k.Back, k.Quit}
}

type model struct {
	chosen     bool
	selected   int
	challenges []string
	output     []string
	width      int
	height     int
	keys       keyMap
	help       help.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

		return m, nil
	}

	if m.chosen {
		return updateChosen(msg, m)
	}

	return updateChoices(msg, m)
}

func (m model) View() string {
	s := ""

	if m.chosen {
		s += viewChosen(m)
	} else {
		s += viewChoices(m)
	}

	return lipgloss.JoinVertical(lipgloss.Left, s, m.help.View(m.keys))
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
}

var (
	highlightColor = lipgloss.AdaptiveColor{Light: "#56f4f1", Dark: "#038a85"}
	borderedBox    = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1)
	selected = lipgloss.NewStyle().
			Foreground(highlightColor).
			Render
)

func updateChosen(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Back):
			m.chosen = false
		}
	}

	return m, nil
}

func updateChoices(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			if m.selected > 0 {
				m.selected--
			}
		case key.Matches(msg, m.keys.Down):
			if m.selected < len(m.challenges)-1 {
				m.selected++
			}
		case key.Matches(msg, m.keys.Select):
			m.output = make([]string, 2)

			m.output[0] = ExecutePartOne(m.challenges[m.selected])
			m.output[1] = ExecutePartTwo(m.challenges[m.selected])

			m.chosen = true
		}
	}

	return m, nil
}

func viewChosen(m model) string {
	columnWidth := (m.width / len(m.output)) - 2

	maxHeight := 0
	for _, content := range m.output {
		height := strings.Count(content, "\n")
		if height > maxHeight {
			maxHeight = height
		}
	}

	var views []string
	for _, content := range m.output {
		view := borderedBox.
			Width(columnWidth).
			Render(content)
		views = append(views, view)
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, views...)
}

func viewChoices(m model) string {
	var views []string

	for i, choice := range m.challenges {
		cursor := " "
		if m.selected == i {
			cursor = ">"
		}

		content := fmt.Sprintf("%s %s\n", cursor, choice)

		if m.selected == i {
			views = append(views, selected(content))
		} else {
			views = append(views, content)
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, views...)
}

func initialModel() model {
	var names []string
	for k := range days {
		names = append(names, k)
	}

	slices.Sort(names)

	return model{
		chosen:     false,
		selected:   0,
		challenges: names,
		keys:       keys,
		help:       help.New(),
	}
}

func main() {
	RegisterDay(day0.Day0{})
	RegisterDay(day1.Day1{})
	RegisterDay(day2.Day2{})
	RegisterDay(day3.Day3{})
	RegisterDay(day4.Day4{})
	RegisterDay(day5.Day5{})

	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Printf("error setting up logging: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err := p.Run(); err != nil {
		fmt.Printf("error running program: %v", err)
		os.Exit(1)
	}
}
