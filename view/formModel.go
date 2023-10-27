package view

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type form struct {
	index int
	questions [4]string
	answerFields [4]textinput.Model
	styles *style
}

func (f form) Init() tea.Cmd {
	return nil
}

func (f form) View() string {
	var s string
	for i := range f.questions {
		s += lipgloss.JoinVertical(
			lipgloss.Center,
			"\n",
			f.questions[i],
			f.styles.InputField.Render(f.answerFields[i].View()),
		)
	}
	return lipgloss.Place(
		width,
		height,
		lipgloss.Center,
		lipgloss.Center,
		s,
	)
}

func (f form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type){
	case tea.KeyMsg:
		switch msg.String(){
		case "enter":
			if f.index == 3 {
				t.addTask(f.createTask())
			}else{
				f.nextForm()
			}
		return f, nil
		}
	}
	f.answerFields[f.index], cmd = f.answerFields[f.index].Update(msg)
	return f, cmd
}

func (f *form) nextForm() {
	if f.index == 3 {
		return
	}	
	f.index++
	f.getCurrentInput().Focus()
}

func (f *form) getCurrentInput() *textinput.Model {
	return &f.answerFields[f.index]
}

func newForm() *form{
	s := defaultStyle()
	f := new(form)
	questions := [4]string{
		"Nome da Tarefa",
		"Descrição da Tarefa",
		"Data de Entrega",
		"Duração",
	}
	place := [4]string{
		"",
		"",
		"Use esse formato: dd/MM/aaaa hh/mm",
		"Ex: 10h",
	}
	f.questions = questions
	for i:=0; i<4;i++ {
		f.answerFields[i] = textinput.New()
		f.answerFields[i].Placeholder = place[i]
	}
	f.answerFields[0].Focus()
	f.styles = s
	
	return f
}

func (f *form) createTask() task{
	return task{
		f.answerFields[0].Value(),
		f.answerFields[1].Value(),
		f.answerFields[2].Value(),
		f.answerFields[3].Value(),
	}	
}
