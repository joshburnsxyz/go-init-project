package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	projectName string
	pkgDecl     string
)

func initFlags() {
	flag.StringVar(&projectName, "name", "", "Name for your project")
	flag.StringVar(&pkgDecl, "pkg", projectName, "The package declaration (github.com/username/mypackage)")
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
		fmt.Println(err)
	} else {
		fmt.Println("=> Creating main.go")
		fmt.Println("=> Initializing go modules")
		fmt.Println("=> Tidying modules")
		fmt.Println("=> Done")
	}
}
