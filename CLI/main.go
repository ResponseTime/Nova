// package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"runtime"

// 	"github.com/responsetime/Nova/filehandle"
// )

// func clearScreen() {
// 	var cmd *exec.Cmd
// 	if runtime.GOOS == "windows" {
// 		cmd = exec.Command("cmd", "/c", "cls")
// 	} else {
// 		cmd = exec.Command("clear")
// 	}
// 	cmd.Stdout = os.Stdout
// 	if err := cmd.Run(); err != nil {
// 		fmt.Println("Error clearing screen:", err)
// 	}
// }

//	func main() {
//		var templt string
//		var project_name string
//		var language string
//		templates := []string{"React", "Vanilla", "Blank React"}
//		flag.StringVar(&templt, "template", "vanilla", "Enter the template name")
//		flag.StringVar(&project_name, "project-name", "Default", "Enter the project name")
//		flag.StringVar(&language, "language", "Javascript", "Enter the language to use")
//		flag.Parse()
//		if len(os.Args) < 2 {
//			fmt.Println("Enter the project name")
//			fmt.Scan(&project_name)
//			clearScreen()
//			fmt.Println("Enter the template")
//			for _, i := range templates {
//				fmt.Println(i)
//			}
//			fmt.Scan(&templt)
//			clearScreen()
//			fmt.Println("Enter Which Language to use")
//			fmt.Println("Javascript")
//			fmt.Println("Typescript")
//			fmt.Scan(&language)
//			clearScreen()
//		}
//		dir, err := os.Getwd()
//		if err != nil {
//			fmt.Println("Error getting current directory:", err)
//			return
//		}
//		filehandle.CREATE_PROJECT(dir, project_name, templt, language)
//		fmt.Printf("%s\n", "cd "+project_name)
//		fmt.Printf("%s\n", "npm install")
//		fmt.Printf("%s\n", "Nova run")
//	}
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
	textInput textinput.Model
	templates []string
	languages []string
	index     int
	err       error
}

func initialModel() menu {
	ti := textinput.New()
	ti.Placeholder = "Default"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	m := menu{
		textInput: ti,
		templates: []string{"React", "Vanilla", "Blank React"},
		languages: []string{"Javascript", "Typescript"},
		index:     0,
		err:       nil,
	}
	return m
}

func (m menu) Init() tea.Cmd {
	return textinput.Blink
}

func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	last := len(m.templates) - 1
	if screenCount == 2 {
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
			switch screenCount {
			case 0:
				projectName = m.textInput.Value()
				screenCount++
			case 1:
				template = m.templates[m.index]
				screenCount++
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

	if screenCount == 0 {
		return lipgloss.JoinVertical(lipgloss.Left, headerStyle.Render("Enter project name\n"), m.textInput.View(), helpMenuStyle.Render("\nEnter/space to confirm or Ctrl+C to exit."))

	} else if screenCount == 1 {
		var options []string
		for i, o := range m.templates {
			if i == m.index {
				options = append(options, selectedOptionStyle.Render(fmt.Sprintf(" > %s", o)))
			} else {
				options = append(options, fmt.Sprintf("   %s", o))
			}
		}
		return lipgloss.JoinVertical(lipgloss.Left, headerStyle.Render("Enter the template\n"), strings.Join(options, "\n"), helpMenuStyle.Render("\nEnter/space to select, arrow keys to navigate, or Ctrl+C to exit."))

	} else if screenCount == 2 {
		var options []string
		for i, o := range m.languages {
			if i == m.index {
				options = append(options, selectedOptionStyle.Render(fmt.Sprintf(" > %s", o)))
			} else {
				options = append(options, fmt.Sprintf("   %s", o))
			}
		}
		return lipgloss.JoinVertical(lipgloss.Left, headerStyle.Render("Enter the language\n"), strings.Join(options, "\n"), helpMenuStyle.Render("\nEnter/space to select, arrow keys to navigate, or Ctrl+C to exit."))
	}
	return "nigasoda"
}

var screenCount int = 0
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
	filehandle.CREATE_PROJECT(dir, projectName, template, language)
	fmt.Printf("%s\n", "cd "+projectName)
	fmt.Printf("%s\n", "npm install")
	fmt.Printf("%s\n", "Nova run")

}