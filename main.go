package main

import (
	"log"
	mo "tui-todo/view"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(mo.InitModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
