package main

import (
	"flag"
	"fmt"
	"os"
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
		os.MkdirAll(projectName, 0777)
	}

	// Generate project files
	err := os.Chdir(projectName)
	if err != nil {
    // Inside of ./projectName
	}
}
