package view

import (
	tb "github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/lipgloss"

	sch "tui-todo/scheduling"
)

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
		{Title: "Início", Width: 20},
		{Title: "Entrega", Width: 30},
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

func (t *table) addTask (tk sch.Task) {
	task = append(task, tk)
	schTask := sch.Scheduling(task)

	var newRows []tb.Row
	for _, st := range schTask {
		newRows = append(newRows, tb.Row{
			st.T.Title,
			st.T.Description,
			st.Start.Format(sch.DeadFormat),
			st.End.Format(sch.DeadFormat),

		})
	}
	t.rows = newRows
	t.t.SetRows(t.rows)
}
