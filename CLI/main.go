package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/responsetime/Nova/filehandle"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Error clearing screen:", err)
	}
}

func main() {
	var templt string
	var project_name string
	var language string
	templates := []string{"React", "Vanilla", "Blank React"}
	flag.StringVar(&templt, "template", "vanilla", "Enter the template name")
	flag.StringVar(&project_name, "project-name", "Default", "Enter the project name")
	flag.StringVar(&language, "language", "Javascript", "Enter the language to use")
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("Enter the project name")
		fmt.Scan(&project_name)
		clearScreen()
		fmt.Println("Enter the template")
		for _, i := range templates {
			fmt.Println(i)
		}
		fmt.Scan(&templt)
		clearScreen()
		fmt.Println("Enter Which Language to use")
		fmt.Println("Javascript")
		fmt.Println("Typescript")
		fmt.Scan(&language)
		clearScreen()
	}
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	filehandle.CREATE_PROJECT(dir, project_name, templt, language)
}
