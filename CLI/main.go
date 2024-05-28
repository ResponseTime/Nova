package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/responsetime/Nova/filehandle"
)

type menu struct {
	textInput   textinput.Model
	templates   []string
	languages   []string
	index       int
	screenCount int
	err         error
}

func initialModel() menu {
	ti := textinput.New()
	ti.Placeholder = "Default"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 30

	m := menu{
		textInput:   ti,
		templates:   []string{"React", "Vanilla", "Blank React"},
		languages:   []string{"Javascript", "Typescript"},
		index:       0,
		screenCount: 0,
		err:         nil,
	}
	return m
}

func (m menu) Init() tea.Cmd {
	return textinput.Blink
}

func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	last := len(m.templates) - 1
	if m.screenCount == 2 {
		last = len(m.languages) - 1
	}
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.index > 0 {
				m.index--
			} else if m.index == 0 {
				m.index = last
			}
		case "down", "j":
			if m.index < last {
				m.index++
			} else if m.index == last {
				m.index = 0
			}
		case "enter", " ":
			m.index = 0
			switch m.screenCount {
			case 0:
				projectName = m.textInput.Value()
				m.screenCount++
			case 1:
				template = m.templates[m.index]
				m.screenCount++
			case 2:
				language = m.languages[m.index]
				return m, tea.Quit
			}
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

var selectedOptionStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#faa"))

var helpMenuStyle = lipgloss.NewStyle().
	Faint(true)

var headerStyle = lipgloss.NewStyle().
	Bold(true)

func (m menu) View() string {

	if m.screenCount == 0 {
		return lipgloss.JoinVertical(lipgloss.Left, headerStyle.Render("Enter a project name\n"), m.textInput.View(), helpMenuStyle.Render("\nEnter/space to confirm or Ctrl+C to exit."))

	} else if m.screenCount == 1 {
		var options []string
		for i, o := range m.templates {
			if i == m.index {
				options = append(options, selectedOptionStyle.Render(fmt.Sprintf(" > %s", o)))
			} else {
				options = append(options, fmt.Sprintf("   %s", o))
			}
		}
		return lipgloss.JoinVertical(lipgloss.Left, headerStyle.Render("Select a template\n"), strings.Join(options, "\n"), helpMenuStyle.Render("\nEnter/space to select, arrow keys to navigate, or Ctrl+C to exit."))

	} else if m.screenCount == 2 {
		var options []string
		for i, o := range m.languages {
			if i == m.index {
				options = append(options, selectedOptionStyle.Render(fmt.Sprintf(" > %s", o)))
			} else {
				options = append(options, fmt.Sprintf("   %s", o))
			}
		}
		return lipgloss.JoinVertical(lipgloss.Left, headerStyle.Render("Select a language\n"), strings.Join(options, "\n"), helpMenuStyle.Render("\nEnter/space to select, arrow keys to navigate, or Ctrl+C to exit."))
	}
	return "nigasoda"
}

var projectName string
var template string
var language string

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	if projectName == "" {
		projectName = "Default"
	}

	filehandle.CREATE_PROJECT(dir, projectName, template, language)
	fmt.Printf("%s\n", "cd "+projectName)
	fmt.Printf("%s\n", "npm install")
	fmt.Printf("%s\n", "Nova run")

}
