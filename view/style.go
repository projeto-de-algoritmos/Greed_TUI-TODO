package view

import "github.com/charmbracelet/lipgloss"

type style struct {
	BorderColor lipgloss.Color
	InputField lipgloss.Style
}

func defaultStyle() *style{
	s := new(style)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().
	BorderForeground(s.BorderColor).
	BorderStyle(lipgloss.NormalBorder()).
	Padding(1).Width(80)
	return s
}

var tableStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))
