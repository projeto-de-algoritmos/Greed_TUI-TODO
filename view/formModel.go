package view

import (
	"time"
	sch "tui-todo/scheduling"

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
		case "ctrl+p":
			f.previousForm()
		case "ctrl+n":
			f.nextForm()
		case "enter":
			if f.index == 3 {
				createTask()
				f.cleanForm()
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

func (f *form) previousForm() {
	if f.index == 0 {
		return
	}
	f.index--
	f.getCurrentInput().Focus()
}

func (f *form) cleanForm() {
	for i:=0; i<4; i++ {
		f.answerFields[i].SetValue("")
	}
	f.index = 0
	f.answerFields[f.index].Focus()
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

func createTask() {
	title := f.answerFields[0].Value()
	desc := f.answerFields[1].Value()
	dead, err1 := time.Parse(sch.DeadFormat, f.answerFields[2].Value())
	dur, err2 := time.ParseDuration(f.answerFields[3].Value())

	if err1 != nil || err2 != nil {
	} 
	task := sch.Task{
		Title: title,
		Description: desc,
		Deadline: dead,
		Duration: dur,
	}
	t.addTask(task)
}
