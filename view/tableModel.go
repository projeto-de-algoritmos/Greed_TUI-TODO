package view

import (
	tb "github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/lipgloss"

	sch "tui-todo/scheduling"
)

type table struct {
	index int
	cols []tb.Column
	rows []tb.Row
	t tb.Model
}

func (t table) Init() tea.Cmd {
	return nil
}

func (t table) View() string {
	return lipgloss.Place(
		width,
		height,
		lipgloss.Center,
		lipgloss.Center,
		tableStyle.Render(t.t.View()) + "\n",
	)
}

func (t table) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			t.nextRow()
		case "k":
			t.previousRow()
		}
	}
	return t, cmd
}

func newTable() *table {
	cols := []tb.Column{
		{Title: "Tarefa", Width: 15},
		{Title: "Descrição", Width: 30},
		{Title: "Início", Width: 20},
		{Title: "Entrega", Width: 20},
		{Title: "Término", Width: 20},
	}

	ttable := tb.New(
		tb.WithColumns(cols),
	)
	t := new(table)
	t.t = ttable
	t.rows = make([]tb.Row, 0, 10)
	
	s := tb.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
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
			st.T.Deadline.Format(sch.DeadFormat),
		})
	}
	t.rows = newRows
	t.t.SetRows(newRows)
}

func (t *table) nextRow() {
	if t.index == len(t.rows)-1 {
		return	
	}
	t.index++
	t.t.SetCursor(t.index)
}

func (t *table) previousRow() {
	if t.index == 0 {
		return
	}
	t.index--
	t.t.SetCursor(t.index)
}
