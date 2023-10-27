package view

import (
	tea "github.com/charmbracelet/bubbletea"
	tb "github.com/charmbracelet/bubbles/table"

	"github.com/charmbracelet/lipgloss"
)

type task struct {
	title string
	description string
	deadline string
	duration string
}

type table struct {
	cols []tb.Column
	rows []tb.Row
	t tb.Model
}

func (t table) Init() tea.Cmd {
	return nil
}

func (t table) View() string {
	return t.t.View()+"\n"
}


func (t table) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}

func newTable() *table {
	cols := []tb.Column{
		{Title: "Tarefa", Width: 10},
		{Title: "Descrição", Width: 20},
		{Title: "Entrega", Width: 30},
		{Title: "Duração", Width: 30},
	}

	ttable := tb.New(
		tb.WithColumns(cols),
	)
	t := new(table)
	t.t = ttable
	t.rows = make([]tb.Row, 0, 10)
	
	s := tb.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder())
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("11")).
		Bold(false)
	t.t.SetStyles(s)
	return t
}

func (t *table) addTask (tk task) {
	t.rows = append(t.rows, tb.Row{
		tk.title,
		tk.description,
		tk.duration,
		tk.deadline,
	})
	t.t.SetRows(t.rows)
}
