package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	cform = iota
	ctable = iota
)

type model struct {
	state int
}

var f = newForm()
var t = newTable()

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	switch m.state {
	case cform:
		return f.View()
	case ctable:
		return t.View()
	}
	s := "model"
	return s
}

var width, height int

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		width = msg.Width
		height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+n":
			m.state = cform
		case "esc":
			m.state = ctable
		}
	}
	switch m.state {
	case cform:
		newView, newCmd := f.Update(msg)
		*f = newView.(form)
		cmd = newCmd
	case ctable:
		newView, newCmd := t.Update(msg)
		*t = newView.(table)
		cmd = newCmd
	}
	return m, cmd
}

func InitModel() model{
	return model{}
}
