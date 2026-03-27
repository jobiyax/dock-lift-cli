package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	return "Dock Lift CLI\n\n(Appuyez sur une touche pour quitter)\n"
}

func main() {
	p := tea.NewProgram(model{})

	if _, err := p.Run(); err != nil {
		fmt.Println("Erreur:", err)
		os.Exit(1)
	}
}
