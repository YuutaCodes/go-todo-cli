package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	choices  *Todos
	cursor   int
	selected map[int]struct{}
	adding   bool
	input    string
}

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	todos := Todos{}
	p := tea.NewProgram(initialModel(&todos))
	if _, err := p.Run(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

func initialModel(todos *Todos) model {
	return model{
		choices:  todos,
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() tea.View {
	var s strings.Builder
	s.WriteString("Todo list\n\n")

	cursor := " "
	if m.cursor == 0 {
		cursor = ">"
	}
	if m.adding {
		s.WriteString(fmt.Sprintf("%s Add task: %s_\n", cursor, m.input))
	} else {
		s.WriteString(fmt.Sprintf("%s [+] Add task\n", cursor))
	}

	for i, choice := range *m.choices {
		cursor := " "
		if m.cursor == i+1 {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Title))
	}
	s.WriteString("\nPress q to quit. ")
	if m.adding {
		s.WriteString("Press esc to go back.")
	}
	return tea.NewView(s.String())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return m, nil
	}

	if m.adding {
		switch key.String() {
		case "enter":
			if m.input != "" {
				m.choices.Add(m.input)
			}
			m.input, m.adding = "", false
		case "esc":
			m.input, m.adding = "", false
		case "backspace":
			if r := []rune(m.input); len(r) > 0 {
				m.input = string(r[:len(r)-1])
			}
		default:
			m.input += key.Text
		}
		return m, nil
	}

	switch key.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(*m.choices) {
			m.cursor++
		}
	case "enter", "space":
		if m.cursor == 0 {
			m.adding = true
			break
		}
		i := m.cursor - 1
		if _, ok := m.selected[i]; ok {
			delete(m.selected, i)
		} else {
			m.selected[i] = struct{}{}
		}
	}
	return m, nil
}
