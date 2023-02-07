package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

  "github.com/bitfield/script"
)

var (
	projectName string
)

func initFlags() {
	flag.StringVar(&projectName, "name", "", "Name for your project")
}

func main() {

	// Parse CLI flags
	initFlags()
	flag.Parse()

	// Create project directory
  if _, err := os.Stat(projectName); os.IsNotExist(err) {
    os.MkdirAll(projectName, 0644)
  }

	// Generate project files
	err := os.Chdir(projectName)
	if err != nil {
		fmt.Println("> ERROR: Could not cd into new project directory, wtf?")
	} else {
		createMainFileCmd := exec.Command("touch", "main.go")
		initGoModuleCmd := exec.Command(fmt.Sprintf("go mod init %s", projectName))

		err := createMainFileCmd.Run()
		if err == nil {
			script.File("main.go").WriteFile("package main")
			err := initGoModuleCmd.Run()
			if err == nil {
				fmt.Println("done")
			} else {
				fmt.Println("> ERROR: Could not initialize go module")
				os.Exit(2)
			}
		} else {
			fmt.Println("> ERROR: Could not create main.go")
			os.Exit(2)
		}
	}
}
